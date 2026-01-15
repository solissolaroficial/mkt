import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { monthlyGoalService } from '../services/monthlyGoalService';
import type {
  RepresentativeMonthlyGoal,
  CreateRepresentativeMonthlyGoalRequest,
  UpdateRepresentativeMonthlyGoalRequest,
  ListRepresentativeMonthlyGoalsRequest,
  ListRepresentativeMonthlyGoalsResponse,
  GetRepresentativeGoalsTableDataRequest,
  GetRepresentativeGoalsTableDataResponse,
} from '../types/monthlyGoal';

/**
 * Hook for getting a single monthly goal by ID
 */
export function useMonthlyGoal(id: string) {
  return useQuery({
    queryKey: ['monthlyGoal', id],
    queryFn: () => monthlyGoalService.getById(id),
    enabled: !!id,
  });
}

/**
 * Hook for listing monthly goals with filters and pagination
 */
export function useMonthlyGoals(params?: ListRepresentativeMonthlyGoalsRequest) {
  return useQuery({
    queryKey: ['monthlyGoals', params],
    queryFn: () => monthlyGoalService.list(params),
  });
}

/**
 * Hook for getting table data (transposed view)
 */
export function useMonthlyGoalsTableData(params?: GetRepresentativeGoalsTableDataRequest) {
  return useQuery({
    queryKey: ['monthlyGoalsTable', params],
    queryFn: () => monthlyGoalService.getTableData(params),
  });
}

/**
 * Hook for creating a monthly goal
 */
export function useCreateMonthlyGoal() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateRepresentativeMonthlyGoalRequest) =>
      monthlyGoalService.create(data),
    onSuccess: () => {
      // Invalidate monthly goals queries
      queryClient.invalidateQueries({ queryKey: ['monthlyGoals'] });
      queryClient.invalidateQueries({ queryKey: ['monthlyGoalsTable'] });
    },
  });
}

/**
 * Hook for updating a monthly goal
 */
export function useUpdateMonthlyGoal() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateRepresentativeMonthlyGoalRequest }) =>
      monthlyGoalService.update(id, data),
    onSuccess: (_, variables) => {
      // Invalidate specific goal query
      queryClient.invalidateQueries({ queryKey: ['monthlyGoal', variables.id] });
      // Invalidate list queries
      queryClient.invalidateQueries({ queryKey: ['monthlyGoals'] });
      queryClient.invalidateQueries({ queryKey: ['monthlyGoalsTable'] });
    },
  });
}

/**
 * Hook for deleting a monthly goal
 */
export function useDeleteMonthlyGoal() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => monthlyGoalService.delete(id),
    onSuccess: () => {
      // Invalidate monthly goals queries
      queryClient.invalidateQueries({ queryKey: ['monthlyGoals'] });
      queryClient.invalidateQueries({ queryKey: ['monthlyGoalsTable'] });
    },
  });
}
