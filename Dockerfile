FROM golang:latest
WORKDIR /app
RUN apt-get update && apt-get -y install vim
COPY ./.vimrc /root
RUN vim +'PlugInstall --sync' +qa
COPY go.mod go.sum ./
RUN go mod download && go mod verify
