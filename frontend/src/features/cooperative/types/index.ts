// ============================================
// Cooperative Types (Snake_Case - Backend Compatible)
// ============================================

// Re-export filters and pagination types
export * from './filters';
export * from './pagination';

// ============================================
// ShowroomItem
// ============================================

export interface ShowroomItem {
  uuid: string;
  pdv: string;
  city?: string;
  contact?: string;
  rep_name: string;
  delivery_forecast?: string; // ISO 8601 (YYYY-MM-DD)
  workshop_date?: string;     // ISO 8601 (YYYY-MM-DD)
  delivered: boolean;
  created_at: string;        // ISO 8601
  updated_at: string;        // ISO 8601
  deleted_at?: string;       // ISO 8601
}

export interface CreateShowroomItemRequest {
  pdv: string;
  city?: string;
  contact?: string;
  rep_name: string;
  delivery_forecast?: string; // ISO 8601
  workshop_date?: string;     // ISO 8601
}

export interface UpdateShowroomItemRequest {
  pdv?: string;
  city?: string;
  contact?: string;
  rep_name?: string;
  delivery_forecast?: string;
  workshop_date?: string;
  delivered?: boolean;
}

// ============================================
// OfflineAction
// ============================================

export interface OfflineAction {
  uuid: string;
  requested_amount: number;
  approved_amount?: string;   // Valor formatado (R$)
  order_number?: string;
  action_date: string;       // ISO 8601 (YYYY-MM-DD)
  departure_date?: string;     // ISO 8601 (YYYY-MM-DD)
  delivery_forecast?: string; // ISO 8601 (YYYY-MM-DD)
  delivery_date?: string;     // ISO 8601 (YYYY-MM-DD)
  city?: string;
  uf?: string;
  category: string;
  pdv: string;
  rep_name: string;
  observation?: string;
  scored?: string;           // "SIM", "NÃO", "AINDA NÃO"
  status: string;            // "pending", "approved", "rejected", "completed"
  month: string;             // "JAN", "FEV", "MAR", etc.
  created_at: string;        // ISO 8601
  updated_at: string;        // ISO 8601
  deleted_at?: string;       // ISO 8601
}

export interface CreateOfflineActionRequest {
  requested_amount: number;
  category: string;
  action_date: string;       // ISO 8601
  pdv: string;
  rep_name: string;
  observation: string;       // OBRIGATÓRIO
}

export interface UpdateOfflineActionRequest {
  approved_amount?: string;
  order_number?: string;
  departure_date?: string;
  delivery_forecast?: string;
  delivery_date?: string;
  city?: string;
  uf?: string;
  scored?: string;
  status?: string;
  observation?: string;
  pdv?: string;
  rep_name?: string;
}

// ============================================
// RepMarketingAction
// ============================================

export interface RepMarketingAction {
  uuid: string;
  rep_name: string;
  date: string;             // ISO 8601 (YYYY-MM-DD)
  description: string;
  month: string;             // "JAN", "FEV", "MAR", etc.
  created_at: string;        // ISO 8601
  updated_at: string;        // ISO 8601
  deleted_at?: string;       // ISO 8601
}

export interface CreateRepMarketingActionRequest {
  rep_name: string;
  date: string;             // ISO 8601
  description: string;
}

export interface UpdateRepMarketingActionRequest {
  rep_name?: string;
  date?: string;
  description?: string;
}
