# read-later-bot
The Telegram bot stores and retrieves random articles or blogs ”read later” from the repository.

The bot is purely written in golang without any framework.

Technologies:
- Short polling and webhook fetching
- Docker


Self-signed certificates generation.
```
openssl req -newkey rsa:2048 -sha256 -nodes -keyout PRIVATE.key -x509 -days 365 -out PUBLIC.pem -subj "/C=US/ST=State/L=City/O=pinkyhi/CN={HERE_IP}"
```
webhok registration
```
curl -F "url=https://{HERE_IP}:8443/telegram_endpoint" -F "certificate=@PUBLIC.pem" https://api.telegram.org/bot{HERE_BOT_TOKEN}/setWebhook?secret_token={HERE_TG_SECRET_TOKEN}
```

To run:
```
go build -o ./app main.go
./app -tg-bot-token {HERE_BOT_TOKEN}
```
