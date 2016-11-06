FROM golang
ADD . /go/src/github.com/gwhn/pgdock
RUN go install github.com/gwhn/pgdock
ENTRYPOINT /go/bin/pgdock
EXPOSE 8080