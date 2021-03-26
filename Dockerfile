# Building
FROM golang:1.15.0-alpine3.12 as builder
# create a working directory
WORKDIR /go/src/app
# add source code
COPY . /go/src/app
# copy docker db configs
# RUN cp config/sample.env .env
# create directory for executable 
RUN apk add git
RUN go build -i main.go


FROM  alpine:3.12.0
# create a working directory
WORKDIR /root/
COPY --from=builder /go/src/app/main .
# run server
CMD ["./main"]  