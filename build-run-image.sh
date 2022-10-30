docker rm -f books-golang
docker rmi books-golang
docker build -t books-golang .
docker run -d --network books-golang --name books-golang -p 9010:9010 books-golang
