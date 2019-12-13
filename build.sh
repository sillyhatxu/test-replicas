CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
docker build -t xushikuan/test-replicas .
docker stack deploy -c docker-compose.yml tr
curl http://localhost:8888/test/address
