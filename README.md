# TGSH

Telegram bot that provides access to the shell.

> [!WARNING]
> You must not use this for fraudulent purposes. All responsibility when using TGSH rests solely with you

## Prerequirements
- Go 1.24

## How to use
1. Clone repo and go to it
```shell
git clone https://github.com/misshanya/tgsh
cd tgsh
```
2. Create Telegram bot in BotFather
3. Fill .env file:
```
BOT_TOKEN=token-of-your-bot
ALLOWED_USER=id-of-your-account
```
ALLOWED_USER is id of your Telegram account which can be obtained via other bots or Telegram clients
4. Build the bot and run it
```shell
go build -o bot .
./bot
```
