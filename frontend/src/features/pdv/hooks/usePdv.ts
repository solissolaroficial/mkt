import { useQuery } from '@tanstack/react-query';
import { pdvService } from '../services/pdvService';
import type { PdvPostFilters, RecurrentPdvFilters } from '../../../shared/types/legacy.types';

// Hook para lista de posts PDV (com filtros)
export const usePdvPosts = (params?: PdvPostFilters) => {
  return useQuery({
    queryKey: ['pdv-posts', params],
    queryFn: () => pdvService.list(params),
    staleTime: 1000 * 60 * 5, // 5 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

// Hook para post PDV individual
export const usePdvPost = (id: string) => {
  return useQuery({
    queryKey: ['pdv-post', id],
    queryFn: () => pdvService.getById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 5,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

// Hook para lista de PDVs recorrentes (com filtros)
export const useRecurrentPdvs = (params?: RecurrentPdvFilters) => {
  return useQuery({
    queryKey: ['recurrent-pdvs', params],
    queryFn: () => pdvService.listRecurrent(params),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

// Hook para PDV recorrente individual
export const useRecurrentPdv = (id: string) => {
  return useQuery({
    queryKey: ['recurrent-pdv', id],
    queryFn: () => pdvService.getRecurrentById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};
