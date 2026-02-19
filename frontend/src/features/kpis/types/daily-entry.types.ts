/**
 * Tipos para entradas diárias de KPI
 */

/**
 * Entrada diária existente
 */
export interface DailyEntry {
  date: string; // YYYY-MM-DD
  value: number;
  context: string; // Campo required, pode ser string vazia
  user: string;
  createdAt: string; // ISO timestamp
}

/**
 * DTO para adicionar entrada diária
 */
export interface AddDailyEntryDTO {
  year: number;
  month: string;
  date: string; // YYYY-MM-DD
  value: number;
  context: string; // Campo required, pode ser string vazia
}

/**
 * DTO para atualizar entrada diária
 */
export interface UpdateDailyEntryDTO extends AddDailyEntryDTO {}

/**
 * DTO para deletar entrada diária
 */
export interface DeleteDailyEntryDTO {
  year: number;
  month: string;
  date: string;
  context?: string;
}
