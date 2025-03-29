#build binary file
FROM golang:1.24.0-alpine as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN cp -n .env.example .env
RUN go build -o /bin/server ./cmd/server/main.go
#run server in container
FROM alpine:latest
WORKDIR /bin
COPY --from=builder /bin/server /bin/server
COPY --from=builder /app/.env /bin/.env
EXPOSE 8080
CMD ["/bin/server"]