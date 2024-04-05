[![](https://img.shields.io/discord/828676951023550495?color=5865F2&logo=discord&logoColor=white)](https://lunish.nl/support)
![](https://img.shields.io/github/repo-size/Luna-devv/githook?maxAge=3600)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/I3I6AFVAP)

**⚠️ In development, breaking changes ⚠️**

## About
GitHook is a discord bot designed to improve GitHub webhooks for discord. Basically just a modern and nicer version of the discord built-in GitHub webhook support.

If you need help developing with this, join **[our Discord Server](https://discord.com/invite/yYd6YKHQZH)**.

![image](https://github.com/Luna-devv/githook/assets/71079641/00fbddee-adc9-4156-9e5b-cfd3845eb6f0)
![image](https://github.com/Luna-devv/githook/assets/71079641/23aed8ca-4577-4c6b-888c-ec7515d7e2cd)
![image](https://github.com/Luna-devv/githook/assets/71079641/0a7962cd-0cde-44b6-b4cc-fb51b6de7a41)

## Setup
Clone this repo with the following commands:

```bash
git clone https://github.com/Luna-devv/githook
```

## Develope
Go goes brrrrrrrr

To run the server run
```bash
go run .
```

## Deploy

Since docker is the best thing that exists for managing deployments, it's the thing we use. If you don't want to use docker, install the go programing language from [go.dev](https://go.dev) and run `go run .` or however you want to run it.

To build the docker container run
```bash
docker build -t githook .
```

To start the docker container (detached) run 
```bash
docker compose up -d
```
