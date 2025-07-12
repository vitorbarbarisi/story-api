FROM golang:latest

ENV GOPROXY https://proxy.golang.org,direct
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o story-api .

EXPOSE 8000
ENTRYPOINT ["./story-api"]