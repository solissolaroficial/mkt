/**
 * Monthly Goal Types for Representatives
 */

export interface RepresentativeMonthlyGoal {
  id: string;
  representativeId: string;
  month: number;
  year: number;
  target: number;
  realized: number;
  percentage: number;
  remaining?: number;
  isTargetMet?: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface RepresentativeMonthlyGoalData {
  id: string;
  representativeId: string;
  month: number;
  year: number;
  target: number;
  realized: number;
  percentage: number;
  remaining: number;
  isTargetMet: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CreateRepresentativeMonthlyGoalRequest {
  representativeId: string;
  month: number;
  year: number;
  target: number;
}

export interface UpdateRepresentativeMonthlyGoalRequest {
  target?: number;
  realized?: number;
}

export interface ListRepresentativeMonthlyGoalsRequest {
  representativeId?: string;
  month?: number;
  year?: number;
  page?: number;
  pageSize?: number;
  sortBy?: string;
  sortOrder?: string;
}

export interface ListRepresentativeMonthlyGoalsResponse {
  data: RepresentativeMonthlyGoalData[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}

export interface GetRepresentativeGoalsTableDataRequest {
  year?: number;
  month?: number;
}

export interface MonthData {
  month: number;
  monthName: string;
  targetSum: number;
  realizedSum: number;
}

export interface MonthValue {
  month: number;
  target: number;
  realized: number;
  percentage: number;
  isMet: boolean;
}

export interface RepresentativeRowData {
  representativeId: string;
  representativeName: string;
  company: string;
  region: string;
  city: string;
  monthValues: MonthValue[];
  totalTarget: number;
  totalRealized: number;
  totalPercentage: number;
}

export interface TableSummaryData {
  totalRepresentatives: number;
  totalTarget: number;
  totalRealized: number;
  overallPercentage: number;
  goalsMetCount: number;
  goalsNotMetCount: number;
}

export interface GetRepresentativeGoalsTableDataResponse {
  year: number;
  months: MonthData[];
  rows: RepresentativeRowData[];
  summary: TableSummaryData;
}
