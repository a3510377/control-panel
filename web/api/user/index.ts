import { AxiosError } from 'axios';
import { ResponseError, RootApi } from '../base';

interface LoginInfo {
  data: {
    Instances: null; // TODO add instances type
    create_at: string;
    id: number;
    name: string;
    permission: number;
    nick?: string;
  };
  token: { token: string; expiration: string };
}

type LoginErrorType = ResponseError & { type: string };

export const Login = (username: string, password: string) => {
  return RootApi?.post('/account/login', { username, password })
    .then((response): LoginInfo => response.data)
    .catch((err: AxiosError<LoginErrorType, any>) => {
      const status = err.request?.status;

      if ([400, 401].includes(status)) return err.response?.data;

      return void 0;
    });
};
