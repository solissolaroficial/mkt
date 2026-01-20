import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  Brand,
  BrandListResponse,
  CreateBrandRequest,
} from '@/shared/types/legacy.types';

export const brandService = {
  /**
   * GET /api/brands
   */
  list: async (): Promise<Brand[]> => {
    console.log('[BRAND SERVICE] list called');
    const response = await apiClient.get<BrandListResponse>(ENDPOINTS.BRANDS.LIST);
    console.log('[BRAND SERVICE] list response.data:', response.data);
    return response.data.brands;
  },

  /**
   * POST /api/brands
   */
  create: async (data: CreateBrandRequest): Promise<Brand> => {
    console.log('[BRAND SERVICE] create called with data:', data);
    const response = await apiClient.post<Brand>(ENDPOINTS.BRANDS.CREATE, data);
    console.log('[BRAND SERVICE] create response.data:', response.data);
    return response.data;
  },

  /**
   * DELETE /api/brands/:id
   */
  delete: async (id: string): Promise<void> => {
    console.log('[BRAND SERVICE] delete called with id:', id);
    await apiClient.delete(ENDPOINTS.BRANDS.DELETE(id));
  },
};
