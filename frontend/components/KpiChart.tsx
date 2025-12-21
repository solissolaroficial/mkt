import React from 'react';
import {
  ResponsiveContainer,
  ComposedChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  Bar,
  Line,
  TooltipProps
} from 'recharts';
import { MonthlyData } from '../types';

interface KpiChartProps {
  data: MonthlyData[];
  title: string;
  color: string;
  onClick?: () => void;
}

const CustomTooltip: React.FC<TooltipProps<number, string>> = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-[#1f2937] border border-gray-700 p-3 rounded shadow-lg text-xs">
        <p className="text-gray-300 font-bold mb-1">{label}</p>
        {payload.map((entry, index) => (
          <p key={index} style={{ color: entry.color }}>
            {entry.name}: {entry.value?.toLocaleString()}
          </p>
        ))}
      </div>
    );
  }
  return null;
};

const KpiChart: React.FC<KpiChartProps> = ({ data, title, color, onClick }) => {
  const barColor = '#1a5145';
  const metaColor = '#ffc233';

  return (
    <div className="flex flex-col h-full bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 p-4" onClick={onClick}>
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-lg font-semibold text-gray-100">{title}</h3>
      </div>
      <div className="flex-grow min-h-[250px]">
        <ResponsiveContainer width="100%" height="100%">
          <ComposedChart data={data} margin={{ top: 10, right: 30, left: 0, bottom: 10 }}>
            <defs>
              <linearGradient id={`colorGradient-${title}`} x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor={barColor} stopOpacity={0.4}/>
                <stop offset="95%" stopColor={barColor} stopOpacity={0}/>
              </linearGradient>
            </defs>
            <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#374151" />
            <XAxis 
              dataKey="month" 
              axisLine={false} 
              tickLine={false} 
              tick={{ fill: '#6b7280', fontSize: 12 }} 
              dy={10}
            />
            <YAxis 
              axisLine={false} 
              tickLine={false} 
              tick={{ fill: '#6b7280', fontSize: 12 }} 
            />
            <Tooltip content={<CustomTooltip />} cursor={{ fill: '#20232b' }} />
            <Legend 
              verticalAlign="top" 
              height={36} 
              iconType="circle"
              wrapperStyle={{ fontSize: '12px', color: '#9ca3af' }}
            />
            <Bar 
              name="Realizado" 
              dataKey="realized" 
              fill={`url(#colorGradient-${title})`} 
              stroke={barColor}
              strokeWidth={1}
              radius={[4, 4, 0, 0]}
              barSize={20}
            />
            <Line 
              type="monotone" 
              name="Meta" 
              dataKey="meta" 
              stroke={metaColor} 
              strokeWidth={2} 
              dot={{ r: 3, fill: metaColor, strokeWidth: 0 }} 
              strokeDasharray="4 4"
            />
          </ComposedChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default KpiChart;