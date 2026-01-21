import { useQuery } from '@tanstack/react-query';
import { kpiService } from '@/features/kpis/services/kpiService';
import { DASHBOARD_KPIS_SLUGS, QUERY_KEYS } from '@/shared/utils/constants';
import { useUIStore } from '@/shared/store/uiStore';

/**
 * Hook para buscar os KPIs principais da Dashboard
 * Usa busca por slugs para otimizar a query
 */
export const useDashboardKpis = () => {
  const { selectedMonth, selectedYear } = useUIStore();

  return useQuery({
    queryKey: QUERY_KEYS.KPIS.BY_SLUGS([...DASHBOARD_KPIS_SLUGS, selectedMonth, selectedYear]),
    queryFn: () => kpiService.getBySlugs(DASHBOARD_KPIS_SLUGS, selectedMonth, selectedYear ? parseInt(selectedYear) : undefined),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });
};
