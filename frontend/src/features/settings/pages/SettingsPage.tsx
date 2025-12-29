import React from 'react';
import SettingsView from '../ui/SettingsView';
import type { Notification, KpiCategory } from '@/shared/types';
import { useKpis } from '@/features/kpis/hooks/useKpis';
import { useUpdateMeta } from '@/features/kpis/hooks/useUpdateMeta';

const SettingsPage: React.FC = () => {
  // Buscar lista de KPIs
  const { data: kpis, isLoading: isLoadingKpis } = useKpis();
  const kpisArray = kpis?.data || [];

  // Hook para atualizar meta
  const updateMeta = useUpdateMeta();

  const handleLogout = () => {
    // TODO: Implement logout logic
    console.log('Logout clicked');
  };

  const handleUpdateNotification = (notifications: Notification[]) => {
    // TODO: Implement notification update logic
    console.log('Notifications updated:', notifications);
  };

  const handleUpdateKpiMeta = (kpiId: string, month: string, year: number, newMeta: number) => {
    updateMeta.mutate({
      kpiId,
      year,
      month,
      meta: newMeta
    });
  };

  // Mock data - TODO: Replace with actual data from hooks
  const notifications: Notification[] = [];

  return (
    <SettingsView
      onLogout={handleLogout}
      notifications={notifications}
      onUpdateNotification={handleUpdateNotification}
      kpis={kpisArray}
      onUpdateKpiMeta={handleUpdateKpiMeta}
    />
  );
};

export default SettingsPage;
