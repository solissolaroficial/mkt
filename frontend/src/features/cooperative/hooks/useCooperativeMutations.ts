import { useMutation, useQueryClient } from '@tanstack/react-query';
import { cooperativeService } from '../services/cooperativeService';
import type { OfflineAction, ShowroomItem, RepMarketingAction } from '../types';
import type {
  CreateOfflineActionRequest,
  UpdateOfflineActionRequest,
  CreateShowroomItemRequest,
  UpdateShowroomItemRequest,
  CreateRepMarketingActionRequest,
  UpdateRepMarketingActionRequest
} from '../types';
import { useBackendErrors } from '@/shared/hooks/useBackendErrors';
import { useToast } from '@/shared/hooks/useToast';

export const useOfflineActionMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  const createMutation = useMutation({
    mutationFn: (data: CreateOfflineActionRequest) =>
      cooperativeService.createOfflineAction(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['offline-actions'] });
      showSuccess('Ação offline criada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar ação offline');
      }
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ uuid, data }: { uuid: string; data: UpdateOfflineActionRequest }) =>
      cooperativeService.updateOfflineAction(uuid, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['offline-actions'] });
      queryClient.invalidateQueries({ queryKey: ['offline-action', variables.uuid] });
      showSuccess('Ação offline atualizada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar ação offline');
      }
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (uuid: string) => cooperativeService.deleteOfflineAction(uuid),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['offline-actions'] });
      showSuccess('Ação offline excluída com sucesso!');
    },
    onError: () => {
      showError('Erro ao excluir ação offline');
    },
  });

  return {
    createOfflineAction: (data: CreateOfflineActionRequest, options?: { onSuccess?: (result: OfflineAction) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateOfflineAction: (variables: { uuid: string; data: UpdateOfflineActionRequest }, options?: { onSuccess?: (result: OfflineAction) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteOfflineAction: (uuid: string) => deleteMutation.mutate(uuid),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};

export const useShowroomItemMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  const createMutation = useMutation({
    mutationFn: (data: CreateShowroomItemRequest) =>
      cooperativeService.createShowroomItem(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['showroom-items'] });
      showSuccess('Item de showroom criado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar item de showroom');
      }
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ uuid, data }: { uuid: string; data: UpdateShowroomItemRequest }) =>
      cooperativeService.updateShowroomItem(uuid, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['showroom-items'] });
      queryClient.invalidateQueries({ queryKey: ['showroom-item', variables.uuid] });
      showSuccess('Item de showroom atualizado com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar item de showroom');
      }
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (uuid: string) => cooperativeService.deleteShowroomItem(uuid),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['showroom-items'] });
      showSuccess('Item de showroom excluído com sucesso!');
    },
    onError: () => {
      showError('Erro ao excluir item de showroom');
    },
  });

  return {
    createShowroomItem: (data: CreateShowroomItemRequest, options?: { onSuccess?: (result: ShowroomItem) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateShowroomItem: (variables: { uuid: string; data: UpdateShowroomItemRequest }, options?: { onSuccess?: (result: ShowroomItem) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteShowroomItem: (uuid: string) => deleteMutation.mutate(uuid),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};

export const useRepMarketingActionMutations = () => {
  const queryClient = useQueryClient();
  const { handleBackendErrors } = useBackendErrors({
    onGlobalError: (msg) => useToast().error(msg),
  });
  const { success: showSuccess, error: showError } = useToast();

  const createMutation = useMutation({
    mutationFn: (data: CreateRepMarketingActionRequest) =>
      cooperativeService.createRepMarketingAction(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['rep-marketing-actions'] });
      showSuccess('Ação de marketing criada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao criar ação de marketing');
      }
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ uuid, data }: { uuid: string; data: UpdateRepMarketingActionRequest }) =>
      cooperativeService.updateRepMarketingAction(uuid, data),
    onSuccess: (result, variables) => {
      queryClient.invalidateQueries({ queryKey: ['rep-marketing-actions'] });
      queryClient.invalidateQueries({ queryKey: ['rep-marketing-action', variables.uuid] });
      showSuccess('Ação de marketing atualizada com sucesso!');
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data);
      } else {
        showError('Erro ao atualizar ação de marketing');
      }
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (uuid: string) => cooperativeService.deleteRepMarketingAction(uuid),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['rep-marketing-actions'] });
      showSuccess('Ação de marketing excluída com sucesso!');
    },
    onError: () => {
      showError('Erro ao excluir ação de marketing');
    },
  });

  return {
    createRepMarketingAction: (data: CreateRepMarketingActionRequest, options?: { onSuccess?: (result: RepMarketingAction) => void; onError?: (error: any) => void }) =>
      createMutation.mutate(data, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    updateRepMarketingAction: (variables: { uuid: string; data: UpdateRepMarketingActionRequest }, options?: { onSuccess?: (result: RepMarketingAction) => void; onError?: (error: any) => void }) =>
      updateMutation.mutate(variables, {
        onSuccess: (result) => {
          options?.onSuccess?.(result);
        },
        onError: (error) => {
          options?.onError?.(error);
        },
      }),
    deleteRepMarketingAction: (uuid: string) => deleteMutation.mutate(uuid),
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
