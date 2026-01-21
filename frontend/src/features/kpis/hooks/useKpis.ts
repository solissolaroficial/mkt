import { useQuery } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { useUIStore } from '@/shared/store/uiStore';
import { QUERY_KEYS } from '@/shared/utils/constants';

/**
 * Hook para buscar lista de KPIs com filtros globais
 */
export const useKpis = () => {
  const { selectedMonth, selectedYear } = useUIStore();

  return useQuery({
    queryKey: QUERY_KEYS.KPIS.LIST({ month: selectedMonth, year: selectedYear }),
    queryFn: () => kpiService.list(selectedMonth, selectedYear ? parseInt(selectedYear) : undefined),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });
};

/**
 * Hook para buscar um KPI específico
 */
export const useKpi = (id: string) => {
  return useQuery({
    queryKey: QUERY_KEYS.KPIS.DETAIL(id),
    queryFn: () => kpiService.getById(id),
    enabled: !!id, // Só executa se tiver ID
  });
};