import { apiClient } from '@/infrastructure/api/client';
import type { AccountPayable } from '../types';

export const financialService = {
  // Get all accounts payable
  list: async (): Promise<AccountPayable[]> => {
    const response = await apiClient.get('/api/financial/accounts-payable');
    return response.data;
  },

  // Create new account payable
  create: async (data: Partial<AccountPayable>): Promise<AccountPayable> => {
    const response = await apiClient.post('/api/financial/accounts-payable', data);
    return response.data;
  },

  // Update account payable
  update: async (id: string, data: Partial<AccountPayable>): Promise<AccountPayable> => {
    const response = await apiClient.put(`/api/financial/accounts-payable/${id}`, data);
    return response.data;
  },

  // Delete account payable
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/financial/accounts-payable/${id}`);
  },

  // Toggle NF arrived status
  toggleNf: async (id: string, arrived: boolean): Promise<AccountPayable> => {
    const response = await apiClient.patch(`/api/financial/accounts-payable/${id}/nf`, { arrived });
    return response.data;
  },

  // Toggle Boleto arrived status
  toggleBoleto: async (id: string, arrived: boolean): Promise<AccountPayable> => {
    const response = await apiClient.patch(`/api/financial/accounts-payable/${id}/boleto`, { arrived });
    return response.data;
  },

  // Send to finance
  sendToFinance: async (id: string): Promise<AccountPayable> => {
    const response = await apiClient.post(`/api/financial/accounts-payable/${id}/send-to-finance`);
    return response.data;
  },
};
