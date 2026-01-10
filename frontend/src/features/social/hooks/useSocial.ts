import { useQuery } from '@tanstack/react-query';
import { socialService } from '../services/socialService';
import type { SocialBenchmarkingFilters } from '@/shared/types/legacy.types';

/**
 * Hook para lista de social benchmarkings (com filtros)
 */
export const useSocialBenchmarkings = (params?: SocialBenchmarkingFilters) => {
  return useQuery({
    queryKey: ['social-benchmarkings', params],
    queryFn: () => socialService.list(params),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

/**
 * Hook para social benchmarking individual
 */
export const useSocialBenchmarking = (id: string) => {
  return useQuery({
    queryKey: ['social-benchmarking', id],
    queryFn: () => socialService.getById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};
