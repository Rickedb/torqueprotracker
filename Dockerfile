
FROM golang:1.22.4-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o torqueprotracker
ENV PORT=8080
EXPOSE 8080

CMD ["/app/torqueprotracker"]