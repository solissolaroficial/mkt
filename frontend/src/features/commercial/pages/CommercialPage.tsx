import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useCommercial } from '../hooks/useCommercial';
import { useUIStore } from '@/shared/store/uiStore';
import { Target, TrendingUp } from 'lucide-react';
import SummaryCard from '../ui/SummaryCard';

const CommercialPage: React.FC = () => {
  const navigate = useNavigate();
  const { commercialKpis, isLoading, error } = useCommercial();
  const { selectedMonth } = useUIStore();

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <div className="w-16 h-16 border-4 border-[#1e5144] border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-400">Carregando KPIs comerciais...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 mb-2">Erro ao carregar KPIs comerciais</p>
          <p className="text-gray-500 text-sm">{error.message}</p>
        </div>
      </div>
    );
  }

  const getKpiSummary = (kpi: any) => {
    const data = kpi.data;
    const validData = data.filter((d: any) => d.realized !== null);

    if (selectedMonth === '---') {
      if (validData.length === 0) return { realized: 0, meta: 0, month: 'ANO' };

      const shouldAverage = kpi.unit === 'percent' || ['conversion_rate'].includes(kpi.id);
      const totalRealized = validData.reduce((acc: number, curr: any) => acc + (curr.realized || 0), 0);

      let totalMeta = 0;
      if (shouldAverage) {
        totalMeta = validData.reduce((acc: number, curr: any) => acc + (curr.meta || 0), 0);
      } else {
        totalMeta = data.reduce((acc: number, curr: any) => acc + (curr.meta || 0), 0);
      }

      if (shouldAverage) {
        return {
          realized: totalRealized / validData.length,
          meta: totalMeta / validData.length,
          month: 'MÉDIA'
        };
      }

      return {
        realized: totalRealized,
        meta: totalMeta,
        month: 'TOTAL'
      };
    }

    const monthData = data.find((d: any) => d.month === selectedMonth);

    if (monthData) {
      return {
        realized: monthData.realized || 0,
        meta: monthData.meta || 0,
        month: monthData.month
      };
    }

    return { realized: 0, meta: 0, month: selectedMonth };
  };

  const getKpiColorClass = (id: string) => {
    if (id.includes('sales') || id.includes('opportunities')) return 'bg-pink-500/10';
    if (id.includes('conversion')) return 'bg-rose-500/10';
    return 'bg-blue-500/10';
  };

  return (
    <div className="space-y-6 animate-in fade-in duration-500">
      <h2 className="text-xl font-bold text-gray-200">Comercial</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {commercialKpis.map((kpi) => {
          const { realized, meta } = getKpiSummary(kpi);

          const handleClick = () => {
            navigate(`/kpis/${kpi.id}`);
          };

          return (
            <SummaryCard
              key={kpi.id}
              title={kpi.title}
              value={realized}
              target={meta}
              icon={TrendingUp}
              colorClass={getKpiColorClass(kpi.id)}
              onClick={() => navigate(`/kpis/${kpi.id}`)}
              unit={kpi.unit}
            />
          );
        })}
      </div>
    </div>
  );
};

export default CommercialPage;
