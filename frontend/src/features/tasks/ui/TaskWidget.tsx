 
import React from 'react';
import type { Task } from '@/shared/types';
import { CheckCircle2, CalendarDays, ArrowRight, CornerDownRight, AlertTriangle, Layers, User } from 'lucide-react';
import type { TaskWidgetProps } from '../types';
import { useUsers } from '@/features/users/hooks';
import { TASK_PRIORITY_LABELS, getPriorityColor } from '../utils/taskHelpers';
import { parseIsoDate, formatShortDate } from '@/shared/utils/dateFormatters';

const TaskWidget: React.FC<TaskWidgetProps> = ({ tasks, onViewAll, onTaskClick }) => {
  const { data: users } = useUsers();

  // Criar mapa de usuários por ID
  const usersMap = React.useMemo(() => {
    if (!users) return new Map();
    return new Map(users.map(u => [u.id, u.name]));
  }, [users]);
  
  // Converter timestamp (ms) ou ISO date string para timestamp
  const getDateValue = (dateStr?: string) => {
    const date = parseIsoDate(dateStr);
    if (!date) return 9999999999999;
    return date.getTime();
  };

  const today = new Date();
  today.setHours(0,0,0,0);
  const todayTime = today.getTime();

  // 1. "Achatar" a estrutura: Extrair todas as subtarefas e tarefas principais pendentes
  const actionableItems = tasks.flatMap(card => {
    if (card.status === 'completed') return [];

    const pendingSubtasks = (card.subtasks || []).filter(sub => sub.status !== 'completed');
    
    // Se houver subtarefas pendentes, mostre-as
    if (pendingSubtasks.length > 0) {
      return pendingSubtasks.map(sub => ({
        id: sub.id,
        title: sub.title,
        dueDate: sub.due_date || card.due_date || 'Sem data',
        assignee: sub.assignee_id || card.assignee_id,
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
      dueDate: card.due_date || 'Sem data',
      assignee: card.assignee_id,
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
            const priorityClass = getPriorityColor(item.parentPriority);

            // Format Date info for box
            let dayStr = '--';
            let weekdayStr = '-';

            if (item.dueDate && item.dueDate !== 'Sem data') {
                 const date = parseIsoDate(item.dueDate);
                 if (date) {
                    dayStr = date.getDate().toString();
                    weekdayStr = date.toLocaleDateString('pt-BR', { weekday: 'short' }).slice(0, 3);
                 }
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
                        {TASK_PRIORITY_LABELS[item.parentPriority] || item.parentPriority}
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
                            <span>{usersMap.get(item.assignee) || item.assignee}</span>
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
