import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type { RepTableData } from '@/shared/types/legacy.types';
import type { RepresentativeProfile } from '../types';

export const representativesService = {
  /**
   * Get representatives table data
   */
  getRepTableData: async (): Promise<RepTableData> => {
    const response = await apiClient.get(ENDPOINTS.REPRESENTATIVES.TABLE);
    return response.data;
  },

  /**
   * Get representative profile by name
   */
  getRepProfile: async (repName: string): Promise<RepresentativeProfile> => {
    const response = await apiClient.get(ENDPOINTS.REPRESENTATIVES.PROFILE(repName));
    return response.data;
  },

  /**
   * Get all representative profiles
   */
  getAllProfiles: async (): Promise<RepresentativeProfile[]> => {
    const response = await apiClient.get(ENDPOINTS.REPRESENTATIVES.PROFILES);
    return response.data.data;
  },
};
