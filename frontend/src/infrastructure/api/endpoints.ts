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
    GET_BY_SLUGS: '/api/kpis/by-slugs',
    UPDATE_MONTHLY: (kpiId: string) =>
      `/api/kpis/${kpiId}/monthly-data`,
  },
  CALENDAR: {
    LIST: '/api/calendar-posts',
    GET: (id: string) => `/api/calendar-posts/${id}`,
    CREATE: '/api/calendar-posts',
    UPDATE: (id: string) => `/api/calendar-posts/${id}`,
    DELETE: (id: string) => `/api/calendar-posts/${id}`,
    UPDATE_STATUS: (id: string) => `/api/calendar-posts/${id}/status`,
    CONFIRM_PUBLISHING: (id: string) => `/api/calendar-posts/${id}/confirm-publishing`,
  } as const,
} as const;