import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { useUIStore } from '@/shared/store/uiStore';
import { QUERY_KEYS } from '@/shared/utils/constants';
import { useToast } from '@/shared/hooks/useToast';

/**
 * Hook para buscar lista de KPIs com filtros globais
 */
export const useKpis = () => {
  const { selectedMonth, selectedYear } = useUIStore();

  return useQuery({
    queryKey: QUERY_KEYS.KPIS.LIST({ month: selectedMonth, year: selectedYear }),
    queryFn: () => kpiService.list(selectedMonth, selectedYear ? parseInt(selectedYear) : undefined),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });
};

/**
 * Hook para buscar um KPI específico
 */
export const useKpi = (id: string) => {
  return useQuery({
    queryKey: QUERY_KEYS.KPIS.DETAIL(id),
    queryFn: () => kpiService.getById(id),
    enabled: !!id, // Só executa se tiver ID
  });
};

/**
 * Hook para deletar dados mensais de um KPI
 */
export const useDeleteMonthlyData = () => {
  const toast = useToast();
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: async ({ kpiId, monthlyDataId }: { kpiId: string; monthlyDataId: string }) => {
      // Delete monthly data via API
      await kpiService.deleteMonthlyData(kpiId, monthlyDataId);
    },
    onSuccess: (_, variables) => {
      // Invalidate queries to refetch data after successful deletion
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.KPIS.DETAIL(variables.kpiId),
        refetchType: 'all',
      });
      toast.success('Dados mensais removidos com sucesso!');
    },
    onError: (error: Error) => {
      toast.error('Erro ao remover dados mensais: ' + error.message);
    },
  });

  return {
    deleteMonthlyData: mutation.mutate,
    isPending: mutation.isPending,
    isError: mutation.isError,
    error: mutation.error,
  };
};
