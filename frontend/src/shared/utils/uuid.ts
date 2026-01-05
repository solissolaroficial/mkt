/**
 * Gera um UUID v4 usando a API nativa crypto.randomUUID()
 * Disponível em browsers modernos e Node.js (v15.6.0+)
 */
export const generateUUID = (): string => {
  if (typeof crypto !== 'undefined' && crypto.randomUUID) {
    return crypto.randomUUID();
  }
  
  // Fallback para ambientes onde crypto.randomUUID não está disponível
  // Gera um UUID v4 compatível com RFC 4122
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
    const r = Math.random() * 16 | 0;
    const v = c === 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
};
