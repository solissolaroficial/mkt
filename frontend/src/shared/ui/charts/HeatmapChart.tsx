import React from 'react';
import type { KpiCategory } from '@/features/kpis';
import { MONTHS } from '@/shared/utils/constants';
import { formatKpiValue } from '@/shared/utils/formatters';

interface HeatmapChartProps {
  kpi: KpiCategory;
}

const HeatmapChart: React.FC<HeatmapChartProps> = ({ kpi }) => {
  // Calcular min/max para escala de cores
  const values = kpi.data.map((d) => d.realized || 0);
  const minValue = Math.min(...values);
  const maxValue = Math.max(...values);

  // Função para calcular intensidade da cor (0-1)
  const getIntensity = (value: number) => {
    if (maxValue === minValue) return 0.5;
    return (value - minValue) / (maxValue - minValue);
  };

  // Converter intensidade em cor (do escuro para o verde)
  const getColor = (intensity: number) => {
    const baseColor = kpi.color || '#1e5144';
    const opacity = 0.2 + intensity * 0.8; // 0.2 a 1.0
    return `${baseColor}${Math.round(opacity * 255).toString(16).padStart(2, '0')}`;
  };

  return (
    <div className="space-y-2">
      <div className="grid grid-cols-6 gap-2">
        {MONTHS.map((month) => {
          const monthData = kpi.data.find((d) => d.month === month);
          const value = monthData?.realized || 0;
          const meta = monthData?.meta || 0;
          const intensity = getIntensity(value);
          const color = getColor(intensity);

          return (
            <div
              key={month}
              className="group relative aspect-square rounded-lg border border-gray-700 hover:border-gray-600 transition-all cursor-pointer"
              style={{ backgroundColor: color }}
            >
              {/* Month Label */}
              <div className="absolute inset-0 flex items-center justify-center">
                <span className="text-xs font-semibold text-white">
                  {month}
                </span>
              </div>

              {/* Tooltip */}
              <div className="absolute bottom-full left-1/2 transform -translate-x-1/2 mb-2 px-3 py-2 bg-[#1a1d24] border border-gray-700 rounded-lg shadow-xl opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap z-10">
                <p className="text-xs text-gray-400">{month}</p>
                <p className="text-sm font-semibold text-white">
                  {formatKpiValue(value, kpi.unit)}
                </p>
                {meta > 0 && (
                  <p className="text-xs text-gray-500">
                    Meta: {formatKpiValue(meta, kpi.unit)}
                  </p>
                )}
                {/* Arrow */}
                <div className="absolute top-full left-1/2 transform -translate-x-1/2 -mt-px">
                  <div className="border-4 border-transparent border-t-gray-700"></div>
                </div>
              </div>
            </div>
          );
        })}
      </div>

      {/* Legend */}
      <div className="flex items-center justify-between text-xs text-gray-500 pt-2">
        <span>Menor: {formatKpiValue(minValue, kpi.unit)}</span>
        <span>Maior: {formatKpiValue(maxValue, kpi.unit)}</span>
      </div>
    </div>
  );
};

export default HeatmapChart;
