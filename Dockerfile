# Stage 1: Build the Go application
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o server .


# Stage 2: Create a minimal image to run the application - minimal
FROM alpine:latest
WORKDIR /root/

#copy files
COPY --from=build /app/server .
COPY --from=build /app/.env .

USER 65534

# set executable
CMD ["./server"]