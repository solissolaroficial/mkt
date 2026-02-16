import { useMutation, useQueryClient } from '@tanstack/react-query';
import { flowService } from '../services/flowService';
import type { CreateFlowRequest, UpdateFlowRequest } from '@/shared/types';

export const useFlowMutations = () => {
  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: (data: CreateFlowRequest) => {
      console.log('useFlowMutations: Creating flow with data:', data);
      return flowService.create(data);
    },
    onSuccess: (response) => {
      console.log('useFlowMutations: Flow created successfully:', response);
      queryClient.invalidateQueries({ queryKey: ['flows'] });
    },
    onError: (error) => {
      console.error('useFlowMutations: Failed to create flow:', error);
      // TODO: Adicionar toast notification
    }
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateFlowRequest }) => {
      console.log('useFlowMutations: Updating flow with id:', id, 'data:', data);
      return flowService.update(id, data);
    },
    onSuccess: (response) => {
      console.log('useFlowMutations: Flow updated successfully:', response);
      queryClient.invalidateQueries({ queryKey: ['flows'] });
    },
    onError: (error) => {
      console.error('useFlowMutations: Failed to update flow:', error);
      // TODO: Adicionar toast notification
    }
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => {
      console.log('useFlowMutations: Deleting flow with id:', id);
      return flowService.delete(id);
    },
    onSuccess: () => {
      console.log('useFlowMutations: Flow deleted successfully');
      queryClient.invalidateQueries({ queryKey: ['flows'] });
    },
    onError: (error) => {
      console.error('useFlowMutations: Failed to delete flow:', error);
      // TODO: Adicionar toast notification
    }
  });

  const reorderMutation = useMutation({
    mutationFn: (flowIds: string[]) => {
      console.log('useFlowMutations: Reordering flows with IDs:', flowIds);
      return flowService.reorder(flowIds);
    },
    onSuccess: () => {
      console.log('useFlowMutations: Flows reordered successfully');
      queryClient.invalidateQueries({ queryKey: ['flows'] });
    },
    onError: (error) => {
      console.error('useFlowMutations: Failed to reorder flows:', error);
      // TODO: Adicionar toast notification
    }
  });

  return {
    createFlow: createMutation.mutate,
    updateFlow: updateMutation.mutate,
    deleteFlow: deleteMutation.mutate,
    reorderFlows: reorderMutation.mutate,
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
