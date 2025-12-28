import { apiClient } from '@/infrastructure/api/client';
import type { RepTableData, RepresentativeProfile } from '../types';

export const representativesService = {
  /**
   * Get representatives table data
   */
  getRepTableData: async (): Promise<RepTableData> => {
    const response = await apiClient.get('/api/representatives/table');
    return response.data;
  },

  /**
   * Get representative profile by name
   */
  getRepProfile: async (repName: string): Promise<RepresentativeProfile> => {
    const response = await apiClient.get(`/api/representatives/profile/${encodeURIComponent(repName)}`);
    return response.data;
  },

  /**
   * Get all representative profiles
   */
  getAllProfiles: async (): Promise<RepresentativeProfile[]> => {
    const response = await apiClient.get('/api/representatives/profiles');
    return response.data;
  },
};
