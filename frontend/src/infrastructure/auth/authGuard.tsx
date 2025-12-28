import { Navigate, Outlet } from 'react-router-dom';
import { tokenStorage } from './tokenStorage';

interface AuthGuardProps {
  redirectTo?: string;
}

export const AuthGuard: React.FC<AuthGuardProps> = ({ redirectTo = '/login' }) => {
  const isAuthenticated = tokenStorage.hasToken();

  if (!isAuthenticated) {
    return <Navigate to={redirectTo} replace />;
  }

  return <Outlet />;
};