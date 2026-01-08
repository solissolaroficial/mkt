import { apiClient } from '../../../infrastructure/api/client';
import { ENDPOINTS } from '../../../infrastructure/api/endpoints';
import type {
  CalendarPost,
  CalendarPostsListResponse,
  CalendarFilters,
  CalendarPostRequest,
  UpdateStatusRequest,
} from '../../../shared/types/legacy.types';

export const calendarService = {
  // GET /api/calendar/posts (com filtros e paginação)
  list: async (params?: CalendarFilters): Promise<CalendarPostsListResponse> => {
    console.log('[CALENDAR SERVICE] list called with params:', params);
    const response = await apiClient.get<CalendarPostsListResponse>(
      ENDPOINTS.CALENDAR.LIST,
      { params }
    );
    console.log('[CALENDAR SERVICE] list response.data:', response.data);
    console.log('[CALENDAR SERVICE] list response.data.data:', response.data?.data);
    console.log('[CALENDAR SERVICE] First post date:', response.data?.data?.[0]?.date);
    return response.data;
  },

  // GET /api/calendar/posts/:id
  getById: async (id: string): Promise<CalendarPost> => {
    console.log('[CALENDAR SERVICE] getById called with id:', id);
    const response = await apiClient.get<CalendarPost>(ENDPOINTS.CALENDAR.GET(id));
    console.log('[CALENDAR SERVICE] getById response.data:', response.data);
    console.log('[CALENDAR SERVICE] getById response.data.date:', response.data?.date);
    return response.data;
  },

  // POST /api/calendar/posts
  create: async (data: CalendarPostRequest): Promise<CalendarPost> => {
    const response = await apiClient.post<CalendarPost>(ENDPOINTS.CALENDAR.CREATE, data);
    return response.data;
  },

  // PUT /api/calendar/posts/:id
  update: async (id: string, data: Partial<CalendarPostRequest>): Promise<CalendarPost> => {
    const response = await apiClient.put<CalendarPost>(ENDPOINTS.CALENDAR.UPDATE(id), data);
    return response.data;
  },

  // PUT /api/calendar/posts/:id/status (NOVO)
  updateStatus: async (
    id: string,
    request: UpdateStatusRequest
  ): Promise<CalendarPost> => {
    const response = await apiClient.put<CalendarPost>(
      ENDPOINTS.CALENDAR.UPDATE_STATUS(id),
      request
    );
    return response.data;
  },

  // POST /api/calendar/posts/:id/confirm-publishing (NOVO)
  confirmPublishing: async (id: string, platforms: string[]): Promise<CalendarPost> => {
    const response = await apiClient.post<CalendarPost>(
      ENDPOINTS.CALENDAR.CONFIRM_PUBLISHING(id),
      { published_platforms: platforms }
    );
    return response.data;
  },

  // DELETE /api/calendar/posts/:id
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(ENDPOINTS.CALENDAR.DELETE(id));
  },
};
