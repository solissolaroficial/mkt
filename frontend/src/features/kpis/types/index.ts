// Importar tipos do legacy (KpiCategory, MonthlyData, etc)
import type { KpiCategory, MonthlyData, KpiLog, BreakdownItem } from '@/shared/types/legacy.types';

// Re-exportar para uso externo
export type { KpiCategory, MonthlyData, KpiLog, BreakdownItem };

// DTOs para API
export interface CreateKpiDTO {
  title: string;
  color: string;
  unit?: string;
}

export interface UpdateKpiDTO {
  title?: string;
  color?: string;
  unit?: string;
}

export interface UpdateMonthlyDataDTO {
  realized?: number;
  meta?: number;
  breakdown?: BreakdownItem[];
}

// Filtros
export interface KpiFilters {
  month?: string;
  year?: string;
  search?: string;
}

// Response da API
export interface KpiListResponse {
  data: KpiCategory[];
  pagination: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}