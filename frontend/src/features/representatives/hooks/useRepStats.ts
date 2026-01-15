import { useQuery } from '@tanstack/react-query';
import { representativeService } from '../services/representativeService';

export const useRepStats = (uuid: string) => {
  return useQuery({
    queryKey: ['representative-stats', uuid],
    queryFn: () => representativeService.getStats(uuid),
    enabled: !!uuid, // Só executar se uuid for fornecido
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};
