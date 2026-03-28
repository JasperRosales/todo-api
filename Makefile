

run:
	@go run cmd/api/main.go

test:
	@go test ./tests/.

push:
	@git add .
	@git commit -m "$(m)"
	@git push