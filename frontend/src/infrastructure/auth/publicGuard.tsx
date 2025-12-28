import { Navigate, Outlet } from 'react-router-dom';
import { tokenStorage } from './tokenStorage';

interface PublicGuardProps {
  redirectTo?: string;
}

export const PublicGuard: React.FC<PublicGuardProps> = ({ redirectTo = '/dashboard' }) => {
  const isAuthenticated = tokenStorage.hasToken();

  if (isAuthenticated) {
    return <Navigate to={redirectTo} replace />;
  }

  return <Outlet />;
};