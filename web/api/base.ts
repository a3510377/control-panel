import axios from 'axios';

export const RootApi = axios.create({
  baseURL: '/',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
    Authorization: localStorage.getItem('token'),
  },
});
