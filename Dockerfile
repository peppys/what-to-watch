FROM golang:1.8 as builder
RUN mkdir /go/src/app
# TODO improve this to explictly copy needed files
COPY . /go/src/app
RUN go install app

FROM debian:jessie-slim
COPY --from=builder /go/bin/app /bin/app
COPY --from=builder /go/src/app/*.pem /
ENTRYPOINT ["/bin/app"]