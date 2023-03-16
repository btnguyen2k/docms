## Sample Dockerfile to package the whole application (backend + frontend) as a docker image.
# Sample build command:
# docker build --rm -t docms:0.1.0 .

FROM node:12-alpine AS builder_fe
RUN mkdir /build
RUN apk add jq sed
COPY . /build
RUN cd /build \
    && export APP_NAME=`jq -r '.name' appinfo.json` \
    && export APP_SHORTNAME=`jq -r '.shortname' appinfo.json` \
    && export APP_INITIAL=`jq -r '.initial' appinfo.json` \
    && export APP_VERSION=`jq -r '.version' appinfo.json` \
    && export APP_DESC=`jq -r '.description' appinfo.json` \
    && cd /build/fe-gui \
    && rm -rf .env \
    && rm -rf dist node_modules \
    && sed -i s/#{pageTitle}/"$APP_NAME v$APP_VERSION"/g public/index.html \
    && sed -i s/#{appName}/"$APP_NAME"/g public/index.html \
    && sed -i s/#{appInitial}/"$APP_INITIAL"/g public/index.html \
    && sed -i s/#{appShortname}/"$APP_SHORTNAME"/g public/index.html \
    && sed -i s/#{appDescription}/"$APP_DESC"/g public/index.html \
    && sed -i s/#{appVersion}/"$APP_VERSION"/g public/index.html \
    && sed -i s/#{appName}/"$APP_NAME"/g src/config.json \
    && sed -i s/#{appInitial}/"$APP_INITIAL"/g src/config.json \
    && sed -i s/#{appShortname}/"$APP_SHORTNAME"/g src/config.json \
    && sed -i s/#{appDescription}/"$APP_DESC"/g src/config.json \
    && sed -i s/#{appVersion}/"$APP_VERSION"/g src/config.json \
    && npm install && npm run build

FROM golang:1.17-alpine AS builder_be
RUN apk add git build-base jq sed
RUN mkdir /build
COPY . /build
RUN cd /build \
    && export APP_NAME=`jq -r '.name' appinfo.json` \
    && export APP_SHORTNAME=`jq -r '.shortname' appinfo.json` \
    && export APP_INITIAL=`jq -r '.initial' appinfo.json` \
    && export APP_VERSION=`jq -r '.version' appinfo.json` \
    && export APP_DESC=`jq -r '.description' appinfo.json` \
    && cd /build/be-api \
    && sed -i s/#{appName}/"$APP_NAME"/g config/application.conf \
    && sed -i s/#{appInitial}/"$APP_INITIAL"/g config/application.conf \
    && sed -i s/#{appShortname}/"$APP_SHORTNAME"/g config/application.conf \
    && sed -i s/#{appDescription}/"$APP_DESC"/g config/application.conf \
    && sed -i s/#{appVersion}/"$APP_VERSION"/g config/application.conf \
    && go build -o main

FROM alpine:3.12
LABEL maintainer="Thanh Nguyen <btnguyen2k@gmail.com>"
RUN mkdir -p /app/frontend
COPY --from=builder_be /build/be-api/main /app/main
COPY --from=builder_be /build/be-api/config /app/config
COPY --from=builder_fe /build/fe-gui/dist /app/frontend
RUN apk add --no-cache -U tzdata bash ca-certificates \
    && update-ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && chmod 711 /app/main \
    && rm -rf /var/cache/apk/*
WORKDIR /app
EXPOSE 8000
CMD ["/app/main"]
#ENTRYPOINT /app/main
