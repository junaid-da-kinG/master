FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

EXPOSE 8080

# Build the Go app
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]

