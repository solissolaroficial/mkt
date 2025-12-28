export const ENDPOINTS = {
  AUTH: {
    LOGIN: '/api/auth/login',
  },
  KPIS: {
    LIST: '/api/kpis',
    CREATE: '/api/kpis',
    GET: (id: string) => `/api/kpis/${id}`,
    UPDATE: (id: string) => `/api/kpis/${id}`,
    DELETE: (id: string) => `/api/kpis/${id}`,
    UPDATE_MONTHLY: (kpiId: string, monthlyId: string) =>
      `/api/kpis/${kpiId}/monthly-data/${monthlyId}`,
  },
} as const;