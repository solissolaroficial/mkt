// Format ISO date string to DD/MM/YYYY for display
export const formatDisplayDate = (dateStr?: string | Date): string => {
  if (!dateStr || dateStr === 'Sem data') return 'Sem data';

  // Convert to Date object
  const date = dateStr instanceof Date ? dateStr : new Date(dateStr);

  // Handle invalid dates
  if (isNaN(date.getTime())) return 'Sem data';

  return new Intl.DateTimeFormat('pt-BR', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric'
  }).format(date);
};

// Format to DD/MM (shorter version) - for Kanban cards
export const formatShortDate = (dateStr?: string): string => {
  if (!dateStr || dateStr === 'Sem data') return '--/--';

  // Extract date part if ISO format (removes "T" separator)
  const cleanDateStr = dateStr.includes('T') ? dateStr.split('T')[0] : dateStr;

  if (cleanDateStr.match(/^\d{4}-\d{2}-\d{2}$/)) {
    const [y, m, d] = cleanDateStr.split('-');
    return `${d}/${m}`;
  }

  return formatDisplayDate(cleanDateStr).split('/').slice(0, 2).join('/');
};

// Parse ISO to Date object safely
export const parseIsoDate = (dateStr?: string): Date | null => {
  if (!dateStr || dateStr === 'Sem data') return null;

  const date = new Date(dateStr);
  return isNaN(date.getTime()) ? null : date;
};

// Format date for input value (YYYY-MM-DD)
export const formatDateForInput = (dateStr?: string | Date): string => {
  if (!dateStr || dateStr === 'Sem data') return '';

  const date = dateStr instanceof Date ? dateStr : new Date(dateStr);
  if (isNaN(date.getTime())) return '';

  return date.toISOString().split('T')[0];
};

// Format date for API request (RFC3339 with UTC midnight)
export const formatDateForAPI = (dateStr?: string | Date): string | undefined => {
  if (!dateStr || dateStr === 'Sem data') return undefined;

  const date = dateStr instanceof Date ? dateStr : new Date(dateStr);
  if (isNaN(date.getTime())) return undefined;

  // Set to midnight UTC to ensure consistent behavior
  const midnightUTC = new Date(
    Date.UTC(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0)
  );

  return midnightUTC.toISOString();
};
