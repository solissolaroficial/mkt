export interface Flow {
  uuid: string;
  name: string;
  description?: string;
  color?: string;
  sort_order: number;
  created_at: string;
  updated_at: string;
}

export interface CreateFlowRequest {
  name: string;
  description?: string;
  color?: string;
}

export interface UpdateFlowRequest {
  name?: string;
  description?: string;
  color?: string;
  sort_order?: number;
}

export interface FlowsListResponse {
  data: Flow[];
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}