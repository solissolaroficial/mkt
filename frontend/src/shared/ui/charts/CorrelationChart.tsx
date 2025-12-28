import React from 'react';
import {
  ResponsiveContainer,
  ScatterChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Scatter,
  Label,
  TooltipProps
} from 'recharts';

interface CorrelationChartProps {
  data: { x: number; y: number }[];
  xLabel: string;
  yLabel: string;
  color: string;
}

const CustomTooltip: React.FC<any> = ({ active, payload }) => {
    if (active && payload && payload.length) {
      return (
        <div className="bg-[#1f2937] border border-gray-700 p-3 rounded shadow-lg text-xs text-white">
          <p className="text-gray-300">
              X: {payload[0].value}<br/>
              Y: {payload[1].value}
          </p>
        </div>
      );
    }
    return null;
};

const CorrelationChart: React.FC<CorrelationChartProps> = ({ data, xLabel, yLabel, color }) => {
  return (
    <div className="flex flex-col h-full bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 p-4">
      <div className="flex-grow min-h-[300px]">
        <ResponsiveContainer width="100%" height="100%">
          <ScatterChart margin={{ top: 20, right: 30, bottom: 20, left: 20 }}>
            <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
            <XAxis 
              type="number" 
              dataKey="x" 
              name={xLabel} 
              unit="R$" 
              tick={{ fontSize: 12, fill: '#6b7280' }}
              axisLine={{ stroke: '#4b5563' }}
              tickLine={{ stroke: '#4b5563' }}
            >
              <Label value={xLabel} offset={0} position="bottom" style={{ fill: '#9ca3af', fontSize: '12px' }} />
            </XAxis>
            <YAxis 
              type="number" 
              dataKey="y" 
              name={yLabel} 
              tick={{ fontSize: 12, fill: '#6b7280' }}
              axisLine={{ stroke: '#4b5563' }}
              tickLine={{ stroke: '#4b5563' }}
            >
              <Label value={yLabel} angle={90} position="insideLeft" style={{ fill: '#9ca3af', fontSize: '12px' }} />
            </YAxis>
            <Tooltip content={<CustomTooltip />} cursor={{ strokeDasharray: '3 3', stroke: '#4b5563' }} />
            <Scatter name="Dados" data={data} fill={color} shape="circle" />
          </ScatterChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default CorrelationChart;
