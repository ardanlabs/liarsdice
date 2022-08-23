import { AxiosRequestConfig } from 'axios'

export const token = () => {
  // Returns the token produced when you join a game
  return (window.localStorage.getItem('token') as string)
    ? (window.localStorage.getItem('token') as string)
    : ''
}

export const axiosConfig: AxiosRequestConfig = {
  headers: {
    authorization: token(),
  },
}
