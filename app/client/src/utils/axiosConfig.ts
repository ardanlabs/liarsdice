import { AxiosRequestConfig } from 'axios'

const token = () => {
  // Returns the token produced when you join a game
  return window.sessionStorage.getItem('token') ?? ''
}

export const axiosConfig: AxiosRequestConfig = {
  headers: {
    authorization: token(),
  },
}
