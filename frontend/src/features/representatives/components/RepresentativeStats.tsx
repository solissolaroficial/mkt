import React from 'react';
import { ArrowLeft, Activity, Package, Megaphone, TrendingUp } from 'lucide-react';
import type { RepresentativeStats, Representative } from '../types';

interface RepresentativeStatsProps {
  stats: RepresentativeStats;
  representative?: Representative;
  isLoading: boolean;
  error: unknown;
  onRefetch: () => void;
  onBack: () => void;
}

const RepresentativeStats: React.FC<RepresentativeStatsProps> = ({
  stats,
  representative,
  isLoading,
  error,
  onRefetch,
  onBack,
}) => {
  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-gray-400">Carregando estatísticas...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col items-center justify-center h-full gap-4">
        <div className="text-rose-400">Erro ao carregar estatísticas</div>
        <button
          onClick={onRefetch}
          className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg"
        >
          Tentar novamente
        </button>
      </div>
    );
  }

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header */}
      <div className="flex items-center gap-4 mb-6">
        <button onClick={onBack} className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white">
          <ArrowLeft size={24} />
        </button>
        <div>
          <h1 className="text-2xl font-bold text-gray-100">
            Estatísticas de {representative?.name || 'Representante'}
          </h1>
          <p className="text-gray-500 text-sm">
            Código: {representative?.code || '-'} | Empresa: {representative?.company || '-'}
          </p>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {/* Total Actions */}
        <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <div className="p-3 bg-[#1e5144]/20 rounded-xl">
              <Activity className="text-[#1e5144]" size={24} />
            </div>
            <span className="text-3xl font-bold text-gray-100">
              {stats.totalActions}
            </span>
          </div>
          <div>
            <h3 className="text-lg font-semibold text-gray-300 mb-1">
              Total de Ações
            </h3>
            <p className="text-sm text-gray-500">
              Todas as ações registradas
            </p>
          </div>
        </div>

        {/* Offline Actions */}
        <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <div className="p-3 bg-blue-500/20 rounded-xl">
              <Package className="text-blue-400" size={24} />
            </div>
            <span className="text-3xl font-bold text-gray-100">
              {stats.offlineActionCount}
            </span>
          </div>
          <div>
            <h3 className="text-lg font-semibold text-gray-300 mb-1">
              Ações Offline
            </h3>
            <p className="text-sm text-gray-500">
              Ações presenciais realizadas
            </p>
          </div>
        </div>

        {/* Showroom Items */}
        <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <div className="p-3 bg-amber-500/20 rounded-xl">
              <TrendingUp className="text-amber-400" size={24} />
            </div>
            <span className="text-3xl font-bold text-gray-100">
              {stats.showroomItemCount}
            </span>
          </div>
          <div>
            <h3 className="text-lg font-semibold text-gray-300 mb-1">
              Itens de Showroom
            </h3>
            <p className="text-sm text-gray-500">
              Produtos em showroom
            </p>
          </div>
        </div>

        {/* Rep Marketing Actions */}
        <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <div className="p-3 bg-purple-500/20 rounded-xl">
              <Megaphone className="text-purple-400" size={24} />
            </div>
            <span className="text-3xl font-bold text-gray-100">
              {stats.repMarketingCount}
            </span>
          </div>
          <div>
            <h3 className="text-lg font-semibold text-gray-300 mb-1">
              Ações de Marketing
            </h3>
            <p className="text-sm text-gray-500">
              Ações de marketing realizadas
            </p>
          </div>
        </div>
      </div>

      {/* Representative Info */}
      {representative && (
        <div className="mt-6 bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl p-6">
          <h2 className="text-xl font-bold text-gray-100 mb-4">
            Informações do Representante
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <p className="text-sm text-gray-500 mb-1">Nome</p>
              <p className="text-base font-medium text-gray-200">{representative.name}</p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Email</p>
              <p className="text-base font-medium text-gray-200">{representative.email}</p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Empresa</p>
              <p className="text-base font-medium text-gray-200">{representative.company}</p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Região</p>
              <p className="text-base font-medium text-gray-200">{representative.region}</p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Cidade</p>
              <p className="text-base font-medium text-gray-200">{representative.city}</p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Status</p>
              <span className={`inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium ${
                representative.active
                  ? 'bg-emerald-500/10 text-emerald-400 border border-emerald-500/20'
                  : 'bg-rose-500/10 text-rose-400 border border-rose-500/20'
              }`}>
                {representative.active ? 'Ativo' : 'Inativo'}
              </span>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default RepresentativeStats;
