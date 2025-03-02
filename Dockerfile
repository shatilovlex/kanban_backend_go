#build binary file
FROM golang:1.24.0-alpine as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /bin/server ./cmd/server/main.go
#run shop server in container
FROM alpine:latest
COPY --from=builder /bin/server /bin/server
EXPOSE 8080
CMD ["/bin/server"]