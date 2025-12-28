import type { GiftItem, GiftTransaction, Notification } from '@/shared/types';

export type { GiftItem, GiftTransaction, Notification };

export interface GiftsViewProps {
  onAddNotification?: (notification: Notification) => void;
}

export type TransactionTab = 'out' | 'in';
