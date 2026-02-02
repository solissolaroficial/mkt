import React from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Bell,
  Search,
  ChevronDown,
  Settings,
  LogOut,
  Menu,
} from 'lucide-react';
import { useUIStore } from '@/shared/store/uiStore';
import { useAuth } from '@/features/auth';
import { useNotifications } from '@/features/notifications/hooks/useNotifications';
import NotificationPanel from '@/shared/ui/notifications/NotificationPanel';
import type { Notification } from '@/shared/types/legacy.types';
import { MONTHS } from '@/shared/utils/constants';

const Header: React.FC = () => {
  const navigate = useNavigate();
  const { user, logout } = useAuth();
  const {
    showNotifications,
    toggleNotifications,
    showUserMenu,
    toggleUserMenu,
    setShowUserMenu,
    selectedMonth,
    setSelectedMonth,
    selectedYear,
    setSelectedYear,
  } = useUIStore();

  // Hook de notificações
  const {
    activeNotifications,
    unreadCount,
    markAsRead,
    markAllAsRead,
    archiveNotification,
  } = useNotifications();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const handleNotificationClick = (notification: Notification) => {
    markAsRead(notification.id);
    // Adicionar lógica para navegar para a tarefa relacionada, se existir
    if (notification.task_id) {
      // Navegar para a tarefa
    }
  };

  return (
    <header className="h-16 bg-[#1a1d24] border-b border-gray-800 flex items-center justify-between px-6 sticky top-0 z-30">
      {/* Left side - Search */}
      <div className="flex items-center gap-4 flex-1">
        <div className="relative w-96">
          <Search
            className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500"
            size={18}
          />
          <input
            type="text"
            placeholder="Buscar..."
            className="w-full pl-10 pr-4 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent placeholder-gray-600"
          />
        </div>
      </div>

      {/* Right side - Filters + Notifications + User */}
      <div className="flex items-center gap-4">
        {/* Month Filter */}
        <select
          value={selectedMonth}
          onChange={(e) => setSelectedMonth(e.target.value)}
          className="px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
        >
          <option value="---">Ano Completo</option>
          {MONTHS.map((month) => (
            <option key={month} value={month}>
              {month}
            </option>
          ))}
        </select>

        {/* Year Filter */}
        <select
          value={selectedYear}
          onChange={(e) => setSelectedYear(e.target.value)}
          className="px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
        >
          <option value="2026">2026</option>
          <option value="2027">2027</option>
        </select>

        {/* Notifications */}
        <div className="relative">
          <button
            onClick={toggleNotifications}
            className="p-2 hover:bg-gray-800 rounded-lg transition-colors relative"
          >
            <Bell size={20} className="text-gray-400" />
            {/* CORREÇÃO: Indicador verde condicional */}
            {unreadCount > 0 && (
              <span className="absolute top-1 right-1 w-2 h-2 bg-emerald-500 rounded-full"></span>
            )}
          </button>

          {/* Notifications Dropdown */}
          {showNotifications && (
            <NotificationPanel
              notifications={activeNotifications}
              onMarkAsRead={markAsRead}
              onArchive={archiveNotification}
              onClearAll={markAllAsRead}
              onNotificationClick={handleNotificationClick}
            />
          )}
        </div>

        {/* User Menu */}
        <div className="relative">
          <button
            onClick={toggleUserMenu}
            className="flex items-center gap-2 px-3 py-2 hover:bg-gray-800 rounded-lg transition-colors"
          >
            <div className="w-8 h-8 bg-gradient-to-br from-[#1e5144] to-emerald-600 rounded-full flex items-center justify-center">
              <span className="text-white text-sm font-semibold">
                {user?.name?.charAt(0) || 'U'}
              </span>
            </div>
            <span className="text-gray-200 text-sm font-medium hidden md:block">
              {user?.name || 'Usuário'}
            </span>
            <ChevronDown size={16} className="text-gray-400" />
          </button>

          {/* User Dropdown */}
          {showUserMenu && (
            <div className="absolute right-0 mt-2 w-56 bg-[#1a1d24] border border-gray-800 rounded-xl shadow-2xl py-2">
              <div className="px-4 py-3 border-b border-gray-800">
                <p className="font-semibold text-white text-sm">{user?.name || 'Usuário'}</p>
                <p className="text-gray-500 text-xs">{user?.email || 'usuario@solis.com.br'}</p>
              </div>

              <button
                onClick={() => {
                  setShowUserMenu(false);
                  navigate('/settings');
                }}
                className="w-full flex items-center gap-3 px-4 py-2 hover:bg-gray-800 transition-colors"
              >
                <Settings size={16} className="text-gray-400" />
                <span className="text-gray-300 text-sm">Configurações</span>
              </button>

              <button
                onClick={handleLogout}
                className="w-full flex items-center gap-3 px-4 py-2 hover:bg-gray-800 transition-colors text-red-400"
              >
                <LogOut size={16} />
                <span className="text-sm">Sair</span>
              </button>
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
