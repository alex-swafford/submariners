import {Table, TableHead, TableRow, TableCell, Button, TableBody, ThemeProvider} from '@material-ui/core'
import { useRef, useEffect } from 'react'

import useKeyPress from '../components/UseKeyPress'
import submarineImage from '../images/SubmarineGreen.png'
import MainTheme from '../themes/MainTheme'

import './GamePage.css'
function GamePage(props) {
    let submarineImageRef = useRef()
    // Turn controls
    const isAPressed = useKeyPress({key : "a"})
    const isDPressed = useKeyPress({key : "d"})
    // Accelerator / Brake
    const isWPressed = useKeyPress({key: "w"})
    const isSPressed = useKeyPress({key: "s"})
    useEffect(() => {
        const interval = setInterval(() => {
            let players = props.gameState.players
            let player = players[0]
            let nowTime = Date.now()
            let elapsedTime = nowTime - props.gameState.previousAnimationTime
            player.x += player.speed * Math.cos(player.angleRadians) * elapsedTime
            player.y += player.speed * Math.sin(player.angleRadians) * elapsedTime
            console.log("Player x:", player.x, "Player y:", player.y, "Player speed:", player.speed)
            let newPlayerAngle = player.angleRadians
            if(isAPressed && !isDPressed) {
                newPlayerAngle -= Math.PI / 600 * elapsedTime
            }
            else if(isDPressed && !isAPressed) {
                newPlayerAngle += Math.PI / 600 * elapsedTime
            }
            if(isWPressed && !isSPressed && player.speed < .3) {
                player.speed += .001 * elapsedTime
                if(player.speed > .3) {
                    player.speed = .3
                }
            }
            else if(isSPressed && !isWPressed && player.speed > 0) {
                player.speed -= .001 * elapsedTime
                if(player.speed < 0) {
                    player.speed = 0
                }
            }
            else if(player.speed > 0) {
                player.speed -= .0002 * elapsedTime
                if (player.speed < 0) {
                    player.speed = 0
                }
            }
            props.gameStateSetter({
            players: [ // Note: this will break as soon as another player is added.
                        {
                            x: player.x,
                            y: player.y,
                            angleRadians: newPlayerAngle,
                            speed: player.speed
                        }
                    ],
            previousAnimationTime: nowTime
        })}, 12)
        return () => {clearInterval(interval)}
    }, [props, isWPressed, isAPressed, isSPressed, isDPressed])
    const canvasRef = useRef()
    useEffect(() => {
        let players = props.gameState.players
        let player = players[0]
        let canvas = canvasRef.current
        let ctx = canvas.getContext('2d')
        ctx.clearRect(0, 0, canvas.width, canvas.height)
        ctx.beginPath()
        ctx.fillStyle = '#000'
        ctx.save()
        try{
            ctx.save()
            ctx.translate(player.x+50, player.y+50)
            ctx.rotate(player.angleRadians)
            ctx.translate(-(player.x+50),-(player.y+50))
            ctx.drawImage(submarineImageRef.current, player.x, player.y)
            ctx.restore()
        }
        catch(e) {
            ctx.rect(player.x, player.y, 50, 50)
            console.log("Failed to draw image:", e)
        }
        ctx.restore()
        ctx.fill()
    }, [props.gameState])
    return <ThemeProvider theme={MainTheme}>
    <img src={submarineImage} ref={submarineImageRef} alt='failed to load.' hidden={true}/>
    <div className='GamePage'>
        <div className='Scoreboard SubWindow'>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Player Name</TableCell>
                        <TableCell>Score</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow>
                        <TableCell>Julius Caesar</TableCell>
                        <TableCell>-44</TableCell>
                    </TableRow>
                    <TableRow>
                        <TableCell>Napolean Bonaparte</TableCell>
                        <TableCell>1821</TableCell>
                    </TableRow>
                    <TableRow>
                        <TableCell>Douglas MacArthur</TableCell>
                        <TableCell>1964</TableCell>
                    </TableRow>
                </TableBody>
                
            </Table>
        </div>
        <div className='MainView SubWindow'>
            <canvas width="1000" height="600" ref={canvasRef} className='Canvas'></canvas>
        </div>
            
        <div className='DataPanel SubWindow'>Data Panel</div>
            
        <div className='ControlBar SubWindow'>
            {props.playerName}'s Submarine
            <div className='right-align'>
                <Button variant='contained' color='primary' className='Button'>ping</Button>
                <Button variant='contained' color='primary' className='Button'>grab</Button>
                <Button variant='contained' color='primary' className='Button'>drop</Button>
            </div>
        </div>

    </div>
    </ThemeProvider> 
    
    
}

export default GamePage