import { apiClient } from '@/infrastructure/api/client';
import type { AppUser } from '@/shared/types/user.types';

export const userService = {
  async listUsers(): Promise<AppUser[]> {
    const response = await apiClient.get('/api/users/all');
    return response.data;
  }
};
