import { useMutation, useQueryClient } from '@tanstack/react-query';
import { settingsService } from '../services/settingsService';
import { ChangePasswordRequest, ChangePasswordResponse } from '../types';

export function useChangePassword() {
  const queryClient = useQueryClient();

  const mutation = useMutation<ChangePasswordResponse, Error, ChangePasswordRequest>({
    mutationFn: (data: ChangePasswordRequest) => settingsService.updatePassword(data),
    onSuccess: () => {
      // Invalidar queries relacionados ao usuário se necessário
      queryClient.invalidateQueries({ queryKey: ['user'] });
    },
  });

  return {
    changePassword: mutation.mutate,
    isLoading: mutation.isPending,
    error: mutation.error,
    isSuccess: mutation.isSuccess,
    reset: mutation.reset,
  };
}
