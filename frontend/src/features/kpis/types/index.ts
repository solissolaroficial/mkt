// Importar tipos do legacy (KpiCategory, MonthlyData, etc)
import type { KpiCategory, MonthlyData, KpiLog, BreakdownItem } from '@/shared/types/legacy.types';

// Re-exportar para uso externo
export type { KpiCategory, MonthlyData, KpiLog, BreakdownItem };

// Tipos de entradas diárias
export type {
  DailyEntry,
  AddDailyEntryDTO,
  UpdateDailyEntryDTO,
  DeleteDailyEntryDTO,
} from './daily-entry.types';

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
  year?: number; // Year (e.g., 2024, 2025)
  realized?: number;
  meta?: number;
  breakdown?: BreakdownItem[];
  month?: string;
  context?: string; // Context of the change (for audit log)
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