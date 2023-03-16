## Sample Dockerfile to package Backend-API as a docker image.
# Sample build command:
# docker build --force-rm --squash -t docms_be:0.1.0 .

FROM golang:1.17-alpine AS builder
LABEL maintainer="Thanh Nguyen <btnguyen2k@gmail.com>"
RUN apk add git build-base \
    && mkdir /build
COPY . /build
RUN cd /build && go build -o main

FROM alpine:3.12
RUN mkdir /app
COPY --from=builder /build/main /app/main
COPY --from=builder /build/config /app/config
COPY --from=builder /build/frontend /app/frontend
RUN apk add --no-cache -U tzdata bash ca-certificates \
    && update-ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && chmod 711 /app/main \
    && rm -rf /var/cache/apk/*
WORKDIR /app
EXPOSE 8000
CMD ["/app/main"]
#ENTRYPOINT /app/main
