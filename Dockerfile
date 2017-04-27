FROM golang:1.8.1

MAINTAINER xiaoyusilen <debug@xiaoyu.fail>

# Create directory
RUN mkdir -p /usr/go/src/realtime-fetch
RUN mkdir -p /usr/go/src/realtime-fetch/log
WORKDIR /usr/go/src/realtime-fetch

# Copy to directory
COPY main /usr/go/src/realtime-fetch

CMD ["./main"]