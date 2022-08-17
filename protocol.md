# Communications

## WebSockets

### Connect Messages

Regardless of how many clients are being served the web page, players are not considered connected until they have sent a 'connect' message over a WebSocket channel to the server and have received the proper response.

Connect message format:

```json
{
    "messageType": "connect", 
    "playerName": "<username>"
    "data": {}
}
```

Once the server has received a client's valid connect message, it must send either a 'connected' message or a 'connection refused' message to the client who sent the connect message.

'Connected' message: 

```json
{
    "messageType": "connected"
    "data": {
        "connectedPlayers": [
            "<first connected player>", 
            "<second connected player>", 
            ..., 
            "<player who just connected>"]
    }
}
```

'Connection Refused' message: 

```json
{
    "messageType": "connectionRefused",
    "data": {}
}
```