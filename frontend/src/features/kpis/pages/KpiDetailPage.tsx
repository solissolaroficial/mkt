import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useKpiById } from '../hooks/useKpiById';
import { useKpiDetailDependencies } from '../hooks/useKpiDetailDependencies';
import { useUIStore } from '@/shared/store/uiStore';
import KpiDetailView from '../ui/KpiDetailView';

/**
 * Página wrapper para exibir detalhes de um KPI específico
 * Obtém o kpiId da URL e busca os dados do KPI
 * Otimizado para buscar apenas os KPIs necessários via slugs
 */
const KpiDetailPage: React.FC = () => {
  const { kpiId } = useParams<{ kpiId: string }>();
  const navigate = useNavigate();
  const { selectedMonth } = useUIStore();
  
  // Buscar KPI específico pelo ID (busca otimizada por ID)
  const { data: kpi, isLoading, error } = useKpiById(kpiId || '');
  
  // Buscar apenas os KPIs necessários para referência cruzada via slugs
  // Otimizado para buscar apenas: opportunities (necessário para Ad Spend cruzar com Opportunities)
  const { data: dependentKpis } = useKpiDetailDependencies();
  const allKpis = dependentKpis || [];

  // Loading state
  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <div className="w-16 h-16 border-4 border-[#1e5144] border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-400">Carregando KPI...</p>
        </div>
      </div>
    );
  }

  // Error state
  if (error || !kpi) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 mb-2">Erro ao carregar KPI</p>
          <p className="text-gray-500 text-sm mb-4">
            {error?.message || 'KPI não encontrado'}
          </p>
          <button 
            onClick={() => navigate('/kpis')}
            className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg transition-colors"
          >
            Voltar para Lista
          </button>
        </div>
      </div>
    );
  }

  // Renderizar KpiDetailView com os dados carregados
  return (
    <KpiDetailView
      kpi={kpi}
      allKpis={allKpis}
      selectedMonth={selectedMonth}
      onBack={() => navigate('/kpis')}
    />
  );
};

export default KpiDetailPage;
