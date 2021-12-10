import {TextField, Button, ThemeProvider} from '@material-ui/core'

import MainTheme from '../themes/MainTheme'

import './LoginPage.css'

/**
 * 
 * @param {pageSetter, nameSetter, currentName} props 
 * @returns a login page. 
 */
function LoginPage(props) {
    return <ThemeProvider theme={MainTheme} >
        <div className='LoginPage'>
            <div className='Title'>
                Submarine Shenanigans
            </div>
            <TextField label="name" onChange={event => props.nameSetter(event.target.value)}/>
            <Button variant='contained' disabled={props.currentName === ''} color='primary' onClick={() => {
                props.pageSetter('game') 
                }}>
                    Connect
            </Button>
        </div>
    </ThemeProvider>
    
}

export default LoginPage