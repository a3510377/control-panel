import { RootApi } from './base';

export const test = () => {
  RootApi.get('/api/instance');
};
