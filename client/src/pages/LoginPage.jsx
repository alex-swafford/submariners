import {TextField, Button, ThemeProvider} from '@material-ui/core'

import MainTheme from '../themes/MainTheme'

import './LoginPage.css'

/**
 * 
 * @param {pageSetter, nameSetter, currentName} props 
 * @returns a login page. 
 */
function LoginPage(props) {
    var isButtonDisabled = props.gameState.players.length === 0 ||
                           props.gameState.players[0].name === ""
    return <ThemeProvider theme={MainTheme} >
        <div className='LoginPage'>
            <div className='Title'>
                Submarine Shenanigans
            </div>
            <TextField label="name" onChange={event => props.gameStateSetter(
                {
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
                }
            )}/>
            <Button variant='contained' disabled={isButtonDisabled} color='primary' onClick={() => {
                props.pageSetter('game') 
                }}>
                    Connect
            </Button>
        </div>
    </ThemeProvider>
    
}

export default LoginPage