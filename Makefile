run-dev:
	go run cmd/service/main.go -env-mode=development


# Swag application
swag:
	@swag init --parseDependency --parseInternal --parseDepth 3 -g cmd/service/main.go
