import { useMutation, useQueryClient } from '@tanstack/react-query';
import { socialService } from '../services/socialService';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';
import type {
  SocialBenchmarking,
  CreateSocialBenchmarkingRequest,
  UpdateSocialBenchmarkingRequest,
} from '@/shared/types/legacy.types';

export const useSocialMutations = (setError?: any) => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    setError,
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  // Create mutation
  const createMutation = useMutation({
    mutationFn: (data: CreateSocialBenchmarkingRequest) => socialService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['social-benchmarkings'] });
      showSuccess('Benchmarking criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar benchmarking');
      }
    },
  });

  // Update mutation
  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateSocialBenchmarkingRequest }) =>
      socialService.update(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['social-benchmarkings'] });
      queryClient.invalidateQueries({ queryKey: ['social-benchmarking', variables.id] });
      showSuccess('Benchmarking atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar benchmarking');
      }
    },
  });

  // Delete mutation
  const deleteMutation = useMutation({
    mutationFn: (id: string) => socialService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['social-benchmarkings'] });
      showSuccess('Benchmarking deletado com sucesso!');
    },
    onError: () => {
      showError('Erro ao deletar benchmarking');
    },
  });

  return {
    create: (data: CreateSocialBenchmarkingRequest, options?: { onSuccess?: (result: SocialBenchmarking) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    update: (variables: { id: string; data: UpdateSocialBenchmarkingRequest }, options?: { onSuccess?: (result: SocialBenchmarking) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    delete: (id: string) => deleteMutation.mutate(id),
    // Loading states
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
