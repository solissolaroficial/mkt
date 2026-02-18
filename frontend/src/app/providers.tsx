import React from 'react';
import { QueryClient, QueryCache, MutationCache } from '@tanstack/react-query';
import { QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider } from 'react-router-dom';
import { toastActions } from '@/shared/store/toastStore';
import { GlobalToastContainer } from '@/shared/components/GlobalToastContainer';
import { router } from '@/infrastructure/router/routes';

// Configuração do React Query
const queryClient = new QueryClient({
  queryCache: new QueryCache({
    onError: (error, query) => {
      console.error('[Query Error]', error, query.queryKey);

      // Mostrar toast apenas para erros em background
      if (query.state.data !== undefined) {
        const errorMessage = error instanceof Error
          ? error.message
          : 'Erro ao atualizar dados';
        toastActions.error(errorMessage);
      }
    },
  }),
  mutationCache: new MutationCache({
    onError: (error) => {
      console.error('[Mutation Error]', error);
      const errorMessage = error instanceof Error
        ? error.message
        : 'Operação falhou';
      toastActions.error(errorMessage);
    },
  }),
  defaultOptions: {
    queries: {
      retry: (failureCount, error) => {
        if (error instanceof Error && 'response' in error) {
          const status = (error as any).response?.status;
          if (status && status >= 400 && status < 500) {
            return false;
          }
        }
        return failureCount < 3;
      },
      refetchOnWindowFocus: false,
      staleTime: 1000 * 60 * 5,
    },
    mutations: {
      retry: false,
    },
  },
});

export const AppProviders: React.FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <GlobalToastContainer />
      <RouterProvider router={router} />
    </QueryClientProvider>
  );
};
