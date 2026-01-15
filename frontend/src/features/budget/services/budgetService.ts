import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  BudgetItem,
  CreateBudgetItemRequest,
  UpdateBudgetItemRequest,
  BudgetItemListResponse,
  BudgetSummaryResponse,
  BudgetYearsResponse,
  BudgetItemFilters,
} from '../types';

const DEBUG = import.meta.env.DEV;

export const budgetService = {
  // ==================== CRUD ====================

  /**
   * Get all budget items with filters and pagination
   */
  list: async (filters?: BudgetItemFilters): Promise<BudgetItemListResponse> => {
    if (DEBUG) console.log('[BUDGET SERVICE] list called with filters:', filters);
    const response = await apiClient.get<BudgetItemListResponse>(
      ENDPOINTS.BUDGET.LIST,
      { params: filters }
    );
    if (DEBUG) console.log('[BUDGET SERVICE] list response.data:', response.data);
    return response.data;
  },

  /**
   * Get budget item by UUID
   */
  getById: async (uuid: string): Promise<BudgetItem> => {
    if (DEBUG) console.log('[BUDGET SERVICE] getById called with uuid:', uuid);
    const response = await apiClient.get<BudgetItem>(
      ENDPOINTS.BUDGET.GET(uuid)
    );
    if (DEBUG) console.log('[BUDGET SERVICE] getById response.data:', response.data);
    return response.data;
  },

  /**
   * Create a new budget item
   */
  create: async (data: CreateBudgetItemRequest): Promise<BudgetItem> => {
    if (DEBUG) console.log('[BUDGET SERVICE] create called with data:', data);
    const response = await apiClient.post<BudgetItem>(
      ENDPOINTS.BUDGET.CREATE,
      data
    );
    if (DEBUG) console.log('[BUDGET SERVICE] create response.data:', response.data);
    return response.data;
  },

  /**
   * Update budget item
   */
  update: async (uuid: string, data: UpdateBudgetItemRequest): Promise<BudgetItem> => {
    if (DEBUG) console.log('[BUDGET SERVICE] update called with uuid:', uuid, 'data:', data);
    const response = await apiClient.put<BudgetItem>(
      ENDPOINTS.BUDGET.UPDATE(uuid),
      data
    );
    if (DEBUG) console.log('[BUDGET SERVICE] update response.data:', response.data);
    return response.data;
  },

  /**
   * Delete budget item (soft delete)
   */
  delete: async (uuid: string): Promise<void> => {
    if (DEBUG) console.log('[BUDGET SERVICE] delete called with uuid:', uuid);
    await apiClient.delete(ENDPOINTS.BUDGET.DELETE(uuid));
  },

  // ==================== BATCH OPERATIONS ====================

  /**
   * Create multiple budget items in batch
   */
  batchCreate: async (items: CreateBudgetItemRequest[]): Promise<BudgetItem[]> => {
    if (DEBUG) console.log('[BUDGET SERVICE] batchCreate called with items count:', items.length);
    const response = await apiClient.post<BudgetItem[]>(
      ENDPOINTS.BUDGET.BATCH,
      { items }
    );
    if (DEBUG) console.log('[BUDGET SERVICE] batchCreate response.data:', response.data);
    return response.data;
  },

  // ==================== ADDITIONAL ENDPOINTS ====================

  /**
   * Get budget summary
   */
  getSummary: async (filters?: BudgetItemFilters): Promise<BudgetSummaryResponse[]> => {
    if (DEBUG) console.log('[BUDGET SERVICE] getSummary called with filters:', filters);
    const response = await apiClient.get<BudgetSummaryResponse[]>(
      ENDPOINTS.BUDGET.SUMMARY,
      { params: filters }
    );
    if (DEBUG) console.log('[BUDGET SERVICE] getSummary response.data:', response.data);
    return response.data;
  },

  /**
   * Get distinct years
   */
  getDistinctYears: async (): Promise<number[]> => {
    if (DEBUG) console.log('[BUDGET SERVICE] getDistinctYears called');
    const response = await apiClient.get<BudgetYearsResponse>(
      ENDPOINTS.BUDGET.YEARS
    );
    if (DEBUG) console.log('[BUDGET SERVICE] getDistinctYears response.data:', response.data);
    return response.data.years;
  },
};
