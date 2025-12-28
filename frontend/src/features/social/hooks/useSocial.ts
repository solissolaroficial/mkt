import { useQuery } from '@tanstack/react-query';
import { socialService } from '../services/socialService';
import { SOCIAL_MEDIA_RANKING } from '@/shared/utils/legacy.constants';

/**
 * Hook to get social benchmarking data
 */
export const useSocialBenchmarking = () => {
  return useQuery({
    queryKey: ['social', 'benchmarking'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return SOCIAL_MEDIA_RANKING;
    },
    staleTime: 1000 * 60 * 10, // 10 minutes
  });
};
