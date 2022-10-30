FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
EXPOSE 9010
CMD ["go", "run", "cmd/main/main.go"]
