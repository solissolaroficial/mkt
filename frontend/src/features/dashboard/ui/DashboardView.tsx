import React from 'react';
import { Target, Youtube, GraduationCap, Megaphone, TrendingUp } from 'lucide-react';
import { useDashboardKpis } from '@/features/dashboard';
import { useKpiCalculations } from '@/features/kpis/hooks/useKpiCalculations';
import { useUIStore } from '@/shared/store/uiStore';
import SummaryCard from './SummaryCard';
import { type KpiCategory } from '@/features/kpis';
import { TaskWidget } from '@/features/tasks';
import { UpcomingPostsWidget } from '@/features/calendar';
import { useNavigate } from 'react-router-dom';
import { useTasks } from '@/features/tasks';
import { useCalendarPosts } from '@/features/calendar';

const DashboardView: React.FC = () => {
  const { data, isLoading, error } = useDashboardKpis();
  const { selectedMonth } = useUIStore();
  const navigate = useNavigate();
  const { data: tasksData } = useTasks();
  const { data: postsData } = useCalendarPosts();
  
  // Aplicar cálculos derivados (CPL, etc)
  const kpis = useKpiCalculations(data || []);
  
  // KPIs já filtrados pelo backend via slugs
  const mainKpis = kpis;

  // Calcular estatísticas para cada KPI
  const getKpiStats = (kpi: KpiCategory) => {
    const currentMonthData = kpi.data.find((d) => d.month === selectedMonth);
    const realized = currentMonthData?.realized || 0;
    const meta = currentMonthData?.meta || 0;
    return { realized, meta };
  };

  // Mapear ícones por título do KPI
  const getIconForKpi = (title: string) => {
    if (title.includes('Youtube')) return Youtube;
    if (title.includes('Treinamentos')) return GraduationCap;
    if (title.includes('Marketing')) return Megaphone;
    if (title.includes('Autoridade')) return TrendingUp;
    return Target;
  };

  // Mapear classe de cor por título do KPI
  const getColorClassForKpi = (title: string) => {
    if (title.includes('Youtube')) return 'bg-red-500/10';
    if (title.includes('Treinamentos')) return 'bg-emerald-500/10';
    if (title.includes('Marketing')) return 'bg-purple-500/10';
    if (title.includes('Visitantes')) return 'bg-amber-500/10';
    if (title.includes('Autoridade')) return 'bg-indigo-500/10';
    return 'bg-blue-500/10';
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <div className="w-16 h-16 border-4 border-[#1e5144] border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-400">Carregando dashboard...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 mb-2">Erro ao carregar dados</p>
          <p className="text-gray-500 text-sm">{error.message}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6 animate-in fade-in duration-500">
      {/* Header */}
      <div className="flex justify-between items-end">
        <h2 className="text-xl font-bold text-gray-200">Visão Geral - {selectedMonth === '---' ? 'Ano Completo' : selectedMonth}</h2>
      </div>
      
      {/* KPI Summary Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {mainKpis.map((kpi) => {
          const stats = getKpiStats(kpi);
          const Icon = getIconForKpi(kpi.title);
          const colorClass = getColorClassForKpi(kpi.title);
          
          return (
            <SummaryCard
              key={kpi.id}
              title={kpi.title}
              value={stats.realized}
              target={stats.meta}
              icon={Icon}
              colorClass={colorClass}
              onClick={() => navigate(`/kpis/${kpi.id}`)}
              unit={kpi.unit}
            />
          );
        })}
      </div>

      {/* Widgets Row - Empilhados verticalmente */}
      <div className="flex flex-col gap-6">
        <div className="h-[400px]">
          <TaskWidget
            tasks={tasksData?.data || []}
            onViewAll={() => navigate('/tasks')}
            onTaskClick={(taskId) => navigate('/tasks')}
          />
        </div>
        <div className="h-auto">
          <UpcomingPostsWidget
            posts={postsData || []}
            onViewCalendar={() => navigate('/calendar')}
          />
        </div>
      </div>
    </div>
  );
};

export default DashboardView;
