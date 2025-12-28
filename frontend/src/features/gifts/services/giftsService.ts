import { apiClient } from '@/infrastructure/api/client';
import type { GiftItem, GiftTransaction } from '../types';

export const giftsService = {
  /**
   * Get all gift items
   */
  getItems: async (): Promise<GiftItem[]> => {
    const response = await apiClient.get('/api/gifts/items');
    return response.data;
  },

  /**
   * Get all gift transactions
   */
  getTransactions: async (): Promise<GiftTransaction[]> => {
    const response = await apiClient.get('/api/gifts/transactions');
    return response.data;
  },

  /**
   * Create a new gift item
   */
  createItem: async (item: Omit<GiftItem, 'stock'>): Promise<GiftItem> => {
    const response = await apiClient.post('/api/gifts/items', item);
    return response.data;
  },

  /**
   * Create a new transaction
   */
  createTransaction: async (transaction: Omit<GiftTransaction, 'id'>): Promise<GiftTransaction> => {
    const response = await apiClient.post('/api/gifts/transactions', transaction);
    return response.data;
  },

  /**
   * Update gift item
   */
  updateItem: async (id: string, item: Partial<GiftItem>): Promise<GiftItem> => {
    const response = await apiClient.put(`/api/gifts/items/${id}`, item);
    return response.data;
  },
};
