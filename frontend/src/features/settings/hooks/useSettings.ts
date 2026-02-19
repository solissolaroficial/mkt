import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { settingsService } from '../services/settingsService';
import { INTERNAL_CONTACTS } from '@/shared/utils/legacy.constants';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';
import type { CreateCredentialRequest, UpdateCredentialRequest } from '../types';

/**
 * Hook to get program credentials
 */
export const useCredentials = () => {
  return useQuery({
    queryKey: ['settings', 'credentials'],
    queryFn: async () => {
      const response = await settingsService.getCredentials();
      return response.credentials || [];
    },
    staleTime: 1000 * 60 * 10, // 10 minutes
  });
};

/**
 * Hook to get internal contacts
 */
export const useContacts = () => {
  return useQuery({
    queryKey: ['settings', 'contacts'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return INTERNAL_CONTACTS;
    },
    staleTime: 1000 * 60 * 10,
  });
};

/**
 * Mutations for program credentials CRUD
 */
export const useCredentialMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  // ==================== CREATE ====================

  const createMutation = useMutation({
    mutationFn: (data: CreateCredentialRequest) =>
      settingsService.createCredential(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['settings', 'credentials'] });
      showSuccess('Credencial criada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar credencial');
      }
    },
  });

  // ==================== UPDATE ====================

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateCredentialRequest }) =>
      settingsService.updateCredential(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['settings', 'credentials'] });
      showSuccess('Credencial atualizada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar credencial');
      }
    },
  });

  // ==================== DELETE ====================

  const deleteMutation = useMutation({
    mutationFn: (id: string) => settingsService.deleteCredential(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['settings', 'credentials'] });
      showSuccess('Credencial excluída com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao excluir credencial');
      }
    },
  });

  return {
    createItem: (data: CreateCredentialRequest, options?: { onSuccess?: () => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: () => {
          options?.onSuccess?.();
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateItem: (variables: { id: string; data: UpdateCredentialRequest }, options?: { onSuccess?: () => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: () => {
          options?.onSuccess?.();
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
