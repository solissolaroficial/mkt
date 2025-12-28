import { useQuery } from '@tanstack/react-query';
import { budgetService } from '../services/budgetService';
import type { BudgetItem } from '../types';
import { RAW_BUDGET_ITEMS } from '@/shared/utils/legacy.constants';

export const useBudgetItems = () => {
  return useQuery({
    queryKey: ['budget-items'],
    queryFn: async () => {
      try {
        return await budgetService.list();
      } catch (error) {
        console.warn('Using mock data for budget items:', error);
        return RAW_BUDGET_ITEMS as BudgetItem[];
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};
