import axios, { Axios, AxiosError, AxiosResponse } from 'axios'
import { game } from '../types/index.d'
import { axiosConfig } from '../utils/axiosConfig'

const apiUrl = process.env.REACT_APP_GO_HOST
  ? process.env.REACT_APP_GO_HOST
  : 'localhost:3000/v1/game'

const setNewGame = (data: game) => {
  let newGame = data
  newGame = newGame.claims ? newGame : { ...newGame, claims: [] }
  newGame = newGame.cups ? newGame : { ...newGame, cups: [] }
  newGame = newGame.player_order ? newGame : { ...newGame, player_order: [] }
  return newGame
}

function getStatus() {
  axios
    .get(`http://${apiUrl}/status`, axiosConfig)
    .then(function (response: AxiosResponse) {
      if (response.data) {
        return response.data
      }
    })
    .catch(function (error: AxiosError) {
      return error
    })
}

export { getStatus }
