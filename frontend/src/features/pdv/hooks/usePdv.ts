import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { pdvService } from '../services/pdvService';
import { MOCK_PDV_POSTS, RECURRENT_PDVS, REP_NAMES } from '@/shared/utils/legacy.constants';
import type { PdvPost, RecurrentPdv } from '@/shared/types';

/**
 * Hook to get all PDV posts
 */
export const usePdvPosts = () => {
  return useQuery({
    queryKey: ['pdv', 'posts'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return MOCK_PDV_POSTS;
    },
    staleTime: 1000 * 60 * 5,
  });
};

/**
 * Hook to get all recurrent PDVs
 */
export const useRecurrentPdvs = () => {
  return useQuery({
    queryKey: ['pdv', 'recurrent'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return RECURRENT_PDVS;
    },
    staleTime: 1000 * 60 * 10,
  });
};

/**
 * Hook to create a new PDV post
 */
export const useCreatePdvPost = () => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: pdvService.createPost,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['pdv', 'posts'] });
    },
  });
};

/**
 * Hook to create a new recurrent PDV
 */
export const useCreateRecurrentPdv = () => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: pdvService.createRecurrentPdv,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['pdv', 'recurrent'] });
    },
  });
};

/**
 * Hook to update PDV post status
 */
export const useUpdatePdvPostStatus = () => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ id, status }: { id: string; status: 'verified' | 'pending' }) => 
      pdvService.updatePostStatus(id, status),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['pdv', 'posts'] });
    },
  });
};
