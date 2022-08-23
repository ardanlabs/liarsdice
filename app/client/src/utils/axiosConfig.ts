import { AxiosRequestConfig } from 'axios'

const token = () => {
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
