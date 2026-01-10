import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  SocialBenchmarking,
  SocialBenchmarkingListResponse,
  SocialBenchmarkingFilters,
  CreateSocialBenchmarkingRequest,
  UpdateSocialBenchmarkingRequest,
} from '@/shared/types/legacy.types';

export const socialService = {
  /**
   * GET /api/social/benchmarking (com filtros e paginação)
   */
  list: async (params?: SocialBenchmarkingFilters): Promise<SocialBenchmarkingListResponse> => {
    console.log('[SOCIAL SERVICE] list called with params:', params);
    const response = await apiClient.get<SocialBenchmarkingListResponse>(
      ENDPOINTS.SOCIAL.LIST,
      { params }
    );
    console.log('[SOCIAL SERVICE] list response.data:', response.data);
    return response.data;
  },

  /**
   * GET /api/social/benchmarking/:id
   */
  getById: async (id: string): Promise<SocialBenchmarking> => {
    console.log('[SOCIAL SERVICE] getById called with id:', id);
    const response = await apiClient.get<SocialBenchmarking>(ENDPOINTS.SOCIAL.GET(id));
    console.log('[SOCIAL SERVICE] getById response.data:', response.data);
    return response.data;
  },

  /**
   * POST /api/social/benchmarking
   */
  create: async (data: CreateSocialBenchmarkingRequest): Promise<SocialBenchmarking> => {
    console.log('[SOCIAL SERVICE] create called with data:', data);
    const response = await apiClient.post<SocialBenchmarking>(ENDPOINTS.SOCIAL.CREATE, data);
    console.log('[SOCIAL SERVICE] create response.data:', response.data);
    return response.data;
  },

  /**
   * PUT /api/social/benchmarking/:id
   */
  update: async (id: string, data: UpdateSocialBenchmarkingRequest): Promise<SocialBenchmarking> => {
    console.log('[SOCIAL SERVICE] update called with id:', id, 'data:', data);
    const response = await apiClient.put<SocialBenchmarking>(ENDPOINTS.SOCIAL.UPDATE(id), data);
    console.log('[SOCIAL SERVICE] update response.data:', response.data);
    return response.data;
  },

  /**
   * DELETE /api/social/benchmarking/:id
   */
  delete: async (id: string): Promise<void> => {
    console.log('[SOCIAL SERVICE] delete called with id:', id);
    await apiClient.delete(ENDPOINTS.SOCIAL.DELETE(id));
  },
};
