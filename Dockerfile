
#Get golang
FROM golang:latest

WORKDIR /go/src/github.com/laughingcabbage/kevingentile.com
COPY . .

RUN go-wrapper download

RUN go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE $PORT
