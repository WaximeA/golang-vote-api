FROM golang:1.13

WORKDIR /go/src/github.com/WaximeA/golang-vote-api
COPY . /go/src/github.com/WaximeA/golang-vote-api

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["golang-vote-api"]