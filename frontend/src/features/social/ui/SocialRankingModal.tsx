import React from 'react';
import type { SocialBenchmarking } from '@/shared/types';
import { X, ThumbsUp, MessageCircle, BarChart2 } from 'lucide-react';
import type { SocialRankingModalProps } from '../types';

const SocialRankingModal: React.FC<SocialRankingModalProps> = ({ isOpen, onClose, data }) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={onClose}>
      <div 
        className="bg-[#1a1d24] rounded-2xl shadow-2xl border border-gray-700 w-full max-w-2xl overflow-hidden animate-in zoom-in-95 duration-200"
        onClick={e => e.stopPropagation()}
      >
        {/* Header */}
        <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
          <div className="flex items-center gap-3">
             <div className="p-2 bg-pink-500/10 rounded-lg border border-pink-500/20 text-pink-400">
                <BarChart2 size={24} />
             </div>
             <div>
                <h3 className="text-xl font-bold text-gray-100">Benchmarking Social</h3>
                <p className="text-sm text-gray-400">Comparativo de Engajamento vs Concorrentes</p>
             </div>
          </div>
          <button onClick={onClose} className="text-gray-500 hover:text-white p-2 hover:bg-gray-800 rounded-full transition-colors">
            <X size={20} />
          </button>
        </div>

        {/* Content */}
        <div className="p-6 bg-[#0f1115]">
            <div className="bg-[#1a1d24] rounded-xl border border-gray-800 overflow-hidden">
                <table className="w-full text-sm text-left">
                    <thead className="bg-[#20232b] text-xs text-gray-400 uppercase font-semibold">
                        <tr>
                            <th className="px-6 py-4">Ranking</th>
                            <th className="px-6 py-4">Marca</th>
                            <th className="px-6 py-4 text-right">Média Likes</th>
                            <th className="px-6 py-4 text-right">Média Comentários</th>
                            <th className="px-6 py-4 text-center">Seguidores</th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-800">
                        {data.map((item, index) => {
                            const isSolis = item.brand === 'Solis Solar';
                            return (
                                <tr 
                                    key={index} 
                                    className={`
                                        transition-colors 
                                        ${isSolis ? 'bg-purple-900/10 hover:bg-purple-900/20' : 'hover:bg-[#20232b]'}
                                    `}
                                >
                                    <td className="px-6 py-4 text-gray-500 font-mono">#{index + 1}</td>
                                    <td className="px-6 py-4">
                                        <span className={`font-semibold ${isSolis ? 'text-purple-400' : 'text-gray-300'}`}>
                                            {item.brand}
                                        </span>
                                        {isSolis && <span className="ml-2 px-2 py-0.5 rounded text-[10px] bg-purple-500 text-white font-bold uppercase">Você</span>}
                                    </td>
                                    <td className="px-6 py-4 text-right">
                                        <div className="flex items-center justify-end gap-2 text-gray-300">
                                            {item.avgLikes.toFixed(2)}
                                            <ThumbsUp size={14} className="text-gray-600" />
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-right">
                                        <div className="flex items-center justify-end gap-2 text-gray-300">
                                            {item.avgComments.toFixed(2)}
                                            <MessageCircle size={14} className="text-gray-600" />
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-center text-gray-600">
                                        -
                                    </td>
                                </tr>
                            );
                        })}
                    </tbody>
                </table>
            </div>
            
            <p className="text-xs text-gray-500 mt-4 text-center">
                * Dados coletados automaticamente das últimas 30 postagens. 
            </p>
        </div>
      </div>
    </div>
  );
};

export default SocialRankingModal;
