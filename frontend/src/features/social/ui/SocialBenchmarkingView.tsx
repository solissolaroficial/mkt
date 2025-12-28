import React from 'react';
import type { SocialBenchmarking } from '@/shared/types';
import { ThumbsUp, MessageCircle, BarChart2, TrendingUp, Users, ArrowLeft } from 'lucide-react';
import type { SocialBenchmarkingViewProps } from '../types';

const SocialBenchmarkingView: React.FC<SocialBenchmarkingViewProps> = ({ data, onBack }) => {
  const hasData = data && data.length > 0;
  // Fallback object to prevent crashes if data is empty
  const solisData = hasData ? data[0] : { avgLikes: 0, avgComments: 0, followers: 0 };
 
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
                <p className="text-xl font-bold text-white">{solisData.avgLikes.toFixed(2)}</p>
            </div>
        </div>
        <div className="bg-[#1a1d24] p-4 rounded-xl border border-white/5 shadow-lg flex items-center gap-4">
            <div className="p-3 bg-pink-500/10 rounded-lg text-pink-400 border border-pink-500/20">
                <MessageCircle size={24} />
            </div>
            <div>
                <p className="text-sm text-gray-500">Média de Comentários</p>
                <p className="text-xl font-bold text-white">{solisData.avgComments.toFixed(2)}</p>
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
                    </tr>
                </thead>
                <tbody className="divide-y divide-gray-800">
                    {!hasData ? (
                        <tr>
                            <td colSpan={6} className="px-6 py-8 text-center text-gray-500">
                                Nenhum dado disponível para exibição.
                            </td>
                        </tr>
                    ) : (
                        data.map((item, index) => {
                            const isSolis = item.brand === 'Solis Solar';
                            
                            // New Calculation: ((Likes + Comments) / Followers) * 100
                            const totalInteractions = item.avgLikes + item.avgComments;
                            const engagementRate = item.followers && item.followers > 0 
                                ? (totalInteractions / item.followers) * 100 
                                : 0;
                                
                            // Prevent division by zero if solis avg likes is 0
                            const relativePerformance = solisData.avgLikes > 0 
                                ? (item.avgLikes / solisData.avgLikes) * 100 
                                : 0; 
                            
                            return (
                                <tr 
                                    key={index} 
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
                                                {item.brand}
                                            </span>
                                            {isSolis && <span className="px-2 py-0.5 rounded text-[10px] bg-[#1e5144] text-white font-bold uppercase shadow-lg shadow-[#1e5144]/20">Você</span>}
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-right">
                                        <div className="flex items-center justify-end gap-2 text-gray-300">
                                            <span className="font-mono text-lg">{item.avgLikes.toFixed(2)}</span>
                                            <ThumbsUp size={16} className={isSolis ? 'text-[#1e5144]' : 'text-gray-600'} />
                                        </div>
                                        <div className="w-full bg-gray-800 h-1 mt-2 rounded-full overflow-hidden flex justify-end">
                                            <div className="bg-gray-600 h-full rounded-full" style={{ width: `${Math.min(relativePerformance, 100)}%` }}></div>
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-right">
                                        <div className="flex items-center justify-end gap-2 text-gray-300">
                                            <span className="font-mono text-lg">{item.avgComments.toFixed(2)}</span>
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
                                </tr>
                            );
                        })
                    )}
                </tbody>
            </table>
        </div>
        
        <div className="p-4 border-t border-gray-800 bg-[#15171c] text-center text-xs text-gray-500">
            * Dados coletados automaticamente das últimas 30 postagens públicas no Instagram. Atualizado em 21/11/2025.
        </div>
      </div>
    </div>
  );
};

export default SocialBenchmarkingView;
