import { apiClient } from '@/infrastructure/api/client';
import type {
  RepresentativeMonthlyGoal,
  CreateRepresentativeMonthlyGoalRequest,
  UpdateRepresentativeMonthlyGoalRequest,
  ListRepresentativeMonthlyGoalsRequest,
  ListRepresentativeMonthlyGoalsResponse,
  GetRepresentativeGoalsTableDataRequest,
  GetRepresentativeGoalsTableDataResponse,
} from '../types/monthlyGoal';

const BASE_URL = '/api/representative-monthly-goals';

/**
 * Service for managing Representative Monthly Goals
 */
export const monthlyGoalService = {
  /**
   * Create a new monthly goal
   */
  create: async (data: CreateRepresentativeMonthlyGoalRequest): Promise<RepresentativeMonthlyGoal> => {
    const response = await apiClient.post<RepresentativeMonthlyGoal>(BASE_URL, data);
    return response.data;
  },

  /**
   * Get a monthly goal by ID
   */
  getById: async (id: string): Promise<RepresentativeMonthlyGoal> => {
    const response = await apiClient.get<RepresentativeMonthlyGoal>(`${BASE_URL}/${id}`);
    return response.data;
  },

  /**
   * Update a monthly goal
   */
  update: async (id: string, data: UpdateRepresentativeMonthlyGoalRequest): Promise<RepresentativeMonthlyGoal> => {
    const response = await apiClient.put<RepresentativeMonthlyGoal>(`${BASE_URL}/${id}`, data);
    return response.data;
  },

  /**
   * Delete a monthly goal (soft delete)
   */
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`${BASE_URL}/${id}`);
  },

  /**
   * List monthly goals with pagination and filters
   */
  list: async (params?: ListRepresentativeMonthlyGoalsRequest): Promise<ListRepresentativeMonthlyGoalsResponse> => {
    const response = await apiClient.get<ListRepresentativeMonthlyGoalsResponse>(BASE_URL, { params });
    return response.data;
  },

  /**
   * Get table data for transposed view (representatives as rows, months as columns)
   */
  getTableData: async (params?: GetRepresentativeGoalsTableDataRequest): Promise<GetRepresentativeGoalsTableDataResponse> => {
    const response = await apiClient.get<GetRepresentativeGoalsTableDataResponse>(`${BASE_URL}/table/data`, { params });
    return response.data;
  },
};
