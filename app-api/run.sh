
go build -o ./wait connect/main.go
go build -o ./db_create create/main.go
go build -o ./api cmd/api/main.go
./wait && ./db_create && ./api -address 0.0.0.0:8082 -database-url db
