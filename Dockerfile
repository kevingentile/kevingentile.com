
#Get golang
FROM golang:latest

WORKDIR /go/src/github.com/kevingentile/kevingentile.com
COPY . .

RUN go install

CMD ["kevingentile.com"]
