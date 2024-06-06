coverage:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
run:
	@PORT=8000 DATABASE_URL="postgres://social:socialzaa007@127.0.0.1:5432/social?sslmode=disable" go run cmd/main.go
