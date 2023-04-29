FROM alpine:latest


RUN apk add --no-cache \
    unzip \
    ca-certificates

ADD ./main /pb/pocketbase

EXPOSE 8080

# start PocketBase
CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]