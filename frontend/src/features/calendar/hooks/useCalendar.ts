import { useQuery } from '@tanstack/react-query';
import { calendarService } from '../services/calendarService';
import type { CalendarPost } from '../../../shared/types/legacy.types';
import { MOCK_CALENDAR_POSTS } from '../../../shared/utils/legacy.constants';

export const useCalendarPosts = () => {
  return useQuery({
    queryKey: ['calendar-posts'],
    queryFn: async () => {
      try {
        return await calendarService.list();
      } catch (error) {
        console.warn('Using mock data for calendar posts:', error);
        return MOCK_CALENDAR_POSTS;
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};

export const useCalendarPost = (id: string) => {
  return useQuery({
    queryKey: ['calendar-post', id],
    queryFn: () => calendarService.getById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 5,
  });
};
