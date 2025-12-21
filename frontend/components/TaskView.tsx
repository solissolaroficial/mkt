
import React, { useState } from 'react';
import { Task } from '../types';
import { 
  CheckCircle2, 
  Circle, 
  Clock, 
  Calendar, 
  Plus, 
  Filter,
  Search,
  MoreVertical
} from 'lucide-react';

interface TaskViewProps {
  initialTasks: Task[];
}

const TaskView: React.FC<TaskViewProps> = ({ initialTasks }) => {
  const [tasks, setTasks] = useState<Task[]>(initialTasks);
  const [filter, setFilter] = useState<'all' | 'pending' | 'completed'>('all');

  const toggleTask = (id: string) => {
    setTasks(tasks.map(t => 
      t.id === id ? { ...t, status: t.status === 'done' ? 'todo' : 'done' } : t
    ));
  };

  const filteredTasks = tasks.filter(t => {
    if (filter === 'pending') return t.status !== 'done';
    if (filter === 'completed') return t.status === 'done';
    return true;
  });

  // Grouping
  const groupedTasks = {
    'Hoje': filteredTasks.filter(t => t.dueDate === 'Hoje'),
    'Próximos': filteredTasks.filter(t => t.dueDate !== 'Hoje' && t.dueDate !== 'Amanhã'),
    'Amanhã': filteredTasks.filter(t => t.dueDate === 'Amanhã'),
  };

  const getPriorityColor = (p: string) => {
    switch (p) {
      case 'alta': return 'text-rose-700 bg-rose-50 border-rose-200';
      case 'media': return 'text-amber-700 bg-amber-50 border-amber-200';
      default: return 'text-blue-700 bg-blue-50 border-blue-200';
    }
  };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header & Controls */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Minhas Tarefas</h1>
          <p className="text-gray-500">Gerencie suas prioridades diárias</p>
        </div>
        
        <div className="flex items-center gap-3 w-full md:w-auto">
          <div className="relative flex-grow md:flex-grow-0">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" size={18} />
            <input 
              type="text" 
              placeholder="Buscar tarefa..." 
              className="w-full md:w-64 pl-10 pr-4 py-2 bg-white rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
            />
          </div>
          <button className="bg-blue-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-blue-700 transition-colors flex items-center gap-2">
            <Plus size={18} /> Nova
          </button>
        </div>
      </div>

      {/* Tabs / Filters */}
      <div className="flex items-center gap-4 border-b border-gray-200 mb-6">
        <button 
          onClick={() => setFilter('all')}
          className={`pb-3 text-sm font-medium border-b-2 transition-colors ${filter === 'all' ? 'border-blue-600 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700'}`}
        >
          Todas
        </button>
        <button 
          onClick={() => setFilter('pending')}
          className={`pb-3 text-sm font-medium border-b-2 transition-colors ${filter === 'pending' ? 'border-blue-600 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700'}`}
        >
          Pendentes
        </button>
        <button 
          onClick={() => setFilter('completed')}
          className={`pb-3 text-sm font-medium border-b-2 transition-colors ${filter === 'completed' ? 'border-blue-600 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700'}`}
        >
          Concluídas
        </button>
      </div>

      {/* Task List */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {['Hoje', 'Amanhã', 'Próximos'].map(groupKey => {
           const tasksInGroup = groupedTasks[groupKey as keyof typeof groupedTasks] || [];
           if (tasksInGroup.length === 0) return null;

           return (
             <div key={groupKey} className="flex flex-col gap-3">
               <h3 className="font-semibold text-gray-700 flex items-center gap-2">
                 <Calendar size={16} /> {groupKey}
                 <span className="bg-gray-100 text-gray-600 text-xs px-2 py-0.5 rounded-full">{tasksInGroup.length}</span>
               </h3>
               
               {tasksInGroup.map(task => {
                 const isCompleted = task.status === 'done';
                 return (
                 <div 
                    key={task.id} 
                    className={`bg-white p-4 rounded-xl border transition-all hover:shadow-md ${isCompleted ? 'border-gray-100 opacity-60' : 'border-gray-200'}`}
                 >
                   <div className="flex items-start gap-3">
                     <button 
                        onClick={() => toggleTask(task.id)}
                        className={`mt-1 flex-shrink-0 transition-colors ${isCompleted ? 'text-green-500' : 'text-gray-300 hover:text-blue-500'}`}
                     >
                       {isCompleted ? <CheckCircle2 size={20} /> : <Circle size={20} />}
                     </button>
                     
                     <div className="flex-grow">
                       <h4 className={`text-sm font-medium mb-1 ${isCompleted ? 'text-gray-500 line-through' : 'text-gray-900'}`}>
                         {task.title}
                       </h4>
                       <div className="flex flex-wrap items-center gap-2">
                         <span className={`text-[10px] px-2 py-0.5 rounded-full border font-semibold uppercase tracking-wide ${getPriorityColor(task.priority)}`}>
                           {task.priority}
                         </span>
                         <span className="text-xs text-gray-400 bg-gray-50 px-2 py-0.5 rounded-full border border-gray-100 capitalize">
                           {task.category}
                         </span>
                         {task.dueDate !== groupKey && (
                            <span className="text-xs text-gray-400 flex items-center gap-1 ml-auto">
                                <Clock size={12} /> {task.dueDate}
                            </span>
                         )}
                       </div>
                     </div>
                     
                     <button className="text-gray-300 hover:text-gray-500">
                        <MoreVertical size={16} />
                     </button>
                   </div>
                 </div>
               )})}
             </div>
           );
        })}
      </div>
    </div>
  );
};

export default TaskView;
