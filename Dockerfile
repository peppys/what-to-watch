FROM golang:1.8 as builder
RUN mkdir -p /go/src/github.com/PeppyS/personal-site-api
# TODO improve this to explictly copy needed files
COPY . /go/src/github.com/PeppyS/personal-site-api/
RUN go install github.com/PeppyS/personal-site-api/cmd/api

FROM debian:jessie-slim
COPY --from=builder /go/bin/api /bin/api
ENTRYPOINT ["/bin/api"]