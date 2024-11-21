SERVER=candy-server
CLIENT=candy-client
PORT=3333
TLS_CERT=candy.tld/cert.pem
TLS_KEY=candy.tld/key.pem
TLS_CA=minica.pem
SERVER_CERT_DIR=candy.tld
CA_CERT=minica.pem
CA_KEY=minica-key.pem
CLIENT_CERT_DIR=client.tld
CLIENT_KEY=${CLIENT_CERT_DIR}/client-key.pem
CLIENT_CSR=${CLIENT_CERT_DIR}/client.csr
CLIENT_CERT=${CLIENT_CERT_DIR}/client-cert.pem

# Сборка сервера и клиента
all: clean build-server build-client

# Сборка и запуск сервера
start-server: clean build-server run-server

# Сборка и запуск клиента
start-client: clean build-client run-client

# Генерация сертификатов для сервера и клиента
generate-certs: generate-server-certs generate-client-certs

# Генерация сертификатов для сервера через minica
generate-server-certs:
	@echo "Generating server certificates with minica..."
	@mkdir -p ${SERVER_CERT_DIR}
	@minica --domains candy.tld
	@echo "Server certificates generated in ${SERVER_CERT_DIR}"

# Генерация сертификатов для клиента
generate-client-certs:
	@echo "Generating client certificates..."
	@mkdir -p ${CLIENT_CERT_DIR}
	@openssl genrsa -out ${CLIENT_KEY} 2048
	@openssl req -new -key ${CLIENT_KEY} -out ${CLIENT_CSR} -subj "/C=RU/ST=State/L=City/O=Organization/CN=client.tld"
	@openssl x509 -req -in ${CLIENT_CSR} -CA ${CA_CERT} -CAkey ${CA_KEY} -CAcreateserial -out ${CLIENT_CERT} -days 365
	@echo "Client certificates generated:"
	@echo "  - Key:  ${CLIENT_KEY}"
	@echo "  - CSR:  ${CLIENT_CSR}"
	@echo "  - Cert: ${CLIENT_CERT}"

# Сборка сервера
build-server:
	@echo "Building server..."
	@go build -o ${SERVER} cmd/candy-server-server/main.go

# сборка клиента
build-client:
	@echo "Building client..."
	@go build -o ${CLIENT} candy-client.go

# Запуск сервера
run-server:
	@echo "Starting server on port ${PORT}..."
	@./${SERVER} --tls-port=${PORT} --tls-certificate=${TLS_CERT} --tls-key=${TLS_KEY} --tls-ca=${TLS_CA}

# Запуск клиента
run-client:
	@echo "Starting client..."
	@./${CLIENT} -k AA -c 2 -m 50

# Очистка бинарных файлов
clean:
	@echo "Cleaning up binaries..."
	@rm -f ${SERVER} ${CLIENT}

# Очистка всех сертификатов
clean-certs:
	@echo "Cleaning up all certificates..."
	@rm -rf *.pem *.srl ${CLIENT_CERT_DIR} ${SERVER_CERT_DIR}
	@echo "Certificates removed."