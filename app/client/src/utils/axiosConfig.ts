import { AxiosRequestConfig } from 'axios'

export const token = () => {
  // Returns the token produced when you join a game
  return (window.sessionStorage.getItem('token') as string)
    ? (window.sessionStorage.getItem('token') as string)
    : ''
}

export const axiosConfig: AxiosRequestConfig = {
  headers: {
    authorization: token(),
  },
}
