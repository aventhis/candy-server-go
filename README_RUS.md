
# Candy API Server

Этот проект является результатом выполнения задания **"Day 04 - Go Boot Camp"**. Основная цель — разработка API сервера для работы с торговыми автоматами по продаже конфет. Задача включает:

- Реализацию API, отвечающего требованиям протокола, описанного в Swagger.
- Поддержку функционала проверки данных, выдачи сдачи и обработки ошибок.
- Внедрение двухсторонней TLS-аутентификации для обеспечения безопасности.
- Интеграцию внешней программы для генерации ASCII-искусства.

## Технологии

- Go
- Swagger (go-swagger)
- OpenSSL
- TLS/SSL

## Доступные конфеты и их стоимость

| Код | Название             | Стоимость (центов) |
|-----|----------------------|--------------------|
| CE  | Cool Eskimo          | 10                |
| AA  | Apricot Aardvark     | 15                |
| NT  | Natural Tiger        | 17                |
| DE  | Dazzling Elderberry  | 21                |
| YR  | Yellow Rambutan      | 23                |


## Функциональность

1. **Покупка конфет**:
    - Клиент отправляет запрос на сервер с данными о количестве денег, типе конфет и количестве.
    - Сервер рассчитывает сдачу или сообщает об ошибке:
        - Если денег достаточно, возвращается сообщение "Thank you!" и сдача.
        - Если денег недостаточно, возвращается ошибка с указанием недостающей суммы.
        - Если данные неверны (отрицательное количество или несуществующий тип конфет), возвращается ошибка.

2. **Безопасность**:
    - Реализована взаимная TLS-аутентификация с использованием самоподписанных сертификатов.
    - Сервер и клиент используют сертификаты, подписанные локальным Центром Сертификации (CA), созданным с помощью [Minica](https://github.com/jsha/minica).

3. **Генерация кода**:
    - Серверный код и модели генерируются автоматически из спецификации Swagger 2.0 с помощью [go-swagger](https://github.com/go-swagger/go-swagger).

## Установка и запуск

### Шаг 1. Загрузка проекта
Клонируйте репозиторий:
```bash
git clone git@github.com:aventhis/candy-server-go.git
cd candy-server-go/src
```

### Шаг 2. Установка зависимостей
Убедитесь, что у вас установлен Go, затем выполните:
```bash
go mod tidy
```

### Шаг 3. Сгенерируйте серверный код из спецификации Swagger (при необходимости):
   ```bash
   swagger generate server -f swagger.yaml -a candy-server
   ```

### Шаг 4. Генерация сертификатов
Если у вас ещё нет сертификатов, выполните:

# Генерация ключей и сертификатов для сервера
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

- `cert.pem` — сертификат сервера.
- `key.pem` — закрытый ключ сервера.
Центр сертификации в папке src
Файлы CA:
- `minica.pem` — сертификат центра сертификации.
- `minica-key.pem` — закрытый ключ CA.

# Генерация ключей и сертификатов для клиента
В папке src/client.tld cгенерируйте клиентские сертификаты:
   ```bash
openssl genrsa -out client-key.pem 2048
openssl req -new -key client-key.pem -out client.csr
openssl x509 -req -in client.csr -CA minica.pem -CAkey minica-key.pem -CAcreateserial -out client-cert.pem -days 365
```
- `client-cert.pem` — сертификат клиента.
- `client-key.pem` — закрытый ключ клиента.
- `client.csr` — запрос на сертификат.

### Шаг 4. Запуск сервера
Запустите сервер с указанием путей к сертификатам:
```bash
go run cmd/candy-server-server/main.go --tls-port=3333 --tls-certificate=candy.tld/cert.pem --tls-key=candy.tld/key.pem --tls-ca=minica.pem
```

## Клиентское приложение

Для тестирования используйте `candy-client`:
```bash
./candy-client -k <тип_конфеты> -c <количество> -m <сумма_денег>
```

Пример:
```bash
./candy-client -k AA -c 2 -m 50
```

## API

Методы и формат запросов описаны в Swagger спецификации. Основной эндпоинт:
- `POST /buy_candy` - обработка покупки конфет.
