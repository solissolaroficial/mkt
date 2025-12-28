import { create } from 'zustand';
import { tokenStorage } from '@/infrastructure/auth/tokenStorage';
import type { AuthUser } from '../types';

interface AuthState {
  user: AuthUser | null;
  isAuthenticated: boolean;
  setUser: (user: AuthUser | null) => void;
  logout: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  isAuthenticated: tokenStorage.hasToken(),

  setUser: (user) => set({ user, isAuthenticated: !!user }),

  logout: () => {
    tokenStorage.clearTokens();
    set({ user: null, isAuthenticated: false });
    // Redirecionar para login será feito no componente
  },
}));