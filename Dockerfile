FROM golang:latest
WORKDIR /app
RUN apt-get update && apt-get -y install vim
