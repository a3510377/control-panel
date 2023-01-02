import { RootApi } from '@/base';

export interface User {
  id: BigInt;
  name: string;
  nick: string;
  permission: BigInt;
  create_at: string;
}

export const GetUsers = () => {
  return RootApi.get('/users')
    .then((d) => d.data)
    .then((users: { data: User[] }) => users.data);
};
