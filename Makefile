run:
	go run cmd/main.go

docker-build-image:
	docker build -t ricanalista/go-expert-stress-test:latest -f Dockerfile .