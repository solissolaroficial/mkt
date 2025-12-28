import { useQuery } from '@tanstack/react-query';
import { settingsService } from '../services/settingsService';
import { PROGRAM_CREDENTIALS, INTERNAL_CONTACTS } from '@/shared/utils/legacy.constants';

/**
 * Hook to get program credentials
 */
export const useCredentials = () => {
  return useQuery({
    queryKey: ['settings', 'credentials'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return PROGRAM_CREDENTIALS;
    },
    staleTime: 1000 * 60 * 10, // 10 minutes
  });
};

/**
 * Hook to get internal contacts
 */
export const useContacts = () => {
  return useQuery({
    queryKey: ['settings', 'contacts'],
    queryFn: async () => {
      // TODO: Replace with actual API call when backend is ready
      return INTERNAL_CONTACTS;
    },
    staleTime: 1000 * 60 * 10,
  });
};
