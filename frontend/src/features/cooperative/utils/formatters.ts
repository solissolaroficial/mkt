// ============================================
// Formatters for Cooperative Integration
// ============================================

/**
 * Converte formato DD/MM/YYYY para ISO 8601 (YYYY-MM-DD)
 */
export const fromDDMMYYYYtoISO = (dateStr: string): string => {
  if (!dateStr) return '';
  
  const parts: string[] = dateStr.split('/');
  if (parts.length !== 3) {
    throw new Error(`Data inválida: ${dateStr}. Formato esperado: DD/MM/YYYY`);
  }
  
  const [day, month, year] = parts;
  
  // Validar se são números
  if (isNaN(parseInt(day)) || isNaN(parseInt(month)) || isNaN(parseInt(year))) {
    throw new Error(`Data inválida: ${dateStr}. Todos os campos devem ser números.`);
  }
  
  // Validar ranges
  const dayNum: number = parseInt(day);
  const monthNum: number = parseInt(month);
  const yearNum: number = parseInt(year);
  
  if (monthNum < 1 || monthNum > 12) {
    throw new Error(`Mês inválido: ${monthNum}. Deve estar entre 1 e 12.`);
  }
  
  if (dayNum < 1 || dayNum > 31) {
    throw new Error(`Dia inválido: ${dayNum}. Deve estar entre 1 e 31.`);
  }
  
  if (yearNum < 1900 || yearNum > 2100) {
    throw new Error(`Ano inválido: ${yearNum}. Deve estar entre 1900 e 2100.`);
  }
  
  return `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
};

/**
 * Converte formato ISO 8601 (YYYY-MM-DD) para DD/MM/YYYY
 */
export const fromISOtoDDMMYYYY = (dateStr: string): string => {
  if (!dateStr) return '';
  const [year, month, day] = dateStr.split('-');
  return `${day}/${month}/${year}`;
};

/**
 * Formata valor numérico como moeda brasileira (R$)
 */
export const formatCurrency = (value: number): string => {
  return value.toLocaleString('pt-BR', {
    style: 'currency',
    currency: 'BRL'
  });
};

/**
 * Remove formatação de moeda e retorna número
 */
export const parseCurrency = (value: string): number => {
  if (!value) return 0;
  
  // Remover formatação
  const cleaned = value
    .replace('R$', '')
    .replace(/\./g, '')
    .replace(',', '.')
    .trim();
  
  // Tentar converter para número
  const parsed = parseFloat(cleaned);
  
  // Validar se é um número válido
  if (isNaN(parsed)) {
    throw new Error(`Valor monetário inválido: ${value}`);
  }
  
  // Validar se é positivo
  if (parsed < 0) {
    throw new Error(`Valor monetário não pode ser negativo: ${value}`);
  }
  
  return parsed;
};

/**
 * Verifica se uma data é válida
 */
export const isValidDate = (dateStr: string): boolean => {
  if (!dateStr) return false;
  const date = new Date(dateStr);
  return !isNaN(date.getTime());
};
