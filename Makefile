build:
	docker compose build

run:
	docker compose up --build

test:
	go test -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "dto.go\|service.go\|handler.go\|mocks\|.pb.go\|.pb\|middleware.go\|/cmd\|/model\|docs\|repository.go\|postgres.go\|redis.go\|minio.go\|db\|configs" > cover.out
	go tool cover -func=cover.out

swag:
	swag init -g cmd/app/main.go