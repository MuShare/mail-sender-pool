docker-build:
	docker build --build-arg VERSION=$(VERSION) -t mushare/mail-sender-pool:latest .
	docker tag mushare/mail-sender-pool:latest mushare/mail-sender-pool:$(VERSION)

docker-build-staging:
	docker build --build-arg VERSION=staging -t mushare/mail-sender-pool:master .

docker-push:
	docker push mushare/mail-sender-pool:latest
	docker push mushare/mail-sender-pool:$(VERSION)

docker-push-staging:
	docker push mushare/mail-sender-pool:master

docker-clean:
	docker rmi mushare/mail-sender-pool:latest || true
	docker rmi mushare/mail-sender-pool:$(VERSION) || true
	docker rm -v $(shell docker ps --filter status=exited -q 2>/dev/null) 2>/dev/null || true
	docker rmi $(shell docker images --filter dangling=true -q 2>/dev/null) 2>/dev/null || true

docker-clean-staging:
	docker rmi mushare/mail-sender-pool:master || true
	docker rm -v $(shell docker ps --filter status=exited -q 2>/dev/null) 2>/dev/null || true
	docker rmi $(shell docker images --filter dangling=true -q 2>/dev/null) 2>/dev/null || true

binary-build:
	mkdir -p bin
	GO111MODULE=on go build -ldflags="-X 'main.VERSION=$(VERSION)'" -o bin/mail-sender-pool main.go

ci-build-production: docker-build docker-push docker-clean

ci-build-staging: docker-build-staging docker-push-staging docker-clean-staging
