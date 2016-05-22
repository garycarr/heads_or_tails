FROM golang:1.6

COPY . /go/src/github.com/garycarr/heads_or_tails

RUN cd /go/src/github.com/garycarr/heads_or_tails &&\
    go install

ENTRYPOINT /go/bin/heads_or_tails
