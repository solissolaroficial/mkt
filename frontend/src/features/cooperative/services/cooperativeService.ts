import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type {
  ShowroomItem,
  OfflineAction,
  RepMarketingAction,
  CreateShowroomItemRequest,
  UpdateShowroomItemRequest,
  CreateOfflineActionRequest,
  UpdateOfflineActionRequest,
  CreateRepMarketingActionRequest,
  UpdateRepMarketingActionRequest,
  ShowroomItemFilters,
  OfflineActionFilters,
  RepMarketingActionFilters,
  PaginatedResponse
} from '../types';

export const cooperativeService = {
  // ============================================
  // Offline Actions
  // ============================================
  
  getOfflineActions: async (params?: OfflineActionFilters): Promise<PaginatedResponse<OfflineAction>> => {
    console.log('[COOPERATIVE SERVICE] getOfflineActions called with params:', params);
    const response = await apiClient.get<PaginatedResponse<OfflineAction>>(
      ENDPOINTS.COOPERATIVE.OFFLINE_ACTIONS_LIST,
      { params }
    );
    console.log('[COOPERATIVE SERVICE] getOfflineActions response.data:', response.data);
    return response.data;
  },

  getOfflineActionById: async (uuid: string): Promise<OfflineAction> => {
    console.log('[COOPERATIVE SERVICE] getOfflineActionById called with uuid:', uuid);
    const response = await apiClient.get<OfflineAction>(
      ENDPOINTS.COOPERATIVE.OFFLINE_ACTIONS_GET(uuid)
    );
    console.log('[COOPERATIVE SERVICE] getOfflineActionById response.data:', response.data);
    return response.data;
  },

  createOfflineAction: async (data: CreateOfflineActionRequest): Promise<OfflineAction> => {
    console.log('[COOPERATIVE SERVICE] createOfflineAction called with data:', data);
    const response = await apiClient.post<OfflineAction>(
      ENDPOINTS.COOPERATIVE.OFFLINE_ACTIONS_CREATE,
      data
    );
    console.log('[COOPERATIVE SERVICE] createOfflineAction response.data:', response.data);
    return response.data;
  },

  updateOfflineAction: async (uuid: string, data: UpdateOfflineActionRequest): Promise<OfflineAction> => {
    console.log('[COOPERATIVE SERVICE] updateOfflineAction called with uuid:', uuid, 'data:', data);
    const response = await apiClient.put<OfflineAction>(
      ENDPOINTS.COOPERATIVE.OFFLINE_ACTIONS_UPDATE(uuid),
      data
    );
    console.log('[COOPERATIVE SERVICE] updateOfflineAction response.data:', response.data);
    return response.data;
  },

  deleteOfflineAction: async (uuid: string): Promise<void> => {
    console.log('[COOPERATIVE SERVICE] deleteOfflineAction called with uuid:', uuid);
    await apiClient.delete(ENDPOINTS.COOPERATIVE.OFFLINE_ACTIONS_DELETE(uuid));
  },

  // ============================================
  // Showroom Items
  // ============================================
  
  getShowroomItems: async (params?: ShowroomItemFilters): Promise<PaginatedResponse<ShowroomItem>> => {
    console.log('[COOPERATIVE SERVICE] getShowroomItems called with params:', params);
    const response = await apiClient.get<PaginatedResponse<ShowroomItem>>(
      ENDPOINTS.COOPERATIVE.SHOWROOM_ITEMS_LIST,
      { params }
    );
    console.log('[COOPERATIVE SERVICE] getShowroomItems response.data:', response.data);
    return response.data;
  },

  getShowroomItemById: async (uuid: string): Promise<ShowroomItem> => {
    console.log('[COOPERATIVE SERVICE] getShowroomItemById called with uuid:', uuid);
    const response = await apiClient.get<ShowroomItem>(
      ENDPOINTS.COOPERATIVE.SHOWROOM_ITEMS_GET(uuid)
    );
    console.log('[COOPERATIVE SERVICE] getShowroomItemById response.data:', response.data);
    return response.data;
  },

  createShowroomItem: async (data: CreateShowroomItemRequest): Promise<ShowroomItem> => {
    console.log('[COOPERATIVE SERVICE] createShowroomItem called with data:', data);
    const response = await apiClient.post<ShowroomItem>(
      ENDPOINTS.COOPERATIVE.SHOWROOM_ITEMS_CREATE,
      data
    );
    console.log('[COOPERATIVE SERVICE] createShowroomItem response.data:', response.data);
    return response.data;
  },

  updateShowroomItem: async (uuid: string, data: UpdateShowroomItemRequest): Promise<ShowroomItem> => {
    console.log('[COOPERATIVE SERVICE] updateShowroomItem called with uuid:', uuid, 'data:', data);
    const response = await apiClient.put<ShowroomItem>(
      ENDPOINTS.COOPERATIVE.SHOWROOM_ITEMS_UPDATE(uuid),
      data
    );
    console.log('[COOPERATIVE SERVICE] updateShowroomItem response.data:', response.data);
    return response.data;
  },

  deleteShowroomItem: async (uuid: string): Promise<void> => {
    console.log('[COOPERATIVE SERVICE] deleteShowroomItem called with uuid:', uuid);
    await apiClient.delete(ENDPOINTS.COOPERATIVE.SHOWROOM_ITEMS_DELETE(uuid));
  },

  // ============================================
  // Rep Marketing Actions
  // ============================================
  
  getRepMarketingActions: async (params?: RepMarketingActionFilters): Promise<PaginatedResponse<RepMarketingAction>> => {
    console.log('[COOPERATIVE SERVICE] getRepMarketingActions called with params:', params);
    const response = await apiClient.get<PaginatedResponse<RepMarketingAction>>(
      ENDPOINTS.COOPERATIVE.REP_MARKETING_ACTIONS_LIST,
      { params }
    );
    console.log('[COOPERATIVE SERVICE] getRepMarketingActions response.data:', response.data);
    return response.data;
  },

  getRepMarketingActionById: async (uuid: string): Promise<RepMarketingAction> => {
    console.log('[COOPERATIVE SERVICE] getRepMarketingActionById called with uuid:', uuid);
    const response = await apiClient.get<RepMarketingAction>(
      ENDPOINTS.COOPERATIVE.REP_MARKETING_ACTIONS_GET(uuid)
    );
    console.log('[COOPERATIVE SERVICE] getRepMarketingActionById response.data:', response.data);
    return response.data;
  },

  createRepMarketingAction: async (data: CreateRepMarketingActionRequest): Promise<RepMarketingAction> => {
    console.log('[COOPERATIVE SERVICE] createRepMarketingAction called with data:', data);
    const response = await apiClient.post<RepMarketingAction>(
      ENDPOINTS.COOPERATIVE.REP_MARKETING_ACTIONS_CREATE,
      data
    );
    console.log('[COOPERATIVE SERVICE] createRepMarketingAction response.data:', response.data);
    return response.data;
  },

  updateRepMarketingAction: async (uuid: string, data: UpdateRepMarketingActionRequest): Promise<RepMarketingAction> => {
    console.log('[COOPERATIVE SERVICE] updateRepMarketingAction called with uuid:', uuid, 'data:', data);
    const response = await apiClient.put<RepMarketingAction>(
      ENDPOINTS.COOPERATIVE.REP_MARKETING_ACTIONS_UPDATE(uuid),
      data
    );
    console.log('[COOPERATIVE SERVICE] updateRepMarketingAction response.data:', response.data);
    return response.data;
  },

  deleteRepMarketingAction: async (uuid: string): Promise<void> => {
    console.log('[COOPERATIVE SERVICE] deleteRepMarketingAction called with uuid:', uuid);
    await apiClient.delete(ENDPOINTS.COOPERATIVE.REP_MARKETING_ACTIONS_DELETE(uuid));
  },
};
