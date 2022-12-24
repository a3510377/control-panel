import axios from 'axios';

export type ResponseError = { error: string };

export const RootApi =
  (typeof window !== 'undefined' || void 0) &&
  axios.create({
    baseURL: process.env.API_BASE_URL + '/api',
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
      Authorization: localStorage.getItem('token') || '',
    },
    params: {
      l: navigator.language || localStorage.getItem('lang') || 'zh-TW',
    },
  });
