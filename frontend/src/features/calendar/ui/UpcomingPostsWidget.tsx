import React from 'react';
import { Calendar, Clock, Instagram, Youtube, Facebook, Linkedin, Star } from 'lucide-react';
import type { CalendarPost } from '../../../shared/types/legacy.types';
import type { UpcomingPostsWidgetProps } from '../types';

const UpcomingPostsWidget: React.FC<UpcomingPostsWidgetProps> = ({ posts, onViewCalendar }) => {
  const today = new Date();
  const nextWeek = new Date();
  nextWeek.setDate(today.getDate() + 7);

  const todayStr = today.toISOString().split('T')[0];
  const nextWeekStr = nextWeek.toISOString().split('T')[0];

  // Filter posts:
  // 1. Date between today and today + 7 days
  // 2. Status not published (pending, approved, adjust, review, in_progress)
  const upcomingPosts = posts
    .filter(p => {
        const isWithinWeek = p.date >= todayStr && p.date <= nextWeekStr;
        const isNotPublished = p.status !== 'published';
        return isWithinWeek && isNotPublished;
    })
    .sort((a, b) => {
        // Sort by date and then by time
        if (a.date !== b.date) return a.date.localeCompare(b.date);
        return a.time.localeCompare(b.time);
    });

  const getPlatformIcon = (platforms?: string[]) => {
      if (!platforms || platforms.length === 0) return <Star size={14} className="text-gray-400" />;
      // Return icon of first platform found for summary display
      if (platforms.includes('instagram')) return <Instagram size={14} className="text-pink-400" />;
      if (platforms.includes('youtube')) return <Youtube size={14} className="text-red-400" />;
      if (platforms.includes('linkedin')) return <Linkedin size={14} className="text-blue-400" />;
      if (platforms.includes('facebook')) return <Facebook size={14} className="text-blue-500" />;
      return <Star size={14} className="text-gray-400" />;
  };

  const getCategoryLabel = (cat: string) => {
      switch(cat) {
          case 'official': return 'Oficial';
          case 'solis_voce': return 'Solis e Você';
          case 'leonardo': return 'Leonardo';
          case 'luiz': return 'Luiz';
          default: return cat;
      }
  };

  const getStatusColor = (status: string) => {
      switch(status) {
          case 'approved': return 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20';
          case 'review': return 'bg-amber-500/10 text-amber-400 border-amber-500/20';
          case 'adjust': return 'bg-rose-500/10 text-rose-400 border-rose-500/20';
          default: return 'bg-blue-500/10 text-blue-400 border-blue-500/20';
      }
  };

  const getStatusLabel = (status: string) => {
      switch(status) {
          case 'approved': return 'Pronto';
          case 'review': return 'Revisão';
          case 'adjust': return 'Ajuste';
          case 'in_progress': return 'Em andamento';
          default: return status;
      }
  };

  return (
    <div className="bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 p-6 flex flex-col h-full">
      <div className="flex justify-between items-center mb-4">
        <div>
          <h3 className="text-lg font-semibold text-gray-100">Postagens a ser feitas</h3>
          <p className="text-sm text-gray-500">Programação da semana</p>
        </div>
        <button 
          onClick={onViewCalendar}
          className="text-xs text-gray-400 hover:text-white border border-gray-700 hover:bg-gray-800 px-3 py-1 rounded-lg transition-colors"
        >
          Ver Calendário
        </button>
      </div>

      <div className="flex-grow space-y-3 overflow-y-auto custom-scrollbar max-h-[300px] pr-1">
        {upcomingPosts.length === 0 ? (
            <div className="text-center py-8 text-gray-600">
                <Calendar size={32} className="mx-auto mb-2 opacity-30" />
                <p className="text-sm">Nenhuma postagem pendente para os próximos 7 dias.</p>
            </div>
        ) : (
            upcomingPosts.map(post => {
                const isToday = post.date === todayStr;
                return (
                    <div key={post.id} className="flex items-center gap-3 p-3 bg-[#20232b]/50 border border-gray-800 rounded-xl hover:border-gray-700 transition-colors group">
                        {/* Date Box */}
                        <div className={`flex flex-col items-center justify-center w-12 h-12 rounded-lg border flex-shrink-0 ${isToday ? 'bg-[#1e5144]/10 border-[#1e5144]/30 text-emerald-400' : 'bg-gray-800/50 border-gray-700 text-gray-400'}`}>
                            <span className="text-[10px] uppercase font-bold">{new Date(post.date).toLocaleDateString('pt-BR', { weekday: 'short' }).slice(0, 3)}</span>
                            <span className="text-lg font-bold leading-none">{post.date.split('-')[2]}</span>
                        </div>

                        {/* Content */}
                        <div className="flex-grow min-w-0">
                            <div className="flex justify-between items-start mb-1">
                                <h4 className="text-sm font-medium text-gray-200 truncate pr-2 group-hover:text-white transition-colors">{post.title}</h4>
                                <span className={`text-[10px] px-1.5 py-0.5 rounded border font-bold uppercase tracking-wider ${getStatusColor(post.status)}`}>
                                    {getStatusLabel(post.status)}
                                </span>
                            </div>
                            
                            <div className="flex items-center gap-3 text-xs text-gray-500">
                                <div className="flex items-center gap-1">
                                    <Clock size={12} />
                                    {post.time}
                                </div>
                                <div className="flex items-center gap-1">
                                    {getPlatformIcon(post.platforms)}
                                    <span className="capitalize">{getCategoryLabel(post.category)}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                );
            })
        )}
      </div>
    </div>
  );
};

export default UpcomingPostsWidget;
