# Communications

## WebSockets

Regardless of how many clients are being served the web page, players are not considered connected until they have sent a 'connect' message over a WebSocket channel to the server.

Connect message format:

```
{
    "messageType": "connect", 
    "player": "<username>"
    "data": {}
}
```

