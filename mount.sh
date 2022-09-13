docker rm -f books-api-golang
docker rmi books-api-golang
docker build -t books-api-golang .
docker run --name books-api-golang --network books-app -p 9010:9010 -d books-api-golang
