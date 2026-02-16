 
import React, { useState, useEffect, useMemo } from 'react';
import type { Task, Subtask, Comment, AllowedUser, TaskStatus } from '@/shared/types';
import type { AppUser } from '@/shared/types/user.types';
import { useUsers } from '@/features/users/hooks';
import { useAuth } from '@/features/auth';
import { useComments, useSubtasks } from '@/features/tasks/hooks';
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
  ChevronDown,
  Check,
  Loader2
} from 'lucide-react';
import { formatDisplayDate, formatDateForInput, formatDateForAPI } from '@/shared/utils/dateFormatters';
import type { TaskModalProps } from '../types';

const TaskModal: React.FC<TaskModalProps> = ({ task, isOpen, onClose, onUpdate, onDelete, onMention, onAddSubtask, onUpdateSubtask, onDeleteSubtask, onAddComment, onUpdateComment, onDeleteComment }) => {
// Load users from API
const { data: appUsers, isLoading: isLoadingUsers } = useUsers();

// Fetch comments and subtasks for the current task
const { data: commentsData, isLoading: isLoadingComments } = useComments(task.id);
const { data: subtasksData, isLoading: isLoadingSubtasks } = useSubtasks(task.id);

// Use fetched data if available, otherwise fall back to task data (for optimistic updates)
const displayComments = commentsData || task.comments || [];
const [displaySubtasks, setDisplaySubtasks] = useState<Subtask[]>(
    subtasksData || task.subtasks || []
);

// Sincronizar displaySubtasks quando subtasksData mudar
useEffect(() => {
    if (subtasksData) {
        setDisplaySubtasks(subtasksData);
    }
}, [subtasksData]);

const [newSubtaskTitle, setNewSubtaskTitle] = useState('');

// Estado para edição local do título de subtarefas (com debounce)
const [editingSubtaskId, setEditingSubtaskId] = useState<string | null>(null);
const [editingSubtaskTitle, setEditingSubtaskTitle] = useState('');
// Initialize with first user UUID from appUsers, or empty string if not loaded
const firstUserId = appUsers?.[0]?.id || '';
const [newSubtaskAssignee, setNewSubtaskAssignee] = useState<string>(firstUserId);
const [newSubtaskDate, setNewSubtaskDate] = useState('');

const [newComment, setNewComment] = useState('');
// Use authenticated user if available, otherwise use first user from appUsers
const { user } = useAuth();
const [activeUser] = useState<string>(user?.id || firstUserId);
  
  // Mention Autocomplete State
  const [showMentions, setShowMentions] = useState(false);
  const [mentionFilter, setMentionFilter] = useState('');


  // Debounce para atualizar título de subtarefa
  useEffect(() => {
    if (!editingSubtaskId || editingSubtaskTitle === '') return;

    const timer = setTimeout(() => {
      const subtask = displaySubtasks.find(st => st.id === editingSubtaskId);
      if (subtask && subtask.title !== editingSubtaskTitle) {
        updateSubtaskField(editingSubtaskId, 'title', editingSubtaskTitle);
      }
      setEditingSubtaskId(null);
    }, 500);

    return () => clearTimeout(timer);
  }, [editingSubtaskTitle, editingSubtaskId]);

  // Show loading while users, comments, or subtasks are being fetched
  if (isLoadingUsers || isLoadingComments || isLoadingSubtasks) {
    return (
      <div className="fixed inset-0 bg-black/80 z-50 flex justify-center items-center p-4 backdrop-blur-sm">
        <div className="bg-[#1a1d24] w-full max-w-5xl h-[85vh] rounded-2xl shadow-2xl border border-gray-800 flex items-center justify-center">
          <div className="flex flex-col items-center gap-4">
            <Loader2 size={48} className="text-pink-500 animate-spin" />
            <p className="text-gray-400">Carregando dados...</p>
          </div>
        </div>
      </div>
    );
  }

  // Handlers
  const handleDescriptionChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    onUpdate({ ...task, description: e.target.value });
  };

  const toggleSubtask = (subtaskId: string) => {
    const subtask = displaySubtasks.find(st => st.id === subtaskId);
    if (!subtask) return;

    const newStatus = subtask.status === 'completed' ? 'pending' : 'completed';

    // Atualizar estado local imediatamente
    const updatedSubtasks = (displaySubtasks || []).map(st =>
      st.id === subtaskId ? { ...st, status: newStatus as TaskStatus } : st
    );
    setDisplaySubtasks(updatedSubtasks);

    // Chamar onUpdateSubtask com { id, data }
    if (onUpdateSubtask) {
      onUpdateSubtask(subtaskId, { status: newStatus });
    }
  };

  const updateSubtaskField = (subtaskId: string, field: keyof Subtask, value: any) => {
    // Atualizar estado local imediatamente (feedback visual rápido)
    const updatedSubtasks = displaySubtasks.map(st =>
      st.id === subtaskId ? { ...st, [field]: value } : st
    );
    setDisplaySubtasks(updatedSubtasks);

    // Chamar onUpdateSubtask com { id, data }
    if (onUpdateSubtask) {
      onUpdateSubtask(subtaskId, { [field]: value });
    }
  };

  const addSubtask = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newSubtaskTitle.trim()) return;
    const newSub: Subtask = {
      id: Date.now().toString(),
      task_id: task.id,
      title: newSubtaskTitle,
      description: '',
      priority: 'low',
      status: 'pending',
      assignee_id: newSubtaskAssignee || '',
      due_date: newSubtaskDate || new Date().toISOString().split('T')[0],
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString()
    };
    if (onAddSubtask) {
      onAddSubtask(newSub);
    }
    setNewSubtaskTitle('');
    setNewSubtaskDate('');
  };

  const deleteSubtask = (subtaskId: string) => {
    if (onDeleteSubtask) {
      onDeleteSubtask(subtaskId);
    }
    const updatedSubtasks = (displaySubtasks || []).filter(st => st.id !== subtaskId);
    onUpdate({ ...task, subtasks: updatedSubtasks });
  };

  const addComment = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newComment.trim()) return;
    
    const lowerComment = newComment.toLowerCase();
    
    if (onMention) {
        // Extract all mentions from the comment
        const mentionRegex = /@(\w+)/g;
        const matches = lowerComment.match(mentionRegex) || [];
        
        // Convert to names (remove the @)
        const mentionedNames = matches.map(m => m.substring(1).toLowerCase());
        
        // Check if any of the mentions correspond to an existing user
        const hasValidMention = mentionedNames.some(name =>
            appUsers?.some(u => u.name.toLowerCase() === name)
        );
        
        if (hasValidMention) {
            onMention(task.id, task.title);
        }
    }

    const comment: Comment = {
      id: Date.now().toString(),
      task_id: task.id,
      user_id: activeUser,
      text: newComment,
      timestamp: new Date().toISOString()
    };
    if (onAddComment) {
      onAddComment(comment);
    }
    setNewComment('');
    setShowMentions(false);
  };

  const toggleArchive = () => {
      onUpdate({ ...task, archived: !task.archived });
  };

  const completedSubtasks = (displaySubtasks || []).filter(s => s.status === 'completed').length || 0;
  const totalSubtasks = (displaySubtasks || []).length || 0;
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


  const filteredUsers = (appUsers || []).filter(u => u.name.toLowerCase().includes(mentionFilter));

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
                        {task.status === 'pending' ? 'A Fazer' : task.status === 'in_progress' ? 'Em Andamento' : 'Concluído'}
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
                        {(displaySubtasks || []).map(st => (
                            <div key={st.id} className="group flex items-center gap-3 bg-[#0f1115] border border-gray-800 hover:border-gray-700 p-2 pr-3 rounded-lg transition-all">
                                <button
                                    onClick={() => toggleSubtask(st.id)}
                                    className={`flex-shrink-0 w-5 h-5 rounded border flex items-center justify-center transition-colors ml-1 ${st.status === 'completed' ? 'bg-emerald-500 border-emerald-500 text-[#0f1115]' : 'border-gray-600 hover:border-emerald-500'}`}
                                >
                                    {st.status === 'completed' && <CheckSquare size={14} />}
                                </button>
                                
                                <input
                                    type="text"
                                    value={editingSubtaskId === st.id ? editingSubtaskTitle : st.title}
                                    onChange={(e) => {
                                        setEditingSubtaskId(st.id);
                                        setEditingSubtaskTitle(e.target.value);
                                    }}
                                    onBlur={() => {
                                        if (editingSubtaskId === st.id) {
                                            updateSubtaskField(st.id, 'title', editingSubtaskTitle || st.title);
                                            setEditingSubtaskId(null);
                                        }
                                    }}
                                    onKeyDown={(e) => {
                                        if (e.key === 'Enter') {
                                            e.preventDefault();
                                            if (editingSubtaskId === st.id) {
                                                updateSubtaskField(st.id, 'title', editingSubtaskTitle || st.title);
                                                setEditingSubtaskId(null);
                                            }
                                        }
                                    }}
                                    className={`flex-grow bg-transparent border-none p-0 text-sm focus:ring-0 ${st.status === 'completed' ? 'line-through text-gray-500' : 'text-gray-200'}`}
                                />

                                <div className="flex items-center gap-2 opacity-60 group-hover:opacity-100 transition-opacity">
                                    <div className="relative">
                                        <div className="flex items-center gap-1.5 bg-[#1a1d24] px-2 py-1 rounded text-xs border border-gray-700">
                                            <span className="w-4 h-4 rounded-full bg-gray-600 flex items-center justify-center text-[8px] font-bold text-white">
                                                {st.assignee_id ? appUsers?.find(u => u.id === st.assignee_id)?.name[0] : 'U'}
                                            </span>
                                            <select
                                                value={st.assignee_id || ''}
                                                onChange={(e) => updateSubtaskField(st.id, 'assignee_id', e.target.value)}
                                                className="absolute inset-0 opacity-0 cursor-pointer w-full"
                                            >
                                                {appUsers?.map(u => <option key={u.id} value={u.id}>{u.name}</option>)}
                                            </select>
                                            <span className="truncate max-w-[60px]">{st.assignee_id ? appUsers?.find(u => u.id === st.assignee_id)?.name : ''}</span>
                                        </div>
                                    </div>

                                    <div className="relative">
                                        <div className="flex items-center gap-1.5 bg-[#1a1d24] px-2 py-1 rounded text-xs border border-gray-700">
                                            <Calendar size={10} />
                                            <span>{st.due_date ? formatDisplayDate(st.due_date) : '--/--'}</span>
                                            <input
                                                type="date"
                                                value={formatDateForInput(st.due_date)}
                                                onChange={(e) => updateSubtaskField(st.id, 'due_date', formatDateForAPI(e.target.value))}
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
                               <span className="text-xs text-gray-300 truncate">
                                 {appUsers?.find(u => u.id === newSubtaskAssignee)?.name || 'Selecione'}
                               </span>
                               <select
                                   value={newSubtaskAssignee}
                                   onChange={(e) => setNewSubtaskAssignee(e.target.value)}
                                   className="absolute inset-0 opacity-0 cursor-pointer"
                               >
                                   {appUsers?.map(u => <option key={u.id} value={u.id}>{u.name}</option>)}
                               </select>
                             </div>
                             <div className="relative w-24 bg-[#0f1115] border border-gray-800 rounded-lg flex items-center px-2">
                                <Calendar size={14} className="text-gray-500 mr-2" />
                                <span className="text-xs text-gray-300 truncate">{newSubtaskDate ? formatDisplayDate(newSubtaskDate) : 'Data'}</span>
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
                        {(displayComments || []).map(comment => (
                            <div key={comment.id} className="flex gap-4 group">
                                <div className="w-8 h-8 rounded-full bg-gradient-to-br from-gray-700 to-gray-800 flex items-center justify-center text-xs font-bold border border-gray-600 flex-shrink-0 text-gray-300">
                                    {appUsers?.find(u => u.id === comment.user_id)?.name[0] || '?'}
                                </div>
                                <div>
                                    <div className="flex items-baseline gap-2 mb-1">
                                        <span className="font-semibold text-gray-200 text-sm">{appUsers?.find(u => u.id === comment.user_id)?.name || comment.user_id}</span>
                                        <span className="text-xs text-gray-600">{new Date(comment.timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}</span>
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
                            {appUsers?.find(u => u.id === activeUser)?.name[0] || '?'}
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
                                            key={u.id}
                                            onClick={() => insertMention(u.name)}
                                            className="w-full text-left px-3 py-2 hover:bg-[#2a2e37] text-gray-300 text-sm flex items-center gap-2"
                                        >
                                            <span className="w-5 h-5 rounded-full bg-gray-700 flex items-center justify-center text-[10px]">{u.name[0]}</span>
                                            {u.name}
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
                                <span className="text-sm text-gray-300 truncate">
                                    {task.start_date ? formatDisplayDate(task.start_date) : 'Selecione'}
                                </span>
                            </div>
                            <input
                                type="date"
                                value={formatDateForInput(task.start_date)}
                                onChange={(e) => onUpdate({...task, start_date: formatDateForAPI(e.target.value)})}
                                className="absolute inset-0 opacity-0 cursor-pointer w-full h-full appearance-none"
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
                                <span className="text-sm text-gray-300 truncate">{task.due_date ? formatDisplayDate(task.due_date) : 'Selecione'}</span>
                            </div>
                            <input
                                type="date"
                                value={formatDateForInput(task.due_date)}
                                onChange={(e) => onUpdate({...task, due_date: formatDateForAPI(e.target.value)})}
                                className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                                onClick={openDatePicker}
                            />
                        </div>
                    </div>
                </div>

                {/* PRIORIDADE GRID */}
                <div className="grid grid-cols-2 gap-3">

                    {/* PRIORIDADE DROPDOWN */}
                    <div>
                        <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Prioridade</label>
                        <div className="relative">
                            <div className="bg-[#1a1d24] border border-gray-700 rounded-lg flex items-center px-3 py-2.5 hover:border-gray-600 transition-colors h-[42px]">
                                <div className={`w-2 h-2 rounded-full mr-2 flex-shrink-0 ${
                                    task.priority === 'high' ? 'bg-rose-500 shadow-[0_0_6px_rgba(244,63,94,0.4)]' :
                                    task.priority === 'medium' ? 'bg-amber-500' : 'bg-blue-500'
                                }`}></div>
                                <span className="text-sm text-gray-200 capitalize truncate">
                                    {task.priority === 'high' ? 'Alta' : task.priority === 'medium' ? 'Média' : 'Baixa'}
                                </span>
                            </div>
                            <select
                                value={task.priority}
                                onChange={(e) => onUpdate({...task, priority: e.target.value as any})}
                                className="absolute inset-0 opacity-0 cursor-pointer w-full h-full appearance-none"
                            >
                                <option value="low">Baixa</option>
                                <option value="medium">Média</option>
                                <option value="high">Alta</option>
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
                                {task.assignee_id ? appUsers?.find(u => u.id === task.assignee_id)?.name[0] : 'U'}
                            </div>
                            <span className="text-sm text-gray-200">{task.assignee_id ? appUsers?.find(u => u.id === task.assignee_id)?.name : ''}</span>
                        </div>
                        <select
                            value={task.assignee_id || ''}
                            onChange={(e) => onUpdate({...task, assignee_id: e.target.value})}
                            className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                        >
                            {appUsers?.map(u => <option key={u.id} value={u.id}>{u.name}</option>)}
                        </select>
                    </div>
                </div>

                {/* STATUS */}
                <div>
                    <label className="text-[10px] text-gray-500 uppercase font-bold mb-1.5 block">Status</label>
                    <div className="relative">
                        <div className="bg-[#1a1d24] border border-gray-700 rounded-lg flex items-center px-3 py-2.5 hover:border-gray-600 transition-colors">
                            <span className="text-sm text-gray-200">
                                {task.status === 'pending' ? 'A Fazer' : task.status === 'in_progress' ? 'Em Andamento' : 'Concluído'}
                            </span>
                        </div>
                        <select
                            value={task.status}
                            onChange={(e) => onUpdate({...task, status: e.target.value as any})}
                            className="absolute inset-0 opacity-0 cursor-pointer w-full h-full"
                        >
                            <option value="pending">A Fazer</option>
                            <option value="in_progress">Em Andamento</option>
                            <option value="completed">Concluído</option>
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
                <p>Criado em {formatDisplayDate(task.created_at) || 'N/A'}</p>
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
