import { apiClient } from '@/infrastructure/api/client';
import type {
  Flow,
  FlowsListResponse,
  CreateFlowRequest,
  UpdateFlowRequest,
} from '@/shared/types';

export const flowService = {
  // List all flows with pagination
  list: async (page = 1, limit = 100): Promise<FlowsListResponse> => {
    try {
      console.log('flowService: Getting flows with page:', page, 'limit:', limit);
      const response = await apiClient.get('/api/flows', { params: { page, limit } });
      console.log('flowService: Flows list received:', response.data);
      return response.data;
    } catch (error) {
      console.error('flowService: Error getting flows:', error);
      throw error;
    }
  },

  // Get flow by ID
  getById: async (id: string): Promise<Flow> => {
    try {
      console.log('flowService: Getting flow with id:', id);
      const response = await apiClient.get(`/api/flows/${id}`);
      console.log('flowService: Flow received:', response.data);
      return response.data;
    } catch (error) {
      console.error('flowService: Error getting flow:', error);
      throw error;
    }
  },

  // Create new flow
  create: async (data: CreateFlowRequest): Promise<Flow> => {
    try {
      console.log('flowService: Sending POST request to /api/flows with data:', data);
      const response = await apiClient.post<Flow>('/api/flows', data);
      console.log('flowService: Response received:', response);
      return response.data;
    } catch (error) {
      console.error('flowService: Error creating flow:', error);
      throw error;
    }
  },

  // Update flow
  update: async (id: string, data: UpdateFlowRequest): Promise<Flow> => {
    try {
      console.log('flowService: Sending PUT request to /api/flows/:id with data:', data);
      const response = await apiClient.put<Flow>(`/api/flows/${id}`, data);
      console.log('flowService: Response received:', response);
      return response.data;
    } catch (error) {
      console.error('flowService: Error updating flow:', error);
      throw error;
    }
  },

  // Delete flow
  delete: async (id: string): Promise<void> => {
    try {
      console.log('flowService: Sending DELETE request to /api/flows/:id');
      await apiClient.delete(`/api/flows/${id}`);
      console.log('flowService: Flow deleted successfully');
    } catch (error) {
      console.error('flowService: Error deleting flow:', error);
      throw error;
    }
  },

  // Reorder flows
  reorder: async (flowIds: string[]): Promise<void> => {
    try {
      console.log('flowService: Sending POST request to /api/flows/reorder with flow_ids:', flowIds);
      await apiClient.post('/api/flows/reorder', { flow_ids: flowIds });
      console.log('flowService: Flows reordered successfully');
    } catch (error) {
      console.error('flowService: Error reordering flows:', error);
      throw error;
    }
  },
};
