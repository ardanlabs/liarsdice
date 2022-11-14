import { AxiosError, AxiosResponseHeaders, AxiosRequestConfig } from 'axios'

export interface connectResponse {
  data: {
    address: string
    token: string
  }
}

export interface usd2weiResponse {
  data: {
    usd: number
    wei: number
  }
}

export interface defaultApiError<D = any> extends AxiosError {
  response:  {
    status: number;
    statusText: string;
    headers: AxiosResponseHeaders;
    config: AxiosRequestConfig<D>;
    request?: any;
    data: {
      error: string
    }
  }
}