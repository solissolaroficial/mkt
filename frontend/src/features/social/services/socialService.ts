import { apiClient } from '@/infrastructure/api/client';
import type { SocialBenchmarking } from '../types';

export const socialService = {
  /**
   * Get social benchmarking data
   */
  getBenchmarking: async (): Promise<SocialBenchmarking[]> => {
    const response = await apiClient.get('/api/social/benchmarking');
    return response.data;
  },
};
