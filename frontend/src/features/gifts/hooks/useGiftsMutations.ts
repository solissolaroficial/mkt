import { useMutation, useQueryClient } from '@tanstack/react-query';
import { giftsService } from '../services/giftsService';
import type {
  GiftItem,
  GiftTransaction,
  CreateGiftItemRequest,
  UpdateGiftItemRequest,
  CreateGiftTransactionRequest,
  UpdateGiftTransactionRequest,
} from '../types';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';

// ==================== GIFT ITEMS ====================

export const useGiftItemMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  const createMutation = useMutation({
    mutationFn: (data: CreateGiftItemRequest) =>
      giftsService.createItem(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] });
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

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateGiftItemRequest }) =>
      giftsService.updateItem(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] });
      queryClient.invalidateQueries({ queryKey: ['gifts', 'item', variables.id] });
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

  const deleteMutation = useMutation({
    mutationFn: (id: string) => giftsService.deleteItem(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] });
      showSuccess('Item excluído com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.status === 409) {
        showError('Item possui transações associadas e não pode ser excluído');
      } else {
        showError('Erro ao excluir item');
      }
    },
  });

  return {
    createItem: (data: CreateGiftItemRequest, options?: { onSuccess?: (result: GiftItem) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateItem: (variables: { id: string; data: UpdateGiftItemRequest }, options?: { onSuccess?: (result: GiftItem) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteItem: (id: string) => deleteMutation.mutate(id),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};

// ==================== GIFT TRANSACTIONS ====================

export const useGiftTransactionMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  const createMutation = useMutation({
    mutationFn: (data: CreateGiftTransactionRequest) =>
      giftsService.createTransaction(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'transactions'] });
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] }); // Invalidate items because stock changes
      showSuccess('Transação criada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar transação');
      }
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateGiftTransactionRequest }) =>
      giftsService.updateTransaction(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'transactions'] });
      queryClient.invalidateQueries({ queryKey: ['gifts', 'transaction', variables.id] });
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] });
      showSuccess('Transação atualizada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar transação');
      }
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => giftsService.deleteTransaction(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'transactions'] });
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] });
      showSuccess('Transação excluída com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao excluir transação');
      }
    },
  });

  return {
    createTransaction: (data: CreateGiftTransactionRequest, options?: { onSuccess?: (result: GiftTransaction) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateTransaction: (variables: { id: string; data: UpdateGiftTransactionRequest }, options?: { onSuccess?: (result: GiftTransaction) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteTransaction: (id: string) => deleteMutation.mutate(id),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
