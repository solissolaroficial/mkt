// ============================================
// Pagination Types (Snake_Case - Backend Compatible)
// ============================================

export interface PaginationResponse {
  page: number;
  page_size: number;
  total: number;
  total_pages: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  pagination: PaginationResponse;
}
