import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  KpiCategory,
  CreateKpiDTO,
  UpdateKpiDTO,
  UpdateMonthlyDataDTO,
  KpiListResponse
} from '../types';

export const kpiService = {
  /**
   * Lista todos os KPIs
   */
  list: async (): Promise<KpiListResponse> => {
    const response = await apiClient.get<KpiListResponse>(ENDPOINTS.KPIS.LIST);
    return response.data;
  },

  /**
   * Busca um KPI por ID
   */
  getById: async (id: string): Promise<KpiCategory> => {
    const response = await apiClient.get<KpiCategory>(ENDPOINTS.KPIS.GET(id));
    return response.data;
  },

  /**
   * Cria um novo KPI
   */
  create: async (data: CreateKpiDTO): Promise<KpiCategory> => {
    const response = await apiClient.post<KpiCategory>(ENDPOINTS.KPIS.CREATE, data);
    return response.data;
  },

  /**
   * Atualiza um KPI
   */
  update: async (id: string, data: UpdateKpiDTO): Promise<KpiCategory> => {
    const response = await apiClient.put<KpiCategory>(ENDPOINTS.KPIS.UPDATE(id), data);
    return response.data;
  },

  /**
   * Deleta um KPI
   */
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(ENDPOINTS.KPIS.DELETE(id));
  },

  /**
   * Atualiza dados mensais de um KPI
   */
  updateMonthlyData: async (
    kpiId: string,
    data: UpdateMonthlyDataDTO
  ): Promise<KpiCategory> => {
    const response = await apiClient.put<KpiCategory>(
      ENDPOINTS.KPIS.UPDATE_MONTHLY(kpiId),
      data
    );
    return response.data;
  },

  /**
   * Atualiza apenas a meta de um mês específico
   */
  updateMeta: async (
    kpiId: string,
    year: number,
    month: string,
    meta: number
  ): Promise<KpiCategory> => {
    const response = await apiClient.put<KpiCategory>(
      ENDPOINTS.KPIS.UPDATE_MONTHLY(kpiId),
      { year, month, meta }
    );
    return response.data;
  },
};