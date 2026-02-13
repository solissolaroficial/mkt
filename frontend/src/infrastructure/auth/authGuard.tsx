import { Navigate, Outlet } from 'react-router-dom';
import { useEffect } from 'react';
import { tokenStorage } from './tokenStorage';
import { useAuthStore } from '@/features/auth/store/authStore';
import { useGetProfile } from '@/features/settings/hooks/useGetProfile';

interface AuthGuardProps {
  redirectTo?: string;
}

export const AuthGuard: React.FC<AuthGuardProps> = ({ redirectTo = '/login' }) => {
  const isAuthenticated = tokenStorage.hasToken();
  const user = useAuthStore((state) => state.user);
  const setUser = useAuthStore((state) => state.setUser);

  const { data: profileData, isLoading } = useGetProfile({
    enabled: isAuthenticated && !user,
  });

  useEffect(() => {
    if (!user && !isLoading && profileData) {
      setUser(profileData);
    }
  }, [user, isLoading, profileData, setUser]);

  if (!isAuthenticated) {
    return <Navigate to={redirectTo} replace />;
  }

  return <Outlet />;
};