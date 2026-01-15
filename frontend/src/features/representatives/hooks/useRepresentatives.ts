import { useQuery } from '@tanstack/react-query';
import { representativeService } from '../services/representativeService';
import type { Representative, RepresentativeStats, ListRepresentativesRequest, ListRepresentativesResponse } from '../types';

export interface UseRepresentativesOptions extends ListRepresentativesRequest {}

export const useRepresentatives = (filters?: UseRepresentativesOptions) => {
  return useQuery<ListRepresentativesResponse>({
    queryKey: ['representatives', filters],
    queryFn: () => representativeService.list(filters),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

export const useRepresentative = (uuid: string) => {
  return useQuery({
    queryKey: ['representative', uuid],
    queryFn: () => representativeService.getById(uuid),
    enabled: !!uuid, // Só executar se uuid for fornecido
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};

export const useRepresentativeStats = (uuid: string) => {
  return useQuery({
    queryKey: ['representative-stats', uuid],
    queryFn: () => representativeService.getStats(uuid),
    enabled: !!uuid, // Só executar se uuid for fornecido
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};
