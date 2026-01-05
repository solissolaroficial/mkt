import { useQuery } from '@tanstack/react-query';
import { userService } from '../services/userService';
import type { AppUser } from '@/shared/types/user.types';
import { useToast } from '@/shared/hooks/useToast';
import { useEffect } from 'react';

export function useUsers() {
  const toast = useToast();
  
  const query = useQuery<AppUser[]>({
    queryKey: ['users'],
    queryFn: userService.listUsers,
    staleTime: 5 * 60 * 1000, // 5 minutos
    retry: 3,
  });

  useEffect(() => {
    if (query.error) {
      console.error('Failed to load users:', query.error);
      toast.error('Erro ao carregar usuários');
    }
  }, [query.error, toast]);

  return query;
}
