import { useQuery } from '@tanstack/react-query';
import { representativesService } from '../services/representativesService';
import { REPS_TRAINING_DATA, REPS_MARKETING_DATA, REP_PROFILES } from '@/shared/utils/legacy.constants';

/**
 * Hook to get representatives training data
 */
export const useRepTrainingData = () => {
  return useQuery({
    queryKey: ['representatives', 'training'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return REPS_TRAINING_DATA;
    },
    staleTime: 1000 * 60 * 5,
  });
};

/**
 * Hook to get representatives marketing data
 */
export const useRepMarketingData = () => {
  return useQuery({
    queryKey: ['representatives', 'marketing'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return REPS_MARKETING_DATA;
    },
    staleTime: 1000 * 60 * 5,
  });
};

/**
 * Hook to get representative profile
 */
export const useRepProfile = (repName: string) => {
  return useQuery({
    queryKey: ['representatives', 'profile', repName],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      const profile = REP_PROFILES[repName];
      if (!profile) {
        throw new Error(`Representative profile not found: ${repName}`);
      }
      return profile;
    },
    staleTime: 1000 * 60 * 10,
    enabled: !!repName,
  });
};

/**
 * Hook to get all representative profiles
 */
export const useAllRepProfiles = () => {
  return useQuery({
    queryKey: ['representatives', 'profiles'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return Object.values(REP_PROFILES);
    },
    staleTime: 1000 * 60 * 10,
  });
};
