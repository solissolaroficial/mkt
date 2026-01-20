import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  SocialBenchmarking,
  SocialBenchmarkingListResponse,
  SocialBenchmarkingFilters,
  CreateSocialBenchmarkingRequest,
  UpdateSocialBenchmarkingRequest,
  SocialPost,
  SocialPostListResponse,
  SocialPostFilters,
  CreateSocialPostRequest,
  UpdateSocialPostRequest,
  SocialDailyAggregation,
  SocialDailyAggregationListResponse,
  SocialDailyAggregationFilters,
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

  // ============================================
  // Social Posts
  // ============================================

  /**
   * GET /api/social/posts (com filtros e paginação)
   */
  listPosts: async (params?: SocialPostFilters): Promise<SocialPostListResponse> => {
    console.log('[SOCIAL SERVICE] listPosts called with params:', params);
    const response = await apiClient.get<SocialPostListResponse>(
      ENDPOINTS.SOCIAL.LIST_POSTS,
      { params }
    );
    console.log('[SOCIAL SERVICE] listPosts response.data:', response.data);
    return response.data;
  },

  /**
   * GET /api/social/posts/:id
   */
  getPostById: async (id: string): Promise<SocialPost> => {
    console.log('[SOCIAL SERVICE] getPostById called with id:', id);
    const response = await apiClient.get<SocialPost>(ENDPOINTS.SOCIAL.GET_POST(id));
    console.log('[SOCIAL SERVICE] getPostById response.data:', response.data);
    return response.data;
  },

  /**
   * POST /api/social/posts
   */
  createPost: async (data: CreateSocialPostRequest): Promise<SocialPost> => {
    console.log('[SOCIAL SERVICE] createPost called with data:', data);
    const response = await apiClient.post<SocialPost>(ENDPOINTS.SOCIAL.CREATE_POST, data);
    console.log('[SOCIAL SERVICE] createPost response.data:', response.data);
    return response.data;
  },

  /**
   * PUT /api/social/posts/:id
   */
  updatePost: async (id: string, data: UpdateSocialPostRequest): Promise<SocialPost> => {
    console.log('[SOCIAL SERVICE] updatePost called with id:', id, 'data:', data);
    const response = await apiClient.put<SocialPost>(ENDPOINTS.SOCIAL.UPDATE_POST(id), data);
    console.log('[SOCIAL SERVICE] updatePost response.data:', response.data);
    return response.data;
  },

  /**
   * DELETE /api/social/posts/:id
   */
  deletePost: async (id: string): Promise<void> => {
    console.log('[SOCIAL SERVICE] deletePost called with id:', id);
    await apiClient.delete(ENDPOINTS.SOCIAL.DELETE_POST(id));
  },

  // ============================================
  // Social Daily Aggregations
  // ============================================

  /**
   * GET /api/social/daily-aggregations (com filtros e paginação)
   */
  listAggregations: async (params?: SocialDailyAggregationFilters): Promise<SocialDailyAggregationListResponse> => {
    console.log('[SOCIAL SERVICE] listAggregations called with params:', params);
    const response = await apiClient.get<SocialDailyAggregationListResponse>(
      ENDPOINTS.SOCIAL.LIST_AGGREGATIONS,
      { params }
    );
    console.log('[SOCIAL SERVICE] listAggregations response.data:', response.data);
    return response.data;
  },

  /**
   * GET /api/social/daily-aggregations/:id
   */
  getAggregationById: async (id: string): Promise<SocialDailyAggregation> => {
    console.log('[SOCIAL SERVICE] getAggregationById called with id:', id);
    const response = await apiClient.get<SocialDailyAggregation>(ENDPOINTS.SOCIAL.GET_AGGREGATION(id));
    console.log('[SOCIAL SERVICE] getAggregationById response.data:', response.data);
    return response.data;
  },

  /**
   * POST /api/social/daily-aggregations/recalculate/:brandID/:date
   */
  recalculateAggregations: async (brandID: string, date: string): Promise<SocialDailyAggregation> => {
    console.log('[SOCIAL SERVICE] recalculateAggregations called with brandID:', brandID, 'date:', date);
    const response = await apiClient.post<SocialDailyAggregation>(
      ENDPOINTS.SOCIAL.RECALCULATE_AGGREGATIONS(brandID, date)
    );
    console.log('[SOCIAL SERVICE] recalculateAggregations response.data:', response.data);
    return response.data;
  },
};
