import { apiClient } from '../../../infrastructure/api/client';
import type { CalendarPost } from '../../../shared/types/legacy.types';

export const calendarService = {
  // Get all calendar posts
  list: async (): Promise<CalendarPost[]> => {
    const response = await apiClient.get('/api/calendar/posts');
    return response.data;
  },

  // Get a single post by id
  getById: async (id: string): Promise<CalendarPost> => {
    const response = await apiClient.get(`/api/calendar/posts/${id}`);
    return response.data;
  },

  // Create new post
  create: async (data: Partial<CalendarPost>): Promise<CalendarPost> => {
    const response = await apiClient.post('/api/calendar/posts', data);
    return response.data;
  },

  // Update post
  update: async (id: string, data: Partial<CalendarPost>): Promise<CalendarPost> => {
    const response = await apiClient.put(`/api/calendar/posts/${id}`, data);
    return response.data;
  },

  // Delete post
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/calendar/posts/${id}`);
  },
};
