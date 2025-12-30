import { useQuery } from '@tanstack/react-query';
import { kpiService } from '@/features/kpis/services/kpiService';
import { TRAINING_KPIS_SLUGS, QUERY_KEYS } from '@/shared/utils/constants';

export const useTraining = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: QUERY_KEYS.KPIS.BY_SLUGS(TRAINING_KPIS_SLUGS),
    queryFn: () => kpiService.getBySlugs(TRAINING_KPIS_SLUGS),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });

  return {
    trainingKpis: data || [],
    isLoading,
    error
  };
};
