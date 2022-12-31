import { useId } from 'react';
import {
  Area,
  AreaChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';
import Title from './title';

export default function Chart(props: {
  type: string;
  data: any[];
  title?: string;
  color?: string;
}) {
  const color = props.color || '#8884d8';
  const colorID = useId();
  return (
    <>
      <Title>{props.title}</Title>
      <ResponsiveContainer width="100%" height="100%">
        <AreaChart
          width={500}
          height={300}
          data={props.data}
          margin={{
            top: 5,
            right: 30,
            left: 20,
            bottom: 5,
          }}
        >
          <defs>
            <linearGradient id={colorID} x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor={color} stopOpacity={0.8} />
              <stop offset="95%" stopColor={color} stopOpacity={0} />
            </linearGradient>
          </defs>
          <XAxis dataKey="time" />
          <Tooltip
            formatter={(value) => `${value}%`}
            labelStyle={{ color: 'black' }}
          />
          <CartesianGrid strokeDasharray="3 3" />
          <YAxis tickFormatter={(tick) => `${tick}%`} />
          <Area
            name={props.type}
            type="monotone"
            dataKey={props.type}
            stroke={color}
            fillOpacity={1}
            fill={`url(#${colorID})`}
          />
        </AreaChart>
      </ResponsiveContainer>
    </>
  );
}
