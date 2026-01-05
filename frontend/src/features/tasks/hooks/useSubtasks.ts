import { useQuery } from '@tanstack/react-query';
import { subtaskService } from '../services/taskService';
import type { Subtask } from '../types';

export const useSubtasks = (taskId: string) => {
  return useQuery({
    queryKey: ['subtasks', taskId],
    queryFn: async () => {
      return await subtaskService.list(taskId);
    },
    enabled: !!taskId,
    staleTime: 1000 * 60 * 5,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};

export const useSubtask = (id: string) => {
  return useQuery({
    queryKey: ['subtasks', id],
    queryFn: async () => {
      return await subtaskService.getById(id);
    },
    enabled: !!id,
    staleTime: 1000 * 60 * 5,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};
