import React from 'react';
import {
  ResponsiveContainer,
  FunnelChart as RechartsFunnelChart,
  Funnel,
  LabelList,
  Tooltip,
  Cell,
  TooltipProps
} from 'recharts';

interface FunnelData {
  step: string;
  value: number;
  fill: string;
}

interface FunnelChartProps {
  data: FunnelData[];
}

const CustomTooltip: React.FC<TooltipProps<number, string>> = ({ active, payload }) => {
    if (active && payload && payload.length) {
        const data = payload[0].payload;
        return (
            <div className="bg-[#1f2937] border border-gray-700 p-2 rounded shadow-lg text-xs text-white">
                <span className="font-bold">{data.step}:</span> {data.value.toLocaleString()}
            </div>
        );
    }
    return null;
};

const FunnelChart: React.FC<FunnelChartProps> = ({ data }) => {
  return (
    <div className="flex flex-col h-full bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 p-4">
        <h3 className="text-lg font-semibold text-gray-100 mb-4">Funil de Conversão</h3>
        <div className="flex-grow min-h-[300px]">
            <ResponsiveContainer width="100%" height="100%">
            <RechartsFunnelChart margin={{ top: 20, right: 50, bottom: 20, left: 20 }}>
                <Tooltip content={<CustomTooltip />} cursor={{ fill: 'transparent' }} />
                <Funnel
                dataKey="value"
                data={data}
                isAnimationActive
                >
                <LabelList 
                    position="right" 
                    fill="#9ca3af" 
                    stroke="none" 
                    dataKey="step" 
                    content={(props: any) => {
                        const { x, y, width, value, index } = props;
                        if (!data[index]) return null;
                        const stepName = data[index].step;
                        const percent = index === 0 ? 100 : Math.round((value / data[0].value) * 100);
                        return (
                            <text x={x + width + 10} y={y + 20} fill="#9ca3af" textAnchor="start" fontSize={12} fontWeight="bold">
                                {stepName}: {value.toLocaleString()} ({percent}%)
                            </text>
                        );
                    }}
                />
                {data.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.fill} stroke="none" />
                ))}
                </Funnel>
            </RechartsFunnelChart>
            </ResponsiveContainer>
        </div>
    </div>
  );
};

export default FunnelChart;