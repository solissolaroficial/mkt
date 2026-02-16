import { useQuery } from '@tanstack/react-query';
import { flowService } from '../services/flowService';

export const useFlows = (page = 1, limit = 100) => {
  return useQuery({
    queryKey: ['flows', page, limit],
    queryFn: () => flowService.list(page, limit),
    staleTime: 1000 * 60 * 5,
  });
};

export const useFlow = (id: string) => {
  return useQuery({
    queryKey: ['flows', id],
    queryFn: () => flowService.getById(id),
    enabled: !!id,
  });
};
