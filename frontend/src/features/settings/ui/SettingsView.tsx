import React, { useState } from 'react';
import { 
  User, 
  Mail, 
  Lock, 
  Camera, 
  Save, 
  LogOut, 
  Shield, 
  BellRing,
  CheckCircle2,
  Archive,
  History,
  RefreshCcw,
  Clock,
  Target
} from 'lucide-react';
import type { Notification, KpiCategory } from '@/shared/types';
import type { SettingsViewProps, SettingsSection, UserFormData } from '../types';

const SettingsView: React.FC<SettingsViewProps> = ({ 
    onLogout, 
    notifications = [], 
    onUpdateNotification,
    kpis,
    onUpdateKpiMeta 
}) => {
  const [activeSection, setActiveSection] = useState<SettingsSection>('profile');
  const [isSaving, setIsSaving] = useState(false);
  const [showSuccess, setShowSuccess] = useState(false);

  // Goal Editing State
  const [selectedGoalKpi, setSelectedGoalKpi] = useState<string>(kpis && kpis.length > 0 ? kpis[0].id : '');

  // Mock User State
  const [formData, setFormData] = useState<UserFormData>({
    name: 'Jackson',
    role: 'Coordenador de Marketing',
    email: 'jackson@solis.com.br',
    phone: '(19) 99999-9999',
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  });

  const handleSave = (e: React.FormEvent) => {
    e.preventDefault();
    setIsSaving(true);
    
    // Simulate API call
    setTimeout(() => {
      setIsSaving(false);
      setShowSuccess(true);
      setTimeout(() => setShowSuccess(false), 3000);
    }, 1500);
  };

  const handleUnarchive = (id: string) => {
    if (onUpdateNotification) {
        // Ao desarquivar, definimos read: false para que ela volte ao topo da lista e notifique o usuário
        const updated = notifications.map(n => n.id === id ? { ...n, archived: false, read: false } : n);
        onUpdateNotification(updated);
    }
  };

  const archivedNotifications = notifications.filter(n => n.archived);

  // Filter KPI for goals
  const currentKpiForGoals = kpis?.find(k => k.id === selectedGoalKpi);

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 max-w-5xl mx-auto pb-10">
      
      {/* Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-100 mb-2">Configurações da Conta</h1>
        <p className="text-gray-500">Gerencie suas informações pessoais e segurança.</p>
      </div>

      <div className="flex flex-col lg:flex-row gap-8">
        
        {/* Sidebar Navigation (Local) */}
        <div className="w-full lg:w-64 flex-shrink-0 space-y-2">
            <button 
                onClick={() => setActiveSection('profile')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'profile' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <User size={18} /> Perfil Público
            </button>
            <button 
                onClick={() => setActiveSection('security')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'security' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <Shield size={18} /> Segurança
            </button>
            <button 
                onClick={() => setActiveSection('history')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'history' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <History size={18} /> Histórico de Notificações
            </button>
            {kpis && onUpdateKpiMeta && (
                <button 
                    onClick={() => setActiveSection('goals')}
                    className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'goals' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
                >
                    <Target size={18} /> Metas do Sistema
                </button>
            )}
            
            <div className="pt-8 mt-8 border-t border-gray-800">
                <button 
                    onClick={onLogout}
                    className="w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 text-rose-400 hover:bg-rose-500/10 hover:text-rose-300 transition-colors border border-transparent hover:border-rose-500/20"
                >
                    <LogOut size={18} /> Sair do Sistema
                </button>
            </div>
        </div>

        {/* Main Content Form */}
        <div className="flex-grow">
            {activeSection === 'profile' && (
                <form onSubmit={handleSave} className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden">
                        <div className="p-8 space-y-8">
                            {/* Avatar Section */}
                            <div className="flex items-center gap-6 pb-8 border-b border-gray-800">
                                <div className="relative group cursor-pointer">
                                    <div className="w-24 h-24 rounded-full bg-gradient-to-tr from-gray-700 to-gray-600 border-4 border-[#1a1d24] flex items-center justify-center text-3xl font-bold text-white shadow-lg overflow-hidden">
                                        {formData.name.substring(0, 2).toUpperCase()}
                                    </div>
                                    <div className="absolute inset-0 bg-black/50 rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                                        <Camera size={24} className="text-white" />
                                    </div>
                                </div>
                                <div>
                                    <h3 className="text-lg font-bold text-gray-100">Sua Foto</h3>
                                    <p className="text-sm text-gray-500 mb-3">Isso será exibido no seu perfil.</p>
                                    <div className="flex gap-3">
                                        <button type="button" className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white text-xs font-semibold rounded-lg transition-colors">Alterar</button>
                                        <button type="button" className="px-4 py-2 bg-gray-800 hover:bg-gray-700 text-gray-300 text-xs font-semibold rounded-lg transition-colors">Remover</button>
                                    </div>
                                </div>
                            </div>

                            {/* Form Fields */}
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Nome Completo</label>
                                    <div className="relative">
                                        <User size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input 
                                            type="text" 
                                            value={formData.name}
                                            onChange={(e) => setFormData({...formData, name: e.target.value})}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Cargo</label>
                                    <input 
                                        type="text" 
                                        value={formData.role}
                                        onChange={(e) => setFormData({...formData, role: e.target.value})}
                                        className="w-full px-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                    />
                                </div>
                                <div className="md:col-span-2">
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">E-mail</label>
                                    <div className="relative">
                                        <Mail size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input 
                                            type="email" 
                                            value={formData.email}
                                            onChange={(e) => setFormData({...formData, email: e.target.value})}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                        />
                                    </div>
                                </div>
                            </div>
                        </div>
                        {/* Footer Actions */}
                        <div className="p-6 bg-[#20232b] border-t border-gray-800 flex justify-between items-center">
                            <div>
                                {showSuccess && (
                                    <span className="flex items-center gap-2 text-emerald-400 text-sm font-medium animate-in fade-in slide-in-from-left-2">
                                        <CheckCircle2 size={16} /> Alterações salvas com sucesso!
                                    </span>
                                )}
                            </div>
                            <button 
                                type="submit"
                                disabled={isSaving}
                                className="px-6 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white font-semibold rounded-xl flex items-center gap-2 shadow-lg shadow-[#1e5144]/20 transition-all disabled:opacity-70 disabled:cursor-not-allowed"
                            >
                                {isSaving ? (
                                    <div className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                                ) : (
                                    <Save size={18} />
                                )}
                                Salvar Alterações
                            </button>
                        </div>
                </form>
            )}

            {activeSection === 'security' && (
                <form onSubmit={handleSave} className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden">
                        <div className="p-8 space-y-8">
                            <div className="pb-6 border-b border-gray-800">
                                <h3 className="text-lg font-bold text-gray-100">Alterar Senha</h3>
                                <p className="text-sm text-gray-500">Mantenha sua conta segura alterando sua senha regularmente.</p>
                            </div>

                            <div className="space-y-4 max-w-md">
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Senha Atual</label>
                                    <div className="relative">
                                        <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input 
                                            type="password" 
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                            placeholder="••••••••"
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Nova Senha</label>
                                    <div className="relative">
                                        <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input 
                                            type="password" 
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                            placeholder="••••••••"
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Confirmar Nova Senha</label>
                                    <div className="relative">
                                        <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input 
                                            type="password" 
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                            placeholder="••••••••"
                                        />
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div className="p-6 bg-[#20232b] border-t border-gray-800 flex justify-end">
                            <button 
                                type="submit"
                                className="px-6 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white font-semibold rounded-xl flex items-center gap-2 shadow-lg shadow-[#1e5144]/20 transition-all"
                            >
                                <Save size={18} /> Salvar Senha
                            </button>
                        </div>
                </form>
            )}

            {activeSection === 'history' && (
                <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden h-full flex flex-col">
                    <div className="p-6 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                        <div>
                            <h3 className="text-lg font-bold text-gray-100 flex items-center gap-2">
                                <Archive size={20} className="text-gray-400" />
                                Notificações Arquivadas
                            </h3>
                            <p className="text-sm text-gray-500">Histórico de alertas que você arquivou.</p>
                        </div>
                    </div>
                    
                    <div className="flex-grow overflow-y-auto custom-scrollbar">
                        {archivedNotifications.length === 0 ? (
                            <div className="p-12 text-center text-gray-500">
                                <Archive size={48} className="mx-auto mb-4 opacity-20" />
                                <p className="text-lg font-medium">Histórico vazio</p>
                                <p className="text-sm">Você não tem notificações arquivadas.</p>
                            </div>
                        ) : (
                            <div className="divide-y divide-gray-800">
                                {archivedNotifications.map(notification => (
                                    <div key={notification.id} className="p-4 hover:bg-[#20232b] transition-colors flex gap-4">
                                        <div className="w-10 h-10 rounded-full bg-gray-800 border border-gray-700 flex items-center justify-center flex-shrink-0 text-gray-400">
                                            {notification.type === 'mention' && <User size={18} />}
                                            {notification.type === 'deadline' && <Clock size={18} />}
                                            {notification.type === 'system' && <BellRing size={18} />}
                                        </div>
                                        <div className="flex-grow">
                                            <div className="flex justify-between items-start">
                                                <h4 className="text-sm font-semibold text-gray-200">{notification.title}</h4>
                                                <span className="text-[10px] text-gray-500">{notification.timestamp}</span>
                                            </div>
                                            <p className="text-sm text-gray-400 mt-1">{notification.message}</p>
                                        </div>
                                        <button 
                                            onClick={() => handleUnarchive(notification.id)}
                                            className="p-2 text-gray-500 hover:text-emerald-400 hover:bg-emerald-500/10 rounded-lg transition-colors self-center"
                                            title="Desarquivar"
                                        >
                                            <RefreshCcw size={18} />
                                        </button>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                </div>
            )}

            {activeSection === 'goals' && currentKpiForGoals && onUpdateKpiMeta && (
                <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden h-full flex flex-col">
                    <div className="p-6 border-b border-gray-800 bg-[#20232b]">
                        <h3 className="text-lg font-bold text-gray-100 flex items-center gap-2 mb-4">
                            <Target size={20} className="text-emerald-400" />
                            Ajuste de Metas
                        </h3>
                        
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Selecione o Indicador (KPI)</label>
                            <select 
                                value={selectedGoalKpi}
                                onChange={(e) => setSelectedGoalKpi(e.target.value)}
                                className="w-full md:w-1/2 bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                            >
                                {kpis?.map(k => (
                                    <option key={k.id} value={k.id}>{k.title}</option>
                                ))}
                            </select>
                        </div>
                    </div>

                    <div className="p-6 overflow-y-auto custom-scrollbar flex-grow">
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                            {currentKpiForGoals.data.map((monthlyData, index) => (
                                <div key={monthlyData.month} className="bg-[#0f1115] p-4 rounded-xl border border-gray-800 hover:border-gray-700 transition-colors">
                                    <div className="flex justify-between items-center mb-2">
                                        <span className="font-bold text-gray-300">{monthlyData.month}</span>
                                        <span className="text-xs text-gray-500">Mês {index + 1}</span>
                                    </div>
                                    <div className="relative">
                                        <label className="block text-[10px] text-gray-500 uppercase font-bold mb-1">Meta Definida</label>
                                        <input 
                                            type="number"
                                            value={monthlyData.meta !== null ? monthlyData.meta : ''}
                                            onChange={(e) => {
                                                const val = parseFloat(e.target.value);
                                                onUpdateKpiMeta(currentKpiForGoals.id, monthlyData.month, isNaN(val) ? 0 : val);
                                            }}
                                            className="w-full bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144] font-mono text-sm"
                                            placeholder="Definir meta"
                                        />
                                        <span className="absolute right-3 top-[26px] text-gray-500 text-xs pointer-events-none">
                                            {currentKpiForGoals.unit === 'currency' ? 'R$' : currentKpiForGoals.unit === 'percent' ? '%' : 'un'}
                                        </span>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                    
                    <div className="p-4 border-t border-gray-800 bg-[#20232b] text-center">
                        <p className="text-xs text-gray-500">
                            As alterações são salvas automaticamente no sistema.
                        </p>
                    </div>
                </div>
            )}
        </div>
      </div>
    </div>
  );
};

export default SettingsView;
