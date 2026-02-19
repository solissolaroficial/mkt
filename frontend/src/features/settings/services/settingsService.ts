import { apiClient } from '@/infrastructure/api/client';
import type {
  ProgramCredential,
  InternalContact,
  ChangePasswordRequest,
  ChangePasswordResponse,
  CreateCredentialRequest,
  UpdateCredentialRequest,
  CredentialResponse
} from '../types';

// Backend response type for credentials
interface CredentialsResponse {
  credentials: ProgramCredential[];
}

export const settingsService = {
  /**
   * Get program credentials
   */
  getCredentials: async (): Promise<CredentialsResponse> => {
    const response = await apiClient.get<CredentialsResponse>('/api/settings/credentials');
    return response.data;
  },

  /**
   * Create a new credential
   */
  createCredential: async (data: CreateCredentialRequest): Promise<CredentialResponse> => {
    const response = await apiClient.post<CredentialResponse>('/api/settings/credentials', data);
    return response.data;
  },

  /**
   * Update a credential
   */
  updateCredential: async (id: string, data: UpdateCredentialRequest): Promise<CredentialResponse> => {
    const response = await apiClient.put<CredentialResponse>(`/api/settings/credentials/${id}`, data);
    return response.data;
  },

  /**
   * Delete a credential
   */
  deleteCredential: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/settings/credentials/${id}`);
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

  /**
   * Upload profile photo
   */
  uploadProfilePhoto: async (file: File): Promise<{ key: string }> => {
    const formData = new FormData();
    formData.append('photo', file);
    
    const response = await apiClient.post('/api/settings/profile-photo', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },

  /**
   * Remove profile photo
   */
  removeProfilePhoto: async (): Promise<{ success: boolean; message: string }> => {
    const response = await apiClient.delete('/api/settings/profile-photo');
    return response.data;
  },
};
