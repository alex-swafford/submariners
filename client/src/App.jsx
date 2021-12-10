import { useState} from 'react';

import './App.css';
import GamePage from './pages/GamePage';
import LoginPage from './pages/LoginPage';


function App() {
  const [playerName, setPlayerName] = useState('')
  const [currentPage, setCurrentPage] = useState('login')
  const [gameState, setGameState] = useState({
    players: [
        {
            x: 200,
            y: 200,
            angleRadians: Math.PI / 4,
            speed: 0,
        }
    ],
    previousAnimationTime: Date.now()
})
  switch(currentPage) {
    case 'game':
      return <GamePage playerName={playerName} gameState={gameState} gameStateSetter={setGameState}/>
    case 'login':
    default:
      return <LoginPage pageSetter={setCurrentPage} currentName={playerName} nameSetter={setPlayerName}/>
  }
}

export default App;
