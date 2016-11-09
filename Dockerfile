FROM golang:onbuild
RUN go get github.com/mattes/migrate
EXPOSE 8080