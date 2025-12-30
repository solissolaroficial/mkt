import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useMarketingKpis } from '../hooks/useMarketing';
import { useUIStore } from '@/shared/store/uiStore';
import { Youtube, GraduationCap, Megaphone, Users } from 'lucide-react';
import SummaryCard from '@/features/commercial/ui/SummaryCard';

const MarketingPage: React.FC = () => {
  const navigate = useNavigate();
  const { marketingKpis, isLoading, error } = useMarketingKpis();
  const { selectedMonth } = useUIStore();

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <div className="w-16 h-16 border-4 border-[#1e5144] border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-400">Carregando KPIs de Marketing...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 mb-2">Erro ao carregar KPIs de Marketing</p>
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

      const shouldAverage = kpi.unit === 'percent';
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

  const getKpiIcon = (slug: string) => {
    if (slug === 'novos_inscritos_no_youtube') return Youtube;
    if (slug.includes('treinamento')) return GraduationCap;
    if (slug.includes('marketing') || slug.includes('inbound')) return Megaphone;
    return Users;
  };

  const getKpiColorClass = (slug: string) => {
    if (slug === 'novos_inscritos_no_youtube') return 'bg-red-500/10 text-red-400 border-red-500/20';
    if (slug.includes('treinamento')) return 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20';
    if (slug.includes('marketing')) return 'bg-purple-500/10 text-purple-400 border-purple-500/20';
    if (slug.includes('inbound')) return 'bg-amber-500/10 text-amber-400 border-amber-500/20';
    return 'bg-blue-500/10 text-blue-400 border-blue-500/20';
  };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div className="flex items-center gap-4">
          <button
            onClick={() => navigate('/kpis')}
            className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
          >
            <Users size={24} />
          </button>
          <div>
            <h1 className="text-2xl font-bold text-gray-100">Marketing</h1>
            <p className="text-gray-500">Visão Geral de Marketing</p>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {marketingKpis.map((kpi) => {
          const { realized, meta } = getKpiSummary(kpi);
          const Icon = getKpiIcon(kpi.slug);

          return (
            <SummaryCard
              key={kpi.id}
              title={kpi.title}
              value={realized}
              target={meta}
              icon={Icon}
              colorClass={getKpiColorClass(kpi.slug)}
              onClick={() => navigate(`/kpis/${kpi.id}`)}
              unit={kpi.unit}
            />
          );
        })}
      </div>
    </div>
  );
};

export default MarketingPage;
