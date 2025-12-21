
import React from 'react';
import { Notification } from '../types';
import { Bell, Check, Clock, AtSign, Settings, Archive } from 'lucide-react';

interface NotificationPanelProps {
  notifications: Notification[];
  onMarkAsRead: (id: string) => void;
  onArchive: (id: string) => void;
  onClearAll: () => void;
  onNotificationClick: (notification: Notification) => void;
}

const NotificationPanel: React.FC<NotificationPanelProps> = ({ notifications, onMarkAsRead, onArchive, onClearAll, onNotificationClick }) => {
  
  // Filtrar apenas as NÃO arquivadas para exibir no painel
  const activeNotifications = notifications.filter(n => !n.archived);

  const getIcon = (type: string) => {
    switch (type) {
        case 'mention': return <AtSign size={16} className="text-blue-400" />;
        case 'deadline': return <Clock size={16} className="text-amber-400" />;
        default: return <Settings size={16} className="text-gray-400" />;
    }
  };

  return (
    <div className="absolute right-0 top-12 w-80 bg-[#1a1d24] border border-gray-700 rounded-xl shadow-2xl z-50 overflow-hidden animate-in fade-in slide-in-from-top-2">
      <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
        <h3 className="font-bold text-gray-100 flex items-center gap-2">
            <Bell size={16} /> Notificações
        </h3>
        {activeNotifications.some(n => !n.read) && (
            <button 
                onClick={(e) => { e.stopPropagation(); onClearAll(); }}
                className="text-xs text-emerald-400 hover:text-emerald-300 font-medium"
            >
                Marcar todas lidas
            </button>
        )}
      </div>
      
      <div className="max-h-[400px] overflow-y-auto custom-scrollbar">
        {activeNotifications.length === 0 ? (
            <div className="p-8 text-center text-gray-500">
                <p className="text-sm">Nenhuma notificação nova.</p>
            </div>
        ) : (
            <div className="divide-y divide-gray-800">
                {activeNotifications.map(notification => (
                    <div 
                        key={notification.id} 
                        onClick={() => onNotificationClick(notification)}
                        className={`p-4 hover:bg-[#20232b] transition-colors relative group cursor-pointer ${!notification.read ? 'bg-[#1e5144]/5' : ''}`}
                    >
                        <div className="flex gap-3">
                            <div className={`mt-1 w-8 h-8 rounded-full flex items-center justify-center border border-gray-700 bg-gray-800 flex-shrink-0`}>
                                {getIcon(notification.type)}
                            </div>
                            <div className="flex-grow min-w-0">
                                <div className="flex justify-between items-start mb-1">
                                    <h4 className={`text-sm font-semibold truncate pr-2 ${!notification.read ? 'text-white' : 'text-gray-400'}`}>
                                        {notification.title}
                                    </h4>
                                    <div className="flex items-center gap-2">
                                        {!notification.read ? (
                                            <span className="w-2 h-2 rounded-full bg-emerald-500 flex-shrink-0"></span>
                                        ) : (
                                            <button
                                                onClick={(e) => { e.stopPropagation(); onArchive(notification.id); }}
                                                className="text-gray-500 hover:text-white transition-colors"
                                                title="Arquivar"
                                            >
                                                <Archive size={14} />
                                            </button>
                                        )}
                                    </div>
                                </div>
                                <p className="text-xs text-gray-400 leading-relaxed mb-2">
                                    {notification.message}
                                </p>
                                <div className="flex justify-between items-center">
                                    <span className="text-[10px] text-gray-600">{notification.timestamp}</span>
                                    {!notification.read && (
                                        <button 
                                            onClick={(e) => { e.stopPropagation(); onMarkAsRead(notification.id); }}
                                            className="text-xs text-emerald-500 hover:text-emerald-400 flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity"
                                        >
                                            <Check size={12} /> Marcar lida
                                        </button>
                                    )}
                                </div>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        )}
      </div>
    </div>
  );
};

export default NotificationPanel;
