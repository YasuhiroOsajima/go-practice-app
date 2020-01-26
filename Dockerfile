FROM golang:1.13.6-alpine3.10 AS build

WORKDIR /
COPY . /go/src/github.com/YasuhiroOsajima/go-practice-app
RUN apk update \
  && apk add --no-cache git
RUN cd /go/src/github.com/YasuhiroOsajima/go-practice-app && go build -o bin/go-practice-app cmd/server.go

FROM alpine:3.10
COPY --from=build /go/src/github.com/YasuhiroOsajima/go-practice-app/bin/go-practice-app /usr/local/bin/
CMD ["go-practice-app"]
