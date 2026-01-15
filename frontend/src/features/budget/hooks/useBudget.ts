import { useQuery } from '@tanstack/react-query';
import { budgetService } from '../services/budgetService';
import type { BudgetItem, BudgetItemFilters } from '../types';

export interface UseBudgetItemsOptions extends BudgetItemFilters {}

export const useBudgetItems = (filters?: UseBudgetItemsOptions) => {
  return useQuery({
    queryKey: ['budget-items', filters],
    queryFn: () => budgetService.list(filters),
    select: (response) => response.data, // Extrair array do objeto paginado
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

export const useBudgetItem = (uuid: string) => {
  return useQuery({
    queryKey: ['budget-item', uuid],
    queryFn: () => budgetService.getById(uuid),
    enabled: !!uuid, // Só executar se uuid for fornecido
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};

export const useBudgetSummary = (filters?: BudgetItemFilters) => {
  return useQuery({
    queryKey: ['budget-summary', filters],
    queryFn: () => budgetService.getSummary(filters),
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};

export const useBudgetYears = () => {
  return useQuery({
    queryKey: ['budget-years'],
    queryFn: () => budgetService.getDistinctYears(),
    staleTime: 1000 * 60 * 60, // 1 hora
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
  });
};
