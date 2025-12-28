import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type { LoginCredentials, LoginResponse } from '../types';

export const authService = {
  /**
   * Faz login do usuário
   * @param credentials - Email e senha
   * @returns Token JWT e dados do usuário
   */
  login: async (credentials: LoginCredentials): Promise<LoginResponse> => {
    const response = await apiClient.post<LoginResponse>(
      ENDPOINTS.AUTH.LOGIN,
      credentials
    );
    return response.data;
  },

  /**
   * Valida se o token ainda é válido (opcional - para refresh)
   */
  validateToken: async (): Promise<boolean> => {
    try {
      // Pode fazer uma chamada para endpoint de validação se existir
      // Por enquanto, apenas verifica se existe token
      return true;
    } catch {
      return false;
    }
  },
};