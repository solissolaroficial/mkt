import { useQuery } from '@tanstack/react-query';
import { representativeService } from '../services/representativeService';
import type { ListRepresentativesRequest, RepresentativeTableData } from '../types';

export interface UseRepTableDataOptions extends ListRepresentativesRequest {}

export const useRepTableData = (filters?: UseRepTableDataOptions) => {
  return useQuery<RepresentativeTableData>({
    queryKey: ['representatives-table', filters],
    queryFn: () => representativeService.getTableData(filters),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};
