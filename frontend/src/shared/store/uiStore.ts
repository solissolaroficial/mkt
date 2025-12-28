import { create } from 'zustand';

interface UIState {
  // Sidebar
  isSidebarOpen: boolean;
  toggleSidebar: () => void;
  setSidebarOpen: (isOpen: boolean) => void;

  // Notificações
  showNotifications: boolean;
  toggleNotifications: () => void;
  setShowNotifications: (show: boolean) => void;

  // User menu
  showUserMenu: boolean;
  toggleUserMenu: () => void;
  setShowUserMenu: (show: boolean) => void;

  // Filtros globais
  selectedMonth: string;
  setSelectedMonth: (month: string) => void;
  selectedYear: string;
  setSelectedYear: (year: string) => void;
}

export const useUIStore = create<UIState>((set) => ({
  // Sidebar state
  isSidebarOpen: true,
  toggleSidebar: () => set((state) => ({ isSidebarOpen: !state.isSidebarOpen })),
  setSidebarOpen: (isOpen) => set({ isSidebarOpen: isOpen }),

  // Notifications state
  showNotifications: false,
  toggleNotifications: () => set((state) => ({ showNotifications: !state.showNotifications })),
  setShowNotifications: (show) => set({ showNotifications: show }),

  // User menu state
  showUserMenu: false,
  toggleUserMenu: () => set((state) => ({ showUserMenu: !state.showUserMenu })),
  setShowUserMenu: (show) => set({ showUserMenu: show }),

  // Filtros globais (mês/ano)
  selectedMonth: 'NOV',
  setSelectedMonth: (month) => set({ selectedMonth: month }),
  selectedYear: '2025',
  setSelectedYear: (year) => set({ selectedYear: year }),
}));