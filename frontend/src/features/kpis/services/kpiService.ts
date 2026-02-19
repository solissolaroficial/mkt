import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  KpiCategory,
  CreateKpiDTO,
  UpdateKpiDTO,
  UpdateMonthlyDataDTO,
  KpiListResponse,
  DailyEntry,
  AddDailyEntryDTO,
  UpdateDailyEntryDTO,
  DeleteDailyEntryDTO,
} from '../types';

export const kpiService = {
  /**
   * Lista todos os KPIs com filtros opcionais de mês e ano
   */
  list: async (month?: string, year?: number): Promise<KpiListResponse> => {
    const params = new URLSearchParams();
    if (month) params.append('month', month);
    if (year) params.append('year', year.toString());

    const url = `${ENDPOINTS.KPIS.LIST}?${params.toString()}`;
    const response = await apiClient.get<KpiListResponse>(url);
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
 
  /**
   * Deleta dados mensais de um KPI
   */
  deleteMonthlyData: async (kpiId: string, monthlyDataId: string): Promise<void> => {
  	await apiClient.delete(`${ENDPOINTS.KPIS.UPDATE_MONTHLY(kpiId)}/${monthlyDataId}`);
  },
 
  /**
   * Busca KPIs por uma lista de slugs com filtros opcionais de mês e ano
   */
  getBySlugs: async (slugs: string[], month?: string, year?: number): Promise<KpiCategory[]> => {
    const params = new URLSearchParams();
    if (month) params.append('month', month);
    if (year) params.append('year', year.toString());

    const url = `${ENDPOINTS.KPIS.GET_BY_SLUGS}?${params.toString()}`;
    const response = await apiClient.post<{ data: KpiCategory[] }>(
      url,
      { slugs }
    );
    return response.data.data;
  },

  /**
   * Métodos para entradas diárias
   */
  dailyEntry: {
    /**
     * Lista entradas diárias de um KPI para um mês/ano específico
     */
    get: async (kpiId: string, month: string, year: number): Promise<DailyEntry[]> => {
      const params = new URLSearchParams({ month, year: year.toString() });
      const response = await apiClient.get<DailyEntry[]>(
        `${ENDPOINTS.KPIS.GET(kpiId)}/daily-entries?${params.toString()}`
      );
      return response.data;
    },

    /**
     * Adiciona uma entrada diária
     */
    add: async (kpiId: string, data: AddDailyEntryDTO): Promise<void> => {
      await apiClient.post(`${ENDPOINTS.KPIS.GET(kpiId)}/daily-entry`, data);
    },

    /**
     * Atualiza uma entrada diária
     */
    update: async (kpiId: string, data: UpdateDailyEntryDTO): Promise<void> => {
      await apiClient.put(`${ENDPOINTS.KPIS.GET(kpiId)}/daily-entry`, data);
    },

    /**
     * Remove uma entrada diária
     */
    delete: async (kpiId: string, data: DeleteDailyEntryDTO): Promise<void> => {
      await apiClient.delete(`${ENDPOINTS.KPIS.GET(kpiId)}/daily-entry`, { data });
    },
  },
};
