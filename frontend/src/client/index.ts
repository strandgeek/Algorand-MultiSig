import { QueryClient } from '@tanstack/react-query'
import axios from 'axios'

export const client = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
})

client.interceptors.request.use(
  config => {
    if (!config.headers) {
      config.headers = {}
    }
    const token = localStorage.getItem("token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    Promise.reject(error);
  }
);

client.interceptors.response.use(response => {
  return response;
}, error => {
 if (error.response.status === 401) {
  window.location.href = '/'
 }
 return error;
});

export const queryClient = new QueryClient()
