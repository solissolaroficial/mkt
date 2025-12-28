export { default as GiftsView } from './ui/GiftsView';

export { useGiftItems, useGiftTransactions, useCreateGiftItem, useCreateTransaction } from './hooks/useGifts';

export type { 
  GiftItem, 
  GiftTransaction, 
  Notification,
  GiftsViewProps,
  TransactionTab
} from './types';
