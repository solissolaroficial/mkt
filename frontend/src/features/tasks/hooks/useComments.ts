import { useQuery } from '@tanstack/react-query';
import { commentService } from '../services/taskService';
import type { Comment } from '../types';

export const useComments = (taskId: string) => {
  return useQuery({
    queryKey: ['comments', taskId],
    queryFn: async () => {
      return await commentService.list(taskId);
    },
    enabled: !!taskId,
    staleTime: 1000 * 60 * 5,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};

export const useComment = (id: string) => {
  return useQuery({
    queryKey: ['comments', id],
    queryFn: async () => {
      return await commentService.getById(id);
    },
    enabled: !!id,
    staleTime: 1000 * 60 * 5,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};
