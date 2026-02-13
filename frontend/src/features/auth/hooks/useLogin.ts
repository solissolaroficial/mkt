import { useMutation } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { authService } from '../services/authService';
import { useAuthStore } from '../store/authStore';
import { tokenStorage } from '@/infrastructure/auth/tokenStorage';
import type { LoginCredentials, AuthUser } from '../types';

export const useLogin = () => {
  const navigate = useNavigate();
  const setUser = useAuthStore((state) => state.setUser);

  return useMutation({
    mutationFn: (credentials: LoginCredentials) => authService.login(credentials),

    onSuccess: (data) => {
      // Salvar tokens (access e refresh)
      tokenStorage.setTokens(data.access_token, data.refresh_token);

      // Transformar e salvar usuário no store
      const authUser: AuthUser = {
        id: data.user.id,
        email: data.user.email,
        name: data.user.name,
        role: data.user.role,
        profilePhotoKey: data.user.profile_photo_key,
        profilePhotoURL: data.user.profile_photo_url,
      };
      setUser(authUser);

      // Redirecionar para dashboard
      navigate('/dashboard');
    },

    onError: (error: any) => {
      console.error('Login failed:', error);
      // Erro será tratado no componente através de mutation.error
    },
  });
};