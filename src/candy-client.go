package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	candyType := flag.String("k", "", "Candy type (e.g., AA)")
	candyCount := flag.Int("c", 0, "Number of candies to buy")
	money := flag.Int("m", 0, "Amount of money given")

	flag.Parse()

	if *candyType == "" || *candyCount <= 0 || *money <= 0 {
		fmt.Println("Usage: ./candy-client -k <CandyType> -c <CandyCount> -m <Money>")
		os.Exit(1)
	}

	// загрузка клиентского сертификата
	cert, err := tls.LoadX509KeyPair("client.tld/client-cert.pem", "client.tld/client-key.pem")
	if err != nil {
		log.Fatalf("Failed to load client key pair: %s", err)
	}

	// загрузка корневого сертификата
	caCert, err := os.ReadFile("minica.pem")
	if err != nil {
		log.Fatalf("Failed to read CA cert: %s", err)
	}

	// cоздание пула сертификатов (CertPool)
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("Failed to append CA cert")
	}

	// настройка TLS-конфигурации
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	// cоздание HTTP-клиента с кастомным транспортом
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// формирование и отправка запроса
	url := "https://localhost:3333/buy_candy"
	jsonData := fmt.Sprintf(`{"money": %d,"candyType": "%s", "candyCount": %d}`, *money, *candyType, *candyCount)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	// читаем тело ответа и выводим результат
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}
	if resp.StatusCode == http.StatusCreated {
		// Парсим JSON и выводим сообщение об успехе
		var successREsponse struct {
			Thanks string `json:"thanks"`
			Change int64  `json:"change"`
		}
		if err := json.Unmarshal(body, &successREsponse); err != nil {
			log.Fatalf("Failed to unmarshal response: %s", err)
		}
		fmt.Printf("%s! Your change is %d\n", successREsponse.Thanks, successREsponse.Change)
	} else {
		// смотрим JSON и выводим сообщение об ошибке
		var errorResponse struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Fatalf("Failed to parse error response: %v", err)
		}
		fmt.Printf("Error: %s\n", errorResponse.Error)
	}
}
