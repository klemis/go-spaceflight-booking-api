FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the API binary
RUN go build -o api ./cmd/api

# Build the migrate binary
RUN go build -o migrate ./cmd/migrate

# Build the schedule binary
RUN go build -o schedule ./cmd/schedule

# Set the default command to run the API binary
CMD ["./api"]