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
  PDV: {
    // Posts PDV
    LIST: '/api/pdv/posts',
    GET: (id: string) => `/api/pdv/posts/${id}`,
    CREATE: '/api/pdv/posts',
    UPDATE: (id: string) => `/api/pdv/posts/${id}`,
    DELETE: (id: string) => `/api/pdv/posts/${id}`,
    UPDATE_STATUS: (id: string) => `/api/pdv/posts/${id}/status`,
    // PDVs Recorrentes
    LIST_RECURRENT: '/api/pdv/recurrent',
    GET_RECURRENT: (id: string) => `/api/pdv/recurrent/${id}`,
    CREATE_RECURRENT: '/api/pdv/recurrent',
    UPDATE_RECURRENT: (id: string) => `/api/pdv/recurrent/${id}`,
    DELETE_RECURRENT: (id: string) => `/api/pdv/recurrent/${id}`,
  } as const,
} as const;