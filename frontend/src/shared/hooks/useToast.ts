import { useToastStore, toastActions, ToastType, Toast } from '@/shared/store/toastStore';
import { useCallback } from 'react';

export const useToast = () => {
  const { toasts, addToast, removeToast, clearToasts } = useToastStore();

  const success = useCallback((message: string, duration?: number) => {
    return addToast('success', message, duration);
  }, [addToast]);

  const error = useCallback((message: string, duration?: number) => {
    return addToast('error', message, duration);
  }, [addToast]);

  const warning = useCallback((message: string, duration?: number) => {
    return addToast('warning', message, duration);
  }, [addToast]);

  const info = useCallback((message: string, duration?: number) => {
    return addToast('info', message, duration);
  }, [addToast]);

  return {
    toasts,
    addToast,
    removeToast,
    clearToasts,
    success,
    error,
    warning,
    info,
  };
};

// Re-exportar tipos e ações não-hook para uso
export type { ToastType, Toast };
export { toastActions };
