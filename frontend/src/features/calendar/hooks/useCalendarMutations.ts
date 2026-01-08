import { useMutation, useQueryClient } from '@tanstack/react-query';
import { calendarService } from '../services/calendarService';
import { useBackendErrors } from '../../../shared/hooks/useBackendErrors';
import { useToast } from '../../../shared/hooks/useToast';
import type {
  CalendarPost,
  CalendarPostRequest,
  UpdateStatusRequest,
} from '../../../shared/types/legacy.types';

export const useCalendarMutations = (setError?: any) => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    setError,
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  // Create post mutation
  const createMutation = useMutation({
    mutationFn: (data: CalendarPostRequest) => calendarService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['calendar-posts'] });
      showSuccess('Post criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar post do calendário');
      }
    },
  });

  // Update post mutation
  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<CalendarPostRequest> }) =>
      calendarService.update(id, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['calendar-posts'] });
      queryClient.invalidateQueries({ queryKey: ['calendar-post', variables.id] });
      showSuccess('Post atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar post do calendário');
      }
    },
  });

  // Update status mutation (NOVO)
  const updateStatusMutation = useMutation({
    mutationFn: ({ id, request }: { id: string; request: UpdateStatusRequest }) =>
      calendarService.updateStatus(id, request),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['calendar-posts'] });
      queryClient.invalidateQueries({ queryKey: ['calendar-post', variables.id] });
      showSuccess('Status atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar status do post');
      }
    },
  });

  // Confirm publishing mutation (NOVO)
  const confirmPublishingMutation = useMutation({
    mutationFn: ({ id, platforms }: { id: string; platforms: string[] }) =>
      calendarService.confirmPublishing(id, platforms),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['calendar-posts'] });
      queryClient.invalidateQueries({ queryKey: ['calendar-post', variables.id] });
      showSuccess('Publicação confirmada com sucesso!');
    },
    onError: () => {
      showError('Erro ao confirmar publicação do post');
    },
  });

  // Delete post mutation
  const deleteMutation = useMutation({
    mutationFn: (id: string) => calendarService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['calendar-posts'] });
      showSuccess('Post deletado com sucesso!');
    },
    onError: () => {
      showError('Erro ao deletar post do calendário');
    },
  });

  return {
    createPost: (data: CalendarPostRequest, options?: { onSuccess?: (result: CalendarPost) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updatePost: (variables: { id: string; data: Partial<CalendarPostRequest> }, options?: { onSuccess?: (result: CalendarPost) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateStatus: (variables: { id: string; request: UpdateStatusRequest }, options?: { onSuccess?: (result: CalendarPost) => void; onError?: (error: any) => void }) =>
      updateStatusMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    confirmPublishing: (id: string, platforms: string[], options?: { onSuccess?: (result: CalendarPost) => void; onError?: (error: any) => void }) =>
      confirmPublishingMutation.mutate({ id, platforms }, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deletePost: (id: string) => deleteMutation.mutate(id),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isUpdatingStatus: updateStatusMutation.isPending,
    isConfirming: confirmPublishingMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
