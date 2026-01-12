import { useQuery } from '@tanstack/react-query';
import { cooperativeService } from '../services/cooperativeService';
import type { ShowroomItem, OfflineAction, RepMarketingAction, PaginatedResponse } from '../types';
import type { OfflineActionFilters, ShowroomItemFilters, RepMarketingActionFilters } from '../types';

export const useOfflineActions = (filters?: OfflineActionFilters) => {
  return useQuery({
    queryKey: ['offline-actions', filters],
    queryFn: () => cooperativeService.getOfflineActions({
      ...filters,
      page: filters?.page || 1,
      limit: filters?.limit || 50,
    }),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
    select: (data: PaginatedResponse<OfflineAction>) => data.data, // Extrair array de dados da resposta paginada
  });
};

export const useShowroomItems = (filters?: ShowroomItemFilters) => {
  return useQuery({
    queryKey: ['showroom-items', filters],
    queryFn: () => cooperativeService.getShowroomItems({
      ...filters,
      page: filters?.page || 1,
      limit: filters?.limit || 50,
    }),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
    select: (data: PaginatedResponse<ShowroomItem>) => data.data, // Extrair array de dados da resposta paginada
  });
};

export const useRepMarketingActions = (filters?: RepMarketingActionFilters) => {
  return useQuery({
    queryKey: ['rep-marketing-actions', filters],
    queryFn: () => cooperativeService.getRepMarketingActions({
      ...filters,
      page: filters?.page || 1,
      limit: filters?.limit || 50,
    }),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
    select: (data: PaginatedResponse<RepMarketingAction>) => data.data, // Extrair array de dados da resposta paginada
  });
};
