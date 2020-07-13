FROM golang:1.13-alpine as builder

WORKDIR /go/src/github.com/nchern/btm
COPY . .

RUN apk update && apk add build-base

# Build the project inside an intermediate container
RUN make install-deps install

FROM golang:1.13-alpine

WORKDIR /

COPY --from=builder /go/bin/* /bin/

ENTRYPOINT btm
