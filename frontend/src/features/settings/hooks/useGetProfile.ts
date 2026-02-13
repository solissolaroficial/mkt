import { useQuery } from '@tanstack/react-query';
import { settingsService } from '../services/settingsService';
import type { AuthUser } from '@/features/auth/types';

interface UseGetProfileOptions {
  enabled?: boolean;
}

export function useGetProfile(options?: UseGetProfileOptions) {
  return useQuery({
    queryKey: ['profile'],
    queryFn: async (): Promise<AuthUser> => {
      const response = await settingsService.getProfile();
      const data = response.data;
      return {
        id: data.id,
        email: data.email,
        name: data.name,
        role: data.role,
        profilePhotoKey: data.profile_photo_key,
        profilePhotoURL: data.profile_photo_url,
      };
    },
    staleTime: 1000 * 60 * 5,
    enabled: options?.enabled ?? true,
  });
}
