# Café Plätschwiesle

![](https://img.shields.io/badge/Framework-Vue3-informational?style=for-the-badge&logo=vuedotjs&color=4FC08D)
![](https://img.shields.io/badge/Language-Typescript-informational?style=for-the-badge&logo=typescript&color=3178C6)
![](https://img.shields.io/badge/Language-Go-informational?style=for-the-badge&logo=go&color=00ADD8)

[![Build Status](https://build.unjx.de/buildStatus/icon?style=flat-square&job=cafe%2Fmain)](https://build.unjx.de/job/cafe/job/main/)

[https://hub.docker.com/r/florianhoss/cafe](https://hub.docker.com/r/florianhoss/cafe)

## docker-compose example:

```
version: '3.9'

services:
  cafe:
    image: florianhoss/cafe:latest
    container_name: cafe
    restart: unless-stopped
    environment:
      - PUID=1000
      - PGID=1000
      - ALLOWED_HOSTS=http://localhost:5000,https://home.example.com
      - SWAGGER=true
      - LOG_LEVEL=info # trace,debug,info,warn,error,fatal,panic
      - AUTH_PASSWORD=SuperSafe
      - AUTH_SECRET=example # create with: openssl rand -base64 20
    volumes:
      - ./storage:/app/storage
    ports:
      - "127.0.0.1:5000:5000"
```

## docker-compose example with MariaDB as database:

```
version: '3.9'

networks:
  net:
    external: false

services:

  cafe-db:
    image: lscr.io/linuxserver/mariadb:latest
    container_name: cafe-db
    restart: unless-stopped
    environment:
      - PUID=1000
      - PGID=1000
      - MYSQL_ROOT_PASSWORD=root
      - TZ=Europe/Berlin
      - MYSQL_DATABASE=db
      - MYSQL_USER=user
      - MYSQL_PASSWORD=password
    volumes:
      - ./db:/config
    expose:
      - 3306
    networks:
      - net

  cafe:
    image: florianhoss/cafe:latest
    container_name: cafe
    restart: unless-stopped
    depends_on:
      - cafe-db
    environment:
      - PUID=1000
      - PGID=1000
      - ALLOWED_HOSTS=http://localhost:5000,https://home.example.com
      - SWAGGER=true
      - LOG_LEVEL=info # trace,debug,info,warn,error,fatal,panic
      - MYSQL_URL=dashboard-db:3306
      - MYSQL_USER=user
      - MYSQL_PASSWORD=password
      - MYSQL_DATABASE=db
      - AUTH_PASSWORD=SuperSafe
      - AUTH_SECRET=example # create with: openssl rand -base64 20
    volumes:
      - ./storage:/app/storage
    ports:
      - "127.0.0.1:5000:5000"
    networks:
      - net
```
