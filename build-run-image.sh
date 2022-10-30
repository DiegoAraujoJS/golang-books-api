docker rm -f books-golang-server
docker rmi books-golang-server
docker build -t books-golang-server .
docker run -d --network books-golang --name books-golang-server -p 9010:9010 books-golang-server
