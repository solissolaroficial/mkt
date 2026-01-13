// ============================================
// Gifts Types (Snake_Case - Backend Compatible)
// ============================================

// ============================================
// GiftItem
// ============================================

export interface GiftItem {
  id: string;
  name: string;
  stock: number;
  price: number;
  created_at: string;        // ISO 8601
  updated_at: string;        // ISO 8601
}

export interface CreateGiftItemRequest {
  name: string;
  stock: number;
  price: number;
}

export interface UpdateGiftItemRequest {
  name?: string;
  stock?: number;
  price?: number;
}

// ============================================
// GiftTransaction
// ============================================

export interface GiftTransaction {
  id: string;
  item_name: string;          // ✅ snake_case
  date: string;               // ISO 8601 (YYYY-MM-DD) - formato retornado pelo backend
  time: string;               // HH:MM
  quantity: number;
  unit: string;
  transaction_type: 'in' | 'out';  // ✅ snake_case
  price?: number;             // Apenas para entradas
  representative?: string;     // Apenas para saídas
  created_at: string;         // ISO 8601
  updated_at: string;         // ISO 8601
}

export interface CreateGiftTransactionRequest {
  item_name: string;         // ✅ snake_case
  quantity: number;
  transaction_type: 'in' | 'out';  // ✅ snake_case
  date: string;               // YYYY-MM-DD (formato esperado pelo backend)
  time?: string;              // HH:MM (opcional)
  price?: number;             // Obrigatório para 'in'
  representative?: string;     // Obrigatório para 'out'
  unit?: string;
}

export interface UpdateGiftTransactionRequest {
  item_name?: string;       // ✅ snake_case
  quantity?: number;
  transaction_type?: 'in' | 'out';  // ✅ snake_case
  date?: string;
  time?: string;
  price?: number;
  representative?: string;
  // ✅ unit NÃO incluído pois o backend não permite atualizar este campo
}

// ============================================
// Paginated Responses
// ============================================

export interface GiftItemListResponse {
  items: GiftItem[];
  meta: {
    total: number;
    page: number;
    limit: number;
  };
}

export interface GiftTransactionListResponse {
  transactions: GiftTransaction[];
  meta: {
    total: number;
    page: number;
    limit: number;
  };
}

// ============================================
// Filters
// ============================================

export interface GiftItemFilters {
  name?: string;
  min_stock?: number;
  max_stock?: number;
  min_price?: number;
  max_price?: number;
  page?: number;
  limit?: number;
}

export interface GiftTransactionFilters {
  item_name?: string;           // ✅ snake_case
  transaction_type?: 'in' | 'out';  // ✅ snake_case
  representative?: string;
  start_date?: string;          // YYYY-MM-DD (ISO 8601)
  end_date?: string;            // YYYY-MM-DD (ISO 8601)
  page?: number;
  limit?: number;
}

// ============================================
// UI Types
// ============================================

export interface GiftsViewProps {}

export type TransactionTab = 'out' | 'in';
