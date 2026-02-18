import React from 'react';
import { useToastStore } from '@/shared/store/toastStore';
import { ToastContainer } from './ToastContainer';

export const GlobalToastContainer: React.FC = () => {
  const { toasts, removeToast } = useToastStore();

  return <ToastContainer toasts={toasts} onRemove={removeToast} />;
};
