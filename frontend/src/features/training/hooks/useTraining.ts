import { useQuery } from '@tanstack/react-query';
import { kpiService } from '@/features/kpis/services/kpiService';
import { TRAINING_KPIS_SLUGS, QUERY_KEYS } from '@/shared/utils/constants';
import { useUIStore } from '@/shared/store/uiStore';

export const useTraining = () => {
  const { selectedMonth, selectedYear } = useUIStore();

  const { data, isLoading, error } = useQuery({
    queryKey: QUERY_KEYS.KPIS.BY_SLUGS([...TRAINING_KPIS_SLUGS, selectedMonth, selectedYear]),
    queryFn: () => kpiService.getBySlugs(TRAINING_KPIS_SLUGS, selectedMonth, selectedYear ? parseInt(selectedYear) : undefined),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });

  return {
    trainingKpis: data || [],
    isLoading,
    error
  };
};
