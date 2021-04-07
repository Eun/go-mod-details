FROM golang:1.16-alpine as builder
ENV GOOS=linux \
    GOARCH=386 \
    CGO_ENABLED=0

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]