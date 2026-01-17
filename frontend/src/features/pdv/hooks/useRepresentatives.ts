import { useQuery } from '@tanstack/react-query';
import { representativesService } from '@/features/representatives/services/representativesService';
import type { RepresentativeProfile } from '@/features/representatives/types';

export const useRepresentatives = () => {
  return useQuery({
    queryKey: ['representatives', 'all-profiles'],
    queryFn: () => representativesService.getAllProfiles(),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
};
