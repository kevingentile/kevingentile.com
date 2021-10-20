
FROM node:lts-alpine as angular

COPY angular/ /angular
WORKDIR /angular/kevingentile-com
RUN npm ci && \ 
    ./node_modules/.bin/ng build --configuration=production --base-href /home/ --deploy-url /home/

FROM golang:alpine

WORKDIR /go/src/github.com/kevingentile/kevingentile.com
COPY . .
COPY --from=angular /angular/kevingentile-com/dist ./angular/kevingentile-com/dist

RUN go install

ENV GIN_MODE release

ENTRYPOINT [ "kevingentile.com" ]
