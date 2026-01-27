import { useMutation, useQueryClient } from '@tanstack/react-query';

import { settingsService } from '../services/settingsService';
import { useAuth } from '@/features/auth/hooks/useAuth';

/**
 * Hook para upload de foto de perfil
 */
export function useUploadProfilePhoto() {
  const queryClient = useQueryClient();
  const { user, setUser } = useAuth();

  return useMutation({
    mutationFn: async (file: File) => {
      return await settingsService.uploadProfilePhoto(file);
    },
    onSuccess: async (data) => {
      // Invalidar queries relacionadas ao usuário para que os dados sejam recarregados
      // Isso inclui a URL completa da foto de perfil gerada pelo backend
      await queryClient.invalidateQueries({ queryKey: ['user'] });
      await queryClient.invalidateQueries({ queryKey: ['profile'] });
    },
  });
}
