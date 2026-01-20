import { useMutation, useQueryClient } from '@tanstack/react-query';
import { socialService } from '../services/socialService';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';
import type {
  SocialBenchmarking,
  CreateSocialBenchmarkingRequest,
  UpdateSocialBenchmarkingRequest,
  SocialPost,
  CreateSocialPostRequest,
  UpdateSocialPostRequest,
  SocialDailyAggregation,
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

  // ============================================
  // Social Posts Mutations
  // ============================================

  // Create post mutation
  const createPostMutation = useMutation({
    mutationFn: (data: CreateSocialPostRequest) => socialService.createPost(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['social-posts'] });
      queryClient.invalidateQueries({ queryKey: ['social-daily-aggregations'] });
      showSuccess('Post criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar post');
      }
    },
  });

  // Update post mutation
  const updatePostMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateSocialPostRequest }) =>
      socialService.updatePost(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['social-posts'] });
      queryClient.invalidateQueries({ queryKey: ['social-post', variables.id] });
      queryClient.invalidateQueries({ queryKey: ['social-daily-aggregations'] });
      showSuccess('Post atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar post');
      }
    },
  });

  // Delete post mutation
  const deletePostMutation = useMutation({
    mutationFn: (id: string) => socialService.deletePost(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['social-posts'] });
      queryClient.invalidateQueries({ queryKey: ['social-daily-aggregations'] });
      showSuccess('Post deletado com sucesso!');
    },
    onError: () => {
      showError('Erro ao deletar post');
    },
  });

  // ============================================
  // Social Daily Aggregations Mutations
  // ============================================

  // Recalculate aggregations mutation
  const recalculateAggregationsMutation = useMutation({
    mutationFn: ({ brandID, date }: { brandID: string; date: string }) =>
      socialService.recalculateAggregations(brandID, date),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['social-daily-aggregations'] });
      queryClient.invalidateQueries({ queryKey: ['social-benchmarkings'] });
      showSuccess('Agregações recalculadas com sucesso!');
    },
    onError: () => {
      showError('Erro ao recalcular agregações');
    },
  });

  return {
    // Social Benchmarking
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

    // Social Posts
    createPost: (data: CreateSocialPostRequest, options?: { onSuccess?: (result: SocialPost) => void; onError?: (error: any) => void }) =>
      createPostMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updatePost: (variables: { id: string; data: UpdateSocialPostRequest }, options?: { onSuccess?: (result: SocialPost) => void; onError?: (error: any) => void }) =>
      updatePostMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deletePost: (id: string) => deletePostMutation.mutate(id),

    // Social Daily Aggregations
    recalculateAggregations: (brandID: string, date: string, options?: { onSuccess?: (result: SocialDailyAggregation) => void; onError?: (error: any) => void }) =>
      recalculateAggregationsMutation.mutate({ brandID, date }, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),

    // Loading states
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
    isCreatingPost: createPostMutation.isPending,
    isUpdatingPost: updatePostMutation.isPending,
    isDeletingPost: deletePostMutation.isPending,
    isRecalculatingAggregations: recalculateAggregationsMutation.isPending,
  };
};
