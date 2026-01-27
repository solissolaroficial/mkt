import { useMutation, useQueryClient } from '@tanstack/react-query';

import { settingsService } from '../services/settingsService';

/**
 * Hook para remoção de foto de perfil
 */
export function useRemoveProfilePhoto() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async () => {
      return await settingsService.removeProfilePhoto();
    },
    onSuccess: async () => {
      // Invalidar queries relacionadas ao usuário para que os dados sejam recarregados
      await queryClient.invalidateQueries({ queryKey: ['user'] });
      await queryClient.invalidateQueries({ queryKey: ['profile'] });
    },
  });
}
