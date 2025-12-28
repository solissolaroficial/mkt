import type { BudgetItem } from '@/shared/types';

export type { BudgetItem };

export interface BudgetViewProps {
  onBack: () => void;
  filterFn?: (item: BudgetItem) => boolean;
  customTitle?: string;
  customSubtitle?: string;
}

export interface EditingCell {
  id: string;
  colIndex: number;
  type: 'budget' | 'realized';
}

export interface BudgetGroup {
  id: string;
  codObj: string;
  name: string;
  items: BudgetItem[];
  subgroups: Record<string, BudgetSubgroup>;
  totals: number[];
  realizedTotals: number[];
}

export interface BudgetSubgroup {
  id: string;
  codGrp: string;
  name: string;
  items: BudgetItem[];
  totals: number[];
  realizedTotals: number[];
}

export interface MarketingBudgetViewProps {
  onBack: () => void;
  selectedMonth: string;
}

export interface BudgetSummaryItem {
  id: string;
  name: string;
  budget: number;
  realized: number;
  requested: number;
  available: number;
  percentUsed: number;
}
