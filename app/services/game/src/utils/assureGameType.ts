import { game } from '../types/index.d'

// Assures that the game that's passed contains all neccesary keys.
function assureGameType(data: game): game {
  let newGame = data
  newGame = newGame.bets ? newGame : { ...newGame, bets: [] }
  newGame = newGame.cups ? newGame : { ...newGame, cups: [] }
  newGame = newGame.playerOrder ? newGame : { ...newGame, playerOrder: [] }
  return newGame
}
export default assureGameType
