web-test
========

This image runs a web service in 8080 port, used for testing. It comes from rawmind/alpine-base.

## Build

```
docker build -t rawmind/web-test:<version> .
```

## Versions

- `0.0.1` [(Dockerfile)](https://github.com/rawmind0/web-test/blob/master/Dockerfile)


## Usage

```
docker run rawmind/web-test:<version> 
```

## Output example

```
Real-Server: proxy-test_web-test_2
GET / HTTP/1.1
Host: web-test.local:9080
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Encoding: gzip, deflate, sdch
Accept-Language: en-US,en;q=0.8,es;q=0.6,pt;q=0.4
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
X-Forwarded-For: 192.168.33.1
X-Forwarded-Host: web-test.local:9080
X-Forwarded-Proto: http
X-Forwarded-Server: traefik_traefik_1
```
