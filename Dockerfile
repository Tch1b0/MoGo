FROM golang:latest

WORKDIR /go/src/app

COPY ./linker ./linker
COPY ./main.go ./main.go
COPY ./commands ./commands
COPY ./main.go ./main.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./token.txt ./token.txt

RUN go mod download
RUN go build github.com/Tch1b0/MoGo

CMD ["./MoGo"]