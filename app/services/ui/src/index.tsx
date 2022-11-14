import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import { BrowserRouter } from 'react-router-dom'
import reportWebVitals from './reportWebVitals'
import axios, { AxiosResponse } from 'axios'
import { apiUrl } from './utils/axiosConfig'

// Entry point of the FE app.
// We get the root element to allow React to render the app inside that.
const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)

// This constant handles the config fetching accross the app.
const getAppConfigFn = (response: AxiosResponse) => {
  const data = response.data
  return data
}

export const getAppConfig = axios
  .get(`http://${apiUrl}/config`)
  .then(getAppConfigFn)

const strictMode = process.env.NODE_ENV === 'production'

// We render on root.
root.render(
  strictMode ? (
    <React.StrictMode>
      <BrowserRouter>
        <App />
      </BrowserRouter>
      ,
    </React.StrictMode>
  ) : (
    <BrowserRouter>
      <App />
    </BrowserRouter>
  ),
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
