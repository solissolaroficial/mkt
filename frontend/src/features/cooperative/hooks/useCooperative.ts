import { useQuery } from '@tanstack/react-query';
import { cooperativeService } from '../services/cooperativeService';
import type { OfflineAction, ShowroomItem } from '../types';
import { OFFLINE_ACTIONS_DATA, MOCK_SHOWROOM_ITEMS } from '@/shared/utils/legacy.constants';

export const useOfflineActions = () => {
  return useQuery({
    queryKey: ['offline-actions'],
    queryFn: async () => {
      try {
        return await cooperativeService.getOfflineActions();
      } catch (error) {
        console.warn('Using mock data for offline actions:', error);
        return OFFLINE_ACTIONS_DATA as OfflineAction[];
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};

export const useShowroomItems = () => {
  return useQuery({
    queryKey: ['showroom-items'],
    queryFn: async () => {
      try {
        return await cooperativeService.getShowroomItems();
      } catch (error) {
        console.warn('Using mock data for showroom items:', error);
        return MOCK_SHOWROOM_ITEMS as ShowroomItem[];
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};
