import React, { useState, useEffect } from 'react';
import type { SocialPost } from '@/shared/types';
import { ThumbsUp, MessageCircle, Share2, Calendar, Clock, Filter, Save, Edit2, Trash2, Plus, ChevronLeft, ChevronRight } from 'lucide-react';
import type { SocialPlatform, SocialPostType } from '@/shared/types';
import { useSocialMutations } from '../hooks/useSocialMutations';
import { useBrands } from '../hooks/useBrands';
import { useToast } from '@/shared/hooks/useToast';

interface SocialPostsViewProps {
  data?: SocialPost[];
  onRefresh?: () => void;
}

const SocialPostsView: React.FC<SocialPostsViewProps> = ({ data, onRefresh }) => {
  const { error: showError } = useToast();

  // Filter States
  const [sortBy, setSortBy] = useState<'post_date' | 'likes' | 'comments' | 'shares' | 'created_at'>('post_date');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');
  const [filterBrand, setFilterBrand] = useState('');
  const [filterPlatform, setFilterPlatform] = useState<SocialPlatform | ''>('');
  const [filterPostType, setFilterPostType] = useState<SocialPostType | ''>('');

  // Helper function to get brand name by brand_id
  const getBrandName = (brandId: string): string => {
    const brand = brands.find(b => b.id === brandId);
    return brand?.name || brandId;
  };
  
  // Pagination States
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(10);

  // Form States
  const [isAdding, setIsAdding] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [brandId, setBrandId] = useState('');
  const [platform, setPlatform] = useState<SocialPlatform>('instagram');
  const [postDate, setPostDate] = useState('');
  const [postTime, setPostTime] = useState('');
  const [likes, setLikes] = useState(0);
  const [comments, setComments] = useState(0);
  const [shares, setShares] = useState<number | undefined>(undefined);
  const [postType, setPostType] = useState<SocialPostType>('image');
  const [caption, setCaption] = useState('');
  const [followersAtPost, setFollowersAtPost] = useState<number | undefined>(undefined);

  // Brands Query
  const { data: brands = [], isLoading: isLoadingBrands } = useBrands();

  // Mutations
  const { createPost, updatePost, deletePost, isCreatingPost, isUpdatingPost, isDeletingPost } = useSocialMutations();

  // Platform options
  const platformOptions: { value: SocialPlatform; label: string }[] = [
    { value: 'instagram', label: 'Instagram' },
    { value: 'facebook', label: 'Facebook' },
    { value: 'linkedin', label: 'LinkedIn' },
    { value: 'tiktok', label: 'TikTok' },
    { value: 'twitter', label: 'Twitter' },
  ];

  // Post type options
  const postTypeOptions: { value: SocialPostType; label: string }[] = [
    { value: 'image', label: 'Imagem' },
    { value: 'video', label: 'Vídeo' },
    { value: 'carousel', label: 'Carrossel' },
    { value: 'reel', label: 'Reel' },
    { value: 'story', label: 'Story' },
  ];

  // Reset form states
  const resetForm = () => {
    setBrandId('');
    setPlatform('instagram');
    setPostDate('');
    setPostTime('');
    setLikes(0);
    setComments(0);
    setShares(undefined);
    setPostType('image');
    setCaption('');
    setFollowersAtPost(undefined);
  };

  // Preencher formulário no modo edição
  useEffect(() => {
    if (editingId && data) {
      const item = data.find(item => item.id === editingId);
      if (item) {
        setBrandId(item.brand_id || '');
        setPlatform(item.platform);
        setPostDate(item.post_date);
        setPostTime(item.post_time || '');
        setLikes(item.likes);
        setComments(item.comments);
        setShares(item.shares);
        setPostType(item.post_type);
        setCaption(item.caption || '');
        setFollowersAtPost(item.followers_at_post);
      }
    }
  }, [editingId, data]);

  // Resetar formulário no modo criação
  useEffect(() => {
    if (isAdding) {
      resetForm();
    }
  }, [isAdding]);

  const handleCreate = (e: React.FormEvent) => {
    e.preventDefault();

    // Form validation
    if (!brandId.trim()) {
      showError('Marca é obrigatória');
      return;
    }
    if (!postDate) {
      showError('Data do post é obrigatória');
      return;
    }
    if (likes < 0) {
      showError('Likes não pode ser negativo');
      return;
    }
    if (comments < 0) {
      showError('Comentários não pode ser negativo');
      return;
    }
    if (shares !== undefined && shares < 0) {
      showError('Shares não pode ser negativo');
      return;
    }
    if (followersAtPost !== undefined && followersAtPost < 0) {
      showError('Seguidores no post não pode ser negativo');
      return;
    }

    createPost({
      brand_id: brandId,
      platform,
      post_date: postDate,
      post_time: postTime || undefined,
      likes,
      comments,
      shares,
      post_type: postType,
      caption: caption || undefined,
      followers_at_post: followersAtPost,
    }, {
      onSuccess: () => {
        setIsAdding(false);
        resetForm();
        onRefresh?.();
      },
      onError: (error) => {
        console.error('Erro ao criar post:', error);
        showError('Erro ao criar post. Tente novamente.');
      }
    });
  };

  const handleUpdate = (e: React.FormEvent) => {
    e.preventDefault();

    if (!editingId) return;

    // Form validation
    if (!brandId.trim()) {
      showError('Marca é obrigatória');
      return;
    }
    if (!postDate) {
      showError('Data do post é obrigatória');
      return;
    }
    if (likes < 0) {
      showError('Likes não pode ser negativo');
      return;
    }
    if (comments < 0) {
      showError('Comentários não pode ser negativo');
      return;
    }
    if (shares !== undefined && shares < 0) {
      showError('Shares não pode ser negativo');
      return;
    }
    if (followersAtPost !== undefined && followersAtPost < 0) {
      showError('Seguidores no post não pode ser negativo');
      return;
    }

    updatePost({
      id: editingId,
      data: {
        brand_id: brandId,
        platform,
        post_date: postDate,
        post_time: postTime || undefined,
        likes,
        comments,
        shares,
        post_type: postType,
        caption: caption || undefined,
        followers_at_post: followersAtPost,
      }
    }, {
      onSuccess: () => {
        setEditingId(null);
        resetForm();
        onRefresh?.();
      },
      onError: (error) => {
        console.error('Erro ao atualizar post:', error);
        showError('Erro ao atualizar post. Tente novamente.');
      }
    });
  };

  const handleDelete = (id: string) => {
    if (!confirm('Tem certeza que deseja deletar este post?')) return;

    deletePost(id);
  };

  // Aplicar filtros localmente
  const filteredData = data ? data.filter(item => {
    const brandName = getBrandName(item.brand_id);
    const matchBrand = filterBrand === '' || brandName.toLowerCase().includes(filterBrand.toLowerCase());
    const matchPlatform = filterPlatform === '' || item.platform === filterPlatform;
    const matchPostType = filterPostType === '' || item.post_type === filterPostType;
    return matchBrand && matchPlatform && matchPostType;
  }) : [];

  // Ordenar localmente
  const sortedData = [...filteredData].sort((a, b) => {
    let comparison = 0;

    switch (sortBy) {
      case 'post_date':
        comparison = new Date(a.post_date).getTime() - new Date(b.post_date).getTime();
        break;
      case 'likes':
        comparison = a.likes - b.likes;
        break;
      case 'comments':
        comparison = a.comments - b.comments;
        break;
      case 'shares':
        comparison = (a.shares || 0) - (b.shares || 0);
        break;
      case 'created_at':
        comparison = new Date(a.created_at).getTime() - new Date(b.created_at).getTime();
        break;
    }

    return sortOrder === 'asc' ? comparison : -comparison;
  });

  // Pagination
  const totalPages = Math.ceil(sortedData.length / itemsPerPage);
  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  const paginatedData = sortedData.slice(startIndex, endIndex);

  // Reset pagination when filters change
  useEffect(() => {
    setCurrentPage(1);
  }, [filterBrand, filterPlatform, filterPostType, sortBy, sortOrder]);

  // Verificar se há dados
  const hasData = sortedData && sortedData.length > 0;

  // Calcular estatísticas (baseado em todos os dados filtrados, não apenas na página atual)
  const totalLikes = sortedData.reduce((sum, item) => sum + item.likes, 0);
  const totalComments = sortedData.reduce((sum, item) => sum + item.comments, 0);
  const totalShares = sortedData.reduce((sum, item) => sum + (item.shares || 0), 0);

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col text-gray-200">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-100">Posts de Social Media</h1>
          <p className="text-gray-500">Gerenciamento de posts individuais por plataforma</p>
        </div>
      </div>

      {/* Unified Control Bar */}
      <div className="flex flex-col md:flex-row justify-between items-center mb-6 gap-4">
        {/* Filters (Left Side) */}
        <div className="flex flex-wrap gap-4 w-full md:w-auto items-center">
          <div className="relative">
            <select
              value={sortBy}
              onChange={(e) => setSortBy(e.target.value as 'post_date' | 'likes' | 'comments' | 'shares' | 'created_at')}
              className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
            >
              <option value="post_date">Data do Post</option>
              <option value="likes">Likes</option>
              <option value="comments">Comentários</option>
              <option value="shares">Shares</option>
              <option value="created_at">Data de Criação</option>
            </select>
            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
          </div>

          <div className="relative">
            <select
              value={sortOrder}
              onChange={(e) => setSortOrder(e.target.value as 'asc' | 'desc')}
              className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
            >
              <option value="desc">Descendente</option>
              <option value="asc">Ascendente</option>
            </select>
            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
          </div>

          <div className="relative">
            <input
              type="text"
              value={filterBrand}
              onChange={(e) => setFilterBrand(e.target.value)}
              placeholder="Buscar marca..."
              className="bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-4 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144] placeholder-gray-600"
            />
            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
          </div>

          <div className="relative">
            <select
              value={filterPlatform}
              onChange={(e) => setFilterPlatform(e.target.value as SocialPlatform | '')}
              className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
            >
              <option value="">Todas Plataformas</option>
              {platformOptions.map(opt => (
                <option key={opt.value} value={opt.value}>{opt.label}</option>
              ))}
            </select>
            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
          </div>

          <div className="relative">
            <select
              value={filterPostType}
              onChange={(e) => setFilterPostType(e.target.value as SocialPostType | '')}
              className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
            >
              <option value="">Todos Tipos</option>
              {postTypeOptions.map(opt => (
                <option key={opt.value} value={opt.value}>{opt.label}</option>
              ))}
            </select>
            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
          </div>
        </div>

        {/* Actions (Right Side) */}
        <div className="flex items-center gap-3 w-full md:w-auto justify-end">
          <button
            onClick={() => setIsAdding(!isAdding)}
            className="bg-[#1e5144] text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-[#163c32] transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/30"
          >
            <Plus size={18} /> Novo Post
          </button>
        </div>
      </div>

      {/* Add/Edit Form Panel */}
      {(isAdding || editingId) && (
        <div className="bg-[#1a1d24] rounded-xl border border-gray-700 shadow-xl mb-6 animate-in slide-in-from-top-4 flex flex-col">
          <div className="p-4 border-b border-gray-800 bg-[#20232b]">
            <h3 className="font-bold text-gray-100">
              {isAdding ? 'Novo Post' : 'Editar Post'}
            </h3>
          </div>

          <form onSubmit={isAdding ? handleCreate : handleUpdate} className="p-6 space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-400 mb-1">
                Marca <span className="text-rose-500">*</span>
              </label>
              {isLoadingBrands ? (
                <div className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-500">
                  Carregando marcas...
                </div>
              ) : (
                <select
                  value={brandId}
                  onChange={(e) => setBrandId(e.target.value)}
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                >
                  <option value="">Selecione uma marca...</option>
                  {brands.map((brand) => (
                    <option key={brand.id} value={brand.id}>
                      {brand.name}
                    </option>
                  ))}
                </select>
              )}
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Plataforma <span className="text-rose-500">*</span>
                </label>
                <select
                  value={platform}
                  onChange={(e) => setPlatform(e.target.value as SocialPlatform)}
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                >
                  {platformOptions.map(opt => (
                    <option key={opt.value} value={opt.value}>{opt.label}</option>
                  ))}
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Tipo de Post <span className="text-rose-500">*</span>
                </label>
                <select
                  value={postType}
                  onChange={(e) => setPostType(e.target.value as SocialPostType)}
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                >
                  {postTypeOptions.map(opt => (
                    <option key={opt.value} value={opt.value}>{opt.label}</option>
                  ))}
                </select>
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Data do Post <span className="text-rose-500">*</span>
                </label>
                <input
                  type="date"
                  value={postDate}
                  onChange={(e) => setPostDate(e.target.value)}
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Hora do Post <span className="text-gray-500 font-normal ml-1">(Opcional)</span>
                </label>
                <input
                  type="time"
                  value={postTime}
                  onChange={(e) => setPostTime(e.target.value)}
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                />
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Likes <span className="text-rose-500">*</span>
                </label>
                <input
                  type="number"
                  min="0"
                  value={likes}
                  onChange={(e) => setLikes(parseInt(e.target.value) || 0)}
                  placeholder="0"
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Comentários <span className="text-rose-500">*</span>
                </label>
                <input
                  type="number"
                  min="0"
                  value={comments}
                  onChange={(e) => setComments(parseInt(e.target.value) || 0)}
                  placeholder="0"
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Shares <span className="text-gray-500 font-normal ml-1">(Opcional)</span>
                </label>
                <input
                  type="number"
                  min="0"
                  value={shares || ''}
                  onChange={(e) => setShares(e.target.value ? parseInt(e.target.value) : undefined)}
                  placeholder="0"
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                />
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Seguidores no Post <span className="text-gray-500 font-normal ml-1">(Opcional)</span>
                </label>
                <input
                  type="number"
                  min="0"
                  value={followersAtPost || ''}
                  onChange={(e) => setFollowersAtPost(e.target.value ? parseInt(e.target.value) : undefined)}
                  placeholder="Número de seguidores"
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-400 mb-1">
                Caption <span className="text-gray-500 font-normal ml-1">(Opcional)</span>
              </label>
              <textarea
                value={caption}
                onChange={(e) => setCaption(e.target.value)}
                placeholder="Caption do post..."
                rows={3}
                className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] resize-none"
              />
            </div>

            <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
              <button
                type="button"
                onClick={() => {
                  setIsAdding(false);
                  setEditingId(null);
                  resetForm();
                }}
                className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
              >
                Cancelar
              </button>
              <button
                type="submit"
                disabled={isCreatingPost || isUpdatingPost}
                className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isCreatingPost || isUpdatingPost ? (
                  <>
                    <div className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin" />
                    {isAdding ? 'Criando...' : 'Atualizando...'}
                  </>
                ) : (
                  <>
                    <Save size={16} />
                    {isAdding ? 'Criar' : 'Atualizar'}
                  </>
                )}
              </button>
            </div>
          </form>
        </div>
      )}

      {/* Stats Cards Row */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-blue-500/10 rounded-lg text-blue-400 border border-blue-500/20">
            <ThumbsUp size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Total de Likes</p>
            <p className="text-xl font-bold text-white">{totalLikes.toLocaleString()}</p>
          </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-pink-500/10 rounded-lg text-pink-400 border border-pink-500/20">
            <MessageCircle size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Total de Comentários</p>
            <p className="text-xl font-bold text-white">{totalComments.toLocaleString()}</p>
          </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-green-500/10 rounded-lg text-green-400 border border-green-500/20">
            <Share2 size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Total de Shares</p>
            <p className="text-xl font-bold text-white">{totalShares.toLocaleString()}</p>
          </div>
        </div>
      </div>

      {/* Main Content Table */}
      <div className="bg-[#1a1d24] rounded-xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col">
        <div className="p-4 border-b border-gray-800 bg-[#20232b]">
          <h2 className="font-semibold text-gray-100">Lista de Posts</h2>
        </div>

        <div className="overflow-x-auto custom-scrollbar flex-grow">
          <table className="w-full text-sm text-left">
            <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
              <tr>
                <th className="px-6 py-4">Marca</th>
                <th className="px-6 py-4">Plataforma</th>
                <th className="px-6 py-4">Tipo</th>
                <th className="px-6 py-4">Data</th>
                <th className="px-6 py-4 text-right">Likes</th>
                <th className="px-6 py-4 text-right">Comentários</th>
                <th className="px-6 py-4 text-right">Shares</th>
                <th className="px-6 py-4 text-center">Ações</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-800">
              {!hasData ? (
                <tr>
                  <td colSpan={8} className="px-6 py-8 text-center text-gray-500">
                    Nenhum post disponível para exibição.
                  </td>
                </tr>
              ) : (
                paginatedData.map((item) => {
                  const platformLabel = platformOptions.find(p => p.value === item.platform)?.label || item.platform;
                  const postTypeLabel = postTypeOptions.find(t => t.value === item.post_type)?.label || item.post_type;

                  return (
                    <tr
                      key={item.id}
                      className="hover:bg-[#20232b] transition-colors"
                    >
                      <td className="px-6 py-4">
                        <span className="font-medium text-gray-300">{getBrandName(item.brand_id)}</span>
                      </td>
                      <td className="px-6 py-4">
                        <span className="px-2 py-1 rounded text-xs font-medium bg-blue-500/10 text-blue-400 border border-blue-500/20">
                          {platformLabel}
                        </span>
                      </td>
                      <td className="px-6 py-4">
                        <span className="px-2 py-1 rounded text-xs font-medium bg-purple-500/10 text-purple-400 border border-purple-500/20">
                          {postTypeLabel}
                        </span>
                      </td>
                      <td className="px-6 py-4">
                        <div className="flex items-center gap-2 text-gray-400">
                          <Calendar size={14} className="text-gray-600" />
                          <span className="font-mono">{new Date(item.post_date).toLocaleDateString('pt-BR')}</span>
                          {item.post_time && (
                            <>
                              <Clock size={14} className="text-gray-600" />
                              <span className="font-mono">{item.post_time}</span>
                            </>
                          )}
                        </div>
                      </td>
                      <td className="px-6 py-4 text-right">
                        <div className="flex items-center justify-end gap-2 text-gray-300">
                          <ThumbsUp size={16} className="text-blue-500" />
                          <span className="font-mono">{item.likes.toLocaleString()}</span>
                        </div>
                      </td>
                      <td className="px-6 py-4 text-right">
                        <div className="flex items-center justify-end gap-2 text-gray-300">
                          <MessageCircle size={16} className="text-pink-500" />
                          <span className="font-mono">{item.comments.toLocaleString()}</span>
                        </div>
                      </td>
                      <td className="px-6 py-4 text-right">
                        <div className="flex items-center justify-end gap-2 text-gray-300">
                          <Share2 size={16} className="text-green-500" />
                          <span className="font-mono">{item.shares ? item.shares.toLocaleString() : '-'}</span>
                        </div>
                      </td>
                      <td className="px-6 py-4 text-center">
                        <div className="flex items-center justify-center gap-2">
                          <button
                            onClick={() => setEditingId(item.id)}
                            className="p-1.5 hover:bg-gray-800 rounded-lg transition-colors text-gray-400 hover:text-white"
                            title="Editar"
                          >
                            <Edit2 size={16} />
                          </button>
                          <button
                            onClick={() => handleDelete(item.id)}
                            disabled={isDeletingPost}
                            className="p-1.5 hover:bg-red-500/10 rounded-lg transition-colors text-gray-400 hover:text-red-400 disabled:opacity-50 disabled:cursor-not-allowed"
                            title="Deletar"
                          >
                            <Trash2 size={16} />
                          </button>
                        </div>
                      </td>
                    </tr>
                  );
                })
              )}
            </tbody>
          </table>
        </div>

        {/* Pagination */}
        {hasData && totalPages > 1 && (
          <div className="p-4 border-t border-gray-800 bg-[#15171c] flex items-center justify-between">
            <div className="text-xs text-gray-500">
              Mostrando {startIndex + 1} - {Math.min(endIndex, sortedData.length)} de {sortedData.length} posts
            </div>
            <div className="flex items-center gap-2">
              <button
                onClick={() => setCurrentPage(prev => Math.max(1, prev - 1))}
                disabled={currentPage === 1}
                className="p-2 hover:bg-gray-800 rounded-lg transition-colors text-gray-400 hover:text-white disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <ChevronLeft size={16} />
              </button>
              <span className="text-sm text-gray-400">
                Página {currentPage} de {totalPages}
              </span>
              <button
                onClick={() => setCurrentPage(prev => Math.min(totalPages, prev + 1))}
                disabled={currentPage === totalPages}
                className="p-2 hover:bg-gray-800 rounded-lg transition-colors text-gray-400 hover:text-white disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <ChevronRight size={16} />
              </button>
            </div>
          </div>
        )}

        <div className="p-4 border-t border-gray-800 bg-[#15171c] text-center text-xs text-gray-500">
          * Dados de posts individuais. As agregações diárias são calculadas automaticamente.
        </div>
      </div>
    </div>
  );
};

export default SocialPostsView;
