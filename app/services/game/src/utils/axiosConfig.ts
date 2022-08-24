import { AxiosRequestConfig } from 'axios'

export const token = () => {
  // Returns the token stored in the session storage of the player, produced when you join a game
  return (window.sessionStorage.getItem('token') as string)
    ? (window.sessionStorage.getItem('token') as string)
    : ''
}

// Default headers for axios.
export const axiosConfig: AxiosRequestConfig = {
  headers: {
    authorization: token(),
  },
}

// Check the .env file inside src/ to see if the player has a GO_HOST configuration
// Prefix REACT_APP is used so react recongnize the variable
export const apiUrl = process.env.REACT_APP_GO_HOST
  ? process.env.REACT_APP_GO_HOST
  : 'localhost:3000/v1/game'
