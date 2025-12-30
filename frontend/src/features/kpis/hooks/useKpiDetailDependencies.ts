import { useQuery } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS, KPI_DETAIL_DEPENDENCIES_SLUGS } from '@/shared/utils/constants';

/**
 * Hook para buscar apenas os KPIs necessários para a página de detalhes
 * Otimizado para buscar apenas: KPI principal + KPI "taxa_de_oportunidades" (para referência cruzada)
 */
export const useKpiDetailDependencies = () => {
  return useQuery({
    queryKey: QUERY_KEYS.KPIS.BY_SLUGS(KPI_DETAIL_DEPENDENCIES_SLUGS),
    queryFn: () => kpiService.getBySlugs(KPI_DETAIL_DEPENDENCIES_SLUGS),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });
};
