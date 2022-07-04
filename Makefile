

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/fsd-csi-driver ./cmd


docker:
	docker build -t realjf/fsd-csi-driver:v0.1 .
	docker push realjf/fsd-csi-driver:v0.1


	

