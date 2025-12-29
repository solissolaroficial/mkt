import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import {
  LayoutDashboard,
  Target,
  GraduationCap,
  Megaphone,
  CheckSquare,
  Calendar,
  MapPin,
  Share2,
  TrendingUp,
  Gift,
  DollarSign,
  KeyRound,
  Phone,
  PieChart,
  ChevronRight,
} from 'lucide-react';
import { useUIStore } from '@/shared/store/uiStore';

const Sidebar: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const { isSidebarOpen, toggleSidebar } = useUIStore();

  const menuItems = [
    { id: 'dashboard', label: 'Dashboard', icon: LayoutDashboard, path: '/dashboard' },
    { id: 'commercial', label: 'Comercial', icon: TrendingUp, path: '/commercial' },
    { id: 'kpis', label: 'KPIs', icon: Target, path: '/kpis' },
    { id: 'marketing', label: 'Marketing', icon: Megaphone, path: '/marketing' },
    { id: 'tasks', label: 'Tarefas', icon: CheckSquare, path: '/tasks' },
    { id: 'calendar', label: 'Calendário', icon: Calendar, path: '/calendar' },
    { id: 'pdv', label: 'PDV', icon: MapPin, path: '/pdv' },
    { id: 'social', label: 'Redes Sociais', icon: Share2, path: '/social' },
    { id: 'cooperative', label: 'Ações Cooperativas', icon: TrendingUp, path: '/cooperative' },
    { id: 'gifts', label: 'Brindes', icon: Gift, path: '/gifts' },
    { id: 'financial', label: 'Financeiro', icon: DollarSign, path: '/financial' },
    { id: 'budget', label: 'Orçamento', icon: PieChart, path: '/budget' },
    { id: 'representatives', label: 'Representantes', icon: GraduationCap, path: '/representatives' },
    { id: 'settings', label: 'Configurações', icon: KeyRound, path: '/settings' },
  ];

  return (
    <aside
      className={`fixed left-0 top-0 h-screen bg-[#1a1d24] border-r border-gray-800 transition-all duration-300 z-40 ${
        isSidebarOpen ? 'w-64' : 'w-16'
      }`}
    >
      {/* Logo/Brand */}
      <div className="h-16 flex items-center justify-between px-4 border-b border-gray-800">
        {isSidebarOpen && (
          <div className="flex items-center gap-2">
            <div className="w-8 h-8 bg-gradient-to-br from-[#1e5144] to-emerald-600 rounded-lg flex items-center justify-center">
              <span className="font-bold text-white text-sm">S</span>
            </div>
            <span className="font-bold text-white text-lg">
              Solis <span className="text-emerald-500">Hub</span>
            </span>
          </div>
        )}
        <button
          onClick={toggleSidebar}
          className="p-2 hover:bg-gray-800 rounded-lg transition-colors"
        >
          <ChevronRight
            size={20}
            className={`text-gray-400 transition-transform ${
              isSidebarOpen ? 'rotate-180' : ''
            }`}
          />
        </button>
      </div>

      {/* Menu Items */}
      <nav className="p-2 mt-2">
        {menuItems.map((item) => {
          const Icon = item.icon;
          const isActive = location.pathname === item.path;

          return (
            <button
              key={item.id}
              onClick={() => navigate(item.path)}
              className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-lg transition-all mb-1 ${
                isActive
                  ? 'bg-[#1e5144] text-white'
                  : 'text-gray-400 hover:bg-gray-800 hover:text-white'
              }`}
            >
              <Icon size={20} className="flex-shrink-0" />
              {isSidebarOpen && (
                <span className="text-sm font-medium">{item.label}</span>
              )}
            </button>
          );
        })}
      </nav>
    </aside>
  );
};

export default Sidebar;
