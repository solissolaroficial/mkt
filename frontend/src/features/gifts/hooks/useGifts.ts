import { useQuery } from '@tanstack/react-query';
import { giftsService } from '../services/giftsService';
import type {
  GiftItem,
  GiftTransaction,
  GiftItemListResponse,
  GiftTransactionListResponse,
  GiftItemFilters,
  GiftTransactionFilters,
} from '../types';

// ==================== GIFT ITEMS ====================

/**
 * Hook to get all gift items
 */
export const useGiftItems = (filters?: GiftItemFilters) => {
  return useQuery({
    queryKey: ['gifts', 'items', filters],
    queryFn: () => giftsService.getItems({
      ...filters,
      page: filters?.page || 1,
      limit: filters?.limit || 50,
    }),
    staleTime: 1000 * 60 * 10, // 10 minutos
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

/**
 * Hook to get gift item by ID
 */
export const useGiftItem = (id: string) => {
  return useQuery({
    queryKey: ['gifts', 'item', id],
    queryFn: () => giftsService.getItemById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    refetchOnWindowFocus: false,
  });
};

// ==================== GIFT TRANSACTIONS ====================

/**
 * Hook to get all gift transactions
 */
export const useGiftTransactions = (filters?: GiftTransactionFilters) => {
  return useQuery({
    queryKey: ['gifts', 'transactions', filters],
    queryFn: () => giftsService.getTransactions({
      ...filters,
      page: filters?.page || 1,
      limit: filters?.limit || 50,
    }),
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    placeholderData: (previousData) => previousData,
    refetchOnWindowFocus: false,
  });
};

/**
 * Hook to get gift transaction by ID
 */
export const useGiftTransaction = (id: string) => {
  return useQuery({
    queryKey: ['gifts', 'transaction', id],
    queryFn: () => giftsService.getTransactionById(id),
    enabled: !!id,
    staleTime: 1000 * 60 * 10,
    retry: 3,
    retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
    refetchOnWindowFocus: false,
  });
};
