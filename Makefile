compile:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main .

runserver:
	gin -i run main.go

dockerpush:
	make compile && \
	docker image build -t muneebabbas/homebot-core . && \
	docker push muneebabbas/homebot-core

removedangling:
	docker image prune
