import { useQuery } from '@tanstack/react-query';
import { taskService } from '../services/taskService';
import type { Task } from '../types';
import { MOCK_TASKS } from '@/shared/utils/legacy.constants';

export const useTasks = () => {
  return useQuery({
    queryKey: ['tasks'],
    queryFn: async () => {
      try {
        return await taskService.list();
      } catch (error) {
        console.warn('Using mock data for tasks:', error);
        return MOCK_TASKS as Task[];
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};

export const useTask = (id: string) => {
  return useQuery({
    queryKey: ['tasks', id],
    queryFn: async () => {
      try {
        return await taskService.getById(id);
      } catch (error) {
        console.warn('Using mock data for task:', error);
        return MOCK_TASKS.find(t => t.id === id) as Task;
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};
