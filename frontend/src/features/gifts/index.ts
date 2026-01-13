/**
 * Gifts Feature - Public API
 * 
 * Este arquivo exporta todos os tipos, serviços, hooks e componentes públicos do módulo Gifts.
 * 
 * Arquitetura:
 * - types/index.ts: Tipos TypeScript em snake_case (compatíveis com backend)
 * - utils/formatters.ts: Utilitários de formatação (data, moeda)
 * - services/giftsService.ts: Serviço de API com todos os endpoints
 * - hooks/useGifts.ts: Hooks de React Query para leitura de dados
 * - hooks/useGiftsMutations.ts: Hooks de React Query para mutações com wrapper functions
 * - ui/GiftsView.tsx: Componente principal da UI
 * - pages/GiftsPage.tsx: Página do módulo
 */

// ==================== TYPES ====================

export type {
  // Gift Items
  GiftItem,
  GiftItemListResponse,
  CreateGiftItemRequest,
  UpdateGiftItemRequest,
  GiftItemFilters,
  
  // Gift Transactions
  GiftTransaction,
  GiftTransactionListResponse,
  CreateGiftTransactionRequest,
  UpdateGiftTransactionRequest,
  GiftTransactionFilters,
  
  // Common
  TransactionTab,
} from './types';

// ==================== UTILITIES ====================

export {
  fromDDMMYYYYtoISO,
  fromISOtoDDMMYYYY,
  formatCurrency,
  parseCurrency,
  isValidDate,
} from './utils/formatters';

// ==================== SERVICES ====================

export { giftsService } from './services/giftsService';

// ==================== HOOKS ====================

// Query Hooks
export {
  useGiftItems,
  useGiftItem,
  useGiftTransactions,
  useGiftTransaction,
} from './hooks/useGifts';

// Mutation Hooks
export {
  useGiftItemMutations,
  useGiftTransactionMutations,
} from './hooks/useGiftsMutations';

// ==================== COMPONENTS ====================

export { default as GiftsView } from './ui/GiftsView';
export { default as GiftsPage } from './pages/GiftsPage';
