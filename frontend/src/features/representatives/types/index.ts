// Representative Types
export interface Representative {
  uuid: string;
  code: number;
  name: string;
  email: string;
  phone: string;
  company: string;
  region: string;
  city: string;
  attendant: string;
  active: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface RepresentativeStats {
  uuid: string;
  onlineCount: number; // Alias para onlineActionCount (compatibilidade com implementação antiga)
  onlineActionCount: number;
  offlineCount: number; // Alias para offlineActionCount (compatibilidade com implementação antiga)
  offlineActionCount: number;
  offlineValue: number; // Alias para offlineActionValue (compatibilidade com implementação antiga)
  offlineActionValue: number;
  showroomItemCount: number;
  repMarketingCount: number;
  totalActions: number;
}

export interface CreateRepresentativeRequest {
  code: number;
  name: string;
  email: string;
  phone: string;
  company: string;
  region: string;
  city: string;
  attendant: string;
}

export interface UpdateRepresentativeRequest {
  name?: string;
  email?: string;
  phone?: string;
  company?: string;
  region?: string;
  city?: string;
  attendant?: string;
  active?: boolean;
}

export interface ListRepresentativesRequest {
  page?: number;
  limit?: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
  name?: string;
  company?: string;
  region?: string;
  city?: string;
  active?: boolean;
}

export interface ListRepresentativesResponse {
  data: Representative[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}

export interface RepresentativeTableData {
  data: Representative[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}

export interface RepProfileModalProps {
  isOpen: boolean;
  onClose: () => void;
  profile: RepresentativeProfile;
}

export interface RepresentativeProfile {
  uuid: string;
  code: number;
  name: string;
  email: string;
  phone: string;
  company: string;
  region: string;
  city: string;
  attendant: string;
  active: boolean;
  createdAt: string;
  updatedAt: string;
  trainingCount: number;
  onlineCount: number;
  offlineCount: number;
  offlineValue: number;
  showroomItemCount: number;
  repMarketingCount: number;
  totalActions: number;
}
