import axios from 'axios';
import jsonBig from 'json-bigint';

export type ResponseError = { error: string };

export const RootApi =
  (typeof window !== 'undefined' || void 0) &&
  axios.create({
    baseURL: process.env.API_BASE_URL + '/api',
    headers: { 'Content-Type': 'application/json', Accept: 'application/json' },
    transformRequest: [
      function (data, headers) {
        this.params = {
          l: navigator.language || localStorage.getItem('lang') || 'zh_TW',
          ...this.params,
        };
        headers.set('Authorization', GetToken());
        return data;
      },
      // @ts-ignore
      ...axios.defaults.transformRequest,
    ],
    transformResponse: (r) => jsonBig().parse(r),
  });

export const GetToken = () => localStorage.getItem('token') || void 0;
