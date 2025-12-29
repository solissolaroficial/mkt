import { useMutation, useQueryClient } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS } from '@/shared/utils/constants';
import type { CreateKpiDTO, UpdateKpiDTO, UpdateMonthlyDataDTO } from '../types';

/**
 * Hook para criar KPI
 */
export const useCreateKpi = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateKpiDTO) => kpiService.create(data),
    onSuccess: () => {
      // Invalida cache para refetch
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
    },
  });
};

/**
 * Hook para atualizar KPI
 */
export const useUpdateKpi = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateKpiDTO }) =>
      kpiService.update(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.DETAIL(variables.id) });
    },
  });
};

/**
 * Hook para deletar KPI
 */
export const useDeleteKpi = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => kpiService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
    },
  });
};

/**
 * Hook para atualizar dados mensais
 */
export const useUpdateMonthlyData = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      kpiId,
      data,
    }: {
      kpiId: string;
      data: UpdateMonthlyDataDTO;
    }) => kpiService.updateMonthlyData(kpiId, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.DETAIL(variables.kpiId) });
    },
  });
};