export { default as SettingsView } from './ui/SettingsView';
export { default as PasswordsView } from './ui/PasswordsView';
export { default as ContactsView } from './ui/ContactsView';

export { useCredentials, useContacts } from './hooks/useSettings';

export type { 
  Notification, 
  KpiCategory, 
  ProgramCredential, 
  InternalContact,
  SettingsViewProps,
  SettingsSection,
  UserFormData
} from './types';
