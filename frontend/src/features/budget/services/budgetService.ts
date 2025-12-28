import { apiClient } from '@/infrastructure/api/client';
import type { BudgetItem } from '../types';

export const budgetService = {
  // Get all budget items
  list: async (): Promise<BudgetItem[]> => {
    const response = await apiClient.get('/api/budget/items');
    return response.data;
  },

  // Update budget item
  update: async (id: string, data: Partial<BudgetItem>): Promise<BudgetItem> => {
    const response = await apiClient.put(`/api/budget/items/${id}`, data);
    return response.data;
  },

  // Export budget data
  export: async (format: 'csv' | 'excel'): Promise<Blob> => {
    const response = await apiClient.get(`/api/budget/export?format=${format}`, {
      responseType: 'blob'
    });
    return response.data;
  },
};
