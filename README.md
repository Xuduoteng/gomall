# Go Gin Boilerplate

<div>

<img src="./assets/golang.png" align="right"/>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Xuduoteng/gomall)
![CI](https://github.com/Xuduoteng/gomall/workflows/Go/badge.svg)
![License](https://img.shields.io/github/license/Xuduoteng/gomall)

## 📘 Documents

- [Project Document](https://Xuduoteng.github.io/gomall/)
- [Swagger Api Document](https://gomall.gganbu.services/swagger/index.html)

</div>

## 🎬 Introduction

A development boilerplate based on the Gin framework, aimed at helping developers quickly build and develop web applications.

- [English](./README.md)
- [中文介绍](./README_zh.md)

## 👀 Keywords

[Go](https://github.com/golang/go)
[Gin](https://github.com/gin-gonic/gin)
[Cobra](https://github.com/spf13/cobra)
[Viper](https://github.com/spf13/viper)
[Gorm](https://github.com/go-gorm/gorm)
[Gin-Swagger](https://github.com/swaggo/gin-swagger)
[Air](https://github.com/cosmtrek/air)
[Logrus](https://github.com/sirupsen/logrus)
[Lumberjack](https://github.com/natefinch/lumberjack)
[Zap](https://github.com/uber-go/zap)
[Jwt](https://github.com/golang-jwt/jwt)

## ✨ Features

- **Fast Development**: Using the **Gin** framework and related tools to speed up the development and iteration process of the project..
- **Easy to use**: Follow the [project-layout](https://github.com/golang-standards/project-layout/tree/master) specification and provide a clear and simple code structure so that even beginners can easily get started.
- **Advanced CLI**: Using **Cobra** to build modern command line tools to simplify project management and operations.
- **Hot Reload**: Using **Air** tool, support hot reload, improve development efficiency.
- **Logging system**: Integrated **Logrus**, **Zap** and **Lumberjack** to achieve all-round log recording and management.
- **Database Support**: Integrated **Gorm**, support mainstream databases such as MySQL, PostgreSQL, etc.
- **Flexible Middleware**: Integrate common middleware to easily implement functions such as logging, authentication, cross-domain, and flow control.
- **API Document**: Use **Gin-Swagger** to generate API documents for easy viewing and debugging of interfaces.

## 🚀 Quick Start

```sh
git clone https://github.com/Xuduoteng/gomall.git
cd gomall
go mod download
```

<!-- ### Create a githooks soft link 【Recommended】

```sh
cd .git/hooks
ln -s ../../githooks/* .
```

> Tips:
>
> - How to confirm that the creation is successful?
>
> Run the following command:
>
> ```sh
> ls -l . # Now your current directory location should be in the .git/hooks directory
> ```
>
> If successful, you will see the output containing the following content:
>
> ```sh
> commit-msg -> ../../githooks/commit-msg
> pre-commit -> ../../githooks/pre-commit
> ``` -->

### Configuration Setup

- Navigate to the `configs` directory, copy the `config.example.yaml` file, and rename it to `config.yaml`.

```sh
cp configs/config.example.yaml configs/config.yaml
```

- Modify the configuration items in the `config.yaml` file.

```sh
vi configs/config.yaml
```

### Run the Application

- Run the project using [air](https://github.com/cosmtrek/air) **【Recommended】**

```sh
air
```

- Run the project using [go run](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program)

```sh
go run main.go server
```

<!--
## 🔨 打包

```sh
make build
``` -->

## 🔨 Build

```sh
make build
```

## 🪤 Deployment

### docker-compose

Deploy the application using **_docker-compose_**.
Ensure that Docker is installed on the server, and you are familiar with the usage of docker compose.

- Copy the `deployments/docker-compose.yml` configuration file to the project root directory.
- Execute the following command to start the application

```sh
docker compose up -d
```

## 🤝 Support

- Star 🌟 the project
- Welcome to submit [issue](https://github.com/Xuduoteng/gomall/issues). Thank you for your support
- Help promote it on social media and recommend it to friends

  [![Twitter](https://img.shields.io/twitter/url?label=Twitter&logo=twitter&style=flat&url=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)
  [![Facebook](https://img.shields.io/twitter/url?label=Facebook&logo=facebook&style=flat&url=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)
  [![WhatsApp](https://img.shields.io/twitter/url?label=WhatsApp&logo=whatsapp&style=flat&url=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)](https://api.whatsapp.com/send?text=Wow:%20https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)
  [![Telegram](https://img.shields.io/twitter/url?label=Telegram&logo=telegram&style=flat&url=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)](https://t.me/share/url?url=https%3A%2F%2Fgithub.com%2FXuduoteng%2Fgomall)

- You can also sponsor a cup of coffee on [Ko-Fi](https://ko-fi.com/dorvan) or [Buy Me A Coffee](https://www.buymeacoffee.com/dorvan)

  <a href='https://ko-fi.com/J3J1T95FG' target='_blank'>
  <img width="145" height="40" src='https://storage.ko-fi.com/cdn/kofi2.png?v=3' border='0' alt='Buy Me a Coffee at ko-fi.com' />
  </a>

  <a href="https://www.buymeacoffee.com/dorvan" target="_blank">
  <img width="145" height="40" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" />
  </a>
