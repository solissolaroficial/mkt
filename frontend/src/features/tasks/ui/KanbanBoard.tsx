 
import React, { useState, useMemo } from 'react';
import type { Task, TaskStatus } from '@/shared/types';
import { 
  Plus, 
  CheckCircle2, 
  Circle, 
  Loader2, 
  User,
  MessageSquare,
  CheckSquare,
  Archive,
  RefreshCcw,
  Inbox,
  CalendarDays,
  BarChart2,
  Layout,
  CornerDownRight,
  Layers,
  ChevronLeft,
  ChevronRight,
  Users
} from 'lucide-react';
import type { KanbanBoardProps } from '../types';

const KanbanBoard: React.FC<KanbanBoardProps> = ({ 
  tasks, 
  onAddTask, 
  onUpdateTask, 
  onDeleteTask, 
  onReorderTasks,
  onTaskClick 
}) => {
  const [newTaskTitle, setNewTaskTitle] = useState('');
  const [isAdding, setIsAdding] = useState(false);
  const [showArchived, setShowArchived] = useState(false);
  const [viewMode, setViewMode] = useState<'board' | 'gantt'>('board');
  const [ganttScale, setGanttScale] = useState<'day' | 'week' | 'month'>('day');
  
  // --- FLOWS STATE ---
  const [flows, setFlows] = useState<string[]>(['Jackson', 'Beatriz', 'Larissa']);
  const [selectedFlow, setSelectedFlow] = useState<string>('Jackson');
  const [isFlowSidebarOpen, setIsFlowSidebarOpen] = useState(true);
  const [isAddingFlow, setIsAddingFlow] = useState(false);
  const [newFlowName, setNewFlowName] = useState('');

  // Drag and Drop State
  const [draggingTaskId, setDraggingTaskId] = useState<string | null>(null);
  const [activeDropColumn, setActiveDropColumn] = useState<TaskStatus | null>(null);
  const [dropIndex, setDropIndex] = useState<number | null>(null);

  const columns: { id: TaskStatus; title: string; color: string; hoverColor: string; icon: any }[] = [
    { id: 'todo', title: 'A Fazer', color: 'bg-gray-800/50', hoverColor: 'bg-gray-800', icon: Circle },
    { id: 'in_progress', title: 'Em Andamento', color: 'bg-blue-900/10', hoverColor: 'bg-blue-900/20', icon: Loader2 },
    { id: 'done', title: 'Concluído', color: 'bg-green-900/10', hoverColor: 'bg-green-900/20', icon: CheckCircle2 }
  ];

  const draggedTask = tasks.find(t => t.id === draggingTaskId);

  const toggleArchiveTask = (taskId: string, e: React.MouseEvent) => {
    e.stopPropagation();
    const task = tasks.find(t => t.id === taskId);
    if (task) {
      onUpdateTask({ ...task, archived: !task.archived });
    }
  };

  const formatDateDisplay = (dateStr?: string) => {
    if (!dateStr || dateStr === 'Sem data') return 'Sem data';
    if (dateStr.match(/^\d{4}-\d{2}-\d{2}$/)) {
        const [y, m, d] = dateStr.split('-');
        return `${d}/${m}`;
    }
    return dateStr;
  }

  const addTask = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTaskTitle.trim()) return;

    const newTask: Task = {
      id: Date.now().toString(),
      title: newTaskTitle,
      description: '',
      dueDate: new Date().toISOString().split('T')[0], // Default today
      priority: 'media',
      status: 'todo',
      category: 'admin',
      assignee: 'Jackson',
      flows: [selectedFlow], // Auto-assign current flow
      subtasks: [],
      comments: [],
      archived: false
    };

    onAddTask(newTask);
    setNewTaskTitle('');
    setIsAdding(false);
  };

  const handleAddFlow = (e: React.FormEvent) => {
      e.preventDefault();
      if (!newFlowName.trim()) return;
      if (!flows.includes(newFlowName.trim())) {
          setFlows([...flows, newFlowName.trim()]);
          setSelectedFlow(newFlowName.trim());
      }
      setNewFlowName('');
      setIsAddingFlow(false);
  };

  const getPriorityColor = (p: string) => {
    switch (p) {
      case 'alta': return 'text-rose-400 bg-rose-500/10 border-rose-500/20';
      case 'media': return 'text-amber-400 bg-amber-500/10 border-amber-500/20';
      default: return 'text-blue-400 bg-blue-500/10 border-blue-500/20';
    }
  };

  // --- Drag and Drop Handlers ---
  const handleDragStart = (e: React.DragEvent, taskId: string) => {
    e.dataTransfer.effectAllowed = 'move';
    e.dataTransfer.setData('text/plain', taskId);

    const target = e.currentTarget as HTMLElement;
    const ghost = target.cloneNode(true) as HTMLElement;
    
    ghost.style.position = 'absolute';
    ghost.style.top = '-1000px'; 
    ghost.style.width = `${target.offsetWidth}px`;
    ghost.style.backgroundColor = '#1a1d24'; 
    ghost.style.color = '#f3f4f6'; 
    ghost.style.borderRadius = '0.5rem';
    ghost.style.padding = '1rem';
    ghost.style.border = '1px solid #374151';
    ghost.style.boxShadow = '0 20px 25px -5px rgba(0, 0, 0, 0.5), 0 10px 10px -5px rgba(0, 0, 0, 0.3)';
    ghost.style.transform = 'rotate(3deg)';
    ghost.style.opacity = '1';
    ghost.style.zIndex = '1000';
    
    document.body.appendChild(ghost);
    e.dataTransfer.setDragImage(ghost, target.offsetWidth / 2, target.offsetHeight / 2);

    setTimeout(() => {
        document.body.removeChild(ghost);
    }, 0);

    setTimeout(() => {
        setDraggingTaskId(taskId);
    }, 10);
  };

  const handleDragOver = (e: React.DragEvent, colId: TaskStatus) => {
    e.preventDefault(); 
    e.stopPropagation();
    e.dataTransfer.dropEffect = 'move';
    
    if (activeDropColumn !== colId) {
        setActiveDropColumn(colId);
    }

    const container = e.currentTarget as HTMLElement;
    const cards = Array.from(container.querySelectorAll('[data-task-id]')).filter((el) => {
      return el.getAttribute('data-dragging') !== 'true'; 
    });
    
    const afterElement = cards.reduce((closest: { offset: number; element: Element | null }, child) => {
      const box = child.getBoundingClientRect();
      const offset = e.clientY - box.top - box.height / 2;
      
      if (offset < 0 && offset > closest.offset) {
        return { offset: offset, element: child };
      } else {
        return closest;
      }
    }, { offset: Number.NEGATIVE_INFINITY, element: null }).element;

    let newIndex;
    if (afterElement) {
        const taskId = afterElement.getAttribute('data-task-id');
        const tasksInColumn = visibleTasks.filter(t => t.status === colId && t.id !== draggingTaskId);
        newIndex = tasksInColumn.findIndex(t => t.id === taskId);
    } else {
        const tasksInColumn = visibleTasks.filter(t => t.status === colId && t.id !== draggingTaskId);
        newIndex = tasksInColumn.length;
    }
    
    if (dropIndex !== newIndex) {
      setDropIndex(newIndex);
    }
  };

  const handleDragLeave = () => {
    // Optional
  };

  const handleDragEnd = () => {
    setDraggingTaskId(null);
    setActiveDropColumn(null);
    setDropIndex(null);
  };

  const handleDrop = (e: React.DragEvent, targetStatus: TaskStatus) => {
    e.preventDefault();
    
    if (draggingTaskId) {
      const draggedTaskItem = tasks.find(t => t.id === draggingTaskId);
      
      if (draggedTaskItem) {
        const updatedTask = { ...draggedTaskItem, status: targetStatus };
        // We only reorder locally relative to other tasks, but since tasks list is global,
        // we might just update status. 
        // For simple Kanban, updating status is usually enough unless strict ordering is persisted.
        // Given complexity, let's just update status for now to satisfy requirement.
        onUpdateTask(updatedTask);
      }
    }
    
    setDraggingTaskId(null);
    setActiveDropColumn(null);
    setDropIndex(null);
  };

  const renderGhostCard = (task: Task) => {
    return (
        <div 
          className="bg-gray-800/30 p-4 rounded-lg border-2 border-dashed border-gray-600 opacity-60 pointer-events-none mb-3 transform transition-all duration-200"
          style={{ height: '140px' }}
        >
            <div className={`w-1 h-8 rounded-r absolute left-0 top-4 ${
                task.priority === 'alta' ? 'bg-rose-500' : 
                task.priority === 'media' ? 'bg-amber-500' : 'bg-blue-500'
            }`}></div>
            <div className="pl-3 opacity-50">
                <h4 className="text-sm font-medium text-gray-400">{task.title}</h4>
                <div className="h-2 w-1/2 bg-gray-700 rounded mt-2"></div>
            </div>
        </div>
    );
  };

  // --- FILTER TASKS ---
  const visibleTasks = tasks.filter(t => {
      const isArchivedMatch = showArchived ? t.archived : !t.archived;
      const isFlowMatch = t.flows?.includes(selectedFlow);
      return isArchivedMatch && isFlowMatch;
  });
  
  const isToday = (dateStr: string) => {
     if (!dateStr || dateStr === 'Sem data') return false;
     const todayStr = new Date().toISOString().split('T')[0];
     return dateStr === todayStr;
  }

  // --- GANTT CHART LOGIC ---
  const parseDate = (str?: string) => {
      if (!str || str === 'Sem data') return null;
      if (str.match(/^\d{4}-\d{2}-\d{2}$/)) {
          const [y, m, d] = str.split('-').map(Number);
          return new Date(y, m - 1, d);
      }
      return null;
  };

  // SCALE CONFIGURATIONS
  const SCALE_CONFIG = {
      day: { 
          pxPerDay: 30, // Compact day width
          tickWidth: 30,
          headerFormat: (d: Date) => d.getDate().toString(),
          subHeaderFormat: (d: Date) => d.toLocaleDateString('pt-BR', { weekday: 'short' }).slice(0,1),
          stepDays: 1
      },
      week: { 
          pxPerDay: 14.28, // 100px per week
          tickWidth: 100, 
          headerFormat: (d: Date) => {
              const endOfWeek = new Date(d);
              endOfWeek.setDate(d.getDate() + 6);
              return `${d.getDate()}/${d.getMonth()+1} - ${endOfWeek.getDate()}/${endOfWeek.getMonth()+1}`;
          },
          subHeaderFormat: (d: Date) => `Sem ${getWeekNumber(d)}`,
          stepDays: 7
      },
      month: { 
          pxPerDay: 5, // 150px per 30 days approx
          tickWidth: 150, 
          headerFormat: (d: Date) => d.toLocaleDateString('pt-BR', { month: 'long', year: 'numeric' }),
          subHeaderFormat: () => '',
          stepDays: 30 // Variable actually
      }
  };

  // Helper to get ISO week number
  function getWeekNumber(d: Date) {
      const date = new Date(d.getTime());
      date.setHours(0, 0, 0, 0);
      date.setDate(date.getDate() + 3 - (date.getDay() + 6) % 7);
      const week1 = new Date(date.getFullYear(), 0, 4);
      return 1 + Math.round(((date.getTime() - week1.getTime()) / 86400000 - 3 + (week1.getDay() + 6) % 7) / 7);
  }

  const ganttRange = useMemo(() => {
      if (visibleTasks.length === 0) return { start: new Date(), end: new Date() };

      let min = new Date(8640000000000000);
      let max = new Date(-8640000000000000);
      
      const tasksWithDates = visibleTasks.filter(t => t.startDate || t.dueDate !== 'Sem data');
      if (tasksWithDates.length === 0) return { start: new Date(), end: new Date() };

      tasksWithDates.forEach(t => {
          const s = parseDate(t.startDate);
          const e = parseDate(t.dueDate);
          
          const start = s || (e ? new Date(e.getTime() - 2*24*60*60*1000) : new Date());
          const end = e || new Date();

          if (start < min) min = start;
          if (end > max) max = end;
      });

      // Buffer based on scale
      min.setDate(min.getDate() - 7);
      max.setDate(max.getDate() + 14);

      return { start: min, end: max };
  }, [visibleTasks]);

  const ganttTicks = useMemo(() => {
      const { start, end } = ganttRange;
      const ticks = [];
      const current = new Date(start);
      
      // Align start date based on scale
      if (ganttScale === 'week') {
          const day = current.getDay();
          const diff = current.getDate() - day + (day === 0 ? -6 : 1); // Adjust when day is sunday
          current.setDate(diff); // Monday
      } else if (ganttScale === 'month') {
          current.setDate(1);
      }

      while (current <= end) {
          ticks.push(new Date(current));
          
          if (ganttScale === 'day') current.setDate(current.getDate() + 1);
          else if (ganttScale === 'week') current.setDate(current.getDate() + 7);
          else if (ganttScale === 'month') current.setMonth(current.getMonth() + 1);
      }
      return ticks;
  }, [ganttRange, ganttScale]);

  const getBarStyles = (task: Task) => {
      const s = parseDate(task.startDate);
      const e = parseDate(task.dueDate);
      
      const start = s || (e ? new Date(e.getTime() - 2*24*60*60*1000) : new Date()); 
      const end = e || start;

      // Align to Gantt Start
      let effectiveStart = new Date(ganttRange.start);
      if (ganttScale === 'week') {
          const day = effectiveStart.getDay();
          effectiveStart.setDate(effectiveStart.getDate() - day + (day === 0 ? -6 : 1));
      } else if (ganttScale === 'month') {
          effectiveStart.setDate(1);
      }

      const diffMs = start.getTime() - effectiveStart.getTime();
      const diffDays = diffMs / (1000 * 60 * 60 * 24);
      
      const durationMs = end.getTime() - start.getTime();
      const durationDays = Math.max(1, durationMs / (1000 * 60 * 60 * 24) + 1);

      const pxPerDay = SCALE_CONFIG[ganttScale].pxPerDay;

      return {
          left: diffDays * pxPerDay,
          width: durationDays * pxPerDay
      };
  };

  const currentConfig = SCALE_CONFIG[ganttScale];
  const totalWidth = ganttTicks.length * currentConfig.tickWidth;

  return (
    <div className="h-full flex overflow-hidden animate-in fade-in slide-in-from-bottom-4 duration-500">
      
      {/* --- FLOWS SIDEBAR --- */}
      <div 
        className={`bg-[#12141a] border-r border-gray-800 flex flex-col transition-all duration-300 relative ${isFlowSidebarOpen ? 'w-64' : 'w-0 border-none'}`}
      >
        <div className={`p-6 overflow-hidden ${isFlowSidebarOpen ? 'opacity-100' : 'opacity-0'} transition-opacity`}>
            <div className="flex items-center justify-between mb-6">
                <h2 className="text-xl font-bold text-gray-100 text-transparent bg-clip-text bg-gradient-to-r from-pink-500 to-rose-500">
                    Fluxos
                </h2>
                <div className="p-1 rounded bg-pink-500/10 text-pink-500">
                    <Layers size={16} />
                </div>
            </div>

            <div className="space-y-2">
                {flows.map(flow => (
                    <button
                        key={flow}
                        onClick={() => setSelectedFlow(flow)}
                        className={`w-full text-left px-4 py-3 rounded-xl transition-all font-medium text-sm flex items-center justify-between group ${
                            selectedFlow === flow 
                            ? 'bg-pink-500/10 text-pink-400 border border-pink-500/20' 
                            : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'
                        }`}
                    >
                        {flow}
                        {selectedFlow === flow && <div className="w-2 h-2 rounded-full bg-pink-500"></div>}
                    </button>
                ))}
            </div>

            {/* Add Flow */}
            {isAddingFlow ? (
                <form onSubmit={handleAddFlow} className="mt-4 animate-in fade-in">
                    <input 
                        type="text" 
                        value={newFlowName}
                        onChange={(e) => setNewFlowName(e.target.value)}
                        placeholder="Nome do Fluxo..."
                        className="w-full px-3 py-2 bg-[#1a1d24] border border-gray-700 rounded-lg text-sm text-white focus:outline-none focus:ring-1 focus:ring-pink-500 mb-2"
                        autoFocus
                    />
                    <div className="flex gap-2">
                        <button type="submit" className="flex-1 bg-pink-600 hover:bg-pink-700 text-white text-xs py-1.5 rounded transition-colors">Salvar</button>
                        <button type="button" onClick={() => setIsAddingFlow(false)} className="flex-1 bg-gray-800 hover:bg-gray-700 text-gray-400 text-xs py-1.5 rounded transition-colors">Cancelar</button>
                    </div>
                </form>
            ) : (
                <button 
                    onClick={() => setIsAddingFlow(true)}
                    className="w-full mt-4 py-2 border border-dashed border-gray-700 rounded-lg text-gray-500 text-xs hover:border-pink-500/50 hover:text-pink-400 transition-colors flex items-center justify-center gap-2"
                >
                    <Plus size={14} /> Novo Fluxo
                </button>
            )}
        </div>
      </div>

      {/* --- TOGGLE BUTTON FOR SIDEBAR --- */}
      <div className="relative z-10">
          <button 
            onClick={() => setIsFlowSidebarOpen(!isFlowSidebarOpen)}
            className="absolute top-6 -right-3 w-6 h-6 bg-[#1a1d24] border border-gray-700 rounded-full flex items-center justify-center text-gray-400 hover:text-white transition-colors z-20"
            style={{ left: isFlowSidebarOpen ? '-12px' : '12px' }}
          >
              {isFlowSidebarOpen ? <ChevronLeft size={14} /> : <ChevronRight size={14} />}
          </button>
      </div>

      {/* --- MAIN CONTENT --- */}
      <div className="flex-grow flex flex-col min-w-0 p-6"> {/* Padding moved here from container */}
        
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-6 gap-4">
            <div>
            <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-2">
                Quadro de Tarefas
                <span className="text-sm font-normal text-pink-400 bg-pink-500/10 px-2 py-0.5 rounded border border-pink-500/20">
                    {selectedFlow}
                </span>
            </h1>
            <p className="text-gray-500">
                {showArchived ? 'Histórico de tarefas concluídas' : 'Gerenciamento visual de fluxo de trabalho'}
            </p>
            </div>
            
            <div className="flex gap-2">
                {viewMode === 'gantt' && (
                    <div className="flex bg-[#1a1d24] border border-gray-700 rounded-lg p-1 mr-2">
                        <button 
                            onClick={() => setGanttScale('day')}
                            className={`px-3 py-1 rounded text-xs font-medium transition-colors ${ganttScale === 'day' ? 'bg-[#1e5144] text-white' : 'text-gray-400 hover:text-white'}`}
                        >
                            Dia
                        </button>
                        <button 
                            onClick={() => setGanttScale('week')}
                            className={`px-3 py-1 rounded text-xs font-medium transition-colors ${ganttScale === 'week' ? 'bg-[#1e5144] text-white' : 'text-gray-400 hover:text-white'}`}
                        >
                            Semana
                        </button>
                        <button 
                            onClick={() => setGanttScale('month')}
                            className={`px-3 py-1 rounded text-xs font-medium transition-colors ${ganttScale === 'month' ? 'bg-[#1e5144] text-white' : 'text-gray-400 hover:text-white'}`}
                        >
                            Mês
                        </button>
                    </div>
                )}

                <button 
                    onClick={() => setViewMode(viewMode === 'board' ? 'gantt' : 'board')}
                    className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 border ${
                        viewMode === 'gantt' 
                        ? 'bg-[#1e5144] text-white border-[#1e5144]' 
                        : 'bg-[#1a1d24] text-gray-400 border-gray-700 hover:bg-gray-800'
                    }`}
                >
                    {viewMode === 'gantt' ? <Layout size={18} /> : <BarChart2 size={18} className="rotate-90" />}
                    {viewMode === 'gantt' ? 'Ver Quadro' : 'Ver Gantt'}
                </button>

                {!showArchived && (
                <button 
                    onClick={() => setIsAdding(!isAdding)}
                    className="bg-[#1e5144] text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-[#163c32] transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/30"
                >
                    <Plus size={18} /> Nova Tarefa
                </button>
                )}
                
                <button 
                    onClick={() => setShowArchived(!showArchived)}
                    className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 border ${
                        showArchived 
                        ? 'bg-gray-700 text-white border-gray-600' 
                        : 'bg-[#1a1d24] text-gray-400 border-gray-700 hover:bg-gray-800'
                    }`}
                >
                    {showArchived ? <Inbox size={18} /> : <Archive size={18} />}
                    {showArchived ? 'Ver Quadro Ativo' : 'Ver Arquivadas'}
                </button>
            </div>
        </div>

        {isAdding && !showArchived && (
            <form onSubmit={addTask} className="mb-6 bg-[#1a1d24] p-4 rounded-xl border border-gray-700 shadow-xl flex gap-4 animate-in fade-in slide-in-from-top-2">
            <input 
                type="text" 
                value={newTaskTitle}
                onChange={(e) => setNewTaskTitle(e.target.value)}
                placeholder="Descreva a nova tarefa..."
                className="flex-grow px-4 py-2 rounded-lg border border-gray-700 focus:outline-none focus:ring-2 focus:ring-[#1e5144] bg-[#0f1115] text-white placeholder-gray-500"
                autoFocus
            />
            <button type="submit" className="px-4 py-2 bg-gray-100 text-gray-900 rounded-lg font-medium hover:bg-white transition-colors">
                Adicionar
            </button>
            <button type="button" onClick={() => setIsAdding(false)} className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg transition-colors">
                Cancelar
            </button>
            </form>
        )}

        {viewMode === 'board' ? (
        <div className="flex-grow overflow-x-auto pb-4">
            <div className="flex gap-6 h-full min-w-[1000px]">
            {columns.map(col => {
                const colTasks = visibleTasks.filter(t => t.status === col.id);
                const Icon = col.icon;
                const isDropTarget = activeDropColumn === col.id;
                
                return (
                <div 
                    key={col.id} 
                    className="flex-1 flex flex-col h-full rounded-xl overflow-hidden"
                    onDragOver={(e) => handleDragOver(e, col.id)}
                    onDragLeave={handleDragLeave}
                    onDrop={(e) => handleDrop(e, col.id)}
                >
                    <div className={`flex items-center justify-between p-3 border-b border-gray-700/50 transition-colors duration-200 ${isDropTarget ? 'bg-[#1e5144]/20' : col.color}`}>
                    <div className="flex items-center gap-2 font-semibold text-gray-300">
                        <Icon size={18} className={col.id === 'in_progress' ? 'animate-spin-slow' : ''} />
                        {col.title}
                        <span className="bg-gray-800/50 text-xs px-2 py-0.5 rounded-full border border-gray-700">
                        {colTasks.length}
                        </span>
                    </div>
                    </div>

                    <div className={`flex-grow p-3 space-y-3 overflow-y-auto custom-scrollbar transition-colors duration-300 ${isDropTarget ? 'bg-[#1e5144]/10' : 'bg-[#15171c]'}`}>
                    
                    {colTasks.map((task, index) => {
                        const completedSubtasks = task.subtasks?.filter(s => s.completed).length || 0;
                        const totalSubtasks = task.subtasks?.length || 0;
                        const isDragging = draggingTaskId === task.id;
                        const taskIsToday = isToday(task.dueDate);
                        
                        const showGhostBefore = draggedTask && isDropTarget && dropIndex === index && draggedTask.id !== task.id;

                        return (
                        <React.Fragment key={task.id}>
                            {showGhostBefore && renderGhostCard(draggedTask)}
                            
                            <div 
                            draggable={!showArchived}
                            data-task-id={task.id}
                            data-dragging={isDragging}
                            onDragStart={(e) => handleDragStart(e, task.id)}
                            onDragEnd={handleDragEnd}
                            onClick={() => onTaskClick(task.id)}
                            className={`
                                bg-[#1a1d24] p-4 rounded-lg border transition-all group relative 
                                ${isDragging 
                                    ? 'hidden' 
                                    : 'opacity-100 border-gray-700 shadow-lg hover:shadow-xl hover:-translate-y-1 hover:border-gray-600'
                                }
                                ${showArchived ? 'opacity-80' : isDragging ? '' : 'cursor-grab active:cursor-grabbing'}
                            `}
                            >
                            <div className={`absolute left-0 top-3 bottom-3 w-1 rounded-r ${
                                task.priority === 'alta' ? 'bg-rose-500' : 
                                task.priority === 'media' ? 'bg-amber-500' : 'bg-blue-500'
                            }`}></div>

                            <div className="pl-3">
                                <div className="flex justify-between items-start mb-2">
                                <span className={`text-[10px] px-2 py-0.5 rounded-full border font-bold uppercase tracking-wide ${getPriorityColor(task.priority)}`}>
                                    {task.priority}
                                </span>
                                
                                {(task.status === 'done' || showArchived) && (
                                    <button 
                                        onClick={(e) => toggleArchiveTask(task.id, e)}
                                        className="text-gray-500 hover:text-emerald-400 p-1 rounded-md hover:bg-gray-800 transition-colors"
                                        title={showArchived ? "Desarquivar" : "Arquivar"}
                                    >
                                        {showArchived ? <RefreshCcw size={14} /> : <Archive size={14} />}
                                    </button>
                                )}
                                </div>

                                <h4 className={`text-sm font-medium mb-3 leading-snug ${task.status === 'done' && !isDragging ? 'text-gray-500 line-through decoration-gray-600' : 'text-gray-200'}`}>
                                {task.title}
                                </h4>

                                <div className="flex items-center justify-between text-xs text-gray-500 border-t border-gray-800 pt-3 mt-3">
                                <div className="flex items-center gap-3">
                                    <div className="flex items-center gap-1.5" title="Responsável">
                                        <div className="w-5 h-5 rounded-full bg-gray-800 flex items-center justify-center border border-gray-700 text-gray-300 font-bold">
                                            {task.assignee ? task.assignee[0] : <User size={12} />}
                                        </div>
                                    </div>
                                    {(task.comments?.length || 0) > 0 && (
                                        <div className="flex items-center gap-1 text-gray-400">
                                            <MessageSquare size={12} /> {task.comments?.length}
                                        </div>
                                    )}
                                    {totalSubtasks > 0 && (
                                        <div className={`flex items-center gap-1 ${completedSubtasks === totalSubtasks ? 'text-emerald-500' : 'text-gray-400'}`}>
                                            <CheckSquare size={12} /> {completedSubtasks}/{totalSubtasks}
                                        </div>
                                    )}
                                </div>
                                
                                <div className="flex items-center gap-1" title="Prazo">
                                    <CalendarDays size={12} className={taskIsToday ? 'text-amber-500' : ''} />
                                    <span className={taskIsToday ? 'text-amber-500 font-medium' : ''}>{formatDateDisplay(task.dueDate)}</span>
                                </div>
                                </div>
                            </div>
                            </div>
                        </React.Fragment>
                        );
                    })}
                    
                    {draggedTask && isDropTarget && dropIndex === colTasks.length && renderGhostCard(draggedTask)}

                    {!showArchived && (
                        <button 
                            onClick={() => { setIsAdding(true); }}
                            className="w-full py-2 border-2 border-dashed border-gray-700 rounded-lg text-gray-500 text-sm hover:border-[#1e5144]/50 hover:text-emerald-400 transition-colors flex items-center justify-center gap-2 group-hover:bg-[#1a1d24]"
                        >
                            <Plus size={16} /> Adicionar
                        </button>
                    )}
                    </div>
                </div>
                );
            })}
            </div>
        </div>
        ) : (
            <div className="flex-grow bg-[#1a1d24] rounded-xl border border-gray-700 overflow-hidden flex flex-col animate-in fade-in">
            {ganttTicks.length > 0 ? (
                <div className="flex-grow overflow-auto custom-scrollbar relative">
                    {/* Header Row */}
                    <div className="flex sticky top-0 z-20 bg-[#20232b] border-b border-gray-700 min-w-max">
                        <div className="sticky left-0 z-30 w-64 p-3 bg-[#20232b] border-r border-gray-700 font-semibold text-gray-300 shadow-md">
                            Tarefa
                        </div>
                        {ganttTicks.map((date, i) => (
                            <div key={i} className="flex-shrink-0 border-r border-gray-700 p-2 text-center flex flex-col justify-center" style={{ width: currentConfig.tickWidth }}>
                                <div className="text-[10px] text-gray-500 uppercase">{currentConfig.subHeaderFormat(date)}</div>
                                <div className={`text-xs font-bold text-gray-300`}>
                                    {currentConfig.headerFormat(date)}
                                </div>
                            </div>
                        ))}
                    </div>
                    
                    {/* Task Rows */}
                    <div className="min-w-max">
                        {visibleTasks.filter(t => t.startDate || t.dueDate !== 'Sem data').map(task => {
                            const { left, width } = getBarStyles(task);
                            return (
                                <div key={task.id} className="flex border-b border-gray-800 hover:bg-[#20232b]/50 transition-colors group">
                                    <div 
                                        className="sticky left-0 z-10 w-64 p-3 border-r border-gray-800 bg-[#1a1d24] group-hover:bg-[#20232b] flex items-center justify-between shadow-sm cursor-pointer"
                                        onClick={() => onTaskClick(task.id)}
                                    >
                                        <div className="truncate pr-2 text-sm text-gray-300 font-medium">{task.title}</div>
                                        <div className={`w-2 h-2 rounded-full flex-shrink-0 ${task.priority === 'alta' ? 'bg-rose-500' : task.priority === 'media' ? 'bg-amber-500' : 'bg-blue-500'}`}></div>
                                    </div>
                                    <div className="relative flex-grow h-10" style={{ width: totalWidth }}>
                                        {/* Grid Lines */}
                                        {ganttTicks.map((date, idx) => (
                                            <div key={idx} className="absolute top-0 bottom-0 border-r border-gray-800/50" style={{ left: idx * currentConfig.tickWidth, width: currentConfig.tickWidth }}></div>
                                        ))}
                                        
                                        {/* Bar */}
                                        <div 
                                            className={`absolute top-2 bottom-2 rounded-md flex items-center px-2 text-xs text-white overflow-hidden whitespace-nowrap shadow-sm cursor-pointer hover:brightness-110 transition-all ${
                                                task.status === 'done' ? 'bg-emerald-500/50' : 
                                                task.priority === 'alta' ? 'bg-rose-500/70' : 
                                                task.priority === 'media' ? 'bg-amber-500/70' : 'bg-blue-500/70'
                                            }`}
                                            style={{ left: `${left}px`, width: `${width}px` }}
                                            onClick={() => onTaskClick(task.id)}
                                            title={`${task.title} (${formatDateDisplay(task.startDate)} - ${formatDateDisplay(task.dueDate)})`}
                                        >
                                            {width > 60 && <span className="drop-shadow-md">{task.assignee}</span>}
                                        </div>
                                    </div>
                                </div>
                            );
                        })}
                    </div>
                </div>
            ) : (
                <div className="flex-grow flex items-center justify-center text-gray-500 p-8">
                    <p>Nenhuma tarefa com data definida para exibir no gráfico.</p>
                </div>
            )}
            </div>
        )}
      </div>
    </div>
  );
};

export default KanbanBoard;
