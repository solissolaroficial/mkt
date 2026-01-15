import { useMutation, useQueryClient } from '@tanstack/react-query';
import { budgetService } from '../services/budgetService';
import type {
  BudgetItem,
  CreateBudgetItemRequest,
  UpdateBudgetItemRequest,
} from '../types';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';

export const useBudgetMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  // ==================== CREATE ====================

  const createMutation = useMutation({
    mutationFn: (data: CreateBudgetItemRequest) =>
      budgetService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['budget-items'] });
      queryClient.invalidateQueries({ queryKey: ['budget-summary'] });
      queryClient.invalidateQueries({ queryKey: ['budget-years'] });
      showSuccess('Item criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar item');
      }
    },
  });

  // ==================== UPDATE ====================

  const updateMutation = useMutation({
    mutationFn: ({ uuid, data }: { uuid: string; data: UpdateBudgetItemRequest }) =>
      budgetService.update(uuid, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['budget-items'] });
      queryClient.invalidateQueries({ queryKey: ['budget-item', variables.uuid] });
      queryClient.invalidateQueries({ queryKey: ['budget-summary'] });
      showSuccess('Item atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar item');
      }
    },
  });

  // ==================== DELETE ====================

  const deleteMutation = useMutation({
    mutationFn: (uuid: string) => budgetService.delete(uuid),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['budget-items'] });
      queryClient.invalidateQueries({ queryKey: ['budget-summary'] });
      showSuccess('Item excluído com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.status === 409) {
        showError('Item possui dependências e não pode ser excluído');
      } else if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao excluir item');
      }
    },
  });

  // ==================== BATCH CREATE ====================

  const batchCreateMutation = useMutation({
    mutationFn: (items: CreateBudgetItemRequest[]) =>
      budgetService.batchCreate(items),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['budget-items'] });
      queryClient.invalidateQueries({ queryKey: ['budget-summary'] });
      queryClient.invalidateQueries({ queryKey: ['budget-years'] });
      showSuccess('Itens criados com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar itens');
      }
    },
  });

  return {
    createItem: (data: CreateBudgetItemRequest, options?: { onSuccess?: (result: BudgetItem) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateItem: (variables: { uuid: string; data: UpdateBudgetItemRequest }, options?: { onSuccess?: (result: BudgetItem) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteItem: (uuid: string) => deleteMutation.mutate(uuid),
    batchCreateItems: (items: CreateBudgetItemRequest[], options?: { onSuccess?: (result: BudgetItem[]) => void; onError?: (error: any) => void }) =>
      batchCreateMutation.mutate(items, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
    isBatchCreating: batchCreateMutation.isPending,
  };
};
