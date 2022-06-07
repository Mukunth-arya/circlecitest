FROM golang:1.17.9-alpine3.10
RUN mkdir /sampleapp
ADD . /sampleapp
WORKDIR /sampleapp

CMD ["go run" "main.go" ]
