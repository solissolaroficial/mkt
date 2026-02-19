import { useQueryWithError } from '@/shared/hooks/useQueryWithError';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS } from '@/shared/utils/constants';
import type { DailyEntry } from '../types';

/**
 * Hook customizado para buscar entradas diárias de um KPI
 * @param kpiId - ID do KPI
 * @param month - Mês (ex: 'Janeiro', 'Fevereiro')
 * @param year - Ano (ex: 2024, 2025)
 * @param enabled - Se a query deve ser executada
 * @returns Objeto com data, isLoading, error e refetch do React Query
 */
export function useDailyEntries(
  kpiId: string,
  month: string,
  year: number,
  enabled: boolean = false
) {
  return useQueryWithError<DailyEntry[]>({
    queryKey: QUERY_KEYS.KPIS.DAILY_ENTRIES(kpiId, month, year),
    queryFn: () => kpiService.dailyEntry.get(kpiId, month, year),
    enabled: enabled && !!kpiId && !!month && !!year,
    staleTime: 5 * 60 * 1000, // 5 minutos
    errorMessage: 'Falha ao carregar entradas diárias',
  });
}
