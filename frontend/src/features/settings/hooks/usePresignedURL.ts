import { useQuery } from '@tanstack/react-query';
import { apiClient } from '@/infrastructure/api/client';

export interface PresignedURLResponse {
  url: string;
  expires_in: number;
}

export const usePresignedURL = () => {
  return useQuery({
    queryKey: ['profile-photo-presigned-url'],
    queryFn: async (): Promise<PresignedURLResponse> => {
      const response = await apiClient.get<PresignedURLResponse>('/api/settings/profile-photo/url');
      return response.data;
    },
    enabled: true, // Sempre buscar presigned URL
    staleTime: 1000 * 60 * 50, // 50 minutos (presigned URL expira em 1 hora)
    refetchInterval: 1000 * 60 * 50, // Refetch a cada 50 minutos
  });
};
