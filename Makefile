compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o main .

runserver:
	gin -i run main.go

dockerpush:
	docker image build -t muneebabbas/homebot-core . && \
	docker push muneebabbas/homebot-core
