FROM alpine:3.6
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

ADD . /tmp/ninefs

RUN set -eux && \
	apk add --no-cache --virtual .build-deps \
		gcc \
		musl-dev \
		make && \
	cd /tmp/ninefs && \
	make clean && make PREFIX=/usr install && \
	cd /tmp && \
	rm -rf /tmp/ninefs && \
	apk del .build-deps

ENTRYPOINT ["ninefs"]
