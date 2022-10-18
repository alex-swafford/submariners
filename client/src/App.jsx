import { useState} from 'react';

import './App.css';
import { w3cwebsocket as W3CWebSocket } from "websocket";
import GamePage from './pages/GamePage';
import LoginPage from './pages/LoginPage';
import { useEffect } from 'react';

const isDevelopMode = true;

var wsURL = isDevelopMode ? `ws://localhost:3000/ws` : `ws://${window.location.href.split('://')[1]}ws`

console.log(`ws URL: ${wsURL}`);

var wsClient = new W3CWebSocket(wsURL);;

/**
 * Exposes ws message send function to child components.
 * @param {array<byte>} message 
 */
function sendWSMessage(message) {
  wsClient.send(message);
}

function wireWSReceiver(handler) {
  wsClient.onmessage = handler;
}


function App() {

  useEffect(() => {
    wsClient.onopen = () => {
      console.log("WS Connection established.");
    }
    wsClient.onerror = (error) => {
      console.log(`Error from websocket: ${error.message}`)
    }
  }, [])

  

  const [currentPage, setCurrentPage] = useState('login')
  const [gameState, setGameState] = useState({
    players: [],
    previousAnimationTime: Date.now()
})
  switch(currentPage) {
    case 'game':
      return <GamePage gameState={gameState} gameStateSetter={setGameState}/>
    case 'login':
    default:
      return <LoginPage pageSetter={setCurrentPage} gameState={gameState} gameStateSetter={setGameState} 
        wsClient ={wsClient} sendMessage={sendWSMessage} wireMessageHandler={wireWSReceiver}/>
  }
}

export default App;
