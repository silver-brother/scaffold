.PHONY:docs
docs:
	@swag init -dir cmd/server

.PHONY:build
build:
	@go build -o ./tmp/server -tags=jsoniter cmd/server/main.go
	@upx ./tmp/server

.PHONY:model
model:
	sqlc generate