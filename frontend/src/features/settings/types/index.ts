import type { Notification, KpiCategory, ProgramCredential, InternalContact } from '@/shared/types';

export type { Notification, KpiCategory, ProgramCredential, InternalContact };

export interface SettingsViewProps {
  onLogout: () => void;
  notifications?: Notification[];
  onUpdateNotification?: (notifications: Notification[]) => void;
  kpis?: KpiCategory[];
  onUpdateKpiMeta?: (kpiId: string, month: string, newMeta: number) => void;
}

export type SettingsSection = 'profile' | 'security' | 'history' | 'goals';

export interface UserFormData {
  name: string;
  role: string;
  email: string;
  phone: string;
  currentPassword: string;
  newPassword: string;
  confirmPassword: string;
}
