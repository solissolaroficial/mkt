// ============================================
// Filter Types (Snake_Case)
// ============================================

export interface ShowroomItemFilters {
  representative_uuid?: string;
  city?: string;
  delivered?: boolean;
  page?: number;
  limit?: number;
  sort_by?: "created_at" | "delivery_forecast";
  sort_order?: 'asc' | 'desc';
}

export interface OfflineActionFilters {
  representative_uuid?: string;
  category?: string;
  status?: string;
  month?: string;
  start_date?: string;
  end_date?: string;
  page?: number;
  limit?: number;
  sort_by?: "action_date" | "created_at";
  sort_order?: 'asc' | 'desc';
}

export interface RepMarketingActionFilters {
  representative_uuid?: string;
  month?: string;
  page?: number;
  limit?: number;
  sort_by?: "date" | "created_at";
  sort_order?: 'asc' | 'desc';
}
