import { useMutation, useQueryClient } from '@tanstack/react-query';
import { cooperativeService } from '../services/cooperativeService';
import type { OfflineAction, ShowroomItem } from '../types';

export const useOfflineActionMutations = () => {
  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: (data: Partial<OfflineAction>) => cooperativeService.createOfflineAction(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['offline-actions'] });
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<OfflineAction> }) =>
      cooperativeService.updateOfflineAction(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['offline-actions'] });
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => cooperativeService.deleteOfflineAction(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['offline-actions'] });
    },
  });

  return {
    createOfflineAction: createMutation.mutate,
    updateOfflineAction: updateMutation.mutate,
    deleteOfflineAction: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};

export const useShowroomItemMutations = () => {
  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: (data: Partial<ShowroomItem>) => cooperativeService.createShowroomItem(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['showroom-items'] });
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<ShowroomItem> }) =>
      cooperativeService.updateShowroomItem(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['showroom-items'] });
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => cooperativeService.deleteShowroomItem(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['showroom-items'] });
    },
  });

  return {
    createShowroomItem: createMutation.mutate,
    updateShowroomItem: updateMutation.mutate,
    deleteShowroomItem: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isUpdating: updateMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
