import { useMutation, useQueryClient } from '@tanstack/react-query';
import { kpiService } from '../services/kpiService';
import { QUERY_KEYS } from '@/shared/utils/constants';

export interface UpdateMetaParams {
  kpiId: string;
  year: number;
  month: string;
  meta: number;
}

export function useUpdateMeta() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ kpiId, year, month, meta }: UpdateMetaParams) =>
      kpiService.updateMeta(kpiId, year, month, meta),
    onSuccess: (_, { kpiId }) => {
      // Invalidar cache do KPI para refletir a nova meta
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.KPIS.DETAIL(kpiId)
      });
      // Invalidar lista de KPIs para atualizar os cards de resumo
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.KPIS.LIST()
      });
    }
  });
}
