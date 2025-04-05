FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o go-api-test-app cmd/go-api-test-app/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go-api-test-app .
CMD [ "./go-api-test-app", "--debug" ]
EXPOSE 8080