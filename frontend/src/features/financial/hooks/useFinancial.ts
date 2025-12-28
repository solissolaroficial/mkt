import { useQuery } from '@tanstack/react-query';
import { financialService } from '../services/financialService';
import type { AccountPayable } from '../types';
import { MOCK_ACCOUNTS_PAYABLE } from '@/shared/utils/legacy.constants';

export const useAccountsPayable = () => {
  return useQuery({
    queryKey: ['accounts-payable'],
    queryFn: async () => {
      try {
        return await financialService.list();
      } catch (error) {
        console.warn('Using mock data for accounts payable:', error);
        return MOCK_ACCOUNTS_PAYABLE as AccountPayable[];
      }
    },
    staleTime: 1000 * 60 * 5,
  });
};
