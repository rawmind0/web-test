FROM rawmind/alpine-base:3.7-0
MAINTAINER Raul Sanchez <rawmind@gmail.com>


ENV WEB_HOME=/opt/web-test \
    WEB_PORT=8080 \
    GOMAXPROCS=2 \
    GOROOT=/usr/lib/go \
    GOPATH=/opt/src \
    GOBIN=/gopath/bin \
    PATH=$PATH:/opt/web-test

# Add service files
ADD root /

RUN apk add --update go git musl-dev \ 
  && mkdir -p /opt/src $WEB_HOME; cd /opt/src \
  && cd /opt/src \
  && go build -o web-test web-test.go \
  && mv ./web-test ${WEB_HOME}; cd ${WEB_HOME} \
  && chmod +x ${WEB_HOME}/web-test \
  && apk del go git musl-dev \
  && rm -rf /var/cache/apk/* /opt/src 

EXPOSE 8080

WORKDIR $WEB_HOME

ENTRYPOINT ["/opt/web-test/web-test"]
