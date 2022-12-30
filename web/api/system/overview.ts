import { ResponseError, RootApi } from '@/base';
import { AxiosError } from 'axios';

export type OverviewErrorType = ResponseError;

export const Overview = (limit?: number) => {
  return RootApi?.get<OverviewData[]>('/overview/', { params: { limit } })
    .then((r) => r.data)
    .catch((err: AxiosError<OverviewErrorType>) => err.response?.data);
};

export const BaseOverview = () => {
  return RootApi?.get<OverviewBaseData>('/overview/base/')
    .then((r) => r.data)
    .catch((err: AxiosError<OverviewErrorType>) => err.response?.data);
};

export interface MemData {
  Total: BigInteger;
  Available: BigInteger;
  STotal: string;
  SAvailable: string;
}

export interface OverviewData {
  time: string;
  mem: MemData;
  cpu_usage: number;
  mem_usage: number;
}
export interface OverviewBaseData {
  cpu: number;
  cores: number;
  mode: string;
  mhz: string;
}
