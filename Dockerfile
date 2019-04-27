FROM node:12-alpine as webbuilder

RUN apk update \
  && apk add git \
  && git clone --depth=1 https://github.com/vicanso/forest.git /forest \
  && cd /forest/web \
  && npm i \
  && npm run build \
  && rm -rf node_module

FROM golang:1.12-alpine as builder

ENV GO111MODULE on

COPY --from=webbuilder /forest /forest

RUN apk update \
  && apk add docker git gcc make musl-dev \
  && go get -u github.com/gobuffalo/packr/v2/packr2 \
  && cd /forest \
  && make test \
  && make build

FROM alpine 

EXPOSE 7001

COPY --from=builder /forest/forest /usr/local/bin/forest


CMD ["forest"]

HEALTHCHECK --interval=10s --timeout=3s \
  CMD forest --mode=check || exit 1