FROM golang:1.10 as builder
WORKDIR $GOPATH/src/github.com/PeppyS/what-to-watch/api/

# Install dep
RUN go get -u github.com/golang/dep/...

# Copy code from host and compile
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN go build -o /bin/api cmd/api/main.go

# Copy binary to debian and run
FROM debian:jessie-slim
COPY --from=builder /bin/api /bin/api
ENTRYPOINT ["/bin/api"]
