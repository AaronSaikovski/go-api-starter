#FROM golang:1.21-alpine
FROM golang:1.21
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server .

EXPOSE 8080

CMD ["./server"]
