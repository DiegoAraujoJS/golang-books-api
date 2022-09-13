FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
EXPOSE 9010
CMD ["go", "run", "cmd/main/main.go"]
