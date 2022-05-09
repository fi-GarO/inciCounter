# syntax=docker/dockerfile:1
# https://www.callicoder.com/docker-golang-image-container-example/

FROM golang:1.18-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk update && apk add bash
RUN apk add build-base

COPY *.go ./
COPY *.json ./

RUN go mod tidy

# zbuildí v containeru linux binary s názvem "inci-counter" (název může být libovolný)
RUN go build -o /inci-counter

EXPOSE 8080

# spouští zbuilděnou binárku v containeru (pokud hlásí, že ji nevidí nebo, že neexistuje, pravděpodobně je problém v nekompatibilitě architektury počítače (32bit, 64bit, amd...), potřeba dohledat fix)
CMD [ "/inci-counter" ] 