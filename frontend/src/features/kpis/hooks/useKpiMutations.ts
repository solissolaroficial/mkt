import { useMutation, useQueryClient } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS } from '@/shared/utils/constants';
import type { CreateKpiDTO, UpdateKpiDTO, UpdateMonthlyDataDTO } from '../types';
import { useToast } from '@/shared/hooks/useToast';

/**
 * Hook para criar KPI
 */
export const useCreateKpi = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

  return useMutation({
    mutationFn: (data: CreateKpiDTO) => kpiService.create(data),
    onSuccess: () => {
      // Invalida cache para refetch
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      toast.success('KPI criado com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao criar KPI:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao criar KPI';
      toast.error(errorMessage);
    },
  });
};

/**
 * Hook para atualizar KPI
 */
export const useUpdateKpi = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateKpiDTO }) =>
      kpiService.update(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.DETAIL(variables.id) });
      toast.success('KPI atualizado com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao atualizar KPI:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao atualizar KPI';
      toast.error(errorMessage);
    },
  });
};

/**
 * Hook para deletar KPI
 */
export const useDeleteKpi = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

  return useMutation({
    mutationFn: (id: string) => kpiService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.KPIS.ALL });
      toast.success('KPI excluído com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao deletar KPI:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao deletar KPI';
      toast.error(errorMessage);
    },
  });
};

/**
 * Hook para atualizar dados mensais
 */
export const useUpdateMonthlyData = () => {
  const queryClient = useQueryClient();
  const toast = useToast();

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
      toast.success('Dados mensais atualizados com sucesso');
    },
    onError: (error: Error) => {
      console.error('Falha ao atualizar dados mensais:', error);
      const errorMessage = error instanceof Error ? error.message : 'Falha ao atualizar dados mensais';
      toast.error(errorMessage);
    },
  });
};