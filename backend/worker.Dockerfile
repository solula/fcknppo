FROM golang:1.21.3-alpine3.18
WORKDIR /build

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/worker/main.go

CMD ["./app"]

