import { createBrowserRouter, Navigate } from 'react-router-dom';
import { AuthGuard } from '../auth/authGuard';
import { PublicGuard } from '../auth/publicGuard';
import { tokenStorage } from '../auth/tokenStorage';

// Layouts
import MainLayout from '@/shared/ui/layout/MainLayout';

// Pages - Auth
import { LoginView } from '@/features/auth';

// Pages - Dashboard
import DashboardView from '@/features/dashboard/ui/DashboardView';

// Pages - KPIs
import { KpiListView } from '@/features/kpis';

// Pages - Marketing
import MarketingPage from '@/features/marketing/pages/MarketingPage';

// Pages - Settings
import SettingsPage from '@/features/settings/pages/SettingsPage';

// Pages - Representatives
import RepresentativesPage from '@/features/representatives/pages/RepresentativesPage';

// Pages - Social
import SocialPage from '@/features/social/pages/SocialPage';

// Pages - Gifts
import GiftsPage from '@/features/gifts/pages/GiftsPage';

// Pages - PDV
import PdvPage from '@/features/pdv/pages/PdvPage';

// Pages - Tasks
import TasksPage from '@/features/tasks/pages/TasksPage';

// Pages - Financial
import FinancialPage from '@/features/financial/pages/FinancialPage';

// Pages - Budget
import BudgetPage from '@/features/budget/pages/BudgetPage';
import MarketingBudgetPage from '@/features/budget/pages/MarketingBudgetPage';

// Pages - Cooperative
import CooperativePage from '@/features/cooperative/pages/CooperativePage';

// Pages - Calendar
import CalendarPage from '@/features/calendar/pages/CalendarPage';

// Componente para redirecionar da raiz baseado em autenticação
const RootRedirect: React.FC = () => {
  const isAuthenticated = tokenStorage.hasToken();
  return <Navigate to={isAuthenticated ? '/dashboard' : '/login'} replace />;
};

export const router = createBrowserRouter([
  {
    path: '/',
    element: <RootRedirect />,
  },

  // Rotas públicas (apenas para não autenticados)
  {
    element: <PublicGuard />,
    children: [
      {
        path: '/login',
        element: <LoginView />,
      },
    ],
  },

  // Rotas protegidas (apenas para autenticados)
  {
    element: <AuthGuard />,
    children: [
      {
        element: <MainLayout />,
        children: [
          {
            path: '/dashboard',
            element: <DashboardView />,
          },
          {
            path: '/kpis',
            element: <KpiListView />,
          },
          {
            path: '/marketing',
            element: <MarketingPage />,
          },
          {
            path: '/settings',
            element: <SettingsPage />,
          },
          {
            path: '/representatives',
            element: <RepresentativesPage />,
          },
          {
            path: '/social',
            element: <SocialPage />,
          },
          {
            path: '/gifts',
            element: <GiftsPage />,
          },
          {
            path: '/pdv',
            element: <PdvPage />,
          },
          {
            path: '/tasks',
            element: <TasksPage />,
          },
          {
            path: '/financial',
            element: <FinancialPage />,
          },
          {
            path: '/budget',
            element: <BudgetPage />,
          },
          {
            path: '/budget/marketing',
            element: <MarketingBudgetPage />,
          },
          {
            path: '/cooperative',
            element: <CooperativePage />,
          },
          {
            path: '/calendar',
            element: <CalendarPage />,
          },
        ],
      },
    ],
  },

  // 404
  {
    path: '*',
    element: (
      <div className="min-h-screen bg-[#0f1115] flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-4xl font-bold text-white mb-2">404</h1>
          <p className="text-gray-400">Página não encontrada</p>
        </div>
      </div>
    ),
  },
]);
