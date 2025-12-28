import type { AccountPayable } from '@/shared/types';

export type { AccountPayable };

export type AccountStatus = 'pending' | 'sent_to_finance';

export type RecurrenceType = 'none' | 'monthly' | 'yearly';

export interface FinancialViewProps {
  onBack?: () => void;
}

export interface EmailFormData {
  to: string;
  subject: string;
  body: string;
}

export interface PaymentFormData {
  supplier: string;
  description: string;
  amount: string;
  dueDate: string;
  recurrence: RecurrenceType;
}
