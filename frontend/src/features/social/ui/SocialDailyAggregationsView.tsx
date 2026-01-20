import React, { useState, useEffect } from 'react';
import type { SocialDailyAggregation } from '@/shared/types';
import type { Brand } from '@/shared/types/legacy.types';
import { TrendingUp, ThumbsUp, MessageCircle, Share2, Calendar, Users, Filter, RefreshCw, ChevronLeft, ChevronRight } from 'lucide-react';
import type { SocialPlatform } from '@/shared/types';
import { useSocialMutations } from '../hooks/useSocialMutations';
import { useBrands } from '../hooks/useBrands';
import { useToast } from '@/shared/hooks/useToast';

interface SocialDailyAggregationsViewProps {
  data?: SocialDailyAggregation[];
  onRefresh?: () => void;
}

const SocialDailyAggregationsView: React.FC<SocialDailyAggregationsViewProps> = ({ data, onRefresh }) => {
  const { error: showError } = useToast();

  // Filter States
  const [sortBy, setSortBy] = useState<'aggregation_date' | 'total_posts' | 'total_likes' | 'total_comments' | 'engagement_rate' | 'created_at'>('aggregation_date');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');
  const [filterBrand, setFilterBrand] = useState('');

  // Pagination States
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(10);

  // Mutations
  const { recalculateAggregations, isRecalculatingAggregations } = useSocialMutations();

  // Brands Query
  const { data: brands = [] } = useBrands();

  // Helper function to get brand name by ID
  const getBrandName = (brandId: string): string => {
    const brand = brands.find(b => b.id === brandId);
    return brand?.name || brandId;
  };

  // Aplicar filtros localmente
  const filteredData = data ? data.filter(item => {
    const brandName = getBrandName(item.brand_id);
    const matchBrand = filterBrand === '' || brandName.toLowerCase().includes(filterBrand.toLowerCase());
    return matchBrand;
  }) : [];

  // Ordenar localmente
  const sortedData = [...filteredData].sort((a, b) => {
    let comparison = 0;

    switch (sortBy) {
      case 'aggregation_date':
        comparison = new Date(a.aggregation_date).getTime() - new Date(b.aggregation_date).getTime();
        break;
      case 'total_posts':
        comparison = a.total_posts - b.total_posts;
        break;
      case 'total_likes':
        comparison = a.total_likes - b.total_likes;
        break;
      case 'total_comments':
        comparison = a.total_comments - b.total_comments;
        break;
      case 'engagement_rate':
        comparison = a.engagement_rate - b.engagement_rate;
        break;
      case 'created_at':
        comparison = new Date(a.created_at).getTime() - new Date(b.created_at).getTime();
        break;
    }

    return sortOrder === 'asc' ? comparison : -comparison;
  });

  // Verificar se há dados
  const hasData = sortedData && sortedData.length > 0;

  // Pagination
  const totalPages = Math.ceil(sortedData.length / itemsPerPage);
  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  const paginatedData = sortedData.slice(startIndex, endIndex);

  // Reset pagination when filters change
  useEffect(() => {
    setCurrentPage(1);
  }, [filterBrand, sortBy, sortOrder]);

  // Calcular estatísticas globais (baseado em todos os dados filtrados, não apenas na página atual)
  const globalStats = hasData ? {
    totalPosts: sortedData.reduce((sum, item) => sum + item.total_posts, 0),
    totalLikes: sortedData.reduce((sum, item) => sum + item.total_likes, 0),
    totalComments: sortedData.reduce((sum, item) => sum + item.total_comments, 0),
    avgEngagementRate: sortedData.reduce((sum, item) => sum + item.engagement_rate, 0) / sortedData.length,
  } : {
    totalPosts: 0,
    totalLikes: 0,
    totalComments: 0,
    avgEngagementRate: 0,
  };

  const handleRecalculate = (brandID: string, date: string) => {
    const brandName = getBrandName(brandID);
    if (!confirm(`Deseja recalcular as agregações para ${brandName} em ${date}?`)) return;

    recalculateAggregations(brandID, date, {
      onSuccess: () => {
        onRefresh?.();
      },
      onError: (error) => {
        console.error('Erro ao recalcular agregações:', error);
        showError('Erro ao recalcular agregações. Tente novamente.');
      }
    });
  };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col text-gray-200">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-100">Agregações Diárias</h1>
          <p className="text-gray-500">Dados agregados calculados automaticamente dos posts individuais</p>
        </div>
      </div>

      {/* Unified Control Bar */}
      <div className="flex flex-col md:flex-row justify-between items-center mb-6 gap-4">
        {/* Filters (Left Side) */}
        <div className="flex flex-wrap gap-4 w-full md:w-auto items-center">
          <div className="relative">
            <select
              value={sortBy}
              onChange={(e) => setSortBy(e.target.value as 'aggregation_date' | 'total_posts' | 'total_likes' | 'total_comments' | 'engagement_rate' | 'created_at')}
              className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
            >
              <option value="aggregation_date">Data de Agregação</option>
              <option value="total_posts">Total de Posts</option>
              <option value="total_likes">Total de Likes</option>
              <option value="total_comments">Total de Comentários</option>
              <option value="engagement_rate">Taxa de Engajamento</option>
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
        </div>

        {/* Actions (Right Side) */}
        <div className="flex items-center gap-3 w-full md:w-auto justify-end">
          <button
            onClick={onRefresh}
            className="bg-[#1a1d24] text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-[#20232b] transition-colors flex items-center gap-2 border border-gray-700"
          >
            <RefreshCw size={18} /> Atualizar
          </button>
        </div>
      </div>

      {/* Stats Cards Row */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-purple-500/10 rounded-lg text-purple-400 border border-purple-500/20">
            <TrendingUp size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Total de Posts</p>
            <p className="text-xl font-bold text-white">{globalStats.totalPosts.toLocaleString()}</p>
          </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-blue-500/10 rounded-lg text-blue-400 border border-blue-500/20">
            <ThumbsUp size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Total de Likes</p>
            <p className="text-xl font-bold text-white">{globalStats.totalLikes.toLocaleString()}</p>
          </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-pink-500/10 rounded-lg text-pink-400 border border-pink-500/20">
            <MessageCircle size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Total de Comentários</p>
            <p className="text-xl font-bold text-white">{globalStats.totalComments.toLocaleString()}</p>
          </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
          <div className="p-3 bg-emerald-500/10 rounded-lg text-emerald-400 border border-emerald-500/20">
            <TrendingUp size={24} />
          </div>
          <div>
            <p className="text-sm text-gray-500">Média de Engajamento</p>
            <p className="text-xl font-bold text-white">{globalStats.avgEngagementRate.toFixed(2)}%</p>
          </div>
        </div>
      </div>

      {/* Main Content Table */}
      <div className="bg-[#1a1d24] rounded-xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col">
        <div className="p-4 border-b border-gray-800 bg-[#20232b]">
          <h2 className="font-semibold text-gray-100">Lista de Agregações Diárias</h2>
        </div>

        <div className="overflow-x-auto custom-scrollbar flex-grow">
          <table className="w-full text-sm text-left">
            <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
              <tr>
                <th className="px-6 py-4">Marca</th>
                <th className="px-6 py-4">Data</th>
                <th className="px-6 py-4 text-right">Total de Posts</th>
                <th className="px-6 py-4 text-right">Total de Likes</th>
                <th className="px-6 py-4 text-right">Média de Likes</th>
                <th className="px-6 py-4 text-right">Total de Comentários</th>
                <th className="px-6 py-4 text-right">Média de Comentários</th>
                <th className="px-6 py-4 text-right">Seguidores</th>
                <th className="px-6 py-4 text-center">Engajamento</th>
                <th className="px-6 py-4 text-center">Ações</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-800">
              {!hasData ? (
                <tr>
                  <td colSpan={10} className="px-6 py-8 text-center text-gray-500">
                    Nenhuma agregação disponível para exibição.
                  </td>
                </tr>
              ) : (
                paginatedData.map((item) => (
                  <tr
                    key={item.id}
                    className="hover:bg-[#20232b] transition-colors"
                  >
                    <td className="px-6 py-4">
                      <span className="font-medium text-gray-300">{getBrandName(item.brand_id)}</span>
                    </td>
                    <td className="px-6 py-4">
                      <div className="flex items-center gap-2 text-gray-400">
                        <Calendar size={14} className="text-gray-600" />
                        <span className="font-mono">{new Date(item.aggregation_date).toLocaleDateString('pt-BR')}</span>
                      </div>
                    </td>
                    <td className="px-6 py-4 text-right">
                      <span className="font-mono text-gray-300">{item.total_posts.toLocaleString()}</span>
                    </td>
                    <td className="px-6 py-4 text-right">
                      <div className="flex items-center justify-end gap-2 text-gray-300">
                        <ThumbsUp size={16} className="text-blue-500" />
                        <span className="font-mono">{item.total_likes.toLocaleString()}</span>
                      </div>
                    </td>
                    <td className="px-6 py-4 text-right">
                      <span className="font-mono text-gray-400">{item.avg_likes.toFixed(2)}</span>
                    </td>
                    <td className="px-6 py-4 text-right">
                      <div className="flex items-center justify-end gap-2 text-gray-300">
                        <MessageCircle size={16} className="text-pink-500" />
                        <span className="font-mono">{item.total_comments.toLocaleString()}</span>
                      </div>
                    </td>
                    <td className="px-6 py-4 text-right">
                      <span className="font-mono text-gray-400">{item.avg_comments.toFixed(2)}</span>
                    </td>
                    <td className="px-6 py-4 text-right">
                      {item.followers_at_date ? (
                        <div className="flex items-center justify-end gap-2 text-gray-400">
                          <Users size={14} className="text-gray-600" />
                          <span className="font-mono">{item.followers_at_date.toLocaleString()}</span>
                        </div>
                      ) : '-'}
                    </td>
                    <td className="px-6 py-4 text-center">
                      <span className="font-mono font-medium text-emerald-400">
                        {item.engagement_rate.toFixed(2)}%
                      </span>
                    </td>
                    <td className="px-6 py-4 text-center">
                      <button
                        onClick={() => handleRecalculate(item.brand_id, item.aggregation_date)}
                        disabled={isRecalculatingAggregations}
                        className="p-1.5 hover:bg-emerald-500/10 rounded-lg transition-colors text-gray-400 hover:text-emerald-400 disabled:opacity-50 disabled:cursor-not-allowed"
                        title="Recalcular Agregações"
                      >
                        <RefreshCw size={16} />
                      </button>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>

        {/* Pagination */}
        {hasData && totalPages > 1 && (
          <div className="p-4 border-t border-gray-800 bg-[#15171c] flex items-center justify-between">
            <div className="text-xs text-gray-500">
              Mostrando {startIndex + 1} - {Math.min(endIndex, sortedData.length)} de {sortedData.length} agregações
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
          * As agregações são calculadas automaticamente a partir dos posts individuais. Use o botão de recalcular para atualizar manualmente.
        </div>
      </div>
    </div>
  );
};

export default SocialDailyAggregationsView;
