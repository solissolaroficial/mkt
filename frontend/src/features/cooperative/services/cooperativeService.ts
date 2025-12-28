import { apiClient } from '@/infrastructure/api/client';
import type { OfflineAction, ShowroomItem } from '../types';

export const cooperativeService = {
  // Get all offline actions
  getOfflineActions: async (): Promise<OfflineAction[]> => {
    const response = await apiClient.get('/api/cooperative/offline-actions');
    return response.data;
  },

  // Create new offline action
  createOfflineAction: async (data: Partial<OfflineAction>): Promise<OfflineAction> => {
    const response = await apiClient.post('/api/cooperative/offline-actions', data);
    return response.data;
  },

  // Update offline action
  updateOfflineAction: async (id: string, data: Partial<OfflineAction>): Promise<OfflineAction> => {
    const response = await apiClient.put(`/api/cooperative/offline-actions/${id}`, data);
    return response.data;
  },

  // Delete offline action
  deleteOfflineAction: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/cooperative/offline-actions/${id}`);
  },

  // Get all showroom items
  getShowroomItems: async (): Promise<ShowroomItem[]> => {
    const response = await apiClient.get('/api/cooperative/showroom-items');
    return response.data;
  },

  // Create new showroom item
  createShowroomItem: async (data: Partial<ShowroomItem>): Promise<ShowroomItem> => {
    const response = await apiClient.post('/api/cooperative/showroom-items', data);
    return response.data;
  },

  // Update showroom item
  updateShowroomItem: async (id: string, data: Partial<ShowroomItem>): Promise<ShowroomItem> => {
    const response = await apiClient.put(`/api/cooperative/showroom-items/${id}`, data);
    return response.data;
  },

  // Delete showroom item
  deleteShowroomItem: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/cooperative/showroom-items/${id}`);
  },
};
