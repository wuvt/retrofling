FROM golang:alpine

RUN set -ex \
        && apk add --no-cache --virtual .build-deps git \
        && go get github.com/wuvt/retrofling \
        && apk del .build-deps

EXPOSE 8080
USER nobody
CMD ["/go/bin/retrofling"]
