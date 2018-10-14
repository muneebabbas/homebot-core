FROM alpine:3.8
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD main /
COPY config-prod.yaml config.yaml
CMD ["/main"]
