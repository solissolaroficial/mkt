// Re-exportar constantes legadas
export { MONTHS, FULL_MONTH_MAP, ALLOWED_USERS } from './legacy.constants';

// Constantes da aplicação
export const APP_NAME = 'Solis Hub';
export const ITEMS_PER_PAGE = 10;
export const API_TIMEOUT = 30000; // 30 segundos

// Query keys para React Query
export const QUERY_KEYS = {
  AUTH: {
    USER: ['auth', 'user'],
  },
  KPIS: {
    ALL: ['kpis'],
    LIST: (filters?: unknown) => ['kpis', 'list', filters],
    DETAIL: (id: string) => ['kpis', 'detail', id],
  },
} as const;