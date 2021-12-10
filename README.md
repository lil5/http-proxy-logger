# HTTP proxy logger

This is HTTP proxy which prints http requests and responses to console including http body.
For example:

```log
2021/05/05 03:50:44 ---REQUEST 3---

POST /mocking/contacts HTTP/1.1
Host: demo7704619.mockable.io
User-Agent: PostmanRuntime/7.28.0
Content-Length: 63
Accept: */*
Accept-Encoding: gzip, deflate, br
Cache-Control: no-cache
Content-Type: application/json
X-Forwarded-For: 172.17.0.1

{
    "firstName": "Stanislav",
    "lastName": "Deviatov"
}

2021/05/05 03:50:44 ---RESPONSE 3---

HTTP/1.1 201 Created
Content-Length: 68
Access-Control-Allow-Origin: *
Content-Type: application/json; charset=UTF-8
Date: Wed, 05 May 2021 03:50:45 GMT
Server: Google Frontend
X-Cloud-Trace-Context: 83ac5937ae7ba8f3ef96ee941227b1b0

{
  "salesforceId": "a0C3L0000008ZSNUA2",
  "action": "updated"
}
```

## Pre-Build Binary

You will find those under [Releases](https://github.com/lil5/http-proxy-logger/releases)

Download the `http-proxy-logger` binary and add it to your `/usr/local/bin/`

## Build binary

```sh
go build -o http-proxy-logger main.go
```

## Run

```sh
http-proxy-logger --port 8001 --target=http://localhost:8000
```
