import { apiClient } from '@/infrastructure/api/client';
import type { PdvPost, RecurrentPdv } from '../types';

export const pdvService = {
  /**
   * Get all PDV posts
   */
  getPosts: async (): Promise<PdvPost[]> => {
    const response = await apiClient.get('/api/pdv/posts');
    return response.data;
  },

  /**
   * Get all recurrent PDVs
   */
  getRecurrentPdvs: async (): Promise<RecurrentPdv[]> => {
    const response = await apiClient.get('/api/pdv/recurrent');
    return response.data;
  },

  /**
   * Create a new PDV post
   */
  createPost: async (post: Omit<PdvPost, 'id' | 'status'>): Promise<PdvPost> => {
    const response = await apiClient.post('/api/pdv/posts', post);
    return response.data;
  },

  /**
   * Create a new recurrent PDV
   */
  createRecurrentPdv: async (pdv: Omit<RecurrentPdv, 'id'>): Promise<RecurrentPdv> => {
    const response = await apiClient.post('/api/pdv/recurrent', pdv);
    return response.data;
  },

  /**
   * Update PDV post status
   */
  updatePostStatus: async (id: string, status: 'verified' | 'pending'): Promise<PdvPost> => {
    const response = await apiClient.put(`/api/pdv/posts/${id}/status`, { status });
    return response.data;
  },
};
