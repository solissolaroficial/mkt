import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import { useEffect } from 'react';
import { useToast } from './useToast';

interface UseQueryWithErrorOptions<TData, TError> extends UseQueryOptions<TData, TError> {
  errorMessage?: string;
}

export function useQueryWithError<TData = unknown, TError = Error>({
  errorMessage,
  ...options
}: UseQueryWithErrorOptions<TData, TError>) {
  const query = useQuery<TData, TError>(options);
  const toast = useToast();

  useEffect(() => {
    if (query.isError && errorMessage) {
      console.error(`[Query Error] ${errorMessage}:`, query.error);
      toast.error(errorMessage);
    }
  }, [query.isError, errorMessage, toast]);

  return query;
}
