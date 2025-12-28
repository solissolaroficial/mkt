import { apiClient } from '@/infrastructure/api/client';
import type { ChannelPerformance, MarketingChannelData } from '../types';

export const marketingService = {
  /**
   * Get channel performance data for a specific month
   */
  getChannelData: async (month: string): Promise<ChannelPerformance[]> => {
    const response = await apiClient.get(`/api/marketing/channels/${month}`);
    return response.data;
  },

  /**
   * Get annual channel performance data
   */
  getAnnualChannelData: async (): Promise<ChannelPerformance[]> => {
    const response = await apiClient.get('/api/marketing/channels/annual');
    return response.data;
  },

  /**
   * Get all marketing channels data
   */
  getAllChannelData: async (): Promise<MarketingChannelData> => {
    const response = await apiClient.get('/api/marketing/channels');
    return response.data;
  },
};
