
FROM node:lts-alpine as angular

COPY angular/ /angular
COPY angular/kevingentile-com/karma.conf.ci.js  angular/kevingentile-com/karma.conf.js
WORKDIR /angular/kevingentile-com

RUN apk --no-cache upgrade && apk add --no-cache chromium
ENV CHROME_BIN=/usr/bin/chromium-browser
RUN npm ci && \ 
    ./node_modules/.bin/ng test --browsers=ChromeHeadless && \
    ./node_modules/.bin/ng build --configuration=production --base-href /home/ --deploy-url /home/

FROM golang:alpine

WORKDIR /go/src/github.com/kevingentile/kevingentile.com
COPY . .
COPY --from=angular /angular/kevingentile-com/dist ./angular/kevingentile-com/dist

RUN go install

ENV GIN_MODE release

ENTRYPOINT [ "kevingentile.com" ]
