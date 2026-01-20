import { useMutation, useQueryClient } from '@tanstack/react-query';
import { brandService } from '../services/brandService';
import type { CreateBrandRequest } from '@/shared/types/legacy.types';

export const useBrandMutations = () => {
  const queryClient = useQueryClient();

  const createMutation = useMutation({
    mutationFn: (data: CreateBrandRequest) => brandService.create(data),
    onSuccess: () => {
      // Invalidate brands query to refetch the list
      queryClient.invalidateQueries({ queryKey: ['brands'] });
    },
    onError: (error) => {
      console.error('Erro ao criar marca:', error);
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) => brandService.delete(id),
    onSuccess: () => {
      // Invalidate brands query to refetch the list
      queryClient.invalidateQueries({ queryKey: ['brands'] });
    },
    onError: (error) => {
      console.error('Erro ao deletar marca:', error);
    },
  });

  return {
    createBrand: createMutation.mutate,
    deleteBrand: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
