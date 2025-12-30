import { useQuery } from '@tanstack/react-query';
import { marketingService } from '../services/marketingService';
import { MARKETING_CHANNELS_DATA, ANNUAL_CHANNEL_DATA } from '@/shared/utils/legacy.constants';
import { kpiService } from '@/features/kpis/services/kpiService';
import { MARKETING_KPIS_SLUGS, QUERY_KEYS } from '@/shared/utils/constants';

/**
 * Hook to get channel data for a specific month or annual data
 */
export const useChannelData = (selectedMonth: string) => {
  return useQuery({
    queryKey: ['marketing', 'channels', selectedMonth],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      // For now, use mock data
      if (selectedMonth === 'ALL') {
        return ANNUAL_CHANNEL_DATA;
      }
      return MARKETING_CHANNELS_DATA[selectedMonth] || [];
    },
    staleTime: 1000 * 60 * 5, // 5 minutes
  });
};

/**
 * Hook to get all marketing channels data
 */
export const useAllChannelData = () => {
  return useQuery({
    queryKey: ['marketing', 'channels', 'all'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return MARKETING_CHANNELS_DATA;
    },
    staleTime: 1000 * 60 * 5,
  });
};

/**
 * Hook to get Marketing KPIs
 */
export const useMarketingKpis = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: QUERY_KEYS.KPIS.BY_SLUGS(MARKETING_KPIS_SLUGS),
    queryFn: () => kpiService.getBySlugs(MARKETING_KPIS_SLUGS),
    staleTime: 1000 * 60 * 5, // 5 minutos
  });

  return {
    marketingKpis: data || [],
    isLoading,
    error
  };
};
