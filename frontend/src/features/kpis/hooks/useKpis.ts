import { useQuery } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS } from '@/shared/utils/constants';

/**
 * Hook para buscar lista de KPIs
 */
export const useKpis = () => {
  return useQuery({
    queryKey: QUERY_KEYS.KPIS.LIST(),
    queryFn: () => kpiService.list(),
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