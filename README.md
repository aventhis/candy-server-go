# Candy API Server

This project is the result of completing the **"Day 04 - Go Boot Camp"** task. The main goal is to develop an API server for candy vending machines. The task includes:

- Implementing an API that meets the requirements of the Swagger protocol.
- Supporting functionality for input validation, change calculation, and error handling.
- Adding mutual TLS authentication for security.
- Integrating an external ASCII art generator.

## Technologies

- Go
- Swagger (go-swagger)
- OpenSSL
- TLS/SSL

## Available Candies and Their Prices

| Code | Name                | Price (cents)     |
|------|---------------------|-------------------|
| CE   | Cool Eskimo         | 10               |
| AA   | Apricot Aardvark    | 15               |
| NT   | Natural Tiger       | 17               |
| DE   | Dazzling Elderberry | 21               |
| YR   | Yellow Rambutan     | 23               |

## Features

1. **Candy Purchase**:
   - The client sends a request to the server with data about the amount of money, candy type, and quantity.
   - The server calculates the change or returns an error:
      - If the money is sufficient, the response includes "Thank you!" and the change.
      - If the money is insufficient, the response contains an error indicating the amount needed.
      - If the input is invalid (negative quantity or nonexistent candy type), an error is returned.

2. **Security**:
   - Mutual TLS authentication is implemented using self-signed certificates.
   - Both server and client use certificates signed by a local Certificate Authority (CA) created with [Minica](https://github.com/jsha/minica).

3. **Code Generation**:
   - Server code and models are automatically generated from the Swagger 2.0 specification using [go-swagger](https://github.com/go-swagger/go-swagger).

## Installation and Execution

### Step 1. Clone the Repository
Clone the repository:
```bash
git clone git@github.com:aventhis/candy-server-go.git
cd candy-server-go/src
```

### Step 2. Install Dependencies
Ensure Go is installed, then run:
```bash
go mod tidy
```

### Step 3. Generate Server Code from Swagger Specification (if needed)
   ```bash
   swagger generate server -f swagger.yaml -a candy-server
   ```

### Step 4. Using Makefile
For simplifying development and testing, the project includes a `Makefile`. Main commands:

- **Build the server and client:**
  ```bash
  make all
  ```
- **Run the server:**
  ```bash
  make start-server
  ```
- **Run the client:**
  ```bash
  make start-client
  ```
- **Generate certificates for the server and client:**
  ```bash
  make generate-certs
  ```
- **Clean binaries and certificates:**
  ```bash
  make clean clean-certs
  ```

Example full flow:
```bash
make clean clean-certs
make generate-certs
make start-server
make start-client
```

### Step 5. Run the Server (Manual Option)
You can also run the server manually with the certificate paths:
```bash
go run cmd/candy-server-server/main.go --tls-port=3333 --tls-certificate=candy.tld/cert.pem --tls-key=candy.tld/key.pem --tls-ca=minica.pem
```

### Step 6. Run the Client (Manual Option)

Build and run the client manually:
```bash
go build candy-client.go
./candy-client -k <candy_type> -c <count> -m <money_amount>
```

Example:
```bash
./candy-client -k AA -c 2 -m 50
```

## API

Methods and request formats are described in the Swagger specification. Main endpoint:
- `POST /buy_candy` - handles candy purchases.
