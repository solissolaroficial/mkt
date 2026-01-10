import React, { useState, useEffect } from 'react';
import type { SocialBenchmarking } from '@/shared/types';
import { ThumbsUp, MessageCircle, BarChart2, TrendingUp, Users, ArrowLeft, Plus, Filter, Save, Edit2, Trash2 } from 'lucide-react';
import type { SocialBenchmarkingViewProps } from '../types';
import { useSocialMutations } from '../hooks/useSocialMutations';

const SocialBenchmarkingView: React.FC<SocialBenchmarkingViewProps> = ({ data, onBack }) => {
  // Filter States
  const [sortBy, setSortBy] = useState<'engagement_rate' | 'avg_likes' | 'avg_comments' | 'created_at'>('engagement_rate');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');
  const [searchBrand, setSearchBrand] = useState('');

  // Form States
  const [isAdding, setIsAdding] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [brandName, setBrandName] = useState('');
  const [avgLikes, setAvgLikes] = useState(0);
  const [avgComments, setAvgComments] = useState(0);
  const [followers, setFollowers] = useState<number | undefined>(undefined);

  // Mutations
  const { create, update, delete: deleteBenchmarking, isCreating, isUpdating, isDeleting } = useSocialMutations();

  // Reset form states
  const resetForm = () => {
    setBrandName('');
    setAvgLikes(0);
    setAvgComments(0);
    setFollowers(undefined);
  };

  // Preencher formulário no modo edição
  useEffect(() => {
    if (editingId && data) {
      const item = data.find(item => item.id === editingId);
      if (item) {
        setBrandName(item.brand_name);
        setAvgLikes(item.avg_likes);
        setAvgComments(item.avg_comments);
        setFollowers(item.followers);
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

    create({
      brand_name: brandName,
      avg_likes: avgLikes,
      avg_comments: avgComments,
      ...(followers !== undefined && { followers }),
    }, {
      onSuccess: () => {
        setIsAdding(false);
        resetForm();
      },
      onError: (error) => {
        console.error('Erro ao criar benchmarking:', error);
        alert('Erro ao criar benchmarking. Tente novamente.');
      }
    });
  };

  const handleUpdate = (e: React.FormEvent) => {
    e.preventDefault();

    if (!editingId) return;

    update({
      id: editingId,
      data: {
        brand_name: brandName,
        avg_likes: avgLikes,
        avg_comments: avgComments,
        ...(followers !== undefined && { followers }),
      }
    }, {
      onSuccess: () => {
        setEditingId(null);
        resetForm();
      },
      onError: (error) => {
        console.error('Erro ao atualizar benchmarking:', error);
        alert('Erro ao atualizar benchmarking. Tente novamente.');
      }
    });
  };

  const handleDelete = (id: string) => {
    if (!confirm('Tem certeza que deseja deletar este benchmarking?')) return;

    deleteBenchmarking(id, {
      onError: (error) => {
        console.error('Erro ao deletar benchmarking:', error);
        alert('Erro ao deletar benchmarking. Tente novamente.');
      }
    });
  };

  // Aplicar filtros localmente
  const filteredData = data ? data.filter(item => {
    const matchBrand = searchBrand === '' || item.brand_name.toLowerCase().includes(searchBrand.toLowerCase());
    return matchBrand;
  }) : [];

  // Ordenar localmente
  const sortedData = [...filteredData].sort((a, b) => {
    let comparison = 0;

    switch (sortBy) {
      case 'engagement_rate':
        const engagementA = a.followers && a.followers > 0 ? ((a.avg_likes + a.avg_comments) / a.followers) * 100 : 0;
        const engagementB = b.followers && b.followers > 0 ? ((b.avg_likes + b.avg_comments) / b.followers) * 100 : 0;
        comparison = engagementA - engagementB;
        break;
      case 'avg_likes':
        comparison = a.avg_likes - b.avg_likes;
        break;
      case 'avg_comments':
        comparison = a.avg_comments - b.avg_comments;
        break;
      case 'created_at':
        comparison = new Date(a.created_at).getTime() - new Date(b.created_at).getTime();
        break;
    }

    return sortOrder === 'asc' ? comparison : -comparison;
  });

  // Verificar se há dados
  const hasData = sortedData && sortedData.length > 0;

  // Encontrar dados da Solis Solar
  const solisData = hasData ? (sortedData.find(item => item.brand_name === 'Solis Solar') || sortedData[0]) : { avg_likes: 0, avg_comments: 0, followers: 0 };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col text-gray-200">

      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div className="flex items-center gap-4">
          {onBack && (
            <button 
              onClick={onBack}
              className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
              title="Voltar"
            >
              <ArrowLeft size={24} />
            </button>
          )}
          <div>
            <h1 className="text-2xl font-bold text-gray-100">Benchmarking Social</h1>
            <p className="text-gray-500">Comparativo estratégico de engajamento vs concorrentes</p>
          </div>
        </div>
      </div> 

      {/* Unified Control Bar */}
      <div className="flex flex-col md:flex-row justify-between items-center mb-6 gap-4">

        {/* Filters (Left Side) */}
        <div className="flex flex-wrap gap-4 w-full md:w-auto items-center">
          <div className="relative">
            <select
              value={sortBy}
              onChange={(e) => setSortBy(e.target.value as any)}
              className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
            >
              <option value="engagement_rate">Engajamento</option>
              <option value="avg_likes">Likes</option>
              <option value="avg_comments">Comentários</option>
              <option value="created_at">Data</option>
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
              value={searchBrand}
              onChange={(e) => setSearchBrand(e.target.value)}
              placeholder="Buscar marca..."
              className="bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-4 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144] placeholder-gray-600"
            />
            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
          </div>
        </div>

        {/* Actions (Right Side) */}
        <div className="flex items-center gap-3 w-full md:w-auto justify-end">
          <button
            onClick={() => setIsAdding(!isAdding)}
            className="bg-[#1e5144] text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-[#163c32] transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/30"
          >
            <Plus size={18} /> Novo Benchmarking
          </button>
        </div>
      </div>

      {/* Add/Edit Form Panel */}
      {(isAdding || editingId) && (
        <div className="bg-[#1a1d24] rounded-xl border border-gray-700 shadow-xl mb-6 animate-in slide-in-from-top-4 flex flex-col">
          <div className="p-4 border-b border-gray-800 bg-[#20232b]">
            <h3 className="font-bold text-gray-100">
              {isAdding ? 'Novo Benchmarking' : 'Editar Benchmarking'}
            </h3>
          </div>

          <form onSubmit={isAdding ? handleCreate : handleUpdate} className="p-6 space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-400 mb-1">
                Nome da Marca <span className="text-rose-500">*</span>
              </label>
              <input
                type="text"
                value={brandName}
                onChange={(e) => setBrandName(e.target.value)}
                placeholder="Ex: Solis Solar"
                className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                required
                maxLength={200}
              />
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Média de Likes <span className="text-rose-500">*</span>
                </label>
                <input
                  type="number"
                  step="0.01"
                  min="0"
                  value={avgLikes}
                  onChange={(e) => setAvgLikes(parseFloat(e.target.value) || 0)}
                  placeholder="0.00"
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Média de Comentários <span className="text-rose-500">*</span>
                </label>
                <input
                  type="number"
                  step="0.01"
                  min="0"
                  value={avgComments}
                  onChange={(e) => setAvgComments(parseFloat(e.target.value) || 0)}
                  placeholder="0.00"
                  className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                  required
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-400 mb-1">
                Seguidores <span className="text-gray-500 font-normal ml-1">(Opcional)</span>
              </label>
              <input
                type="number"
                min="0"
                value={followers || ''}
                onChange={(e) => setFollowers(e.target.value ? parseInt(e.target.value) : undefined)}
                placeholder="Número de seguidores"
                className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
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
                disabled={isCreating || isUpdating}
                className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isCreating || isUpdating ? (
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
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
            <div className="p-3 bg-[#1e5144]/10 rounded-lg text-emerald-400 border border-[#1e5144]/20">
                <TrendingUp size={24} />
            </div>
            <div>
                <p className="text-sm text-gray-500">Sua Posição</p>
                <p className="text-xl font-bold text-white">{hasData ? '1º Lugar' : '-'}</p>
            </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
            <div className="p-3 bg-blue-500/10 rounded-lg text-blue-400 border border-blue-500/20">
                <ThumbsUp size={24} />
            </div>
            <div>
                <p className="text-sm text-gray-500">Média de Likes (Solis)</p>
                <p className="text-xl font-bold text-white">{solisData.avg_likes.toFixed(2)}</p>
            </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
            <div className="p-3 bg-pink-500/10 rounded-lg text-pink-400 border border-pink-500/20">
                <MessageCircle size={24} />
            </div>
            <div>
                <p className="text-sm text-gray-500">Média de Comentários</p>
                <p className="text-xl font-bold text-white">{solisData.avg_comments.toFixed(2)}</p>
            </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
            <div className="p-3 bg-cyan-500/10 rounded-lg text-cyan-400 border border-cyan-500/20">
                <Users size={24} />
            </div>
            <div>
                <p className="text-sm text-gray-500">Seu Público</p>
                <p className="text-xl font-bold text-white">
                    {solisData.followers ? (solisData.followers / 1000).toFixed(1) + 'k' : '-'}
                </p>
            </div>
        </div>
      </div> 

      {/* Main Content Table */}
      <div className="bg-[#1a1d24] rounded-xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col">
        <div className="p-4 border-b border-gray-800 bg-[#20232b]">
            <h2 className="font-semibold text-gray-100 flex items-center gap-2">
                <BarChart2 size={18} className="text-gray-400" />
                Ranking de Performance
            </h2>
        </div> 

        <div className="overflow-x-auto custom-scrollbar flex-grow">
            <table className="w-full text-sm text-left">
                <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
                    <tr>
                        <th className="px-6 py-4 w-24 text-center">Posição</th>
                        <th className="px-6 py-4">Marca</th>
                        <th className="px-6 py-4 text-right">Média de Likes</th>
                        <th className="px-6 py-4 text-right">Média de Comentários</th>
                        <th className="px-6 py-4 text-center">Seguidores</th>
                        <th className="px-6 py-4 text-center">Taxa de Engajamento</th>
                        <th className="px-6 py-4 text-center">Ações</th>
                    </tr>
                </thead>
                <tbody className="divide-y divide-gray-800">
                    {!hasData ? (
                        <tr>
                            <td colSpan={7} className="px-6 py-8 text-center text-gray-500">
                                Nenhum dado disponível para exibição.
                            </td>
                        </tr>
                    ) : (
                        sortedData.map((item, index) => {
                            const isSolis = item.brand_name === 'Solis Solar';

                            // Calculate engagement rate once
                            const totalInteractions = item.avg_likes + item.avg_comments;
                            const engagementRate = item.followers && item.followers > 0
                                ? (totalInteractions / item.followers) * 100
                                : 0;

                            // Calculate relative performance
                            const relativePerformance = solisData && solisData.avg_likes > 0
                                ? (item.avg_likes / solisData.avg_likes) * 100
                                : 0;

                            return (
                                <tr
                                    key={item.id}
                                    className={`
                                        transition-colors
                                        ${isSolis ? 'bg-[#1e5144]/10 hover:bg-[#1e5144]/20' : 'hover:bg-[#20232b]'}
                                    `}
                                >
                                    <td className="px-6 py-4 text-center">
                                        <span className={`
                                            inline-flex items-center justify-center w-8 h-8 rounded-full font-bold
                                            ${index === 0 ? 'bg-yellow-500/20 text-yellow-400 border border-yellow-500/30' : 
                                            index === 1 ? 'bg-gray-400/20 text-gray-300 border border-gray-400/30' : 
                                            index === 2 ? 'bg-orange-700/20 text-orange-400 border border-orange-700/30' : 
                                            'text-gray-600'}
                                            `}>
                                            #{index + 1}
                                        </span>
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="flex items-center gap-2">
                                            <span className={`font-bold text-lg ${isSolis ? 'text-emerald-400' : 'text-gray-300'}`}>
                                                {item.brand_name}
                                            </span>
                                            {isSolis && <span className="px-2 py-0.5 rounded text-[10px] bg-[#1e5144] text-white font-bold uppercase shadow-lg shadow-[#1e5144]/20">Você</span>}
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-right">
                                        <div className="flex items-center justify-end gap-2 text-gray-300">
                                            <span className="font-mono text-lg">{item.avg_likes.toFixed(2)}</span>
                                            <ThumbsUp size={16} className={isSolis ? 'text-[#1e5144]' : 'text-gray-600'} />
                                        </div>
                                        <div className="w-full bg-gray-800 h-1 mt-2 rounded-full overflow-hidden flex justify-end">
                                            <div className="bg-gray-600 h-full rounded-full" style={{ width: `${Math.min(relativePerformance, 100)}%` }}></div>
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-right">
                                        <div className="flex items-center justify-end gap-2 text-gray-300">
                                            <span className="font-mono text-lg">{item.avg_comments.toFixed(2)}</span>
                                            <MessageCircle size={16} className={isSolis ? 'text-[#1e5144]' : 'text-gray-600'} />
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-center text-gray-400">
                                        {item.followers ? (
                                            <div className="flex items-center justify-center gap-1 font-mono">
                                                <Users size={14} className="text-gray-600" />
                                                {item.followers.toLocaleString()}
                                            </div>
                                        ) : '-'}
                                    </td>
                                    <td className="px-6 py-4 text-center">
                                        <span className={`font-mono font-medium ${isSolis ? 'text-emerald-400' : 'text-gray-500'}`}>
                                            {engagementRate.toFixed(2)}%
                                        </span>
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
                                                disabled={isDeleting}
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

        <div className="p-4 border-t border-gray-800 bg-[#15171c] text-center text-xs text-gray-500">
            * Dados gerenciados pelo administrador. Atualizado em tempo real.
        </div>
      </div>
    </div>
  );
};

export default SocialBenchmarkingView;
