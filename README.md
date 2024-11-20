## Генерация сертификатов

### Сертификаты для сервера через Minica:

Убедитесь, что у вас установлен Minica. Скачать Minica можно с официального репозитория: [Minica на GitHub](https://github.com/jsha/minica).

### Шаги для генерации сертификатов:
1. Запустите Minica с вашим доменом:
   ```bash
   ./minica --domains candy.tld
   ```
2. После выполнения команды, Minica создаст следующие файлы:
   - minica.pem — корневой сертификат.
   - minica-key.pem — приватный ключ CA.
   - В папке candy.tld:
     - cert.pem — сертификат сервера.
     -   key.pem — приватный ключ сервера.

3. Переместите файлы cert.pem и key.pem в папку проекта 
(например, src/candy.tld).

### Сертификаты для клиента:
В папке src/client.tld cгенерируйте клиентские сертификаты: 
```bash
    openssl genrsa -out client-key.pem 2048
    openssl req -new -key client-key.pem -out client-req.pem
    openssl x509 -req -in client-req.pem -CA minica.pem -CAkey minica-key.pem -CAcreateserial -out client-cert.pem -days 365
  ```

Запустите сервер:
```bash
go run cmd/candy-server-server/main.go --tls-certificate=server-cert.pem --tls-key=server-key.pem
  ```
Запустите клиент:
```bash
./candy-client -k AA -c 2 -m 50
  ```