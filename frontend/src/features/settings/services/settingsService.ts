import { apiClient } from '@/infrastructure/api/client';
import type { ProgramCredential, InternalContact, ChangePasswordRequest, ChangePasswordResponse } from '../types';

export const settingsService = {
  /**
   * Get program credentials
   */
  getCredentials: async (): Promise<ProgramCredential[]> => {
    const response = await apiClient.get('/api/settings/credentials');
    return response.data;
  },

  /**
   * Get internal contacts
   */
  getContacts: async (): Promise<InternalContact[]> => {
    const response = await apiClient.get('/api/settings/contacts');
    return response.data;
  },

  /**
   * Get user profile
   */
  getProfile: async () => {
    const response = await apiClient.get('/api/settings/profile');
    return response.data;
  },

  /**
   * Update user profile
   */
  updateProfile: async (data: { name: string; email: string; role: string }): Promise<void> => {
    const response = await apiClient.put('/api/settings/profile', data);
    return response.data;
  },

  /**
   * Update user password
   */
  updatePassword: async (data: ChangePasswordRequest): Promise<ChangePasswordResponse> => {
    const response = await apiClient.put('/api/settings/password', data);
    return response.data;
  },
};
