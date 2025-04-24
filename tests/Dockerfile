FROM golang:1.23.5-alpine

RUN apk update && apk upgrade && \
    apk add build-base && apk add --no-cache inotify-tools && \
    apk add --no-cache bash

EXPOSE 3456

COPY tests/entrypoint.sh /usr/local/bin

RUN chmod +x /usr/local/bin/entrypoint.sh

COPY . /ki/

WORKDIR /ki

ENTRYPOINT ["entrypoint.sh"]