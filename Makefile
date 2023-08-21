.PHONY:docs
docs:
	@rm -rf ./docs/*
	@swag init -g ./cmd/server/main.go -o ./docs > /dev/null 2>&1
	@echo "docs generated successfully"

.PHONY:build
build:
	@go build -o ./tmp/server -tags=jsoniter cmd/server/main.go
	@upx ./tmp/server

.PHONY:model
model:
	sqlc generate