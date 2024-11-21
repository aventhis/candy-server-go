package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	candyType := flag.String("k", "", "Candy type (e.g., AA)")
	candyCount := flag.Int("c", 0, "Number of candies to buy")
	money := flag.Int("m", 0, "Amount of money given")

	if *candyType == "" || *candyCount <= 0 || *money <= 0 {
		fmt.Println("Usage: ./candy-client -k <CandyType> -c <CandyCount> -m <Money>")
		os.Exit(1)
	}
	//
	//// загрузка клиентского сертификата
	//cert, err := tls.LoadX509KeyPair("client.tld/client-cert.pem", "client.tld/client-key.pem")
	//if err != nil {
	//	log.Fatalf("Failed to load client key pair: %s", err)
	//}
	//
	//// загрузка корневого сертификата
	//caCert, err := os.ReadFile("minica.pem")
	//if err != nil {
	//	log.Fatalf("Failed to read CA cert: %s", err)
	//}
	//
	//// cоздание пула сертификатов (CertPool)
	//caCertPool := x509.NewCertPool()
	//if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
	//	log.Fatalf("Failed to append CA cert")
	//}
	//
	//// настройка TLS-конфигурации
	//tlsConfig := &tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	RootCAs:      caCertPool,
	//}
	//
	//// cоздание HTTP-клиента с кастомным транспортом
	//transport := &http.Transport{TLSClientConfig: tlsConfig}
	//client := &http.Client{Transport: transport}
	//
	//// формирование и отправка запроса

}
