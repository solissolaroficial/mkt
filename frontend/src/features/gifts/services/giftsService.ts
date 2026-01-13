import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  GiftItem,
  GiftTransaction,
  GiftItemListResponse,
  GiftTransactionListResponse,
  GiftItemFilters,
  GiftTransactionFilters,
  CreateGiftItemRequest,
  UpdateGiftItemRequest,
  CreateGiftTransactionRequest,
  UpdateGiftTransactionRequest,
} from '../types';

const DEBUG = import.meta.env.DEV;

export const giftsService = {
  // ==================== GIFT ITEMS ====================

  /**
   * Get all gift items with filters and pagination
   */
  getItems: async (params?: GiftItemFilters): Promise<GiftItemListResponse> => {
    if (DEBUG) console.log('[GIFTS SERVICE] getItems called with params:', params);
    const response = await apiClient.get<GiftItemListResponse>(
      ENDPOINTS.GIFTS.ITEMS_LIST,
      { params }
    );
    if (DEBUG) console.log('[GIFTS SERVICE] getItems response.data:', response.data);
    return response.data;
  },

  /**
   * Get gift item by ID
   */
  getItemById: async (id: string): Promise<GiftItem> => {
    if (DEBUG) console.log('[GIFTS SERVICE] getItemById called with id:', id);
    const response = await apiClient.get<GiftItem>(
      ENDPOINTS.GIFTS.ITEMS_GET(id)
    );
    if (DEBUG) console.log('[GIFTS SERVICE] getItemById response.data:', response.data);
    return response.data;
  },

  /**
   * Create a new gift item
   */
  createItem: async (data: CreateGiftItemRequest): Promise<GiftItem> => {
    if (DEBUG) console.log('[GIFTS SERVICE] createItem called with data:', data);
    const response = await apiClient.post<GiftItem>(
      ENDPOINTS.GIFTS.ITEMS_CREATE,
      data
    );
    if (DEBUG) console.log('[GIFTS SERVICE] createItem response.data:', response.data);
    return response.data;
  },

  /**
   * Update gift item
   */
  updateItem: async (id: string, data: UpdateGiftItemRequest): Promise<GiftItem> => {
    if (DEBUG) console.log('[GIFTS SERVICE] updateItem called with id:', id, 'data:', data);
    const response = await apiClient.put<GiftItem>(
      ENDPOINTS.GIFTS.ITEMS_UPDATE(id),
      data
    );
    if (DEBUG) console.log('[GIFTS SERVICE] updateItem response.data:', response.data);
    return response.data;
  },

  /**
   * Delete gift item (soft delete)
   */
  deleteItem: async (id: string): Promise<void> => {
    if (DEBUG) console.log('[GIFTS SERVICE] deleteItem called with id:', id);
    await apiClient.delete(ENDPOINTS.GIFTS.ITEMS_DELETE(id));
  },

  // ==================== GIFT TRANSACTIONS ====================

  /**
   * Get all gift transactions with filters and pagination
   */
  getTransactions: async (params?: GiftTransactionFilters): Promise<GiftTransactionListResponse> => {
    if (DEBUG) console.log('[GIFTS SERVICE] getTransactions called with params:', params);
    const response = await apiClient.get<GiftTransactionListResponse>(
      ENDPOINTS.GIFTS.TRANSACTIONS_LIST,
      { params }
    );
    if (DEBUG) console.log('[GIFTS SERVICE] getTransactions response.data:', response.data);
    return response.data;
  },

  /**
   * Get gift transaction by ID
   */
  getTransactionById: async (id: string): Promise<GiftTransaction> => {
    if (DEBUG) console.log('[GIFTS SERVICE] getTransactionById called with id:', id);
    const response = await apiClient.get<GiftTransaction>(
      ENDPOINTS.GIFTS.TRANSACTIONS_GET(id)
    );
    if (DEBUG) console.log('[GIFTS SERVICE] getTransactionById response.data:', response.data);
    return response.data;
  },

  /**
   * Create a new gift transaction
   */
  createTransaction: async (data: CreateGiftTransactionRequest): Promise<GiftTransaction> => {
    if (DEBUG) console.log('[GIFTS SERVICE] createTransaction called with data:', data);
    const response = await apiClient.post<GiftTransaction>(
      ENDPOINTS.GIFTS.TRANSACTIONS_CREATE,
      data
    );
    if (DEBUG) console.log('[GIFTS SERVICE] createTransaction response.data:', response.data);
    return response.data;
  },

  /**
   * Update gift transaction
   */
  updateTransaction: async (id: string, data: UpdateGiftTransactionRequest): Promise<GiftTransaction> => {
    if (DEBUG) console.log('[GIFTS SERVICE] updateTransaction called with id:', id, 'data:', data);
    const response = await apiClient.put<GiftTransaction>(
      ENDPOINTS.GIFTS.TRANSACTIONS_UPDATE(id),
      data
    );
    if (DEBUG) console.log('[GIFTS SERVICE] updateTransaction response.data:', response.data);
    return response.data;
  },

  /**
   * Delete gift transaction (soft delete)
   */
  deleteTransaction: async (id: string): Promise<void> => {
    if (DEBUG) console.log('[GIFTS SERVICE] deleteTransaction called with id:', id);
    await apiClient.delete(ENDPOINTS.GIFTS.TRANSACTIONS_DELETE(id));
  },
};
