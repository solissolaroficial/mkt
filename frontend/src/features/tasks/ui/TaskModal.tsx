 
import React, { useState } from 'react';
import type { Task, Subtask, Comment, AllowedUser } from '@/shared/types';
import { ALLOWED_USERS } from '@/shared/utils/legacy.constants';
import { 
  X, 
  CheckSquare, 
  MessageSquare, 
  User, 
  Trash2, 
  Send, 
  AlignLeft,
  CalendarDays,
  Calendar,
  Archive,
  RefreshCcw,
  Plus,
  Layers,
  ChevronDown,
  Check
} from 'lucide-react';
import type { TaskModalProps } from '../types';

// Temporary available flows - in a real app this would come from props or context
const AVAILABLE_FLOWS = ['Jackson', 'Beatriz', 'Larissa', 'Geral'];

const TaskModal: React.FC<TaskModalProps> = ({ task, isOpen, onClose, onUpdate, onDelete, onMention }) => {
  const [newSubtaskTitle, setNewSubtaskTitle] = useState('');
  const [newSubtaskAssignee, setNewSubtaskAssignee] = useState<AllowedUser>('Jackson');
  const [newSubtaskDate, setNewSubtaskDate] = useState('');
  
  const [newComment, setNewComment] = useState('');
  const [activeUser] = useState<AllowedUser>('Jackson'); 
  
  // Mention Autocomplete State
  const [showMentions, setShowMentions] = useState(false);
  const [mentionFilter, setMentionFilter] = useState('');

  // Flow Dropdown State
  const [isFlowDropdownOpen, setIsFlowDropdownOpen] = useState(false);

  if (!isOpen || !task) return null; // Safety check

  // Handlers
  const handleDescriptionChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    onUpdate({ ...task, description: e.target.value });
  };

  const toggleSubtask = (subtaskId: string) => {
    const updatedSubtasks = task.subtasks?.map(st => 
      st.id === subtaskId ? { ...st, completed: !st.completed } : st
    ) || [];
    onUpdate({ ...task, subtasks: updatedSubtasks });
  };

  const updateSubtaskField = (subtaskId: string, field: keyof Subtask, value: any) => {
    const updatedSubtasks = task.subtasks?.map(st => 
      st.id === subtaskId ? { ...st, [field]: value } : st
    ) || [];
    onUpdate({ ...task, subtasks: updatedSubtasks });
  };

  const addSubtask = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newSubtaskTitle.trim()) return;
    const newSub: Subtask = {
      id: Date.now().toString(),
      title: newSubtaskTitle,
      completed: false,
      assignee: newSubtaskAssignee,
      dueDate: newSubtaskDate || new Date().toISOString().split('T')[0]
    };
    onUpdate({ ...task, subtasks: [...(task.subtasks || []), newSub] });
    setNewSubtaskTitle('');
    setNewSubtaskDate('');
  };

  const deleteSubtask = (subtaskId: string) => {
    const updatedSubtasks = task.subtasks?.filter(st => st.id !== subtaskId) || [];
    onUpdate({ ...task, subtasks: updatedSubtasks });
  };

  const addComment = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newComment.trim()) return;
    
    const lowerComment = newComment.toLowerCase();
    
    if (onMention) {
        if (lowerComment.includes('@jackson')) {
            onMention(task.id, task.title);
        }
    }

    const comment: Comment = {
      id: Date.now().toString(),
      user: activeUser,
      text: newComment,
      timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    };
    onUpdate({ ...task, comments: [...(task.comments || []), comment] });
    setNewComment('');
    setShowMentions(false);
  };

  const toggleArchive = () => {
      onUpdate({ ...task, archived: !task.archived });
  };

  const completedSubtasks = task.subtasks?.filter(s => s.completed).length || 0;
  const totalSubtasks = task.subtasks?.length || 0;
  const progress = totalSubtasks === 0 ? 0 : (completedSubtasks / totalSubtasks) * 100;

  // Helper to safely open date picker
  const openDatePicker = (e: React.MouseEvent<HTMLInputElement>) => {
    try {
      if (e.currentTarget && 'showPicker' in e.currentTarget) {
         // @ts-ignore
         e.currentTarget.showPicker();
      }
    } catch (error) {}
  };

  const renderCommentText = (text: string) => {
    return text.split(' ').map((word, index) => {
        if (word.startsWith('@')) {
            return (
                <span key={index} className="text-[#1e5144] font-bold bg-[#1e5144]/10 px-1 rounded mx-0.5">
                    {word}
                </span>
            );
        }
        return <span key={index}>{word} </span>;
    });
  };

  // Mention Autocomplete Logic
  const handleCommentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
      const val = e.target.value;
      setNewComment(val);
      
      const lastWord = val.split(/[\s\n]+/).pop();
      if (lastWord && lastWord.startsWith('@')) {
          setShowMentions(true);
          setMentionFilter(lastWord.substring(1).toLowerCase());
      } else {
          setShowMentions(false);
      }
  };

  const insertMention = (user: string) => {
      const words = newComment.split(/[\s\n]+/);
      words.pop(); // Remove partial mention
      const text = words.join(' ') + (words.length > 0 ? ' ' : '') + `@${user} `;
      setNewComment(text);
      setShowMentions(false);
  };

  const toggleFlow = (flow: string) => {
      const currentFlows = task.flows || [];
      let newFlows;
      if (currentFlows.includes(flow)) {
          newFlows = currentFlows.filter(f => f !== flow);
      } else {
          newFlows = [...currentFlows, flow];
      }
      onUpdate({ ...task, flows: newFlows });
  };

  const filteredUsers = ALLOWED_USERS.filter(u => u.toLowerCase().includes(mentionFilter));

  return (
    <div className="fixed inset-0 bg-black/80 z-50 flex justify-center items-center p-4 backdrop-blur-sm" onClick={onClose}>
      
      {/* Styles for date picker override */}
      <style>{`
        .date-input-wrapper {
          position: relative;
          color-scheme: dark;
        }
        .date-input-wrapper input[type="date"]::-webkit-calendar-picker-indicator {
          position: absolute;
          top: 0;
          left: 0;
          width: 100%;
          height: 100%;
          margin: 0;
          padding: 0;
          opacity: 0;
          cursor: pointer;
        }
      `}</style>

      <div 
        className="bg-[#1a1d24] w-full max-w-5xl h-[85vh] rounded-2xl shadow-2xl border border-gray-800 flex overflow-hidden text-gray-200 relative animate-in zoom-in-95 duration-200" 
        onClick={e => e.stopPropagation()}
      >
        {/* CLOSE BUTTON */}
        <button 
          onClick={onClose} 
          className="absolute top-4 right-4 z-50 text-gray-400 hover:text-white p-2 hover:bg-white/10 rounded-full transition-colors"
        >
          <X size={24} />
        </button>

        {/* MAIN COLUMN */}
        <div className="flex-1 flex flex-col min-w-0 bg-[#1a1d24] overflow-hidden">
            
            {/* Header */}
            <div className="p-8 pb-4">
                <input 
                    type="text" 
                    value={task.title}
                    onChange={(e) => onUpdate({...task, title: e.target.value})}
                    className="text-2xl md:text-3xl font-bold text-white bg-transparent border-none p-0 focus:ring-0 w-full placeholder-gray-600 focus:outline-none"
                    placeholder="Título da tarefa..."
                />
                <div className="flex items-center gap-3 mt-4 text-sm">
                    <span className="bg-[#20232b] text-gray-300 px-3 py-1 rounded-md text-xs uppercase font-bold tracking-wider border border-gray-700">
                        {task.category}
                    </span>
                    <span className="text-gray-500">na lista</span>
                    <span className="font-medium text-emerald-400 border-b border-emerald-400/30 pb-0.5">
                        {task.status === 'todo' ? 'A Fazer' : task.status === 'in_progress' ? 'Em Andamento' : 'Concluído'}
                    </span>
                    {task.archived && <span className="text-amber-500 text-xs font-bold px-2 py-0.5 bg-amber-500/10 rounded border border-amber-500/20">ARQUIVADA</span>}
                </div>
            </div>

            {/* Scrollable Content */}
            <div className="flex-1 overflow-y-auto custom-scrollbar p-8 pt-2 space-y-8">
                
                {/* Description */}
                <div>
                    <div className="flex items-center gap-3 mb-3 text-gray-300">
                        <AlignLeft size={20} />
                        <h3 className="font-semibold text-lg">Descrição</h3>
                    </div>
                    <div className="bg-[#0f1115] border border-gray-800 rounded-xl p-4 transition-colors focus-within:border-gray-700">
                        <textarea 
                            value={task.description || ''}
                            onChange={handleDescriptionChange}
                            placeholder="Adicione uma descrição mais detalhada..."
                            className="w-full min-h-[120px] bg-transparent border-none p-0 text-gray-300 focus:ring-0 resize-none placeholder-gray-600 leading-relaxed text-sm"
                        />
                    </div>
                </div>

                {/* Checklist */}
                <div>
                    <div className="flex items-center justify-between mb-3 text-gray-300">
                        <div className="flex items-center gap-3">
                            <CheckSquare size={20} />
                            <h3 className="font-semibold text-lg">Checklist</h3>
                        </div>
                        {totalSubtasks > 0 && (
                            <span className="text-xs text-gray-500 font-medium">{Math.round(progress)}% concluído</span>
                        )}
                    </div>

                    {totalSubtasks > 0 && (
                        <div className="w-full bg-gray-800/50 rounded-full h-1.5 mb-5 overflow-hidden">
                            <div 
                                className="bg-emerald-500 h-full rounded-full transition-all duration-300"
                                style={{ width: `${progress}%` }}
                            ></div>
                        </div>
                    )}

                    <div className="space-y-2 mb-4">
                        {task.subtasks?.map(st => (
                            <div key={st.id} className="group flex items-center gap-3 bg-[#0f1115] border border-gray-800 hover:border-gray-700 p-2 pr-3 rounded-lg transition-all">
                                <button 
                                    onClick={() => toggleSubtask(st.id)}
                                    className={`flex-shrink-0 w-5 h-5 rounded border flex items-center justify-center transition-colors ml-1 ${st.completed ? 'bg-emerald-500 border-emerald-500 text-[#0f1115]' : 'border-gray-600 hover:border-emerald-500'}`}
                                >
                                    {st.completed && <CheckSquare size={14} />}
                                </button>
                                
                                <input 
                                    type="text" 
                                    value={st.title}
                                    onChange={(e) => updateSubtaskField(st.id, 'title', e.target.value)}
                                    className={`flex-grow bg-transparent border-none p-0 text-sm focus:ring-0 ${st.completed ? 'line-through text-gray-500' : 'text-gray-200'}`}
                                />

                                <div className="flex items-center gap-2 opacity-60 group-hover:opacity-100 transition-opacity">
                                    <div className="relative">
                                        <div className="flex items-center gap-1.5 bg-[#1a1d24] px-2 py-1 rounded text-xs border border-gray-700">
                                            <span className="w-4 h-4 rounded-full bg-gray-600 flex items-center justify-center text-[8px] font-bold text-white">
                                                {st.assignee ? st.assignee[0] : 'U'}
                                            </span>
                                            <select
                                                value={st.assignee || 'Jackson'}
                                                onChange={(e) => updateSubtaskField(st.id, 'assignee', e.target.value)}
                                                className="absolute inset-0 opacity-0 cursor-pointer w-full"
                                            >
                                                {ALLOWED_USERS.map(u => <option key={u} value={u}>{u}</option>)}
                                            </select>
                                            <span className="truncate max-w-[60px]">{st.assignee}</span>
                                        </div>
                                    </div>

                                    <div className="relative">
                                        <div className="flex items-center gap-1.5 bg-[#1a1d24] px-2 py-1 rounded text-xs border border-gray-700">
                                            <Calendar size={10} />
                                            <span>{st.dueDate ? st.dueDate.split('-').slice(1).reverse().join('/') : '--/--'}</span>
                                            <input 
                                                type="date"
                                                value={st.dueDate || ''}
                                                onChange={(e) => updateSubtaskField(st.id, 'dueDate', e.target.value)}
                                                className="absolute inset-0 opacity-0 cursor-pointer w-full"
                                                onClick={openDatePicker}
                                            />
                                        </div>
                                    </div>

                                    <button onClick={() => deleteSubtask(st.id)} className="p-1 hover:text-rose-500 transition-colors">
                                        <Trash2 size={14} />
                                    </button>
                                </div>
                            </div>
                        ))}
                    </div>

                    {/* Add Item Row */}
                    <form onSubmit={addSubtask} className="flex gap-2">
                        <div className="flex-grow flex items-center bg-[#0f1115] border border-gray-800 rounded-lg px-3 focus-within:border-gray-600 transition-colors">
                            <Plus size={16} className="text-gray-500 mr-2" />
                            <input 
                                type="text" 
                                value={newSubtaskTitle}
                                onChange={(e) => setNewSubtaskTitle(e.target.value)}
                                placeholder="Adicionar um item..."
                                className="flex-grow bg-transparent border-none py-2.5 text-sm focus:ring-0 text-gray-200 placeholder-gray-600"
                            />
                        </div>
                        <div className="flex gap-2">
                             <div className="relative w-28 bg-[#0f1115] border border-gray-800 rounded-lg flex items-center px-2">
                                <User size={14} className="text-gray-500 mr-2" />
                                <span className="text-xs text-gray-300 truncate">{newSubtaskAssignee}</span>
                                <select
                                    value={newSubtaskAssignee}
                                    onChange={(e) => setNewSubtaskAssignee(e.target.value as AllowedUser)}
                                    className="absolute inset-0 opacity-0 cursor-pointer"
                                >
                                    {ALLOWED_USERS.map(u => <option key={u} value={u}>{u}</option>)}
                                </select>
                             </div>
                             <div className="relative w-24 bg-[#0f1115] border border-gray-800 rounded-lg flex items-center px-2">
                                <Calendar size={14} className="text-gray-500 mr-2" />
                                <span className="text-xs text-gray-300 truncate">{newSubtaskDate ? newSubtaskDate.split('-').slice(1).reverse().join('/') : 'Data'}</span>
                                <input 
                                    type="date"
                                    value={newSubtaskDate}
                                    onChange={(e) => setNewSubtaskDate(e.target.value)}
                                    className="absolute inset-0 opacity-0 cursor-pointer"
                                    onClick={openDatePicker}
                                />
                             </div>
                             <button 
                                type="submit"
                                disabled={!newSubtaskTitle.trim()}
                                className="px-4 bg-[#20232b] text-gray-300 text-sm font-medium rounded-lg hover:bg-[#2a2e37] disabled:opacity-50 transition-colors"
                             >
                                Adicionar
                             </button>
                        </div>
                    </form>
                </div>

                {/* Comments */}
                <div>
                    <div className="flex items-center gap-3 mb-4 text-gray-300">
                        <MessageSquare size={20} />
                        <h3 className="font-semibold text-lg">Comentários</h3>
                    </div>

                    <div className="space-y-4 mb-5 pl-2">
                        {task.comments?.map(comment => (
                            <div key={comment.id} className="flex gap-4 group">
                                <div className="w-8 h-8 rounded-full bg-gradient-to-br from-gray-700 to-gray-800 flex items-center justify-center text-xs font-bold border border-gray-600 flex-shrink-0 text-gray-300">
                                    {comment.user[0]}
                                </div>
                                <div>
                                    <div className="flex items-baseline gap-2 mb-1">
                                        <span className="font-semibold text-gray-200 text-sm">{comment.user}</span>
                                        <span className="text-xs text-gray-600">{comment.timestamp}</span>
                                    </div>
                                    <div className="bg-[#20232b] px-4 py-2.5 rounded-2xl rounded-tl-none border border-gray-800 text-sm text-gray-300 leading-relaxed max-w-2xl">
                                        {renderCommentText(comment.text)}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>

                    <div className="flex gap-3">
                        <div className="w-8 h-8 rounded-full bg-gray-800 flex items-center justify-center text-gray-500 font-bold text-xs flex-shrink-0 border border-gray-700">
                            {activeUser[0]}
                        </div>
                        <div className="flex-grow bg-[#0f1115] border border-gray-800 rounded-xl flex items-center pr-2 focus-within:border-gray-600 transition-colors relative">
                            <textarea 
                                value={newComment}
                                onChange={handleCommentChange}
                                placeholder="Escreva um comentário..."
                                className="flex-grow bg-transparent border-none py-3 px-4 text-sm text-gray-200 focus:ring-0 resize-none h-[46px] placeholder-gray-600"
                                onKeyDown={(e) => {
                                    if(e.key === 'Enter' && !e.shiftKey) {
                                        e.preventDefault();
                                        addComment(e);
                                    }
                                }}
                            />
                            <button 
                                onClick={addComment}
                                disabled={!newComment.trim()}
                                className="p-2 text-gray-500 hover:text-emerald-400 transition-colors disabled:opacity-50"
                            >
                                <Send size={18} />
                            </button>

                            {showMentions && (
                                <div className="absolute bottom-full left-0 mb-2 w-48 bg-[#20232b] border border-gray-700 rounded-lg shadow-xl overflow-hidden z-10 animate-in fade-in slide-in-from-bottom-2">
                                    {filteredUsers.map(u => (
                                        <button
                                            key={u}
                                            onClick={() => insertMention(u)}
                                            className="w-full text-left px-3 py-2 hover:bg-[#2a2e37] text-gray-300 text-sm flex items-center gap-2"
                                        >
                                            <span className="w-5 h-5 rounded-full bg-gray-700 flex items-center justify-center text-[10px]">{u[0]}</span>
                                            {u}
                                        </button>
                                    ))}
                                </div>
                            )}
                        </div>
                    </div>
                </div>

            </div>
        </div>

        {/* SIDEBAR */}
        <div className="w-[300px] bg-[#0f1115] border-l border-gray-800 flex flex-col">
            <div className="p-5 space-y-4">
                {/* DATA DE ENTRADA & PRAZO FINAL GRID */}
                <div className="grid grid-cols-2 gap-3">
                    {/* DATA DE ENTRADA */}
                    <div className="group">
                        <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Data de Entrada</label>
                        <div className="relative bg-[#1a1d24] border border-gray-700 rounded-lg hover:border-gray-600 transition-colors group-focus-within:border-emerald-500/50">
                            <div className="flex items-center px-3 py-2.5">
                                <Calendar size={14} className="text-gray-500 mr-2 flex-shrink-0" />
                                <span className="text-sm text-gray-300 truncate">{task.startDate ? task.startDate.split('-').reverse().join('/') : 'Selecione'}</span>
                            </div>
                            <input 
                                type="date"
                                value={task.startDate || ''}
                                onChange={(e) => onUpdate({...task, startDate: e.target.value})}
                                className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                                onClick={openDatePicker}
                            />
                        </div>
                    </div>

                    {/* PRAZO FINAL */}
                    <div className="group">
                        <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Prazo Final</label>
                        <div className="relative bg-[#1a1d24] border border-gray-700 rounded-lg hover:border-gray-600 transition-colors group-focus-within:border-emerald-500/50">
                            <div className="flex items-center px-3 py-2.5">
                                <CalendarDays size={14} className="text-gray-500 mr-2 flex-shrink-0" />
                                <span className="text-sm text-gray-300 truncate">{task.dueDate ? task.dueDate.split('-').reverse().join('/') : 'Selecione'}</span>
                            </div>
                            <input 
                                type="date"
                                value={task.dueDate}
                                onChange={(e) => onUpdate({...task, dueDate: e.target.value})}
                                className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                                onClick={openDatePicker}
                            />
                        </div>
                    </div>
                </div>

                {/* FLUXOS & PRIORIDADE GRID */}
                <div className="grid grid-cols-2 gap-3">
                    {/* FLUXOS DROPDOWN */}
                    <div className="relative">
                        <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block flex items-center gap-1">
                            <Layers size={10} /> Fluxos
                        </label>
                        <button 
                            type="button"
                            onClick={() => setIsFlowDropdownOpen(!isFlowDropdownOpen)}
                            className="w-full bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-2.5 text-left text-sm text-gray-200 flex items-center justify-between hover:border-gray-600 transition-colors h-[42px]"
                        >
                            <span className="truncate">
                                {task.flows && task.flows.length > 0 
                                    ? (task.flows.length > 1 ? `${task.flows.length} Selec.` : task.flows[0]) 
                                    : 'Selecione'}
                            </span>
                            <ChevronDown size={14} className="text-gray-500 flex-shrink-0" />
                        </button>
                        
                        {isFlowDropdownOpen && (
                            <>
                                <div className="fixed inset-0 z-10" onClick={() => setIsFlowDropdownOpen(false)}></div>
                                <div className="absolute top-full left-0 mt-1 w-48 bg-[#20232b] border border-gray-700 rounded-lg shadow-xl z-20 overflow-hidden animate-in fade-in slide-in-from-top-2">
                                    {AVAILABLE_FLOWS.map(flow => {
                                        const isActive = task.flows?.includes(flow);
                                        return (
                                            <button
                                                key={flow}
                                                type="button"
                                                onClick={() => toggleFlow(flow)}
                                                className={`w-full text-left px-3 py-2 text-xs flex items-center justify-between hover:bg-[#1a1d24] transition-colors ${isActive ? 'text-pink-400 font-medium bg-pink-500/5' : 'text-gray-400'}`}
                                            >
                                                {flow}
                                                {isActive && <Check size={12} />}
                                            </button>
                                        )
                                    })}
                                </div>
                            </>
                        )}
                    </div>

                    {/* PRIORIDADE DROPDOWN */}
                    <div>
                        <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Prioridade</label>
                        <div className="relative">
                            <div className="bg-[#1a1d24] border border-gray-700 rounded-lg flex items-center px-3 py-2.5 hover:border-gray-600 transition-colors h-[42px]">
                                <div className={`w-2 h-2 rounded-full mr-2 flex-shrink-0 ${
                                    task.priority === 'alta' ? 'bg-rose-500 shadow-[0_0_6px_rgba(244,63,94,0.4)]' : 
                                    task.priority === 'media' ? 'bg-amber-500' : 'bg-blue-500'
                                }`}></div>
                                <span className="text-sm text-gray-200 capitalize truncate">
                                    {task.priority === 'alta' ? 'Alta' : task.priority === 'media' ? 'Média' : 'Baixa'}
                                </span>
                            </div>
                            <select
                                value={task.priority}
                                onChange={(e) => onUpdate({...task, priority: e.target.value as any})}
                                className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                            >
                                <option value="baixa">Baixa</option>
                                <option value="media">Média</option>
                                <option value="alta">Alta</option>
                            </select>
                        </div>
                    </div>
                </div>

                {/* RESPONSÁVEL */}
                <div>
                    <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Responsável</label>
                    <div className="relative">
                        <div className="bg-[#1a1d24] border border-gray-700 rounded-lg flex items-center px-3 py-2.5 hover:border-gray-600 transition-colors">
                            <div className="w-5 h-5 rounded-full bg-gray-600 flex items-center justify-center text-[10px] font-bold text-white mr-2">
                                {task.assignee ? task.assignee[0] : 'U'}
                            </div>
                            <span className="text-sm text-gray-200">{task.assignee}</span>
                        </div>
                        <select
                            value={task.assignee || 'Jackson'}
                            onChange={(e) => onUpdate({...task, assignee: e.target.value as any})}
                            className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                        >
                            {ALLOWED_USERS.map(u => <option key={u} value={u}>{u}</option>)}
                        </select>
                    </div>
                </div>

                {/* STATUS */}
                <div>
                    <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Status</label>
                    <div className="relative">
                        <div className="bg-[#1a1d24] border border-gray-700 rounded-lg flex items-center px-3 py-2.5 hover:border-gray-600 transition-colors">
                            <span className="text-sm text-gray-200">
                                {task.status === 'todo' ? 'A Fazer' : task.status === 'in_progress' ? 'Em Andamento' : 'Concluído'}
                            </span>
                        </div>
                        <select
                            value={task.status}
                            onChange={(e) => onUpdate({...task, status: e.target.value as any})}
                            className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                        >
                            <option value="todo">A Fazer</option>
                            <option value="in_progress">Em Andamento</option>
                            <option value="done">Concluído</option>
                        </select>
                    </div>
                </div>

                <button 
                    onClick={toggleArchive}
                    className="w-full flex items-center justify-center gap-2 px-3 py-2.5 border border-gray-700 rounded-lg text-sm text-gray-400 hover:bg-[#1a1d24] hover:text-white transition-all hover:border-gray-600 mt-2"
                >
                    {task.archived ? <RefreshCcw size={16} /> : <Archive size={16} />}
                    {task.archived ? 'Desarquivar Tarefa' : 'Arquivar Tarefa'}
                </button>
            </div>

            <div className="mt-auto p-5 border-t border-gray-800 text-xs text-gray-600 space-y-2">
                <p>Criado em 21/11/2025</p>
                <p>ID: {task.id}</p>
                
                <button 
                    onClick={() => onDelete(task.id)}
                    className="flex items-center gap-1.5 text-rose-500/70 hover:text-rose-400 mt-2 transition-colors font-medium hover:underline"
                >
                    <Trash2 size={14} /> Excluir Tarefa
                </button>
            </div>
        </div>

      </div>
    </div>
  );
};

export default TaskModal;
