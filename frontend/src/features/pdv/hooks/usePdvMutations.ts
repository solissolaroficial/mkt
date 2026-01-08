import { useMutation, useQueryClient } from '@tanstack/react-query';
import { pdvService } from '../services/pdvService';
import { useBackendErrors } from '../../../shared/hooks/useBackendErrors';
import { useToast } from '../../../shared/hooks/useToast';
import type {
  PdvPost,
  RecurrentPdv,
  CreatePdvPostRequest,
  UpdatePdvPostRequest,
  CreateRecurrentPdvRequest,
  UpdateRecurrentPdvRequest,
} from '../../../shared/types/legacy.types';

export const usePdvMutations = (setError?: any) => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    setError,
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  // ============================================
  // Posts PDV Mutations
  // ============================================

  // Create post mutation
  const createPostMutation = useMutation({
    mutationFn: (data: CreatePdvPostRequest) => pdvService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['pdv-posts'] });
      showSuccess('Post PDV criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar post PDV');
      }
    },
  });

  // Update post mutation
  const updatePostMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdatePdvPostRequest }) =>
      pdvService.update(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['pdv-posts'] });
      queryClient.invalidateQueries({ queryKey: ['pdv-post', variables.id] });
      showSuccess('Post PDV atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar post PDV');
      }
    },
  });

  // Update status mutation
  const updateStatusMutation = useMutation({
    mutationFn: ({ id, status }: { id: string; status: 'pending' | 'approved' | 'rejected' | 'published' | 'cancelled' }) =>
      pdvService.updateStatus(id, status),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['pdv-posts'] });
      queryClient.invalidateQueries({ queryKey: ['pdv-post', variables.id] });
      showSuccess('Status do post PDV atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar status do post PDV');
      }
    },
  });

  // Delete post mutation
  const deletePostMutation = useMutation({
    mutationFn: (id: string) => pdvService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['pdv-posts'] });
      showSuccess('Post PDV deletado com sucesso!');
    },
    onError: () => {
      showError('Erro ao deletar post PDV');
    },
  });

  // ============================================
  // PDVs Recorrentes Mutations
  // ============================================

  // Create recurrent PDV mutation
  const createRecurrentMutation = useMutation({
    mutationFn: (data: CreateRecurrentPdvRequest) => pdvService.createRecurrent(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['recurrent-pdvs'] });
      showSuccess('PDV recorrente criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar PDV recorrente');
      }
    },
  });

  // Update recurrent PDV mutation
  const updateRecurrentMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateRecurrentPdvRequest }) =>
      pdvService.updateRecurrent(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['recurrent-pdvs'] });
      queryClient.invalidateQueries({ queryKey: ['recurrent-pdv', variables.id] });
      showSuccess('PDV recorrente atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar PDV recorrente');
      }
    },
  });

  // Delete recurrent PDV mutation
  const deleteRecurrentMutation = useMutation({
    mutationFn: (id: string) => pdvService.deleteRecurrent(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['recurrent-pdvs'] });
      showSuccess('PDV recorrente deletado com sucesso!');
    },
    onError: () => {
      showError('Erro ao deletar PDV recorrente');
    },
  });

  return {
    // Posts PDV
    createPost: (data: CreatePdvPostRequest, options?: { onSuccess?: (result: PdvPost) => void; onError?: (error: any) => void }) =>
      createPostMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updatePost: (variables: { id: string; data: UpdatePdvPostRequest }, options?: { onSuccess?: (result: PdvPost) => void; onError?: (error: any) => void }) =>
      updatePostMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateStatus: (variables: { id: string; status: 'pending' | 'approved' | 'rejected' | 'published' | 'cancelled' }, options?: { onSuccess?: (result: PdvPost) => void; onError?: (error: any) => void }) =>
      updateStatusMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deletePost: (id: string) => deletePostMutation.mutate(id),
    // PDVs Recorrentes
    createRecurrent: (data: CreateRecurrentPdvRequest, options?: { onSuccess?: (result: RecurrentPdv) => void; onError?: (error: any) => void }) =>
      createRecurrentMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateRecurrent: (variables: { id: string; data: UpdateRecurrentPdvRequest }, options?: { onSuccess?: (result: RecurrentPdv) => void; onError?: (error: any) => void }) =>
      updateRecurrentMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteRecurrent: (id: string) => deleteRecurrentMutation.mutate(id),
    // Loading states
    isCreatingPost: createPostMutation.isPending,
    isUpdatingPost: updatePostMutation.isPending,
    isUpdatingStatus: updateStatusMutation.isPending,
    isDeletingPost: deletePostMutation.isPending,
    isCreatingRecurrent: createRecurrentMutation.isPending,
    isUpdatingRecurrent: updateRecurrentMutation.isPending,
    isDeletingRecurrent: deleteRecurrentMutation.isPending,
  };
};
