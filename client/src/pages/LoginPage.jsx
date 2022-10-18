import {TextField, Button, ThemeProvider} from '@material-ui/core'
import { useEffect } from 'react'

import MainTheme from '../themes/MainTheme'

import './LoginPage.css'

/**
 * 
 * @param {pageSetter, nameSetter, currentName, sendMessage, wireMessageHandler} props 
 * @returns a login page. 
 */
function LoginPage(props) {

    useEffect(() => {
        props.wireMessageHandler((event) => {
            console.log(`Received message from server: ${event.data}`);
            var decodedMessage = JSON.parse(event.data);
            if(decodedMessage.MessageType === "connected") {
                props.gameStateSetter({
                    players: decodedMessage.Data.connectedPlayers,
                    previousAnimationTime: Date.now()
                });
                props.pageSetter("game");
            }
        })
    }, [])

    var isButtonDisabled = props.gameState.players.length === 0 ||
                           props.gameState.players[0].name === ""
    return <ThemeProvider theme={MainTheme} >
        <div className='LoginPage'>
            <div className='Title'>
                Stupendous Submarine Stunts
            </div>
            <TextField label="name" onChange={event => {
                    props.gameStateSetter({
                        players: [
                                    {
                                        x: 203,
                                        y: 200,
                                        angleRadians: Math.PI / 4,
                                        speed: 0,
                                        name: event.target.value
                                    }
                        ],
                        previousAnimationTime: Date.now()
                    });
                }
            }/>
            <Button variant='contained' disabled={isButtonDisabled} color='primary' onClick={() => {
                var messageToSend = `{"messageType": "connect", "playerName": "${props.gameState.players[0].name}", "data": {}}`;
                console.log(`Sending ${JSON.stringify(messageToSend)}`);
                props.sendMessage(messageToSend);
                }}>
                    Connect
            </Button>
        </div>
    </ThemeProvider>
    
}

export default LoginPage