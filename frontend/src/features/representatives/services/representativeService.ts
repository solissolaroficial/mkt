import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  Representative,
  RepresentativeStats,
  CreateRepresentativeRequest,
  UpdateRepresentativeRequest,
  ListRepresentativesRequest,
  ListRepresentativesResponse,
  RepresentativeTableData,
} from '../types';

const DEBUG = import.meta.env.DEV;

export const representativeService = {
  // ==================== CRUD ====================

  /**
   * Get all representatives with filters and pagination
   */
  list: async (filters?: ListRepresentativesRequest): Promise<ListRepresentativesResponse> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] list called with filters:', filters);
    const response = await apiClient.get<any>(
      ENDPOINTS.REPRESENTATIVES.LIST,
      { params: filters }
    );
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] list response.data:', response.data);
    // Map backend response to frontend type
    return {
      data: response.data.data || [],
      total: response.data.total || 0,
      page: response.data.page || 1,
      pageSize: response.data.pageSize || 10,
      totalPages: response.data.totalPages || 1,
    };
  },

  /**
   * Get representative by UUID
   */
  getById: async (uuid: string): Promise<Representative> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] getById called with uuid:', uuid);
    const response = await apiClient.get<Representative>(
      ENDPOINTS.REPRESENTATIVES.GET(uuid)
    );
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] getById response.data:', response.data);
    return response.data;
  },

  /**
   * Create a new representative
   */
  create: async (data: CreateRepresentativeRequest): Promise<Representative> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] create called with data:', data);
    const response = await apiClient.post<Representative>(
      ENDPOINTS.REPRESENTATIVES.CREATE,
      data
    );
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] create response.data:', response.data);
    return response.data;
  },

  /**
   * Update representative
   */
  update: async (uuid: string, data: UpdateRepresentativeRequest): Promise<Representative> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] update called with uuid:', uuid, 'data:', data);
    const response = await apiClient.put<Representative>(
      ENDPOINTS.REPRESENTATIVES.UPDATE(uuid),
      data
    );
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] update response.data:', response.data);
    return response.data;
  },

  /**
   * Delete representative (soft delete)
   */
  delete: async (uuid: string): Promise<void> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] delete called with uuid:', uuid);
    await apiClient.delete(ENDPOINTS.REPRESENTATIVES.DELETE(uuid));
  },

  // ==================== ADDITIONAL ENDPOINTS ====================

  /**
   * Get table data (same as list but for table component)
   */
  getTableData: async (filters?: ListRepresentativesRequest): Promise<RepresentativeTableData> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] getTableData called with filters:', filters);
    const response = await apiClient.get<any>(
      ENDPOINTS.REPRESENTATIVES.TABLE,
      { params: filters }
    );
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] getTableData response.data:', response.data);
    // Map backend response to frontend type
    return {
      data: response.data.data || [],
      total: response.data.total || 0,
      page: response.data.page || 1,
      pageSize: response.data.pageSize || 10,
      totalPages: response.data.totalPages || 1,
    };
  },

  /**
   * Get representative statistics
   */
  getStats: async (uuid: string): Promise<RepresentativeStats> => {
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] getStats called with uuid:', uuid);
    const response = await apiClient.get<RepresentativeStats>(
      ENDPOINTS.REPRESENTATIVES.STATS(uuid)
    );
    if (DEBUG) console.log('[REPRESENTATIVE SERVICE] getStats response.data:', response.data);
    return response.data;
  },
};
