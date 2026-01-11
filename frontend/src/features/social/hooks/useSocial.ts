import { useQuery } from '@tanstack/react-query';
import { socialService } from '../services/socialService';
import type {
  SocialBenchmarkingFilters,
  SocialPostFilters,
  SocialDailyAggregationFilters,
} from '@/shared/types/legacy.types';

/**
 * Hook para lista de social benchmarkings (com filtros)
 */
export const useSocialBenchmarkings = (params?: SocialBenchmarkingFilters) => {
  return useQuery({
    queryKey: ['social-benchmarkings', params],
    queryFn: () => socialService.list(params),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

/**
 * Hook para social benchmarking individual
 */
export const useSocialBenchmarking = (id: string) => {
  return useQuery({
    queryKey: ['social-benchmarking', id],
    queryFn: () => socialService.getById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

// ============================================
// Social Posts Hooks
// ============================================

/**
 * Hook para lista de social posts (com filtros)
 */
export const useSocialPosts = (params?: SocialPostFilters) => {
  return useQuery({
    queryKey: ['social-posts', params],
    queryFn: () => socialService.listPosts(params),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

/**
 * Hook para social post individual
 */
export const useSocialPost = (id: string) => {
  return useQuery({
    queryKey: ['social-post', id],
    queryFn: () => socialService.getPostById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

// ============================================
// Social Daily Aggregations Hooks
// ============================================

/**
 * Hook para lista de social daily aggregations (com filtros)
 */
export const useSocialDailyAggregations = (params?: SocialDailyAggregationFilters) => {
  return useQuery({
    queryKey: ['social-daily-aggregations', params],
    queryFn: () => socialService.listAggregations(params),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

/**
 * Hook para social daily aggregation individual
 */
export const useSocialDailyAggregation = (id: string) => {
  return useQuery({
    queryKey: ['social-daily-aggregation', id],
    queryFn: () => socialService.getAggregationById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};
