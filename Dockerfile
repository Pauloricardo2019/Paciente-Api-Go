FROM golang:1.17

WORKDIR /app

COPY main.go ./
COPY go.mod ./
COPY go.sum ./

RUN go build -o /server

EXPOSE 8081 

CMD ["/server"]