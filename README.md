```
go run ./cmd/api
```

```
migrate -path=./migrations -database="postgres://postgres:12345@localhost:5432/godb?sslmode=disable" up
```
