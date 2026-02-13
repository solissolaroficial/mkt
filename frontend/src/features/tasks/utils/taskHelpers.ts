// Tradução de prioridades para português
export const TASK_PRIORITY_LABELS: Record<string, string> = {
  low: 'Baixa',
  medium: 'Média',
  high: 'Alta',
  urgent: 'Urgente',
};

// Cores de prioridade (padronizadas)
export const getPriorityColor = (priority: string): string => {
  switch (priority) {
    case 'high': return 'text-rose-400 bg-rose-500/10 border-rose-500/20';
    case 'medium': return 'text-amber-400 bg-amber-500/10 border-amber-500/20';
    case 'low': return 'text-blue-400 bg-blue-500/10 border-blue-500/20';
    case 'urgent': return 'text-red-400 bg-red-500/10 border-red-500/20';
    default: return 'text-gray-400 bg-gray-500/10 border-gray-500/20';
  }
};
