import { ResponseError, RootApi } from '@/base';
import { AxiosError } from 'axios';

export type OverviewErrorType = ResponseError;

export const Overview = () => {
  return RootApi?.get<Overview>('/overview')
    .then((r) => r.data)
    .catch((err: AxiosError<OverviewErrorType>) => err.response?.data);
};

export const BaseOverview = () => {
  return RootApi?.get<OverviewBase>('/overview/base')
    .then((r) => r.data)
    .catch((err: AxiosError<OverviewErrorType>) => err.response?.data);
};

export interface MemData {
  Total: BigInteger;
  Available: BigInteger;
  UsedPercent: number;
  STotal: string;
  SAvailable: string;
}

export interface Overview {
  time: string;
  mem: MemData;
  cpu_usage: number;
}
export interface OverviewBase {
  cpu: number;
  cores: number;
  mode: string;
  mhz: string;
}
