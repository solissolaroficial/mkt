import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useKpis } from '../hooks/useKpis';
import KpiChart from './KpiChart';
import { useUIStore } from '@/shared/store/uiStore';
import { Target, TrendingUp, TrendingDown, ArrowRight } from 'lucide-react';
import { formatKpiValue } from '@/shared/utils/formatters';

const KpiListView: React.FC = () => {
  const navigate = useNavigate();
  const { data, isLoading, error } = useKpis();
  const { selectedMonth } = useUIStore();

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <div className="w-16 h-16 border-4 border-[#1e5144] border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-400">Carregando KPIs...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 mb-2">Erro ao carregar KPIs</p>
          <p className="text-gray-500 text-sm">{error.message}</p>
        </div>
      </div>
    );
  }

  const kpis = data?.data || [];

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header */}
      <div className="mb-6">
        <h1 className="text-3xl font-bold text-white mb-2">KPIs</h1>
        <p className="text-gray-400">
          Indicadores de Performance - {selectedMonth === '---' ? 'Visão Anual' : selectedMonth}
        </p>
      </div>

      {/* KPIs Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 flex-grow">
        {kpis.map((kpi) => (
          <KpiChart
            key={kpi.id}
            kpi={kpi}
            months={['JAN', 'FEV', 'MAR', 'ABR', 'MAI', 'JUN', 'JUL', 'AGO', 'SET', 'OUT', 'NOV', 'DEZ']}
            onClick={() => navigate(`/kpis/${kpi.id}`)}
          />
        ))}
      </div>

      {/* Empty State */}
      {kpis.length === 0 && (
        <div className="flex flex-col items-center justify-center h-full text-center">
          <Target className="text-gray-600 mb-4" size={48} />
          <h3 className="text-xl font-semibold text-gray-300 mb-2">Nenhum KPI encontrado</h3>
          <p className="text-gray-500">
            Não há indicadores configurados no momento.
          </p>
        </div>
      )}
    </div>
  );
};

export default KpiListView;
