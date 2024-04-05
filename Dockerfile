FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

# Build the Go application binary
RUN go build -o main

CMD ["/app/main"]
