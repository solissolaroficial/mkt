import { create } from 'zustand';

// Helper function to get current month in the correct format (JAN, FEV, MAR, etc.)
const getCurrentMonth = (): string => {
  const months = ['JAN', 'FEV', 'MAR', 'ABR', 'MAI', 'JUN', 'JUL', 'AGO', 'SET', 'OUT', 'NOV', 'DEZ'];
  const currentMonthIndex = new Date().getMonth();
  return months[currentMonthIndex];
};

// Helper function to get current year
const getCurrentYear = (): string => {
  return new Date().getFullYear().toString();
};

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

  // Filtros globais (mês/ano) - Usa mês e ano atual dinamicamente
  selectedMonth: getCurrentMonth(),
  setSelectedMonth: (month) => set({ selectedMonth: month }),
  selectedYear: getCurrentYear(),
  setSelectedYear: (year) => set({ selectedYear: year }),
}));