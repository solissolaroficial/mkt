import { useMutation } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { authService } from '../services/authService';
import { useAuthStore } from '../store/authStore';
import { tokenStorage } from '@/infrastructure/auth/tokenStorage';
import type { LoginCredentials } from '../types';

export const useLogin = () => {
  const navigate = useNavigate();
  const setUser = useAuthStore((state) => state.setUser);

  return useMutation({
    mutationFn: (credentials: LoginCredentials) => authService.login(credentials),

    onSuccess: (data) => {
      // Salvar tokens (access e refresh)
      tokenStorage.setTokens(data.access_token, data.refresh_token);

      // Salvar usuário no store
      setUser(data.user);

      // Redirecionar para dashboard
      navigate('/dashboard');
    },

    onError: (error: any) => {
      console.error('Login failed:', error);
      // Erro será tratado no componente através de mutation.error
    },
  });
};