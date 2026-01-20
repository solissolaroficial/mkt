import { useQuery } from '@tanstack/react-query';
import { brandService } from '../services/brandService';
import type { Brand } from '@/shared/types/legacy.types';

/**
 * Hook para lista de brands
 */
export const useBrands = () => {
  return useQuery({
    queryKey: ['brands'],
    queryFn: () => brandService.list(),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};
