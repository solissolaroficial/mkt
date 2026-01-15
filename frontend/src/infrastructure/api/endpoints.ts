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
  SOCIAL: {
    // Social Benchmarking
    LIST: '/api/social/benchmarking',
    GET: (id: string) => `/api/social/benchmarking/${id}`,
    CREATE: '/api/social/benchmarking',
    UPDATE: (id: string) => `/api/social/benchmarking/${id}`,
    DELETE: (id: string) => `/api/social/benchmarking/${id}`,
    // Social Posts
    LIST_POSTS: '/api/social/posts',
    GET_POST: (id: string) => `/api/social/posts/${id}`,
    CREATE_POST: '/api/social/posts',
    UPDATE_POST: (id: string) => `/api/social/posts/${id}`,
    DELETE_POST: (id: string) => `/api/social/posts/${id}`,
    // Social Daily Aggregations
    LIST_AGGREGATIONS: '/api/social/daily-aggregations',
    GET_AGGREGATION: (id: string) => `/api/social/daily-aggregations/${id}`,
    RECALCULATE_AGGREGATIONS: (brandName: string, date: string) =>
      `/api/social/daily-aggregations/recalculate/${brandName}/${date}`,
  } as const,
  COOPERATIVE: {
    // Offline Actions
    OFFLINE_ACTIONS_LIST: '/api/cooperative/offline-actions',
    OFFLINE_ACTIONS_GET: (uuid: string) => `/api/cooperative/offline-actions/${uuid}`,
    OFFLINE_ACTIONS_CREATE: '/api/cooperative/offline-actions',
    OFFLINE_ACTIONS_UPDATE: (uuid: string) => `/api/cooperative/offline-actions/${uuid}`,
    OFFLINE_ACTIONS_DELETE: (uuid: string) => `/api/cooperative/offline-actions/${uuid}`,
    // Showroom Items
    SHOWROOM_ITEMS_LIST: '/api/cooperative/showroom-items',
    SHOWROOM_ITEMS_GET: (uuid: string) => `/api/cooperative/showroom-items/${uuid}`,
    SHOWROOM_ITEMS_CREATE: '/api/cooperative/showroom-items',
    SHOWROOM_ITEMS_UPDATE: (uuid: string) => `/api/cooperative/showroom-items/${uuid}`,
    SHOWROOM_ITEMS_DELETE: (uuid: string) => `/api/cooperative/showroom-items/${uuid}`,
    // Rep Marketing Actions
    REP_MARKETING_ACTIONS_LIST: '/api/cooperative/rep-marketing-actions',
    REP_MARKETING_ACTIONS_GET: (uuid: string) => `/api/cooperative/rep-marketing-actions/${uuid}`,
    REP_MARKETING_ACTIONS_CREATE: '/api/cooperative/rep-marketing-actions',
    REP_MARKETING_ACTIONS_UPDATE: (uuid: string) => `/api/cooperative/rep-marketing-actions/${uuid}`,
    REP_MARKETING_ACTIONS_DELETE: (uuid: string) => `/api/cooperative/rep-marketing-actions/${uuid}`,
  } as const,
  GIFTS: {
    // Gift Items
    ITEMS_LIST: '/api/gifts/items',
    ITEMS_GET: (id: string) => `/api/gifts/items/${id}`,
    ITEMS_CREATE: '/api/gifts/items',
    ITEMS_UPDATE: (id: string) => `/api/gifts/items/${id}`,
    ITEMS_DELETE: (id: string) => `/api/gifts/items/${id}`,
    // Gift Transactions
    TRANSACTIONS_LIST: '/api/gifts/transactions',
    TRANSACTIONS_GET: (id: string) => `/api/gifts/transactions/${id}`,
    TRANSACTIONS_CREATE: '/api/gifts/transactions',
    TRANSACTIONS_UPDATE: (id: string) => `/api/gifts/transactions/${id}`,
    TRANSACTIONS_DELETE: (id: string) => `/api/gifts/transactions/${id}`,
  } as const,
  BUDGET: {
    LIST: '/api/budget',
    GET: (uuid: string) => `/api/budget/${uuid}`,
    CREATE: '/api/budget',
    UPDATE: (uuid: string) => `/api/budget/${uuid}`,
    DELETE: (uuid: string) => `/api/budget/${uuid}`,
    BATCH: '/api/budget/batch',
    SUMMARY: '/api/budget/summary',
    YEARS: '/api/budget/years',
  } as const,
  REPRESENTATIVES: {
    LIST: '/api/v1/representatives',
    GET: (uuid: string) => `/api/v1/representatives/${uuid}`,
    CREATE: '/api/v1/representatives',
    UPDATE: (uuid: string) => `/api/v1/representatives/${uuid}`,
    DELETE: (uuid: string) => `/api/v1/representatives/${uuid}`,
    TABLE: '/api/v1/representatives/table',
    STATS: (uuid: string) => `/api/v1/representatives/${uuid}/stats`,
    PROFILE: (name: string) => `/api/v1/representatives/profile/${name}`,
    PROFILES: '/api/v1/representatives/profiles',
  } as const,
} as const;