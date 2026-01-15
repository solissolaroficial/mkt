import type { BudgetItem } from '@/shared/types';

export type { BudgetItem };

// ==================== REQUESTS ====================

export interface CreateBudgetItemRequest {
  cod_obj: string;        // OBRIGATÓRIO
  obj: string;            // OBRIGATÓRIO
  cod_grp: string;        // OBRIGATÓRIO
  grp: string;            // OBRIGATÓRIO
  cod: string;            // OBRIGATÓRIO
  desc: string;           // OBRIGATÓRIO
  vals: number[];         // OBRIGATÓRIO, exatamente 12 valores
  realized_vals: number[]; // OBRIGATÓRIO, exatamente 12 valores
  year: number;           // OBRIGATÓRIO, 2000-2100
}

export interface UpdateBudgetItemRequest {
  cod_obj?: string;
  obj?: string;
  cod_grp?: string;
  grp?: string;
  cod?: string;
  desc?: string;
  vals?: number[];
  realized_vals?: number[];
  year?: number;
}

export interface DeleteBudgetItemRequest {
  uuid: string;
}

export interface BatchCreateBudgetItemsRequest {
  items: CreateBudgetItemRequest[];
}

// ==================== RESPONSES ====================

export interface BudgetItemListResponse {
  data: BudgetItem[];
  total: number;
  page: number;
  limit: number;
}

export interface BudgetSummaryResponse {
  cod_obj: string;
  obj: string;
  cod_grp: string;
  grp: string;
  total_budget: number;
  total_realized: number;
  variance: number;
}

export interface BudgetYearsResponse {
  years: number[];
}

// ==================== FILTERS ====================

export interface BudgetItemFilters {
  // Exact match filters
  cod_obj?: string;
  cod_grp?: string;
  cod?: string;
  year?: number;

  // LIKE search filters (case-insensitive)
  obj?: string;
  grp?: string;
  desc?: string;

  // Pagination
  page?: number;      // Default: 1, Min: 1
  limit?: number;     // Default: 50, Min: 1, Max: 100

  // Sorting
  sort_by?: 'cod_obj' | 'obj' | 'cod_grp' | 'grp' | 'cod' | 'desc' | 'created_at';
  sort_order?: 'asc' | 'desc';
}

// ==================== UI COMPONENTS ====================

export interface BudgetViewProps {
  onBack: () => void;
  dataItems: BudgetItem[];      // Dados do backend
  isLoading?: boolean;           // Estado de carregamento
  error?: Error | null;          // Erro de carregamento
  onRefetch?: () => void;       // Função para recarregar
  years?: number[];             // Anos disponíveis
  selectedYear?: number;        // Ano selecionado
  onYearChange?: (year: number | undefined) => void; // Callback para mudar ano
  filterFn?: (item: BudgetItem) => boolean;
  customTitle?: string;
  customSubtitle?: string;
}

export interface EditingCell {
  id: string;
  colIndex: number;
  type: 'budget' | 'realized';
}

export interface BudgetGroup {
  id: string;
  codObj: string;
  name: string;
  items: BudgetItem[];
  subgroups: Record<string, BudgetSubgroup>;
  totals: number[];
  realizedTotals: number[];
}

export interface BudgetSubgroup {
  id: string;
  codGrp: string;
  name: string;
  items: BudgetItem[];
  totals: number[];
  realizedTotals: number[];
}

export interface MarketingBudgetViewProps {
  onBack: () => void;
  selectedMonth: string;
}

export interface BudgetSummaryItem {
  id: string;
  name: string;
  budget: number;
  realized: number;
  requested: number;
  available: number;
  percentUsed: number;
}
