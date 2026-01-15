import React from 'react';
import type { RepresentativeProfile } from '../types';
import { X, Mail, Phone, MapPin, User, Building, Hash, GraduationCap, Megaphone, DollarSign, Globe, FolderOpen } from 'lucide-react';
import type { RepProfileModalProps } from '../types';

const RepProfileModal: React.FC<RepProfileModalProps> = ({ isOpen, onClose, profile }) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={onClose}>
      <div 
        className="bg-[#1a1d24] w-full max-w-2xl rounded-2xl shadow-2xl border border-gray-700 overflow-hidden animate-in zoom-in-95 duration-200 max-h-[90vh] flex flex-col"
        onClick={e => e.stopPropagation()}
      >
        {/* Header */}
        <div className="relative h-32 bg-gradient-to-r from-[#1e5144] to-emerald-900 shrink-0">
            <button 
                onClick={onClose}
                className="absolute top-4 right-4 p-2 text-white/70 hover:text-white bg-black/20 hover:bg-black/40 rounded-full transition-colors"
            >
                <X size={20} />
            </button>
        </div>

        {/* Content */}
        <div className="px-8 pb-8 -mt-12 relative flex-grow overflow-y-auto custom-scrollbar">
            <div className="flex flex-col items-center mb-6">
                <div className="w-24 h-24 rounded-full border-4 border-[#1a1d24] bg-gray-700 shadow-xl flex items-center justify-center overflow-hidden">
                    {/* Placeholder for Rep Photo */}
                    <User size={48} className="text-gray-400" />
                </div>
                <h2 className="text-2xl font-bold text-white mt-3 text-center">{profile.name}</h2>
                <span className="px-3 py-1 mt-1 bg-gray-800 text-gray-400 text-xs rounded-full border border-gray-700">
                    Código: {profile.code}
                </span>
            </div>

            <div className="space-y-6">
                
                {/* Stats Section */}
                <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
                    <div className="bg-[#20232b] p-3 rounded-xl border border-gray-800 flex flex-col items-center text-center">
                        <div className="p-2 rounded-lg bg-emerald-500/10 text-emerald-400 mb-2">
                            <GraduationCap size={18} />
                        </div>
                        <span className="text-2xl font-bold text-white">{profile.trainingCount}</span>
                        <span className="text-[10px] text-gray-500 uppercase tracking-wide">Treinamentos</span>
                    </div>
                    <div className="bg-[#20232b] p-3 rounded-xl border border-gray-800 flex flex-col items-center text-center">
                        <div className="p-2 rounded-lg bg-blue-500/10 text-blue-400 mb-2">
                            <Globe size={18} />
                        </div>
                        <span className="text-2xl font-bold text-white">{profile.onlineCount}</span>
                        <span className="text-[10px] text-gray-500 uppercase tracking-wide">Ações Online</span>
                    </div>
                    <div className="bg-[#20232b] p-3 rounded-xl border border-gray-800 flex flex-col items-center text-center">
                        <div className="p-2 rounded-lg bg-amber-500/10 text-amber-400 mb-2">
                            <FolderOpen size={18} />
                        </div>
                        <span className="text-2xl font-bold text-white">{profile.offlineCount}</span>
                        <span className="text-[10px] text-gray-500 uppercase tracking-wide">Ações Offline</span>
                    </div>
                    <div className="bg-[#20232b] p-3 rounded-xl border border-gray-800 flex flex-col items-center text-center">
                        <div className="p-2 rounded-lg bg-rose-500/10 text-rose-400 mb-2">
                            <DollarSign size={18} />
                        </div>
                        <span className="text-lg font-bold text-white mt-1">
                            {profile.offlineValue.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL', maximumFractionDigits: 0 })}
                        </span>
                        <span className="text-[10px] text-gray-500 uppercase tracking-wide mt-1">Investimento</span>
                    </div>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div className="bg-[#20232b] p-4 rounded-xl border border-gray-800">
                        <h3 className="text-xs font-semibold text-gray-500 uppercase mb-3 flex items-center gap-2">
                            <Building size={14} /> Dados da Representação
                        </h3>
                        <div className="space-y-2">
                            <div>
                                <p className="text-xs text-gray-500">Nome da Representação</p>
                                <p className="text-sm font-medium text-gray-200">{profile.company}</p>
                            </div>
                            <div className="border-t border-gray-700/50 pt-2 mt-2">
                                <p className="text-xs text-gray-500 flex items-center gap-1 mb-1">
                                    <MapPin size={12} /> Região de Atuação
                                </p>
                                <p className="text-sm text-gray-300 leading-relaxed">{profile.region}</p>
                            </div>
                        </div>
                    </div>

                    <div className="bg-[#20232b] p-4 rounded-xl border border-gray-800">
                        <h3 className="text-xs font-semibold text-gray-500 uppercase mb-3 flex items-center gap-2">
                            <Hash size={14} /> Contatos
                        </h3>
                        <div className="space-y-3">
                            <div className="flex items-start gap-3">
                                <Phone size={16} className="text-emerald-500 mt-0.5" />
                                <div>
                                    <p className="text-xs text-gray-500">Telefone / WhatsApp</p>
                                    <p className="text-sm text-gray-200 whitespace-pre-line">{profile.phone}</p>
                                </div>
                            </div>
                            <div className="flex items-start gap-3">
                                <Mail size={16} className="text-blue-500 mt-0.5" />
                                <div>
                                    <p className="text-xs text-gray-500">E-mail</p>
                                    <p className="text-sm text-gray-200 whitespace-pre-line">{profile.email}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div className="flex items-center justify-between px-4 py-3 bg-[#1e5144]/10 rounded-lg border border-[#1e5144]/20">
                    <span className="text-sm text-emerald-400 font-medium flex items-center gap-2">
                        <User size={16} /> Atendente Solis
                    </span>
                    <span className="text-sm font-bold text-white">{profile.attendant}</span>
                </div>
            </div>
        </div>
      </div>
    </div>
  );
};

export default RepProfileModal;
