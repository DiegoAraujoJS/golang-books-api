FROM golang:latest
WORKDIR /app
RUN apt-get update && apt-get -y install vim
COPY go.mod go.sum ./
RUN go mod download && go mod verify
