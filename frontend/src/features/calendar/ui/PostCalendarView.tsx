import React, { useState, useEffect } from 'react';
import { useQueryClient } from '@tanstack/react-query';
import {
  Calendar,
  ChevronLeft,
  ChevronRight,
  Clock,
  CheckCircle2,
  AlertCircle,
  Star,
  Instagram,
  Youtube,
  Scissors,
  User,
  Users,
  Megaphone,
  X,
  Send,
  MoreVertical,
  Check,
  Facebook,
  Linkedin,
  CloudLightning,
  Plus,
  Filter,
  RefreshCcw,
  Edit,
  CheckSquare,
  Square,
  Image as ImageIcon,
  Paperclip,
  Video,
  Layers,
  FileText,
  BookOpen,
  Info,
  Tag,
  Briefcase,
  Trash2,
  Copy,
  Edit2,
  Save,
  RotateCcw,
  LayoutGrid,
  Smartphone,
  Grid,
  Battery,
  Wifi,
  Signal,
  Search
} from 'lucide-react';
import type { CalendarPost, PostCategory, PostHistoryEvent, PostStatus, PostType, PublicUser } from '@/shared/types/legacy.types';
import type { PostCalendarViewProps } from '../types';
import { useUsers } from '@/features/users/hooks';
import type { AppUser } from '@/shared/types/user.types';
import type { CalendarPostRequest } from '@/shared/types/legacy.types';

const DAYS_OF_WEEK = ['Dom', 'Seg', 'Ter', 'Qua', 'Qui', 'Sex', 'Sáb'];

// --- STYLES CONFIGURATION ---
const CATEGORY_CONFIG: Record<PostCategory, { label: string, color: string, bg: string, border: string }> = {
  official: { 
    label: 'Oficial', 
    color: 'text-amber-400', 
    bg: 'bg-amber-500/10',
    border: 'border-amber-500/30'
  },
  solis_voce: { 
    label: 'Solis e Você', 
    color: 'text-blue-400', 
    bg: 'bg-blue-500/10',
    border: 'border-blue-500/30'
  },
  leonardo: { 
    label: 'Leonardo', 
    color: 'text-purple-400', 
    bg: 'bg-purple-500/10',
    border: 'border-purple-500/30'
  },
  luiz: { 
    label: 'Luiz', 
    color: 'text-emerald-400', 
    bg: 'bg-emerald-500/10',
    border: 'border-emerald-500/30'
  }
};

// --- PROFILE DATA FOR PREVIEW ---
const PROFILE_DATA: Record<string, { handle: string; name: string; bio: string; link: string; initials: string }> = {
  official: {
    handle: 'solissolaroficial',
    name: 'Solis Solar',
    bio: '☀️ A melhor empresa de Energia Solar do Brasil\n⏰ Atendimento de segunda à sexta das 7h às 17h20\n"Tudo posso naquele que me fortalece. Fl. 4:13"',
    link: 'linktr.ee/solissolar',
    initials: 'S'
  },
  solis_voce: {
    handle: 'solisevoce',
    name: 'Solis e Você',
    bio: 'Seu canal de relacionamento com a Solis! 💚☀️\n◦ Para colaboradores, representantes, PDVs e parceiros.\nAcompanhe os bastidores da @solissolaroficial',
    link: 'linktr.ee/solissolar',
    initials: 'SV'
  },
  leonardo: {
    handle: 'leonardochamonecardoso',
    name: 'Leonardo Chamone Cardoso',
    bio: 'Diretor Comercial @solissolaroficial\nCasado com @bia_orsichamone\nAquecedor Solar | Mkt & Vendas | Deus & Família | Esporte & Saúde',
    link: 'solissolar.com.br',
    initials: 'L'
  },
  luiz: {
    handle: 'luizantoniodossantospinto',
    name: 'Luiz Antonio dos Santos Pinto',
    bio: 'CEO @solissolaroficial; presidente da Profort; diretor do CIESP Alta Noroeste; cofundador e ex-presidente da ABRASOL',
    link: 'linkedin.com/in/luiz-antonio...',
    initials: 'L'
  }
};

const TYPE_CONFIG: Record<PostType, { label: string, icon: any }> = {
    video: { label: 'Vídeo', icon: Video },
    static: { label: 'Estático', icon: ImageIcon },
    carousel: { label: 'Carrossel', icon: Layers },
    story: { label: 'Story', icon: Instagram }, // Using Instagram icon for Story representation
    article_linkedin: { label: 'Artigo Linkedin', icon: FileText },
    article_blog: { label: 'Artigo Blog', icon: BookOpen }
};

const TikTokIcon = ({ className }: { className?: string }) => (
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className={className}>
        <path d="M9 12a4 4 0 1 0 4 4V4a5 5 0 0 0 5 5" />
    </svg>
);

const PostCalendarView: React.FC<PostCalendarViewProps> = ({ posts, mutations, selectedCategory = 'all', setSelectedCategory, selectedType = 'all', setSelectedType, selectedAssignee = 'all', setSelectedAssignee }) => {
  const queryClient = useQueryClient();
  const [currentDate, setCurrentDate] = useState(new Date(2025, 11, 1)); // Dec 2025 default
  const [selectedPost, setSelectedPost] = useState<CalendarPost | null>(null);
  const [adjustComment, setAdjustComment] = useState('');
  const [viewMode, setViewMode] = useState<'calendar' | 'feed'>('calendar');
  
  // Load users from system
  const { data: appUsers, isLoading: isLoadingUsers } = useUsers();
  
  // Filters (use props if provided, otherwise use local state for backward compatibility)
  const [localCategory, setLocalCategory] = useState<PostCategory | 'all'>('all');
  const [localType, setLocalType] = useState<PostType | 'all'>('all');
  const [localAssignee, setLocalAssignee] = useState<string>('all');
  
  const categoryFilter = selectedCategory !== undefined ? selectedCategory : localCategory;
  const typeFilter = selectedType !== undefined ? selectedType : localType;
  const assigneeFilter = selectedAssignee !== undefined ? selectedAssignee : localAssignee;
  
  const handleCategoryChange = (value: PostCategory | 'all') => {
    if (setSelectedCategory) {
      setSelectedCategory(value);
    } else {
      setLocalCategory(value);
    }
  };
  
  const handleTypeChange = (value: PostType | 'all') => {
    if (setSelectedType) {
      setSelectedType(value);
    } else {
      setLocalType(value);
    }
  };
  
  const handleAssigneeChange = (value: string) => {
    if (setSelectedAssignee) {
      setSelectedAssignee(value);
    } else {
      setLocalAssignee(value);
    }
  };
  
  // Add Post State
  const [isAddingPost, setIsAddingPost] = useState(false);
  const [newPostTitle, setNewPostTitle] = useState('');
  const [newPostDescription, setNewPostDescription] = useState('');
  const [newPostAssignee, setNewPostAssignee] = useState<string>(''); // UUID instead of name
  const [newPostDate, setNewPostDate] = useState('');
  const [newPostCategory, setNewPostCategory] = useState<PostCategory>('official');
  const [newPostType, setNewPostType] = useState<PostType>('static');
  const [newPostPlatforms, setNewPostPlatforms] = useState<string[]>([]);

  // Publish Modal State
  const [isPublishingModalOpen, setIsPublishingModalOpen] = useState(false);
  const [publishingPlatforms, setPublishingPlatforms] = useState<string[]>([]);
  
  // Edit Mode States
  const [isEditingDesc, setIsEditingDesc] = useState(false);
  const [editDescValue, setEditDescValue] = useState('');
  
  const [isEditingTitle, setIsEditingTitle] = useState(false);
  const [editTitleValue, setEditTitleValue] = useState('');

  const [copyFeedback, setCopyFeedback] = useState(false);

  useEffect(() => {
      // Reset edit states when post changes or modal closes
      setIsEditingDesc(false);
      setEditDescValue('');
      
      setIsEditingTitle(false);
      setEditTitleValue('');
      
      setCopyFeedback(false);
  }, [selectedPost]);

  const socialIcons = [
      { id: 'facebook', icon: Facebook, color: 'text-blue-500' },
      { id: 'instagram', icon: Instagram, color: 'text-pink-500' },
      { id: 'linkedin', icon: Linkedin, color: 'text-blue-600' },
      { id: 'youtube', icon: Youtube, color: 'text-red-500' },
      { id: 'tiktok', icon: TikTokIcon, color: 'text-black dark:text-white' }
  ];

  // Get unique assignees for filter
  const uniqueAssignees = Array.from(new Set(posts.map(p => p.assignee?.name).filter(Boolean))) as string[];

  // Filter posts locally (only if filters are not provided via props)
  const filteredPosts = posts.filter(p => {
    const matchCat = categoryFilter === 'all' || p.category === categoryFilter;
    const matchType = typeFilter === 'all' || p.type === typeFilter;
    const matchAssignee = assigneeFilter === 'all' || p.assignee?.name === assigneeFilter;
    return matchCat && matchType && matchAssignee;
  });

  // --- CALENDAR LOGIC ---
  const getDaysInMonth = (date: Date) => {
    const year = date.getFullYear();
    const month = date.getMonth();
    const firstDay = new Date(year, month,1);
    const lastDay = new Date(year, month + 1, 0);
    const days = [];
 
    // Padding days (previous month)
    const startDayOfWeek = firstDay.getDay(); // 0 (Sun) - 6 (Sat)
    for (let i = 0; i < startDayOfWeek; i++) {
        days.push({ day: null, date: null });
    }
 
    // Actual days
    for (let i =1; i <= lastDay.getDate(); i++) {
        days.push({ 
            day: i, 
            date: new Date(year, month, i).toISOString().split('T')[0] 
        });
    }
 
    return days;
  };

  const days = getDaysInMonth(currentDate);
 
  const prevMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() - 1, 1));
  };

  const nextMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 1));
  };

  // --- ACTIONS ---
  const handleStatusChange = async (newStatus: PostStatus, comment?: string) => {
    if (!selectedPost || !mutations) return;
    
    // If publishing, open confirmation modal
    if (newStatus === 'published') {
        setPublishingPlatforms(selectedPost.published_platforms || []);
        setIsPublishingModalOpen(true);
        return;
    }

    // Use backend mutation for status changes
    const result = await new Promise<any>((resolve, reject) => {
        mutations.updateStatus({
            id: selectedPost.id,
            request: {
                status: newStatus,
                text: comment
            }
        }, {
            onSuccess: resolve,
            onError: reject
        });
    });
    
    // Update selectedPost with the returned data (result is already the CalendarPost, not { data: CalendarPost })
    if (result) {
        setSelectedPost(result);
    }
    
    if(newStatus === 'adjust') setAdjustComment('');
  };

  const confirmPublishing = async () => {
    if (!selectedPost || !mutations) return;

    // Use backend mutation to confirm publishing
    const result = await new Promise<any>((resolve, reject) => {
        mutations.confirmPublishing(selectedPost.id, publishingPlatforms, {
            onSuccess: resolve,
            onError: reject
        });
    });
    
    // Update selectedPost with the returned data (result is already the CalendarPost, not { data: CalendarPost })
    if (result) {
        setSelectedPost(result);
    }
    
    setIsPublishingModalOpen(false);
  };

  const attachImage = async () => {
      if (!selectedPost || !mutations) return;
      // Simulate attaching an image
      const result = await new Promise<any>((resolve, reject) => {
          mutations.updatePost({
              id: selectedPost.id,
              data: {
                  image_url: 'https://images.unsplash.com/photo-1611162617474-5b21e879e113?auto=format&fit=crop&q=80&w=800'
              }
          }, {
              onSuccess: resolve,
              onError: reject
          });
      });
      
      // Update selectedPost with the returned data (result is already the CalendarPost, not { data: CalendarPost })
      if (result) {
          setSelectedPost(result);
      }
  };

  const handleChangePhoto = async () => {
      if (!selectedPost || !mutations) return;
      // Mock upload/change
      const newUrl = window.prompt("Insira a URL da nova imagem:", selectedPost.image_url || '');
      if (newUrl) {
          const result = await new Promise<any>((resolve, reject) => {
              mutations.updatePost({
                  id: selectedPost.id,
                  data: { image_url: newUrl }
              }, {
                  onSuccess: resolve,
                  onError: reject
              });
          });
          
          // Update selectedPost with the returned data (result is already the CalendarPost, not { data: CalendarPost })
          if (result) {
              setSelectedPost(result);
          }
      }
  };

  const handleDeletePost = () => {
      if (!selectedPost || !mutations) return;
      if (window.confirm("Tem certeza que deseja excluir esta postagem? Esta ação não pode ser desfeita.")) {
          mutations.deletePost(selectedPost.id);
          setSelectedPost(null);
      }
  };

  // --- DESCRIPTION EDITING ---
  const startEditingDescription = () => {
      if (selectedPost) {
          setEditDescValue(selectedPost.description || '');
          setIsEditingDesc(true);
      }
  };

  const saveDescription = async () => {
      if (selectedPost && mutations) {
          const result = await new Promise<any>((resolve, reject) => {
              mutations.updatePost({
                  id: selectedPost.id,
                  data: { description: editDescValue }
              }, {
                  onSuccess: resolve,
                  onError: reject
              });
          });
          
          // Update selectedPost with the returned data (result is already the CalendarPost, not { data: CalendarPost })
          if (result) {
              setSelectedPost(result);
          }
          
          setIsEditingDesc(false);
      }
  };

  // --- TITLE EDITING ---
  const startEditingTitle = () => {
      if (selectedPost) {
          setEditTitleValue(selectedPost.title);
          setIsEditingTitle(true);
      }
  };

  const saveTitle = async () => {
      if (selectedPost && mutations && editTitleValue.trim() !== '') {
          const result = await new Promise<any>((resolve, reject) => {
              mutations.updatePost({
                  id: selectedPost.id,
                  data: { title: editTitleValue }
              }, {
                  onSuccess: resolve,
                  onError: reject
              });
          });
          
          // Update selectedPost with the returned data (result is already the CalendarPost, not { data: CalendarPost })
          if (result) {
              setSelectedPost(result);
          }
      }
      setIsEditingTitle(false);
  };

  const copyCaptionToClipboard = () => {
      if (selectedPost && selectedPost.caption) {
          navigator.clipboard.writeText(selectedPost.caption);
          setCopyFeedback(true);
          setTimeout(() => setCopyFeedback(false), 2000);
      }
  };

  // Used only when creating post
  const toggleNewPostPlatform = (platform: string) => {
      if (newPostPlatforms.includes(platform)) {
          setNewPostPlatforms(prev => prev.filter(p => p !== platform));
      } else {
          setNewPostPlatforms(prev => [...prev, platform]);
      }
  };
  
  // Used in publishing modal
  const togglePublishingPlatform = (platform: string) => {
      if (publishingPlatforms.includes(platform)) {
          setPublishingPlatforms(prev => prev.filter(p => p !== platform));
      } else {
          setPublishingPlatforms(prev => [...prev, platform]);
      }
  };

  const handleCreatePost = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newPostTitle || !newPostDate) return;

    // Map selected assignee UUID (or undefined if not selected)
    const assigneeID = newPostAssignee || undefined;

    // Use backend mutation to create post
    mutations?.createPost({
        title: newPostTitle,
        description: newPostDescription,
        date: newPostDate,
        time: '12:00', // Default
        caption: '',
        category: newPostCategory,
        type: newPostType,
        platforms: newPostPlatforms,
        assignee_id: assigneeID
    });

    setIsAddingPost(false);
    setNewPostTitle('');
    setNewPostDescription('');
    setNewPostDate('');
    setNewPostPlatforms([]);
    setNewPostCategory('official');
    setNewPostType('static');
    setNewPostAssignee('');
  };

  // Posts for Feed Preview (sorted new to old) - FILTER BY CURRENT MONTH
  const feedPosts = [...filteredPosts]
    .filter(p => {
        const isInsta = p.platforms?.includes('instagram') || p.platforms?.length === 0 || !p.platforms;
        // Parsing date to check month/year match
        const pDate = new Date(p.date + 'T12:00:00');
        const isSameMonth = pDate.getMonth() === currentDate.getMonth() && pDate.getFullYear() === currentDate.getFullYear();
        
        return isInsta && isSameMonth;
    })
    .sort((a, b) => new Date(b.date + 'T' + b.time).getTime() - new Date(a.date + 'T' + a.time).getTime());

  // Determine current profile for preview based on category filter
  const currentProfileKey = categoryFilter === 'all' ? 'official' : categoryFilter;
  const currentProfile = PROFILE_DATA[currentProfileKey] || PROFILE_DATA['official'];
 
  return (
    <div className="h-full flex flex-col animate-in fade-in duration-500">
      
      {/* --- HEADER (Outside Calendar Card) --- */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div>
            <h1 className="text-2xl font-bold text-gray-100">Calendário Solis</h1>
            <p className="text-gray-500">Gestão de conteúdo e aprovações</p>
        </div>
      </div>

      {/* --- CALENDAR CARD CONTAINER --- */}
      <div className="flex-grow flex flex-col bg-[#0f1115] text-gray-100 rounded-xl overflow-hidden shadow-2xl border border-gray-800">
      
        {/* TOOLBAR */}
        <div className="flex flex-col xl:flex-row items-center justify-between p-4 border-b border-gray-800 bg-[#1a1d24] gap-4">
            
            {/* Left Controls: Mode & Filters */}
            <div className="flex flex-wrap items-center gap-3 w-full xl:w-auto">
                {/* View Mode Toggle */}
                <div className="flex bg-[#0f1115] p-1 rounded-lg border border-gray-700 mr-2">
                    <button 
                        onClick={() => setViewMode('calendar')}
                        className={`p-2 rounded-md transition-colors ${viewMode === 'calendar' ? 'bg-[#1e5144] text-white' : 'text-gray-400 hover:text-white hover:bg-gray-800'}`}
                        title="Modo Calendário"
                    >
                        <Calendar size={18} />
                    </button>
                    <button 
                        onClick={() => setViewMode('feed')}
                        className={`p-2 rounded-md transition-colors ${viewMode === 'feed' ? 'bg-[#1e5144] text-white' : 'text-gray-400 hover:text-white hover:bg-gray-800'}`}
                        title="Preview do Feed"
                    >
                        <Smartphone size={18} />
                    </button>
                </div>

                {/* Category Filter */}
                <div className="relative">
                    <select
                        value={categoryFilter}
                        onChange={(e) => handleCategoryChange(e.target.value as any)}
                        className="appearance-none bg-[#0f1115] border border-gray-700 text-gray-300 rounded-lg pl-9 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144] cursor-pointer hover:border-gray-500 transition-colors"
                    >
                        <option value="all">Todas Categorias</option>
                        {Object.entries(CATEGORY_CONFIG).map(([key, config]) => (
                            <option key={key} value={key}>{config.label}</option>
                        ))}
                    </select>
                    <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                </div>
 
                {/* Type Filter */}
                <div className="relative">
                    <select
                        value={typeFilter}
                        onChange={(e) => handleTypeChange(e.target.value as any)}
                        className="appearance-none bg-[#0f1115] border border-gray-700 text-gray-300 rounded-lg pl-9 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144] cursor-pointer hover:border-gray-500 transition-colors"
                    >
                        <option value="all">Todos Tipos</option>
                        {Object.entries(TYPE_CONFIG).map(([key, config]) => (
                            <option key={key} value={key}>{config.label}</option>
                        ))}
                    </select>
                    <Tag size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                </div>
 
                {/* Assignee/Supplier Filter */}
                <div className="relative">
                    <select
                        value={assigneeFilter}
                        onChange={(e) => handleAssigneeChange(e.target.value)}
                        className="appearance-none bg-[#0f1115] border border-gray-700 text-gray-300 rounded-lg pl-9 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144] cursor-pointer hover:border-gray-500 transition-colors"
                    >
                        <option value="all">Todos Fornecedores</option>
                        {uniqueAssignees.map(assignee => (
                            <option key={assignee} value={assignee}>{assignee}</option>
                        ))}
                    </select>
                    <Briefcase size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                </div>
            </div>

            {/* Right Controls: Nav & Add */}
            <div className="flex items-center gap-3 w-full xl:w-auto justify-between xl:justify-end">
                {/* Navigation (Always Visible now) */}
                <>
                    <div className="w-px h-8 bg-gray-700 mx-2 hidden lg:block"></div>
                    <div className="flex items-center gap-4 bg-gray-900 p-1.5 rounded-full border border-gray-800">
                        <button onClick={prevMonth} className="p-2 hover:bg-gray-800 rounded-full text-gray-400 hover:text-white transition-colors">
                            <ChevronLeft size={20} />
                        </button>
                        <span className="text-sm font-semibold min-w-[120px] text-center capitalize text-gray-200">
                            {currentDate.toLocaleString('pt-BR', { month: 'long', year: 'numeric' })}
                        </span>
                        <button onClick={nextMonth} className="p-2 hover:bg-gray-800 rounded-full text-gray-400 hover:text-white transition-colors">
                            <ChevronRight size={20} />
                        </button>
                    </div>
                </>
                
                <button 
                    onClick={() => setIsAddingPost(true)}
                    className="bg-[#1e5144] hover:bg-[#163c32] text-white p-2 rounded-lg transition-colors shadow-lg shadow-[#1e5144]/20"
                    title="Adicionar Postagem"
                >
                    <Plus size={20} />
                </button>
            </div>
        </div>

        {/* --- CONTENT AREA --- */}
        {viewMode === 'calendar' ? (
            <>
                {/* --- MINI LEGEND --- */}
                <div className="px-6 py-2 bg-[#15171c] border-b border-gray-800 flex flex-wrap items-center gap-x-6 gap-y-2 text-[10px] text-gray-500">
                    <div className="flex items-center gap-2">
                        <span className="font-bold uppercase tracking-wider text-gray-600">Categorias:</span>
                        {Object.entries(CATEGORY_CONFIG).map(([key, config]) => (
                            <div key={key} className="flex items-center gap-1">
                                <div className={`w-2 h-2 rounded-full ${config.color.replace('text-', 'bg-')}`}></div>
                                <span>{config.label}</span>
                            </div>
                        ))}
                    </div>
                    <div className="w-px h-3 bg-gray-700 hidden sm:block"></div>
                    <div className="flex items-center gap-2">
                        <span className="font-bold uppercase tracking-wider text-gray-600">Tipos:</span>
                        {Object.entries(TYPE_CONFIG).map(([key, config]) => {
                            const Icon = config.icon;
                            return (
                                <div key={key} className="flex items-center gap-1">
                                    <Icon size={10} />
                                    <span>{config.label}</span>
                                </div>
                            );
                        })}
                    </div>
                </div>

                {/* --- GRID HEADER --- */}
                <div className="grid grid-cols-7 border-b border-gray-800 bg-[#1a1d24]">
                    {DAYS_OF_WEEK.map(day => (
                        <div key={day} className="py-3 text-center text-xs font-semibold text-gray-500 uppercase tracking-wider">
                            {day}
                        </div>
                    ))}
                </div>

                {/* --- CALENDAR GRID --- */}
                <div className="flex-grow overflow-y-auto custom-scrollbar bg-[#0f1115] p-4">
                    <div className="grid grid-cols-7 gap-3">
                        {days.map((item, idx) => {
                            if (!item.date) return <div key={idx} className="aspect-[4/5] opacity-0" />; // Empty spacer
                            
                            const dayPosts = filteredPosts.filter(p => p.date === item.date);
                            const isToday = item.date === new Date().toISOString().split('T')[0];
 
                            return (
                                <div 
                                    key={item.date} 
                                    className={`aspect-[4/5] bg-[#1a1d24]/50 border border-white/5 rounded-lg p-2 flex flex-col relative group transition-all hover:border-white/10 ${isToday ? 'ring-1 ring-[#1e5144]/50' : ''}`}
                                >
                                    {/* Day Number */}
                                    <span className={`text-sm font-medium mb-2 ${dayPosts.length > 0 ? 'text-white' : 'text-gray-600'}`}>
                                        {item.day}
                                    </span>
 
                                    {/* Posts List */}
                                    <div className="flex-grow overflow-y-auto custom-scrollbar space-y-2 pr-1">
                                        {dayPosts.map(post => {
                                            // Category defines Color/Border
                                            const catConfig = CATEGORY_CONFIG[post.category];
                                            // Type defines Icon
                                            const typeConfig = TYPE_CONFIG[post.type];
                                            const TypeIcon = typeConfig.icon;
 
                                            // Check publication status for border color
                                            const isPendingPlatforms = (post.platforms?.length || 0) > (post.published_platforms?.length || 0);
                                            const isPublished = post.status === 'published';
                                            
                                            let statusBorder = catConfig.border;
                                            if (isPublished && isPendingPlatforms) statusBorder = 'border-amber-500/50';
                                            else if (isPublished) statusBorder = 'border-emerald-500/50';
 
                                            return (
                                                <button 
                                                    key={post.id}
                                                    onClick={() => setSelectedPost(post)}
                                                    className={`w-full text-left p-2 rounded border ${catConfig.bg} ${statusBorder} hover:opacity-80 transition-all`}
                                                >
                                                    <div className="flex justify-between items-start mb-1">
                                                        <div className="flex items-center gap-1.5">
                                                            <TypeIcon size={12} className={catConfig.color} />
                                                            <span className="text-[9px] text-gray-500 font-mono">{post.time}</span>
                                                        </div>
                                                        <div className="flex items-center">
                                                            {post.status === 'published' && !isPendingPlatforms && <CheckCircle2 size={12} className="text-[#1e5144]" />}
                                                            {post.status === 'published' && isPendingPlatforms && <AlertCircle size={12} className="text-amber-500" />}
                                                            {post.status === 'approved' && <Check size={12} className="text-green-500" />}
                                                            {post.status === 'adjust' && <AlertCircle size={12} className="text-red-500" />}
                                                            {post.status === 'review' && <Clock size={12} className="text-amber-500" />}
                                                        </div>
                                                    </div>
                                                    <p className="text-[10px] leading-tight font-medium text-gray-300 line-clamp-3">
                                                        {post.title}
                                                    </p>
                                                </button>
                                            );
                                        })}
                                    </div>
                                </div>
                            );
                        })}
                    </div>
                </div>
            </>
        ) : (
            /* --- FEED PREVIEW MODE --- */
            <div className="flex-grow overflow-y-auto custom-scrollbar bg-[#0f1115]">
                <div className="min-h-full w-full flex items-center justify-center py-8">
                    <div className="relative border-8 border-gray-900 rounded-[3rem] overflow-hidden bg-black w-[380px] h-[750px] shadow-2xl flex flex-col shrink-0">
                        {/* Phone Notch/Status Bar */}
                        <div className="h-7 w-full flex justify-between items-center px-6 pt-2 select-none">
                            <span className="text-xs font-semibold text-white">9:41</span>
                            <div className="flex gap-1.5 items-center text-white">
                                <Signal size={10} fill="currentColor" />
                                <Wifi size={10} />
                                <Battery size={10} fill="currentColor" />
                            </div>
                        </div>
 
                        {/* Instagram Header */}
                        <div className="h-12 border-b border-gray-800 flex items-center justify-between px-4 text-white">
                            <div className="flex items-center gap-1 font-bold text-lg cursor-pointer">
                                {currentProfile.handle} <ChevronLeft size={16} className="-rotate-90" />
                            </div>
                            <div className="flex gap-4">
                                <Plus size={24} strokeWidth={1.5} />
                                <div className="relative">
                                    <div className="w-1.5 h-1.5 bg-red-500 rounded-full absolute -top-0.5 -right-0.5"></div>
                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"><path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4z"/><line x1="3" y1="6" x2="21" y2="6"/><path d="M16 10a4 4 0 1-8 0"/></svg>
                                </div>
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="18" x2="21" y2="18"/></svg>
                            </div>
                        </div>
 
                        {/* Profile Section */}
                        <div className="px-4 py-3 flex items-center gap-6">
                            <div className="w-20 h-20 rounded-full p-[2px] bg-gradient-to-tr from-yellow-400 via-red-500 to-purple-500">
                                <div className="w-full h-full rounded-full bg-black border-2 border-black flex items-center justify-center overflow-hidden">
                                    <span className="text-2xl font-bold text-white">{currentProfile.initials}</span>
                                </div>
                            </div>
                            <div className="flex-grow flex justify-around text-center text-white">
                                <div>
                                    <div className="font-bold text-lg">1.2k</div>
                                    <div className="text-xs">Posts</div>
                                </div>
                                <div>
                                    <div className="font-bold text-lg">6.8k</div>
                                    <div className="text-xs">Seguidores</div>
                                </div>
                                <div>
                                    <div className="font-bold text-lg">450</div>
                                    <div className="text-xs">Seguindo</div>
                                </div>
                            </div>
                        </div>
 
                        {/* Bio */}
                        <div className="px-4 pb-4 text-white text-xs leading-tight whitespace-pre-line">
                            <div className="font-bold mb-0.5">{currentProfile.name}</div>
                            <div className="text-gray-300">{currentProfile.bio}</div>
                            <div className="text-blue-400 font-medium mt-0.5">{currentProfile.link}</div>
                        </div>
 
                        {/* Action Buttons */}
                        <div className="px-4 flex gap-2 mb-4">
                            <button className="flex-1 bg-[#262626] text-white py-1.5 rounded-lg text-sm font-semibold">Editar perfil</button>
                            <button className="flex-1 bg-[#262626] text-white py-1.5 rounded-lg text-sm font-semibold">Compartilhar perfil</button>
                            <button className="px-2 bg-[#262626] text-white rounded-lg flex items-center justify-center"><User size={16}/></button>
                        </div>
 
                        {/* Tab Bar */}
                        <div className="flex border-t border-gray-800">
                            <div className="flex-1 py-2.5 flex justify-center border-b border-white text-white">
                                <Grid size={22} />
                            </div>
                            <div className="flex-1 py-2.5 flex justify-center text-gray-500">
                                <Video size={24} />
                            </div>
                            <div className="flex-1 py-2.5 flex justify-center text-gray-500">
                                <Users size={22} />
                            </div>
                        </div>
 
                        {/* Grid Content */}
                        <div className="flex-grow overflow-y-auto custom-scrollbar bg-black">
                            <div className="grid grid-cols-3 gap-0.5">
                                {feedPosts.map(post => {
                                        const TypeIcon = TYPE_CONFIG[post.type].icon;
                                        console.log('[POSTCALENDARVIEW] Rendering post in feed:', post.id, 'date:', post.date);
                                        return (
                                            <div
                                                key={post.id}
                                                className="aspect-square bg-[#121212] relative cursor-pointer group"
                                                onClick={() => setSelectedPost(post)}
                                            >
                                                {post.image_url ? (
                                                    <img src={post.image_url} alt="" className="w-full h-full object-cover" />
                                                ) : (
                                                    <div className="w-full h-full flex flex-col items-center justify-center p-2 text-center">
                                                        <TypeIcon size={24} className={CATEGORY_CONFIG[post.category].color} />
                                                        <span className="text-[8px] text-gray-400 mt-1 line-clamp-2">{post.title}</span>
                                                    </div>
                                                )}
                                                {/* Overlay Date */}
                                                <div className="absolute bottom-1 right-1 bg-black/60 px-1 rounded text-[8px] text-white font-mono opacity-0 group-hover:opacity-100 transition-opacity">
                                                    {post.date ? post.date.split('-').slice(1).reverse().join('/') : 'N/A'}
                                            </div>
                                            {/* Type Icon Overlay */}
                                            <div className="absolute top-1 right-1 text-white drop-shadow-md">
                                                {post.type === 'video' && <Video size={12} fill="white" />}
                                                {post.type === 'carousel' && <Layers size={12} fill="white" />}
                                            </div>
                                        </div>
                                    );
                                })}
                                {/* Placeholder items to fill grid look */}
                                {[...Array(Math.max(0, 9 - feedPosts.length))].map((_, i) => (
                                    <div key={`ph-${i}`} className="aspect-square bg-[#1a1a1a]"></div>
                                ))}
                            </div>
                        </div>
 
                        {/* Bottom Nav Mockup */}
                        <div className="h-12 border-t border-gray-800 flex justify-around items-center px-2 text-white bg-black">
                            <div className="p-2"><svg width="24" height="24" viewBox="0 0 24 24" fill="white" stroke="currentColor" strokeWidth="2"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path><polyline points="9 22 9 12 15 12 15 22"></polyline></svg></div>
                            <div className="p-2 text-gray-500"><Search size={24} /></div>
                            <div className="p-2 text-gray-500"><Plus size={24} className="border-2 border-gray-500 rounded-lg p-0.5" /></div>
                            <div className="p-2 text-gray-500"><Video size={24} /></div>
                            <div className="p-2"><div className="w-6 h-6 rounded-full bg-gray-700 border border-gray-500"></div></div>
                        </div>
                        {/* Home Indicator */}
                        <div className="h-5 flex justify-center items-end pb-2">
                            <div className="w-32 h-1 bg-gray-100 rounded-full"></div>
                        </div>
                    </div>
                </div>
            </div>
        )}
      </div>

      {/* --- NEW POST MODAL --- */}
      {isAddingPost && (
         <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/70 backdrop-blur-sm" onClick={() => setIsAddingPost(false)}>
             <div className="bg-[#1a1d24] w-full max-w-md rounded-2xl border border-gray-700 shadow-2xl overflow-hidden" onClick={e => e.stopPropagation()}>
                 <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                     <h3 className="font-bold text-gray-100">Adicionar Nova Postagem</h3>
                     <button onClick={() => setIsAddingPost(false)}><X size={20} className="text-gray-500 hover:text-white" /></button>
                 </div>
                 <form onSubmit={handleCreatePost} className="p-6 space-y-4">
                     <div>
                         <label className="block text-xs text-gray-400 uppercase mb-1">Título</label>
                         <input 
                             type="text" 
                             value={newPostTitle}
                             onChange={(e) => setNewPostTitle(e.target.value)}
                             className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                             placeholder="Ex: Lançamento Produto X"
                         />
                     </div>
                     <div>
                         <label className="block text-xs text-gray-400 uppercase mb-1">Descrição</label>
                         <textarea 
                             value={newPostDescription}
                             onChange={(e) => setNewPostDescription(e.target.value)}
                             className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144] resize-none h-20"
                             placeholder="Breve descrição do conteúdo..."
                         />
                     </div>
                     <div className="grid grid-cols-2 gap-4">
                        <div>
                             <label className="block text-xs text-gray-400 uppercase mb-1">Data</label>
                             <input 
                                 type="date" 
                                 value={newPostDate}
                                 onChange={(e) => setNewPostDate(e.target.value)}
                                 className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                             />
                        </div>
                        <div>
                             <label className="block text-xs text-gray-400 uppercase mb-1">Responsável</label>
                             <select
                                 value={newPostAssignee}
                                 onChange={(e) => setNewPostAssignee(e.target.value)}
                                 className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                 disabled={isLoadingUsers}
                             >
                                 <option value="">Selecione um responsável</option>
                                 {appUsers?.map(user => (
                                     <option key={user.id} value={user.id}>
                                         {user.name}
                                     </option>
                                 ))}
                             </select>
                        </div>
                     </div>
                     <div className="grid grid-cols-2 gap-4">
                         <div>
                             <label className="block text-xs text-gray-400 uppercase mb-1">Categoria</label>
                             <select 
                                 value={newPostCategory}
                                 onChange={(e) => setNewPostCategory(e.target.value as PostCategory)}
                                 className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                             >
                                 {Object.entries(CATEGORY_CONFIG).map(([key, config]) => (
                                     <option key={key} value={key}>{config.label}</option>
                                 ))}
                             </select>
                         </div>
                         <div>
                             <label className="block text-xs text-gray-400 uppercase mb-1">Tipo</label>
                             <select 
                                 value={newPostType}
                                 onChange={(e) => setNewPostType(e.target.value as PostType)}
                                 className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                             >
                                 {Object.entries(TYPE_CONFIG).map(([key, config]) => (
                                     <option key={key} value={key}>{config.label}</option>
                                 ))}
                             </select>
                         </div>
                     </div>
                     <div>
                        <label className="block text-xs text-gray-400 uppercase mb-2">Plataformas de Publicação</label>
                        <div className="flex flex-wrap gap-2">
                            {socialIcons.map(icon => (
                                <button
                                    key={icon.id}
                                    type="button"
                                    onClick={() => toggleNewPostPlatform(icon.id)}
                                    className={`flex items-center gap-2 px-3 py-1.5 rounded-lg border text-sm transition-colors ${newPostPlatforms.includes(icon.id) ? 'bg-[#1e5144]/20 border-[#1e5144] text-white' : 'bg-[#0f1115] border-gray-700 text-gray-400 hover:border-gray-500'}`}
                                >
                                    <icon.icon size={14} />
                                    <span className="capitalize">{icon.id}</span>
                                </button>
                            ))}
                        </div>
                     </div>
                     <button 
                         type="submit" 
                         className="w-full bg-[#1e5144] hover:bg-[#163c32] text-white font-medium py-2 rounded-lg transition-colors"
                     >
                         Criar Postagem
                     </button>
                 </form>
             </div>
         </div>
      )}

      {/* --- PUBLISHING CONFIRMATION MODAL --- */}
      {isPublishingModalOpen && selectedPost && (
          <div className="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm" onClick={() => setIsPublishingModalOpen(false)}>
              <div className="bg-[#1a1d24] w-full max-w-sm rounded-2xl border border-gray-700 shadow-2xl overflow-hidden animate-in zoom-in-95" onClick={e => e.stopPropagation()}>
                  <div className="p-4 border-b border-gray-800 bg-[#20232b]">
                      <h3 className="font-bold text-gray-100 flex items-center gap-2">
                          <CloudLightning size={18} className="text-[#1e5144]" />
                          Confirmar Publicação
                      </h3>
                      <p className="text-xs text-gray-500 mt-1">Marque as redes onde o post foi publicado:</p>
                  </div>
                  <div className="p-6 space-y-3">
                        {selectedPost.platforms && selectedPost.platforms.length > 0 ? (
                             selectedPost.platforms.map(platform => {
                                const Icon = socialIcons.find(s => s.id === platform)?.icon || CloudLightning;
                                const isChecked = publishingPlatforms.includes(platform);
                                return (
                                    <div 
                                        key={platform} 
                                        className={`flex items-center justify-between p-3 rounded-lg border cursor-pointer transition-colors ${isChecked ? 'bg-[#1e5144]/10 border-[#1e5144]/30' : 'bg-[#0f1115] border-gray-700 hover:bg-gray-800'}`}
                                        onClick={() => togglePublishingPlatform(platform)}
                                    >
                                        <div className="flex items-center gap-3">
                                            <Icon size={18} className="text-gray-300" />
                                            <span className="capitalize text-sm font-medium text-gray-200">{platform}</span>
                                        </div>
                                        <div className={`w-5 h-5 rounded border flex items-center justify-center transition-colors ${isChecked ? 'bg-[#1e5144] border-[#1e5144]' : 'border-gray-600'}`}>
                                            {isChecked && <Check size={14} className="text-white" />}
                                        </div>
                                    </div>
                                );
                             })
                        ) : (
                            <div className="text-center text-gray-500 text-sm py-4">
                                Nenhuma plataforma foi planejada para este post.
                            </div>
                        )}
                  </div>
                  <div className="p-4 border-t border-gray-800 bg-[#15171c] flex justify-end gap-2">
                      <button 
                        onClick={() => setIsPublishingModalOpen(false)}
                        className="px-3 py-1.5 text-sm text-gray-400 hover:text-white"
                      >
                          Cancelar
                      </button>
                      <button 
                        onClick={confirmPublishing}
                        className="px-4 py-1.5 bg-[#1e5144] text-white rounded-lg text-sm font-medium hover:bg-[#163c32]"
                      >
                          Confirmar
                      </button>
                  </div>
              </div>
          </div>
      )}

      {/* --- POST DETAILS MODAL --- */}
      {selectedPost && !isPublishingModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/70 backdrop-blur-sm" onClick={() => setSelectedPost(null)}>
              <div className="w-full max-w-5xl bg-[#1a1d24] border border-gray-700 rounded-2xl shadow-2xl overflow-hidden flex flex-col md:flex-row max-h-[90vh]" onClick={e => e.stopPropagation()}>
                  
                  {/* Left: Preview Image */}
                  <div className="w-full md:w-1/2 bg-black flex items-center justify-center relative group">
                      {selectedPost.image_url ? (
                          <>
                              <img
                                  src={selectedPost.image_url}
                                  alt={selectedPost.title} 
                                  className="max-h-full max-w-full object-contain"
                              />
                              {/* Category Badge */}
                              <div className="absolute top-4 left-4">
                                  <span className={`px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wide border bg-black/50 backdrop-blur-md flex items-center gap-2 ${CATEGORY_CONFIG[selectedPost.category].color} border-gray-700`}>
                                       {CATEGORY_CONFIG[selectedPost.category].label}
                                  </span>
                              </div>
                              {/* Change Photo Overlay Button */}
                              <div className="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                                  <button 
                                      onClick={handleChangePhoto}
                                      className="px-4 py-2 bg-white/20 backdrop-blur-md hover:bg-white/30 text-white rounded-lg text-sm font-medium flex items-center gap-2 transition-colors border border-white/30"
                                  >
                                      <RefreshCcw size={16} /> Trocar Foto
                                  </button>
                              </div>
                          </>
                      ) : (
                          <div className="flex flex-col items-center justify-center text-gray-500 gap-4">
                              <ImageIcon size={64} className="opacity-20" />
                              <p className="text-sm">Nenhuma imagem anexada</p>
                              <button 
                                  onClick={attachImage}
                                  className="px-4 py-2 bg-gray-800 hover:bg-gray-700 text-gray-300 rounded-lg text-sm font-medium flex items-center gap-2 transition-colors border border-gray-700"
                              >
                                  <Paperclip size={16} /> Anexar Imagem
                              </button>
                          </div>
                      )}
                  </div>
 
                  {/* Right: Info & Actions */}
                  <div className="w-full md:w-1/2 flex flex-col h-full border-l border-gray-800">
                      
                      {/* Header */}
                      <div className="p-6 border-b border-gray-800 flex justify-between items-start">
                          <div>
                              <div className="flex items-center gap-3 mb-2">
                                  {isEditingTitle ? (
                                      <input 
                                          type="text" 
                                          value={editTitleValue}
                                          onChange={(e) => setEditTitleValue(e.target.value)}
                                          onBlur={saveTitle}
                                          onKeyDown={(e) => e.key === 'Enter' && saveTitle()}
                                          className="text-xl font-bold text-white bg-[#0f1115] border border-gray-700 rounded px-2 py-1 focus:outline-none focus:ring-1 focus:ring-[#1e5144] w-full"
                                          autoFocus
                                      />
                                  ) : (
                                      <h2 
                                          className="text-xl font-bold text-white cursor-pointer hover:underline decoration-dotted decoration-gray-500"
                                          onDoubleClick={startEditingTitle}
                                          title="Duplo clique para editar"
                                      >
                                          {selectedPost.title}
                                      </h2>
                                  )}
                              </div>
                              <div className="flex items-center gap-4 text-sm text-gray-400 mb-3 flex-wrap">
                                  <span className="flex items-center gap-1">
                                      <Calendar size={14}/>
                                      {selectedPost.date ? selectedPost.date.split('-').reverse().join('/') : 'N/A'}
                                  </span>
                                  <span className="flex items-center gap-1"><Clock size={14}/> {selectedPost.time}</span>
                                  {selectedPost.assignee && (
                                      <span className="flex items-center gap-1 border-l border-gray-700 pl-4"><User size={14}/> {selectedPost.assignee.name}</span>
                                  )}
                              </div>
                              
                              {/* Status Badge */}
                              <div>
                                  {selectedPost.status === 'published' && <span className="px-2 py-1 bg-[#1e5144]/20 text-emerald-400 text-xs rounded border border-[#1e5144]/30 font-bold uppercase">Publicado</span>}
                                  {selectedPost.status === 'approved' && <span className="px-2 py-1 bg-green-500/10 text-green-400 text-xs rounded border border-green-500/20 font-bold uppercase">Aprovado / Agendamento</span>}
                                  {selectedPost.status === 'adjust' && <span className="px-2 py-1 bg-red-500/10 text-red-400 text-xs rounded border border-red-500/20 font-bold uppercase">Solicitado Ajuste</span>}
                                  {selectedPost.status === 'review' && <span className="px-2 py-1 bg-amber-500/10 text-amber-400 text-xs rounded border border-amber-500/20 font-bold uppercase">Em Revisão</span>}
                                  {selectedPost.status === 'in_progress' && <span className="px-2 py-1 bg-blue-500/10 text-blue-400 text-xs rounded border border-blue-500/20 font-bold uppercase">Em Andamento</span>}
                                  <span className="ml-2 px-2 py-1 bg-gray-800 text-gray-400 text-xs rounded border border-gray-700 font-bold uppercase">{TYPE_CONFIG[selectedPost.type].label}</span>
                              </div>
                          </div>
                          <div className="flex gap-2">
                              <button 
                                  onClick={handleDeletePost} 
                                  className="p-2 text-gray-500 hover:text-red-400 hover:bg-red-500/10 rounded-full transition-colors"
                                  title="Excluir Postagem"
                              >
                                  <Trash2 size={20} />
                              </button>
                              <button onClick={() => setSelectedPost(null)} className="p-2 text-gray-500 hover:text-white transition-colors">
                                  <X size={24} />
                              </button>
                          </div>
                      </div>
 
                      {/* Published Platforms Status Display */}
                      <div className="px-6 py-4 border-b border-gray-800 bg-[#0f1115]">
                          <p className="text-xs font-semibold text-gray-500 uppercase mb-2">Plataformas Planejadas</p>
                          <div className="flex gap-4">
                              {selectedPost.platforms && selectedPost.platforms.length > 0 ? (
                                  selectedPost.platforms.map((platform) => {
                                      const iconData = socialIcons.find(s => s.id === platform);
                                      if(!iconData) return null;
                                       
                                      const isPublished = selectedPost.published_platforms?.includes(platform);
                                       
                                      return (
                                          <div 
                                              key={platform}
                                              className="relative group"
                                              title={`${platform}: ${isPublished ? 'Publicado' : 'Pendente'}`}
                                          >
                                              <div className={`w-8 h-8 rounded-full border flex items-center justify-center transition-all ${isPublished ? 'bg-green-500/10 border-green-500/50' : 'bg-gray-800 border-gray-700 opacity-50'}`}>
                                                  <iconData.icon size={16} className={isPublished ? 'text-green-500' : 'text-gray-500'} />
                                              </div>
                                              {isPublished && (
                                                  <div className="absolute -top-1 -right-1 w-3 h-3 bg-green-500 rounded-full border border-[#0f1115] flex items-center justify-center">
                                                      <Check size={8} className="text-white" />
                                                  </div>
                                              )}
                                          </div>
                                      );
                                  })
                              ) : (
                                  <span className="text-sm text-gray-500 italic">Nenhuma plataforma definida</span>
                              )}
                          </div>
                      </div>
 
                      {/* Scrollable Content */}
                      <div className="flex-grow overflow-y-auto custom-scrollbar p-6 space-y-6">
                          
                          {/* Description */}
                          <div>
                              <div className="flex items-center justify-between mb-2">
                                  <h3 className="text-xs font-semibold text-gray-500 uppercase tracking-wider">Descrição</h3>
                                  {isEditingDesc && (
                                      <div className="flex gap-1">
                                          <button onClick={() => setIsEditingDesc(false)} className="text-gray-500 hover:text-white p-1"><X size={14} /></button>
                                          <button onClick={saveDescription} className="text-emerald-400 hover:text-emerald-300 p-1"><Save size={14} /></button>
                                      </div>
                                  )}
                              </div>
                              
                              {isEditingDesc ? (
                                  <textarea 
                                      value={editDescValue}
                                      onChange={(e) => setEditDescValue(e.target.value)}
                                      className="w-full bg-[#0f1115] border border-gray-700 rounded-lg p-3 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144] min-h-[100px]"
                                  />
                              ) : (
                                  <p 
                                      className="text-gray-300 text-sm leading-relaxed p-4 rounded-lg bg-[#20232b]/50 border border-gray-800/50 cursor-pointer hover:border-gray-600 transition-colors"
                                      onDoubleClick={startEditingDescription}
                                      title="Duplo clique para editar"
                                  >
                                      {selectedPost.description || <span className="italic text-gray-600">Sem descrição...</span>}
                                  </p>
                              )}
                          </div>
 
                          {/* Caption */}
                          <div>
                              <div className="flex items-center justify-between mb-2">
                                  <h3 className="text-xs font-semibold text-gray-500 uppercase tracking-wider">Legenda</h3>
                                  <button 
                                      onClick={copyCaptionToClipboard}
                                      className={`flex items-center gap-1 px-2 py-1 rounded text-xs transition-colors ${copyFeedback ? 'text-emerald-400 bg-emerald-500/10' : 'text-gray-500 hover:text-white hover:bg-gray-800'}`}
                                      title="Copiar Legenda"
                                  >
                                      {copyFeedback ? <Check size={14} /> : <Copy size={14} />}
                                      {copyFeedback ? 'Copiado!' : 'Copiar'}
                                  </button>
                              </div>
                              <p className="text-gray-300 text-sm leading-relaxed whitespace-pre-line bg-[#0f1115] p-4 rounded-lg border border-gray-800 min-h-[60px]">
                                  {selectedPost.caption || <span className="text-gray-600 italic">Sem legenda...</span>}
                              </p>
                          </div>
 
                          {/* History Timeline */}
                          <div>
                              <h3 className="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-3">Histórico</h3>
                              <div className="max-h-[200px] overflow-y-auto custom-scrollbar pr-2">
                                  <div className="space-y-4 relative before:absolute before:left-2 before:top-2 before:bottom-0 before:w-0.5 before:bg-gray-800">
                                      {selectedPost.history.map((event, idx) => (
                                          <div key={idx} className="relative pl-8">
                                              <div className={`absolute left-0 top-1 w-4 h-4 rounded-full border-2 border-[#1a1d24] flex items-center justify-center 
                                                  ${event.action === 'published' ? 'bg-[#1e5144]' : event.action === 'approved' ? 'bg-green-500' : event.action === 'adjust_request' ? 'bg-red-500' : 'bg-blue-500'}`
                                              }></div>
                                              <div className="text-xs text-gray-500 mb-0.5 flex justify-between">
                                                  <span>{event.timestamp}</span>
                                                  <span className="font-medium text-gray-400">{event.user}</span>
                                              </div>
                                              <p className="text-sm text-gray-300">
                                                  {event.action === 'upload' && 'Fez upload da mídia'}
                                                  {event.action === 'status_change' && event.text}
                                                  {event.action === 'published' && <span className="text-emerald-400 font-bold">Publicou a postagem!</span>}
                                                  {event.action === 'approved' && <span className="text-green-400">Aprovou para agendamento</span>}
                                                  {event.action === 'adjust_request' && (
                                                      <span className="text-red-300">Solicitou ajuste: "{event.text}"</span>
                                                  )}
                                              </p>
                                          </div>
                                      ))}
                                  </div>
                              </div>
                          </div>
                      </div>
 
                      {/* Footer Actions - Status Workflow */}
                      <div className="p-6 border-t border-gray-800 bg-[#0f1115]">
                          <div className="grid grid-cols-2 gap-2 mb-3">
                              <button 
                                  onClick={() => handleStatusChange('in_progress')}
                                  className={`py-2 rounded-lg text-xs font-bold uppercase transition-colors border ${selectedPost.status === 'in_progress' ? 'bg-blue-500/20 text-blue-400 border-blue-500/50' : 'bg-gray-800 text-gray-400 border-gray-700 hover:bg-gray-700'}`}
                              >
                                  Em Andamento
                              </button>
                              <button 
                                  onClick={() => handleStatusChange('review')}
                                  className={`py-2 rounded-lg text-xs font-bold uppercase transition-colors border ${selectedPost.status === 'review' ? 'bg-amber-500/20 text-amber-400 border-amber-500/50' : 'bg-gray-800 text-gray-400 border-gray-700 hover:bg-gray-700'}`}
                              >
                                  Revisão
                              </button>
                          </div>
 
                          {selectedPost.status === 'adjust' && (
                              <div className="mb-3 bg-red-500/10 border border-red-500/20 p-3 rounded-lg">
                                  <p className="text-xs text-red-400 mb-1 font-bold uppercase">Solicitar Ajuste</p>
                                  <div className="flex gap-2">
                                      <input 
                                          type="text" 
                                          value={adjustComment} 
                                          onChange={(e) => setAdjustComment(e.target.value)}
                                          className="flex-grow bg-[#1a1d24] border border-gray-700 rounded px-2 py-1 text-sm text-white focus:outline-none"
                                          placeholder="Motivo..."
                                      />
                                      <button onClick={() => handleStatusChange('adjust', adjustComment)} className="bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded text-xs">Enviar</button>
                                  </div>
                              </div>
                          )}
 
                          <div className="grid grid-cols-3 gap-2">
                               <button 
                                  onClick={() => { setAdjustComment(''); handleStatusChange('adjust', 'Solicitado Ajuste'); }}
                                  className={`py-2 rounded-lg text-xs font-bold uppercase transition-colors border ${selectedPost.status === 'adjust' ? 'bg-red-500/20 text-red-400 border-red-500/50' : 'bg-gray-800 text-gray-400 border-gray-700 hover:bg-red-900/20 hover:text-red-300'}`}
                              >
                                  Solicitar Ajuste
                              </button>
 
                               <button 
                                  onClick={() => handleStatusChange('approved')}
                                  className={`py-2 rounded-lg text-xs font-bold uppercase transition-colors border ${selectedPost.status === 'approved' ? 'bg-green-500/20 text-green-400 border-green-500/50' : 'bg-gray-800 text-gray-400 border-gray-700 hover:bg-green-900/20 hover:text-green-300'}`}
                              >
                                  Pronta p/ Agendar
                              </button>
                              
                              <button 
                                  onClick={() => handleStatusChange('published')}
                                  className={`py-2 rounded-lg text-xs font-bold uppercase transition-colors border flex items-center justify-center gap-1 ${selectedPost.status === 'published' ? 'bg-[#1e5144] text-white border-[#1e5144]' : 'bg-gray-800 text-gray-400 border-gray-700 hover:bg-[#1e5144]/50 hover:text-white'}`}
                              >
                                  <CloudLightning size={14} />
                                  Publicado
                              </button>
                          </div>
                      </div>
 
                  </div>
              </div>
        </div>
      )}
 
    </div>
  );
};

export default PostCalendarView;
