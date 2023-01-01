import { ResponseError, RootApi } from '@/base';
import { AxiosError } from 'axios';

export type OverviewErrorType = ResponseError;

export const Overview = (limit?: number) => {
  return RootApi?.get<OverviewData[]>('/overview', { params: { limit } })
    .then((r) => r.data)
    .catch((err: AxiosError<OverviewErrorType>) => err.response?.data);
};

export const BaseOverview = () => {
  return RootApi?.get<OverviewBaseData>('/overview/base')
    .then((r) => r.data)
    .catch((err: AxiosError<OverviewErrorType>) => err.response?.data);
};

export interface MemData {
  total: BigInteger;
  available: BigInteger;
  str_total: string;
  str_available: string;
}

export interface OverviewData {
  time: string;
  mem: MemData;
  cpu_usage: number;
  mem_usage: number;
}
export interface OverviewBaseData {
  CPUs: {
    cores: number;
    cpu: number;
    mhz: string;
    mode: string;
  }[];
  host: {
    boot_time: string;
    name: string;
    platform: string;
    version: string;
  };
  system_time: string;
}
