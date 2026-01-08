import { useQuery } from '@tanstack/react-query';
import { calendarService } from '../services/calendarService';
import type { CalendarFilters } from '../../../shared/types/legacy.types';

// Hook para lista de posts (com filtros)
export const useCalendarPosts = (params?: CalendarFilters) => {
  return useQuery({
    queryKey: ['calendar-posts', params],
    queryFn: () => calendarService.list(params),
    staleTime: 1000 * 60 * 5, // 5 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

// Hook para post individual (NOVO)
export const useCalendarPost = (id: string) => {
  return useQuery({
    queryKey: ['calendar-post', id],
    queryFn: () => calendarService.getById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 5,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};
