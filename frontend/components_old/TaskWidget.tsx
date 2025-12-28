
import React from 'react';
import { Task } from '../types';
import { CheckCircle2, CalendarDays, ArrowRight, CornerDownRight, AlertTriangle, Layers, User } from 'lucide-react';

interface TaskWidgetProps {
  tasks: Task[];
  onViewAll: () => void;
  onTaskClick?: (taskId: string) => void;
}

const TaskWidget: React.FC<TaskWidgetProps> = ({ tasks, onViewAll, onTaskClick }) => {
  
  // Converter timestamp (ms) ou ISO date string para timestamp
  const getDateValue = (dateStr?: string) => {
    if (!dateStr || dateStr === 'Sem data') return 9999999999999;
    
    // Tratamento para datas ISO YYYY-MM-DD
    if (dateStr.match(/^\d{4}-\d{2}-\d{2}$/)) {
        // Criar data em UTC e pegar o timestamp ajustado ao fuso horário local para comparação correta de "dia"
        const [y, m, d] = dateStr.split('-').map(Number);
        const date = new Date(y, m - 1, d); // Mês é 0-indexado
        return date.getTime();
    }
    
    // Fallback para legado (se ainda existir)
    const now = new Date();
    now.setHours(0, 0, 0, 0); 
    const cleanDate = dateStr.trim().toLowerCase();

    if (cleanDate === 'ontem') return now.getTime() - 86400000;
    if (cleanDate === 'hoje') return now.getTime();
    if (cleanDate === 'amanhã') return now.getTime() + 86400000;
    
    return 9999999999999; 
  };

  const today = new Date();
  today.setHours(0,0,0,0);
  const todayTime = today.getTime();

  // 1. "Achatar" a estrutura: Extrair todas as subtarefas e tarefas principais pendentes
  const actionableItems = tasks.flatMap(card => {
    if (card.status === 'done') return [];

    const pendingSubtasks = (card.subtasks || []).filter(sub => !sub.completed);
    
    // Se houver subtarefas pendentes, mostre-as
    if (pendingSubtasks.length > 0) {
      return pendingSubtasks.map(sub => ({
        id: sub.id,
        title: sub.title,
        dueDate: sub.dueDate || card.dueDate || 'Sem data',
        assignee: sub.assignee || card.assignee,
        parentId: card.id,
        parentTitle: card.title, // Contexto é o título do Card Pai
        parentPriority: card.priority,
        isSubtask: true
      }));
    }

    // Se NÃO houver subtarefas pendentes (ou não houver subtarefas), mostre o próprio Card como tarefa
    return [{
      id: card.id,
      title: card.title,
      dueDate: card.dueDate || 'Sem data',
      assignee: card.assignee,
      parentId: card.id,
      parentTitle: card.category.charAt(0).toUpperCase() + card.category.slice(1), // Contexto é a categoria
      parentPriority: card.priority,
      isSubtask: false
    }];
  });

  // 2. Ordenar por data
  const sortedItems = actionableItems.sort((a, b) => getDateValue(a.dueDate) - getDateValue(b.dueDate)).slice(0, 10);

  const handleItemClick = (parentId: string) => {
    if (onTaskClick) {
      onTaskClick(parentId);
    } else {
      onViewAll();
    }
  };

  return (
    <div className="bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 p-6 flex flex-col h-full">
      <div className="flex justify-between items-center mb-4">
        <div>
          <h3 className="text-lg font-semibold text-gray-100">Minhas Execuções</h3>
          <p className="text-sm text-gray-500">Tarefas do dia e pendências</p>
        </div>
        <button 
          onClick={onViewAll}
          className="text-xs text-gray-400 hover:text-white border border-gray-700 hover:bg-gray-800 px-3 py-1 rounded-lg transition-colors"
        >
          Ver Quadro
        </button>
      </div>

      <div className="flex-grow space-y-3 overflow-y-auto custom-scrollbar pr-1 max-h-[600px]">
        {sortedItems.length === 0 ? (
          <div className="text-center py-12 text-gray-600">
            <CheckCircle2 size={48} className="mx-auto mb-2 opacity-50 text-emerald-500" />
            <p className="font-medium text-gray-400">Tudo executado!</p>
            <p className="text-xs mt-1">Nenhuma pendência encontrada.</p>
          </div>
        ) : (
          sortedItems.map((item, idx) => {
            const taskDateValue = getDateValue(item.dueDate);
            const isLate = taskDateValue < todayTime;
            const isToday = taskDateValue === todayTime;
            
            // Logic for Date Box Color
            let dateBoxClass = 'bg-gray-800/50 border-gray-700 text-gray-400';
            if (isLate) dateBoxClass = 'bg-rose-500/10 border-rose-500/30 text-rose-400';
            else if (isToday) dateBoxClass = 'bg-amber-500/10 border-amber-500/30 text-amber-400';

            // Priority Badge Color
            let priorityClass = 'bg-blue-500/10 text-blue-400 border-blue-500/20';
            if (item.parentPriority === 'alta') priorityClass = 'bg-rose-500/10 text-rose-400 border-rose-500/20';
            if (item.parentPriority === 'media') priorityClass = 'bg-amber-500/10 text-amber-400 border-amber-500/20';

            // Format Date info for box
            let dayStr = '--';
            let weekdayStr = '-';
            
            if (item.dueDate && item.dueDate !== 'Sem data') {
                 // Using manual parsing to avoid timezone shifts
                 const [y, m, d] = item.dueDate.split('-').map(Number);
                 const dateObj = new Date(y, m - 1, d);
                 dayStr = d.toString();
                 weekdayStr = dateObj.toLocaleDateString('pt-BR', { weekday: 'short' }).slice(0, 3);
            }

            return (
              <div 
                key={`${item.id}-${idx}`}
                onClick={() => handleItemClick(item.parentId)}
                className="flex items-center gap-3 p-3 bg-[#20232b]/50 border border-gray-800 rounded-xl hover:border-gray-700 transition-colors group cursor-pointer"
              >
                {/* Date Box */}
                <div className={`flex flex-col items-center justify-center w-12 h-12 rounded-lg border flex-shrink-0 ${dateBoxClass}`}>
                    <span className="text-[10px] uppercase font-bold">{weekdayStr}</span>
                    <span className="text-lg font-bold leading-none">{dayStr}</span>
                </div>
                
                <div className="flex-grow min-w-0">
                  {/* Row 1: Title + Priority */}
                  <div className="flex justify-between items-start mb-1">
                    <p className="text-sm font-medium text-gray-200 truncate pr-2 group-hover:text-white transition-colors leading-tight">
                      {item.title}
                    </p>
                    <span className={`text-[10px] px-1.5 py-0.5 rounded border font-bold uppercase tracking-wider whitespace-nowrap ${priorityClass}`}>
                        {item.parentPriority}
                    </span>
                  </div>
                  
                  {/* Row 2: Metadata */}
                  <div className="flex items-center gap-3 text-xs text-gray-500">
                    <div className="flex items-center gap-1 truncate max-w-[60%]">
                        {item.isSubtask ? <CornerDownRight size={12} className="flex-shrink-0" /> : <Layers size={12} className="flex-shrink-0" />}
                        <span className="truncate">{item.parentTitle}</span>
                    </div>
                    
                    {item.assignee && (
                        <div className="flex items-center gap-1 border-l border-gray-700 pl-3">
                            <User size={12} />
                            <span>{item.assignee}</span>
                        </div>
                    )}
                  </div>
                </div>
              </div>
            );
          })
        )}
      </div>
    </div>
  );
};

export default TaskWidget;
