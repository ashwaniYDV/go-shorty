# shorten-url-fiber-redis


### Start the server

`docker-compose up -d`

### Request response examples

```
Request:

{
    "url": "https://youtu.be/edCnzelVRlc"
}

Response:

{
    "url": "https://youtu.be/edCnzelVRlc",
    "short": "localhost:3000/34fa1b70-e54c-4a0b-8c30-3c8b6b0c30c6",
    "expiry": 24,
    "rate_limit": 7,
    "rate_limit_reset": 23
}
```

```
Request:

{
    "url": "https://www.youtube.com/channel/UCgMjDy6Y7WISZ529S4VyXUg",
    "short": "tutorial"
}

Response:

{
    "url": "https://www.youtube.com/channel/UCgMjDy6Y7WISZ529S4VyXUg",
    "short": "localhost:3000/tutorial",
    "expiry": 24,
    "rate_limit": 8,
    "rate_limit_reset": 25
}
```

```
Request:

{
    "url": "https://youtu.be/edCnzelVRlc",
    "expiry": 30
}

Response:

{
    "url": "https://youtu.be/edCnzelVRlc",
    "short": "localhost:3000/8d04203d-177d-4c11-b06b-34f3b757e2c9",
    "expiry": 30,
    "rate_limit": 6,
    "rate_limit_reset": 22
}
```

```
Request:

{
    "url": "https://ashwaniYDV.github.io",
    "short": "mywebsite",
    "expiry": 50
}

Response:

{
    "url": "https://ashwaniYDV.github.io",
    "short": "localhost:3000/mywebsite",
    "expiry": 50,
    "rate_limit": 4,
    "rate_limit_reset": 20
}
```


COURTESY:

https://www.youtube.com/watch?v=edCnzelVRlc&list=PL5dTjWUk_cPbXTq9j-3Vaq08rjH1D5cTn