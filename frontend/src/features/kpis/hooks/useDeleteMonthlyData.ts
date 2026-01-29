import { useMutation, useQueryClient } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { toast } from 'sonner';
import { QUERY_KEYS } from '@/shared/utils/constants';
import type { KpiCategory } from '@/shared/types';

/**
 * Hook para deletar dados mensais de um KPI
 */
export const useDeleteMonthlyData = () => {
	const queryClient = useQueryClient();
	const { selectedMonth, selectedYear } = useUIStore();

	const mutation = useMutation({
		mutationFn: async ({ kpiId, monthlyDataId }: { kpiId: string; monthlyDataId: string }) => {
			// Delete monthly data via API
			await kpiService.deleteMonthlyData(kpiId, monthlyDataId);

			// Invalidate queries to refetch data
			queryClient.invalidateQueries({
				queryKey: [QUERY_KEYS.KPIS.DETAIL(kpiId)],
			refetchType: 'all',
			});
		},
		onSuccess: () => {
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
