FROM golang:latest

WORKDIR /app

COPY ./utils ./utils
COPY ./main.go ./main.go