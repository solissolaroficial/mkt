import { useQueryWithError } from '@/shared/hooks/useQueryWithError';
import { representativesService } from '@/features/representatives/services/representativesService';
import type { RepresentativeProfile } from '@/features/representatives/types';

export const useRepresentatives = () => {
  return useQueryWithError({
    queryKey: ['representatives', 'all-profiles'],
    queryFn: () => representativesService.getAllProfiles(),
    staleTime: 5 * 60 * 1000, // 5 minutes
    errorMessage: 'Falha ao carregar representantes',
  });
};
