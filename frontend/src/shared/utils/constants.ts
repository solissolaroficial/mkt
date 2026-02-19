// Re-exportar constantes legadas
export { MONTHS, FULL_MONTH_MAP, ALLOWED_USERS } from './legacy.constants';

// Constantes da aplicação
export const APP_NAME = 'Solis Hub';
export const ITEMS_PER_PAGE = 10;
export const API_TIMEOUT = 30000; // 30 segundos

// Slugs dos KPIs da Dashboard
export const DASHBOARD_KPIS_SLUGS = [
  'novos_inscritos_no_youtube',
  'treinamentos_realizados_reps',
  'pessoas_treinadas_solis',
  'acoes_de_marketing_reps',
  'visitantes_inbound',
  'autoridade_na_internet_da',
];

// Slugs dos KPIs Comerciais
export const COMMERCIAL_KPIS_SLUGS = [
  'taxa_de_oportunidades',
  'taxa_de_conversao_global',
];

// Slugs dos KPIs de Marketing
export const MARKETING_KPIS_SLUGS = [
  'novos_inscritos_no_youtube',
  'treinamentos_realizados_reps',
  'pessoas_treinadas_solis',
  'acoes_de_marketing_reps',
  'visitantes_inbound',
];

// Slugs dos KPIs de Treinamentos
export const TRAINING_KPIS_SLUGS = [
  'treinamentos_realizados_reps',
  'pessoas_treinadas_solis',
];

// Slugs dos KPIs necessários para cálculos cruzados na página de detalhes
export const KPI_DETAIL_DEPENDENCIES_SLUGS = [
  'taxa_de_oportunidades', // Necessário para cálculos de Ad Spend (CPL)
];

// Interface para filtros de listagem de KPIs
export interface KpiListFilters {
  month?: string;
  year?: string;
}

// Query keys para React Query
export const QUERY_KEYS = {
  AUTH: {
    USER: ['auth', 'user'],
  },
  KPIS: {
    ALL: ['kpis'],
    LIST: (filters?: KpiListFilters) => ['kpis', 'list', filters],
    DETAIL: (id: string) => ['kpis', 'detail', id],
    BY_SLUGS: (slugs: string[]) => ['kpis', 'by-slugs', slugs],
    DAILY_ENTRIES: (kpiId: string, month: string, year: number) =>
      ['kpis', 'daily-entries', kpiId, month, year],
  },
} as const;