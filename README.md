# goffredo-go

The Discord bot that introduces you in the vocal channels with a musical theme.

No further updates are planned in this repository but a totally new version is coming.

I was just an experiment to try GoLang.

## How to use

1. Create a Discord bot
2. Create a file named `secrets.conf` in the `conf` folder using this template:
```
[discord]
clientid = 0000000000000
clientsecret = AAAAAAAAAAAAAAAAAA
bottoken = ODU0NDI4OTE3NTc3MTU0NjAw.YMjy8A.uaCGQr8HL_oYfqvbHf7gIJQGaGs
redirecturl = http://yourhost:8000/callback
```
where client id, client secret and bot token are given from Discord
3. Execute `docker-compose up -d --build`
