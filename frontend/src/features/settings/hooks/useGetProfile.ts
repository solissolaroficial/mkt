import { useQuery } from '@tanstack/react-query';
import { settingsService } from '../services/settingsService';
import type { AuthUser } from '@/features/auth/types';

export function useGetProfile() {
  return useQuery({
    queryKey: ['profile'],
    queryFn: async (): Promise<AuthUser> => {
      const response = await settingsService.getProfile();
      return response.data;
    },
    staleTime: 1000 * 60 * 5, // 5 minutos
  });
}
