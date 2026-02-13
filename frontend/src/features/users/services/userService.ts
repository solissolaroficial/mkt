import { apiClient } from '@/infrastructure/api/client';
import type { AppUser } from '@/shared/types/user.types';

// Transforma resposta snake_case do backend para camelCase do frontend
function transformUser(data: any): AppUser {
  return {
    id: data.id,
    email: data.email,
    name: data.name,
    role: data.role,
    active: data.active,
    profilePhotoKey: data.profile_photo_key,
    profilePhotoURL: data.profile_photo_url,
    created_at: data.created_at,
    updated_at: data.updated_at,
  };
}

export const userService = {
  async listUsers(): Promise<AppUser[]> {
    const response = await apiClient.get('/api/users/all');
    return response.data.map(transformUser);
  }
};
