import { useMutation, useQueryClient } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS } from '@/shared/utils/constants';
import type { AddDailyEntryDTO, UpdateDailyEntryDTO, DeleteDailyEntryDTO } from '../types';
import { useToast } from '@/shared/hooks/useToast';

/**
 * Hook para adicionar entrada diária
 */
export const useAddDailyEntry = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

  return useMutation({
    mutationFn: ({
      kpiId,
      data,
    }: {
      kpiId: string;
      data: AddDailyEntryDTO;
    }) => kpiService.dailyEntry.add(kpiId, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.DETAIL(variables.kpiId) });
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.KPIS.DAILY_ENTRIES(variables.kpiId, variables.data.month, variables.data.year)
      });
      toast.success('Entrada diária adicionada com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao adicionar entrada diária:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao adicionar entrada diária';
      toast.error(errorMessage);
    },
  });
};

/**
 * Hook para atualizar entrada diária
 */
export const useUpdateDailyEntry = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

  return useMutation({
    mutationFn: ({
      kpiId,
      data,
    }: {
      kpiId: string;
      data: UpdateDailyEntryDTO;
    }) => kpiService.dailyEntry.update(kpiId, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.DETAIL(variables.kpiId) });
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.KPIS.DAILY_ENTRIES(variables.kpiId, variables.data.month, variables.data.year)
      });
      toast.success('Entrada diária atualizada com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao atualizar entrada diária:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao atualizar entrada diária';
      toast.error(errorMessage);
    },
  });
};

/**
 * Hook para deletar entrada diária
 */
export const useDeleteDailyEntry = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

  return useMutation({
    mutationFn: ({
      kpiId,
      data,
    }: {
      kpiId: string;
      data: DeleteDailyEntryDTO;
    }) => kpiService.dailyEntry.delete(kpiId, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.DETAIL(variables.kpiId) });
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.KPIS.DAILY_ENTRIES(variables.kpiId, variables.data.month, variables.data.year)
      });
      toast.success('Entrada diária removida com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao remover entrada diária:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao remover entrada diária';
      toast.error(errorMessage);
    },
  });
};
