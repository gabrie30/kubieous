FROM golang:1.9.2-alpine


WORKDIR /go/src/github.com/gabrie30/kubieous
COPY . .

RUN apk update
RUN apk add git libc6-compat
RUN go get
RUN go install

ENTRYPOINT /go/bin/kubieous

EXPOSE 8080
