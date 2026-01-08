import { apiClient } from '../../../infrastructure/api/client';
import { ENDPOINTS } from '../../../infrastructure/api/endpoints';
import type {
  PdvPost,
  RecurrentPdv,
  PdvPostListResponse,
  RecurrentPdvListResponse,
  PdvPostFilters,
  RecurrentPdvFilters,
  CreatePdvPostRequest,
  UpdatePdvPostRequest,
  CreateRecurrentPdvRequest,
  UpdateRecurrentPdvRequest,
} from '../../../shared/types/legacy.types';

export const pdvService = {
  // ============================================
  // Posts PDV
  // ============================================

  // GET /api/pdv-posts (com filtros e paginação)
  list: async (params?: PdvPostFilters): Promise<PdvPostListResponse> => {
    console.log('[PDV SERVICE] list called with params:', params);
    const response = await apiClient.get<PdvPostListResponse>(
      ENDPOINTS.PDV.LIST,
      { params }
    );
    console.log('[PDV SERVICE] list response.data:', response.data);
    return response.data;
  },

  // GET /api/pdv-posts/:id
  getById: async (id: string): Promise<PdvPost> => {
    console.log('[PDV SERVICE] getById called with id:', id);
    const response = await apiClient.get<PdvPost>(ENDPOINTS.PDV.GET(id));
    console.log('[PDV SERVICE] getById response.data:', response.data);
    return response.data;
  },

  // POST /api/pdv-posts
  create: async (data: CreatePdvPostRequest): Promise<PdvPost> => {
    console.log('[PDV SERVICE] create called with data:', data);
    const response = await apiClient.post<PdvPost>(ENDPOINTS.PDV.CREATE, data);
    console.log('[PDV SERVICE] create response.data:', response.data);
    return response.data;
  },

  // PUT /api/pdv-posts/:id
  update: async (id: string, data: UpdatePdvPostRequest): Promise<PdvPost> => {
    console.log('[PDV SERVICE] update called with id:', id, 'data:', data);
    const response = await apiClient.put<PdvPost>(ENDPOINTS.PDV.UPDATE(id), data);
    console.log('[PDV SERVICE] update response.data:', response.data);
    return response.data;
  },

  // PATCH /api/pdv/posts/:id/status
  updateStatus: async (
    id: string,
    status: 'pending' | 'approved' | 'rejected' | 'published' | 'cancelled'
  ): Promise<PdvPost> => {
    console.log('[PDV SERVICE] updateStatus called with id:', id, 'status:', status);
    const response = await apiClient.patch<PdvPost>(
      ENDPOINTS.PDV.UPDATE_STATUS(id),
      { status }
    );
    console.log('[PDV SERVICE] updateStatus response.data:', response.data);
    return response.data;
  },

  // DELETE /api/pdv-posts/:id
  delete: async (id: string): Promise<void> => {
    console.log('[PDV SERVICE] delete called with id:', id);
    await apiClient.delete(ENDPOINTS.PDV.DELETE(id));
  },

  // ============================================
  // PDVs Recorrentes
  // ============================================

  // GET /api/recurrent-pdvs (com filtros e paginação)
  listRecurrent: async (params?: RecurrentPdvFilters): Promise<RecurrentPdvListResponse> => {
    console.log('[PDV SERVICE] listRecurrent called with params:', params);
    const response = await apiClient.get<RecurrentPdvListResponse>(
      ENDPOINTS.PDV.LIST_RECURRENT,
      { params }
    );
    console.log('[PDV SERVICE] listRecurrent response.data:', response.data);
    return response.data;
  },

  // GET /api/recurrent-pdvs/:id
  getRecurrentById: async (id: string): Promise<RecurrentPdv> => {
    console.log('[PDV SERVICE] getRecurrentById called with id:', id);
    const response = await apiClient.get<RecurrentPdv>(ENDPOINTS.PDV.GET_RECURRENT(id));
    console.log('[PDV SERVICE] getRecurrentById response.data:', response.data);
    return response.data;
  },

  // POST /api/recurrent-pdvs
  createRecurrent: async (data: CreateRecurrentPdvRequest): Promise<RecurrentPdv> => {
    console.log('[PDV SERVICE] createRecurrent called with data:', data);
    const response = await apiClient.post<RecurrentPdv>(ENDPOINTS.PDV.CREATE_RECURRENT, data);
    console.log('[PDV SERVICE] createRecurrent response.data:', response.data);
    return response.data;
  },

  // PUT /api/recurrent-pdvs/:id
  updateRecurrent: async (id: string, data: UpdateRecurrentPdvRequest): Promise<RecurrentPdv> => {
    console.log('[PDV SERVICE] updateRecurrent called with id:', id, 'data:', data);
    const response = await apiClient.put<RecurrentPdv>(ENDPOINTS.PDV.UPDATE_RECURRENT(id), data);
    console.log('[PDV SERVICE] updateRecurrent response.data:', response.data);
    return response.data;
  },

  // DELETE /api/recurrent-pdvs/:id
  deleteRecurrent: async (id: string): Promise<void> => {
    console.log('[PDV SERVICE] deleteRecurrent called with id:', id);
    await apiClient.delete(ENDPOINTS.PDV.DELETE_RECURRENT(id));
  },
};
