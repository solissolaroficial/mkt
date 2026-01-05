export interface AppUser {
  id: string;
  email: string;
  name: string;
  role: 'admin' | 'marketing' | 'commercial';
  active: boolean;
  created_at: string;
  updated_at: string;
}
