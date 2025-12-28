import React from 'react';
import SettingsView from '../ui/SettingsView';
import type { Notification, KpiCategory } from '@/shared/types';

const SettingsPage: React.FC = () => {
  const handleLogout = () => {
    // TODO: Implement logout logic
    console.log('Logout clicked');
  };

  const handleUpdateNotification = (notifications: Notification[]) => {
    // TODO: Implement notification update logic
    console.log('Notifications updated:', notifications);
  };

  const handleUpdateKpiMeta = (kpiId: string, month: string, newMeta: number) => {
    // TODO: Implement KPI meta update logic
    console.log('KPI meta updated:', kpiId, month, newMeta);
  };

  // Mock data - TODO: Replace with actual data from hooks
  const notifications: Notification[] = [];
  const kpis: KpiCategory[] = [];

  return (
    <SettingsView
      onLogout={handleLogout}
      notifications={notifications}
      onUpdateNotification={handleUpdateNotification}
      kpis={kpis}
      onUpdateKpiMeta={handleUpdateKpiMeta}
    />
  );
};

export default SettingsPage;
