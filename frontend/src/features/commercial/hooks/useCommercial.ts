import { useQuery } from '@tanstack/react-query';
import { kpiService } from '@/features/kpis/services/kpiService';
import { COMMERCIAL_KPIS_SLUGS, QUERY_KEYS } from '@/shared/utils/constants';

export const useCommercial = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: QUERY_KEYS.KPIS.BY_SLUGS(COMMERCIAL_KPIS_SLUGS),
    queryFn: () => kpiService.getBySlugs(COMMERCIAL_KPIS_SLUGS),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });

  return {
    commercialKpis: data || [],
    isLoading,
    error
  };
};
