import { useQuery } from '@tanstack/react-query';
import { taskService } from '../services/taskService';
import type { Task } from '../types';

interface UsePaginatedTasksOptions {
  page?: number;
  pageSize?: number;
  enabled?: boolean;
}

interface PaginatedTasksResult {
  tasks: Task[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
  isLoading: boolean;
  error: Error | null;
}

export const usePaginatedTasks = ({
  page = 1,
  pageSize = 50,
  enabled = true,
}: UsePaginatedTasksOptions = {}): PaginatedTasksResult => {
  const { data, isLoading, error } = useQuery({
    queryKey: ['tasks', page, pageSize],
    queryFn: () => taskService.list(),
    enabled,
  });

  // Calculate pagination metadata
  const tasks = data || [];
  const total = tasks.length;
  const totalPages = Math.ceil(total / pageSize);

  return {
    tasks,
    total,
    page,
    pageSize,
    totalPages,
    isLoading,
    error: error as Error | null,
  };
};
