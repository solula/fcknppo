FROM golang:1.21.3-alpine3.18
WORKDIR /build

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

EXPOSE 8080

CMD ["./app"]

