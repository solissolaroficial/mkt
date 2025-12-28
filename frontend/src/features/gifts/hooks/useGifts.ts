import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { giftsService } from '../services/giftsService';
import { GIFT_ITEMS, GIFT_TRANSACTIONS_MOCK, REP_NAMES } from '@/shared/utils/legacy.constants';
import type { GiftItem, GiftTransaction, Notification } from '@/shared/types';

/**
 * Hook to get all gift items
 */
export const useGiftItems = () => {
  return useQuery({
    queryKey: ['gifts', 'items'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return GIFT_ITEMS;
    },
    staleTime: 1000 * 60 * 5,
  });
};

/**
 * Hook to get all gift transactions
 */
export const useGiftTransactions = () => {
  return useQuery({
    queryKey: ['gifts', 'transactions'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return GIFT_TRANSACTIONS_MOCK;
    },
    staleTime: 1000 * 60 * 5,
  });
};

/**
 * Hook to create a new gift item
 */
export const useCreateGiftItem = () => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: giftsService.createItem,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'items'] });
    },
  });
};

/**
 * Hook to create a new transaction
 */
export const useCreateTransaction = () => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: giftsService.createTransaction,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['gifts', 'transactions'] });
    },
  });
};
