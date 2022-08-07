import { QueryClient } from '@tanstack/react-query'
import axios from 'axios'

export const client = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
})

export const queryClient = new QueryClient()
