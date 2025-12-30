import React from 'react';
import { useNavigate } from 'react-router-dom';
import { GraduationCap, Users } from 'lucide-react';
import { useTraining } from '../hooks/useTraining';
import { useUIStore } from '@/shared/store/uiStore';
import SummaryCard from '@/features/commercial/ui/SummaryCard';
import HeatmapChart from '../ui/HeatmapChart';
import { REPS_TRAINING_DATA, FULL_MONTH_MAP } from '@/shared/utils/legacy.constants';

const TrainingPage: React.FC = () => {
  const navigate = useNavigate();
  const { trainingKpis, isLoading, error } = useTraining();
  const { selectedMonth } = useUIStore();

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-500"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 mb-2">Erro ao carregar KPIs de treinamento</p>
          <p className="text-gray-500 text-sm">{(error as Error).message}</p>
        </div>
      </div>
    );
  }

  const getKpiSummary = (kpi: any) => {
    const data = kpi.data;
    const validData = data.filter((d: any) => d.realized !== null);

    if (selectedMonth === '---') {
      // Ano completo - soma todos os valores
      const totalRealized = validData.reduce((acc: number, curr: any) => acc + (curr.realized || 0), 0);
      const totalMeta = data.reduce((acc: number, curr: any) => acc + (curr.meta || 0), 0);
      return {
        realized: totalRealized,
        meta: totalMeta
      };
    }

    // Mês específico
    const monthData = data.find((d: any) => d.month === selectedMonth);
    return {
      realized: monthData?.realized || 0,
      meta: monthData?.meta || 0
    };
  };

  const getKpiIcon = (slug: string) => {
    if (slug.includes('treinamento')) return GraduationCap;
    return Users;
  };

  const getKpiColorClass = (slug: string) => {
    if (slug.includes('treinamento')) return 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20';
    return 'bg-blue-500/10 text-blue-400 border-blue-500/20';
  };

  // Filter heatmap data based on selected month
  const getFilteredTrainingData = () => {
    if (selectedMonth === '---') {
      return REPS_TRAINING_DATA;
    }

    const fullMonthName = FULL_MONTH_MAP[selectedMonth];
    const filteredRows = REPS_TRAINING_DATA.rows.filter(row => row.month === fullMonthName);
    
    return {
      ...REPS_TRAINING_DATA,
      rows: filteredRows
    };
  };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div className="flex items-center gap-4">
          <button onClick={() => navigate('/kpis')} className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white">
            <GraduationCap size={24} />
          </button>
          <div>
            <h1 className="text-2xl font-bold text-gray-100">Treinamentos</h1>
            <p className="text-gray-500">Visão Geral de Treinamentos</p>
          </div>
        </div>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        {trainingKpis.map((kpi) => {
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
      {/* Heatmap Chart */}
      <HeatmapChart
        data={getFilteredTrainingData()}
        title={`Frequência de Treinamentos (${selectedMonth === '---' ? 'Ano Completo' : FULL_MONTH_MAP[selectedMonth]})`}
      />
    </div>
  );
};

export default TrainingPage;
