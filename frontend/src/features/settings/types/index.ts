import type { Notification, KpiCategory, ProgramCredential, InternalContact } from '@/shared/types';

export type { Notification, KpiCategory, ProgramCredential, InternalContact };

export interface SettingsViewProps {
  onLogout: () => void;
  notifications?: Notification[];
  onUpdateNotification?: (notifications: Notification[]) => void;
  kpis?: KpiCategory[];
  onUpdateKpiMeta?: (kpiId: string, month: string, year: number, newMeta: number) => void;
}

export type SettingsSection = 'profile' | 'security' | 'history' | 'goals' | 'passwords';

export interface UserFormData {
  name: string;
  role: string;
  email: string;
  currentPassword: string;
  newPassword: string;
  confirmPassword: string;
}

export interface ChangePasswordRequest {
  currentPassword: string;
  newPassword: string;
}

export interface ChangePasswordResponse {
  message: string;
}

// ==================== CREDENTIALS CRUD TYPES ====================

export interface CreateCredentialRequest {
  name: string;
  user?: string;
  password?: string;
  access?: string;
  notes?: string;
  active?: boolean;
}

export interface UpdateCredentialRequest {
  name?: string;
  user?: string;
  password?: string;
  access?: string;
  notes?: string;
  active?: boolean;
}

export interface CredentialResponse {
  id: string;
  name: string;
  user?: string;
  password?: string;
  access?: string;
  notes?: string;
  active?: boolean;
  created_at?: string;
  updated_at?: string;
}

export interface CredentialsListResponse {
  credentials: CredentialResponse[];
}
