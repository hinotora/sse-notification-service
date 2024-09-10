# SSE Notification service

!!!!!!!!README STILL IN PROGRESS!!!!!!!

Golang + redis pubsub


## Clone & Build

```bash
    $ git clone https://github.com/hinotora/sse-notification-service.git

    $ cd sse-notification-service

    $ cp .env.example .env

    # Configure env ny you own requirements

    $ make build
```

## Run

```bash
    make up
```

## Authorization

```json

// You need to send following JWT payload with sign (sing must be same in .env and JWT)
// HS256 algo

{
  "sub": "2-user-2",
  "iss": "1-stand-1",
  "iat": 1725895592,
  "exp": 1728476792
}

```

## Usage

1. Start to listen SSE connection 

```bash
$ curl -i  http://localhost:8081/sse -H 'Authorization: Bearer jwt

HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: Content-Type
Cache-Control: no-cache
Connection: keep-alive
Content-Type: text/event-stream
X-Connection-Id: 37f347f6-de41-4e1a-be0d-fa0f420f630a
Date: Tue, 10 Sep 2024 15:49:27 GMT
Transfer-Encoding: chunked

```

2. Explore and find active connections

```bash

    127.0.0.1:6379> KEYS *

    1) "1-stand-1:client:2-user-2" # <--- existing sse connection

    127.0.0.1:6379> HGETALL 1-stand-1:client:2-user-2
    1) "channel_id" 
    2) "1-stand-1:channel:2-user-2" # <--- channel name
    3) "application_id"
    4) "1-stand-1"
    5) "user_id"
    6) "2-user-2"

```

3. Send message into choosen channel

```bash
    # redis-cli
    127.0.0.1:6379> PUBLISH 1-stand-1:channel:2-user-2 '{"id":"2", "type":"event", "data":{"hello":"world"}}'
```

```bash

# curl output

id: cc5d30d2-8148-4868-89ea-ddb428a4fdcc
event: ping
data: null

id: 2
event: event
data: {"hello":"world"}

```

## Stopping

```bash
    make restart # to restart

    make down # to stop and delete containers
```