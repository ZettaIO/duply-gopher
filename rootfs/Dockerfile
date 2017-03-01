FROM golang:1.7-alpine

RUN apk add --update duplicity duply libffi libffi-dev ca-certificates build-base python-dev py-pip linux-headers haveged && \
    apk upgrade && \
    pip install python-swiftclient python-keystoneclient requests && \
    rm -rf /var/cache/apk/* && \
    apk del build-base libffi-dev python-dev linux-headers

WORKDIR /srv/duply-gopher

COPY duply-gopher /srv/duply-gopher/
COPY scripts /srv/duply-gopher/scripts
COPY templates /srv/duply-gopher/templates

CMD ["./duply-gopher", "config.yaml"]
