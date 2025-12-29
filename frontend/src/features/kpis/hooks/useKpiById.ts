import { useQuery } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import type { KpiCategory } from '../types';

/**
 * Hook customizado para buscar um KPI específico pelo ID
 * @param kpiId - ID do KPI a ser buscado
 * @returns Objeto com data, isLoading, error e refetch do React Query
 */
export function useKpiById(kpiId: string) {
  return useQuery({
    queryKey: ['kpi', kpiId],
    queryFn: () => kpiService.getById(kpiId),
    enabled: !!kpiId, // Só busca se tiver ID válido
    staleTime: 5 * 60 * 1000, // 5 minutos
    retry: 1, // Tenta uma vez em caso de erro
  });
}
