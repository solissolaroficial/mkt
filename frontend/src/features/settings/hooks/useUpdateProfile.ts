import { useMutation, useQueryClient } from '@tanstack/react-query';
import { settingsService } from '../services/settingsService';
import { useAuth } from '@/features/auth/hooks/useAuth';
import type { AuthUser } from '@/features/auth/types';

/**
 * Hook para atualizar o perfil do usuário
 */
export function useUpdateProfile() {
  const queryClient = useQueryClient();
  const { user, setUser } = useAuth();

  return useMutation({
    mutationFn: async (data: { name: string; email: string; role: string }) => {
      await settingsService.updateProfile(data);
    },
    onSuccess: async (_, variables) => {
      // Atualizar os dados do usuário no contexto de autenticação
      // Usamos o user atual para manter o id, mas atualizamos name, email e role
      if (user) {
        const updatedUser: AuthUser = {
          id: user.id,
          name: variables.name,
          email: variables.email,
          role: variables.role,
        };
        setUser(updatedUser);
      }

      // Invalidar queries relacionadas ao usuário para que os dados sejam recarregados
      await queryClient.invalidateQueries({ queryKey: ['user'] });
    },
  });
}
