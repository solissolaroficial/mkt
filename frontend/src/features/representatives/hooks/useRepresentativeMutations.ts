import { useMutation, useQueryClient } from '@tanstack/react-query';
import { representativeService } from '../services/representativeService';
import type {
  Representative,
  CreateRepresentativeRequest,
  UpdateRepresentativeRequest,
} from '../types';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';

export const useRepresentativeMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  // ==================== CREATE ====================

  const createMutation = useMutation({
    mutationFn: (data: CreateRepresentativeRequest) =>
      representativeService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['representatives'] });
      showSuccess('Representante criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar representante');
      }
    },
  });

  // ==================== UPDATE ====================

  const updateMutation = useMutation({
    mutationFn: ({ uuid, data }: { uuid: string; data: UpdateRepresentativeRequest }) =>
      representativeService.update(uuid, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['representatives'] });
      queryClient.invalidateQueries({ queryKey: ['representative', variables.uuid] });
      queryClient.invalidateQueries({ queryKey: ['representative-stats', variables.uuid] });
      showSuccess('Representante atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar representante');
      }
    },
  });

  // ==================== DELETE ====================

  const deleteMutation = useMutation({
    mutationFn: (uuid: string) => representativeService.delete(uuid),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['representatives'] });
      showSuccess('Representante excluído com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.status === 409) {
        showError('Representante possui dependências e não pode ser excluído');
      } else if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao excluir representante');
      }
    },
  });

  return {
    createItem: (data: CreateRepresentativeRequest, options?: { onSuccess?: (result: Representative) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateItem: (variables: { uuid: string; data: UpdateRepresentativeRequest }, options?: { onSuccess?: (result: Representative) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteItem: (uuid: string) => deleteMutation.mutate(uuid),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
