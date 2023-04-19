# familia gildo
familia gildo means "trade" in Esperanto.

## outline


## purpose
to improve the motivation of my child for help.

## how to develop
small step

## infra
GCP(Cloud Run)

## env

### OS

```
Mac Book Air M2 2022
Ventura 13.2.1
```

### go
```
❯ go version                     
go version go1.20.2 darwin/arm64
```

### gcloud
```
❯ gcloud version
Google Cloud SDK 422.0.0
```

### gqlgen
```
❯ gqlgen version        
v0.17.28
```

### nodejs
```
❯ node -v
v18.15.0
```

### npm
```
❯ npm -v 
9.5.0
```

### yarn
```
❯ yarn -v
1.22.19
```

### IDE(goland)
```
GoLand 2022.3.4
ビルド #GO-223.8836.56、ビルド日 2023年3月23日
```

### docker
```
❯ docker version
Client: Docker Engine - Community
 Version:           23.0.1
 
   ~~~

Server:
 Engine:
  Version:          23.0.3

```

### docker-compose
```
❯ docker compose version
Docker Compose version v2.17.2
```

## setup

### go module

```
cd src/backend
```

```
go mod init github.com/sky0621/kaubandus
```

### cobra

https://github.com/spf13/cobra

https://github.com/spf13/cobra-cli/blob/main/README.md

```
go install github.com/spf13/cobra-cli@latest
```

```
cd src/backend
```

```
cobra-cli init
```

#### add command

```
cd src/backend
```

```
cobra-cli add server
```

### sqlc

https://docs.sqlc.dev/en/latest/overview/install.html

### wire

```
go install github.com/google/wire/cmd/wire@latest
```

### goverter

https://goverter.jmattheis.de/#/

```
go install github.com/jmattheis/goverter/cmd/goverter@latest
```
