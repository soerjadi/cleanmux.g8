###
# Stage 1: Builder
###
FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

COPY . /app

RUN apk update \
  && apk upgrade \
  && apk add \
  git build-base ca-certificates && \
  update-ca-certificates 2>/dev/null || true

RUN go mod download

ARG CACHEBUILD

# to prevent cache on building our app
RUN cd cmd/rest && \
  GOOS=linux go build -tags release -a -ldflags "-linkmode external -extldflags -static" -o /app/bin/$name$ && \
  echo $CACHEBUILD

###
# Stage 2: Serving
###
FROM busybox

RUN mkdir -p /etc/$name$/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/$name$ /usr/app/$name$
COPY --from=builder /app/files/config-docker.ini /etc/$name$/config.ini

CMD [ "/usr/app/$name$" ]