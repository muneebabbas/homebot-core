FROM alpine:3.8
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD main /
CMD ["/main"]
