FROM golang:1.23.4

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o cmd/main.go

CMD [ "/app/cmd" ]