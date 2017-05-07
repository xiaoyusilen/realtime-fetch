FROM alpine

MAINTAINER xiaoyusilen <debug@xiaoyu.fail>

RUN mkdir -p /usr/go/src/realtime-fetch/log

COPY main /usr/go/src/realtime-fetch
WORKDIR /usr/go/src/realtime-fetch

CMD ["/bin/sh"]
