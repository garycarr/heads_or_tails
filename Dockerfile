FROM golang:1.8

WORKDIR /go/src/github/gcarr/heads_or_tails
COPY . /go/src/github/gcarr/heads_or_tails

RUN go-wrapper install

CMD ["go-wrapper", "run"]
