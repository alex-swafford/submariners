import { useState} from 'react';

import './App.css';
import GamePage from './pages/GamePage';
import LoginPage from './pages/LoginPage';


function App() {
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
      return <LoginPage pageSetter={setCurrentPage} gameState={gameState} gameStateSetter={setGameState}/>
  }
}

export default App;
