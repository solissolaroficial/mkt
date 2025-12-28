 
import React, { useState } from 'react';
import { 
  DollarSign, 
  Calendar, 
  FileText, 
  Barcode, 
  Send, 
  AlertCircle, 
  CheckCircle2, 
  Clock, 
  X,
  Mail,
  ArrowLeft,
  Plus,
  Trash2,
  Edit2,
  Repeat
} from 'lucide-react';
import type { AccountPayable } from '@/shared/types';
import type { FinancialViewProps } from '../types';

const FinancialView: React.FC<FinancialViewProps> = ({ onBack }) => {
  const [activeTab, setActiveTab] = useState<'overdue' | 'today' | 'upcoming'>('today');
  const [accounts, setAccounts] = useState<AccountPayable[]>([]);
  
  // Email Modal State
  const [isEmailModalOpen, setIsEmailModalOpen] = useState(false);
  const [selectedAccount, setSelectedAccount] = useState<AccountPayable | null>(null);
  const [emailTo, setEmailTo] = useState('financeiro@solis.com.br');
  const [emailSubject, setEmailSubject] = useState('');
  const [emailBody, setEmailBody] = useState('');

  // Add/Edit Form Modal State
  const [isFormModalOpen, setIsFormModalOpen] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [formData, setFormData] = useState({
      supplier: '',
      description: '',
      amount: '',
      dueDate: '',
      recurrence: 'none' as 'none' | 'monthly' | 'yearly'
  });

  // Helpers
  const today = new Date().toISOString().split('T')[0];

  const getFilteredAccounts = () => {
    return accounts.filter(acc => {
      if (acc.status === 'sent_to_finance') return false; // Optional: hide sent ones or keep them? Keeping simple filters for now.

      if (activeTab === 'overdue') return acc.dueDate < today;
      if (activeTab === 'today') return acc.dueDate === today;
      if (activeTab === 'upcoming') return acc.dueDate > today;
      return true;
    });
  };

  const toggleStatus = (id: string, field: 'nfArrived' | 'boletoArrived') => {
    setAccounts(prev => prev.map(acc => 
      acc.id === id ? { ...acc, [field]: !acc[field] } : acc
    ));
  };

  // --- Email Logic ---
  const handleOpenEmailModal = (account: AccountPayable) => {
    setSelectedAccount(account);
    setEmailSubject(`Pagamento: ${account.supplier} - Vencimento ${account.dueDate.split('-').reverse().join('/')}`);
    setEmailBody(`Olá Financeiro,\n\nSegue em anexo a Nota Fiscal e Boleto para o pagamento referente a ${account.description}.\n\nFornecedor: ${account.supplier}\nValor: R$ ${account.amount.toLocaleString('pt-BR', { minimumFractionDigits: 2 })}\nVencimento: ${account.dueDate.split('-').reverse().join('/')}\n\nAtenciosamente,\nJackson`);
    setIsEmailModalOpen(true);
  };

  const handleSendEmail = () => {
    if (selectedAccount) {
      setAccounts(prev => prev.map(acc => 
        acc.id === selectedAccount.id ? { ...acc, status: 'sent_to_finance' } : acc
      ));
      setIsEmailModalOpen(false);
    }
  };

  // --- Add/Edit Logic ---
  const handleOpenForm = (account?: AccountPayable) => {
      if (account) {
          setEditingId(account.id);
          setFormData({
              supplier: account.supplier,
              description: account.description,
              amount: account.amount.toString(),
              dueDate: account.dueDate,
              recurrence: account.recurrence || 'none'
          });
      } else {
          setEditingId(null);
          setFormData({
              supplier: '',
              description: '',
              amount: '',
              dueDate: new Date().toISOString().split('T')[0],
              recurrence: 'none'
          });
      }
      setIsFormModalOpen(true);
  };

  const handleDelete = (id: string) => {
      if (window.confirm('Tem certeza que deseja excluir esta conta?')) {
          setAccounts(prev => prev.filter(acc => acc.id !== id));
      }
  };

  const handleSaveForm = (e: React.FormEvent) => {
      e.preventDefault();
      
      const newAccount: AccountPayable = {
          id: editingId || Date.now().toString(),
          supplier: formData.supplier,
          description: formData.description,
          amount: parseFloat(formData.amount),
          dueDate: formData.dueDate,
          recurrence: formData.recurrence,
          nfArrived: editingId ? (accounts.find(a => a.id === editingId)?.nfArrived || false) : false,
          boletoArrived: editingId ? (accounts.find(a => a.id === editingId)?.boletoArrived || false) : false,
          status: 'pending'
      };

      if (editingId) {
          setAccounts(prev => prev.map(acc => acc.id === editingId ? newAccount : acc));
      } else {
          setAccounts(prev => [...prev, newAccount]);
      }
      setIsFormModalOpen(false);
  };

  const filteredAccounts = getFilteredAccounts();

  // Summary counts
  const overdueCount = accounts.filter(a => a.dueDate < today && a.status === 'pending').length;
  const todayCount = accounts.filter(a => a.dueDate === today && a.status === 'pending').length;
  const upcomingCount = accounts.filter(a => a.dueDate > today && a.status === 'pending').length;

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
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
            <h1 className="text-2xl font-bold text-gray-100">Contas a Pagar</h1>
            <p className="text-gray-500">Gestão de pagamentos e envio ao financeiro</p>
          </div>
        </div>

        <button 
            onClick={() => handleOpenForm()}
            className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/20"
        >
            <Plus size={18} /> Novo Pagamento
        </button>
      </div>

      {/* Stats / Tabs */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
        <button 
          onClick={() => setActiveTab('overdue')}
          className={`p-4 rounded-xl border flex items-center justify-between transition-all ${
            activeTab === 'overdue' 
              ? 'bg-rose-900/20 border-rose-500/50 ring-1 ring-rose-500/50' 
              : 'bg-[#1a1d24] border-white/5 hover:border-white/10'
          }`}
        >
          <div className="flex items-center gap-3">
             <div className="p-2 rounded-lg bg-rose-500/10 text-rose-400">
                <AlertCircle size={20} />
             </div>
             <div className="text-left">
                <p className="text-sm text-gray-400">Vencidas</p>
                <p className="text-xl font-bold text-white">{overdueCount}</p>
             </div>
          </div>
        </button>

        <button 
          onClick={() => setActiveTab('today')}
          className={`p-4 rounded-xl border flex items-center justify-between transition-all ${
            activeTab === 'today' 
              ? 'bg-amber-900/20 border-amber-500/50 ring-1 ring-amber-500/50' 
              : 'bg-[#1a1d24] border-white/5 hover:border-white/10'
          }`}
        >
          <div className="flex items-center gap-3">
             <div className="p-2 rounded-lg bg-amber-500/10 text-amber-400">
                <Clock size={20} />
             </div>
             <div className="text-left">
                <p className="text-sm text-gray-400">Vence Hoje</p>
                <p className="text-xl font-bold text-white">{todayCount}</p>
             </div>
          </div>
        </button>

        <button 
          onClick={() => setActiveTab('upcoming')}
          className={`p-4 rounded-xl border flex items-center justify-between transition-all ${
            activeTab === 'upcoming' 
              ? 'bg-emerald-900/20 border-emerald-500/50 ring-1 ring-emerald-500/50' 
              : 'bg-[#1a1d24] border-white/5 hover:border-white/10'
          }`}
        >
          <div className="flex items-center gap-3">
             <div className="p-2 rounded-lg bg-emerald-500/10 text-emerald-400">
                <Calendar size={20} />
             </div>
             <div className="text-left">
                <p className="text-sm text-gray-400">A Vencer</p>
                <p className="text-xl font-bold text-white">{upcomingCount}</p>
             </div>
          </div>
        </button>
      </div>

      {/* Table */}
      <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col">
         <div className="overflow-x-auto custom-scrollbar flex-grow">
            <table className="w-full text-sm text-left">
                <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
                    <tr>
                        <th className="px-6 py-4">Fornecedor</th>
                        <th className="px-6 py-4">Vencimento</th>
                        <th className="px-6 py-4 text-right">Valor</th>
                        <th className="px-6 py-4 text-center">NF</th>
                        <th className="px-6 py-4 text-center">Boleto</th>
                        <th className="px-6 py-4 text-center">Ação</th>
                    </tr>
                </thead>
                <tbody className="divide-y divide-gray-800">
                    {filteredAccounts.length === 0 ? (
                        <tr>
                            <td colSpan={6} className="px-6 py-12 text-center text-gray-500">
                                Nenhuma conta encontrada nesta categoria.
                            </td>
                        </tr>
                    ) : (
                        filteredAccounts.map(acc => (
                            <tr key={acc.id} className="hover:bg-[#20232b] transition-colors group">
                                <td className="px-6 py-4">
                                    <div className="flex items-center gap-2">
                                        <p className="font-medium text-gray-200">{acc.supplier}</p>
                                        {acc.recurrence && acc.recurrence !== 'none' && (
                                            <span className="bg-blue-500/10 text-blue-400 text-[10px] px-1.5 py-0.5 rounded border border-blue-500/20 uppercase font-bold" title={`Recorrência: ${acc.recurrence}`}>
                                                {acc.recurrence === 'monthly' ? 'Mensal' : 'Anual'}
                                            </span>
                                        )}
                                    </div>
                                    <p className="text-xs text-gray-500">{acc.description}</p>
                                </td>
                                <td className="px-6 py-4 text-gray-300">
                                    {acc.dueDate.split('-').reverse().join('/')}
                                </td>
                                <td className="px-6 py-4 text-right font-mono text-gray-200">
                                    {acc.amount.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                                </td>
                                <td className="px-6 py-4 text-center">
                                    <button 
                                        onClick={() => toggleStatus(acc.id, 'nfArrived')}
                                        className={`p-2 rounded-lg transition-colors border ${
                                            acc.nfArrived 
                                            ? 'bg-blue-500/10 text-blue-400 border-blue-500/20' 
                                            : 'bg-gray-800 text-gray-600 border-gray-700 hover:text-gray-400'
                                        }`}
                                        title={acc.nfArrived ? "NF Recebida" : "NF Pendente"}
                                    >
                                        <FileText size={18} />
                                    </button>
                                </td>
                                <td className="px-6 py-4 text-center">
                                    <button 
                                        onClick={() => toggleStatus(acc.id, 'boletoArrived')}
                                        className={`p-2 rounded-lg transition-colors border ${
                                            acc.boletoArrived 
                                            ? 'bg-blue-500/10 text-blue-400 border-blue-500/20' 
                                            : 'bg-gray-800 text-gray-600 border-gray-700 hover:text-gray-400'
                                        }`}
                                        title={acc.boletoArrived ? "Boleto Recebido" : "Boleto Pendente"}
                                    >
                                        <Barcode size={18} />
                                    </button>
                                </td>
                                <td className="px-6 py-4 text-center">
                                    <div className="flex items-center justify-center gap-2">
                                        <button 
                                            onClick={() => handleOpenEmailModal(acc)}
                                            disabled={!acc.nfArrived || !acc.boletoArrived}
                                            className="inline-flex items-center gap-2 px-3 py-1.5 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-xs font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-[#1e5144]/20"
                                        >
                                            <Send size={14} /> Enviar
                                        </button>
                                        
                                        <div className="flex bg-[#0f1115] rounded-lg border border-gray-700 overflow-hidden">
                                            <button 
                                                onClick={() => handleOpenForm(acc)}
                                                className="p-1.5 hover:bg-gray-800 text-gray-400 hover:text-blue-400 transition-colors border-r border-gray-700"
                                                title="Editar"
                                            >
                                                <Edit2 size={14} />
                                            </button>
                                            <button 
                                                onClick={() => handleDelete(acc.id)}
                                                className="p-1.5 hover:bg-gray-800 text-gray-400 hover:text-rose-400 transition-colors"
                                                title="Excluir"
                                            >
                                                <Trash2 size={14} />
                                            </button>
                                        </div>
                                    </div>
                                </td>
                            </tr>
                        ))
                    )}
                </tbody>
            </table>
         </div>
      </div>

      {/* Email Modal */}
      {isEmailModalOpen && selectedAccount && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsEmailModalOpen(false)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-lg rounded-2xl shadow-2xl border border-gray-700 overflow-hidden animate-in zoom-in-95"
                onClick={e => e.stopPropagation()}
            >
                <div className="p-6 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <Mail size={18} /> Enviar ao Financeiro
                    </h3>
                    <button onClick={() => setIsEmailModalOpen(false)} className="text-gray-500 hover:text-white">
                        <X size={20} />
                    </button>
                </div>
                
                <div className="p-6 space-y-4">
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Para:</label>
                        <input 
                            type="email" 
                            value={emailTo}
                            onChange={(e) => setEmailTo(e.target.value)}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                        />
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Assunto:</label>
                        <input 
                            type="text" 
                            value={emailSubject}
                            onChange={(e) => setEmailSubject(e.target.value)}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                        />
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Mensagem:</label>
                        <textarea 
                            value={emailBody}
                            onChange={(e) => setEmailBody(e.target.value)}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144] h-40 resize-none"
                        />
                    </div>

                    <div className="flex items-center gap-2 text-xs text-gray-500 bg-[#15171c] p-2 rounded border border-gray-800">
                        <CheckCircle2 size={14} className="text-emerald-500" />
                        Os arquivos (NF e Boleto) serão anexados automaticamente.
                    </div>
                </div>

                <div className="p-4 border-t border-gray-800 flex justify-end gap-3 bg-[#15171c]">
                    <button 
                        onClick={() => setIsEmailModalOpen(false)}
                        className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                    >
                        Cancelar
                    </button>
                    <button 
                        onClick={handleSendEmail}
                        className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2"
                    >
                        <Send size={16} /> Enviar E-mail
                    </button>
                </div>
            </div>
        </div>
      )}

      {/* Add/Edit Payment Modal */}
      {isFormModalOpen && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsFormModalOpen(false)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-lg rounded-2xl shadow-2xl border border-gray-700 overflow-hidden animate-in zoom-in-95"
                onClick={e => e.stopPropagation()}
            >
                <div className="p-6 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <DollarSign size={18} className="text-emerald-400" /> 
                        {editingId ? 'Editar Pagamento' : 'Novo Pagamento'}
                    </h3>
                    <button onClick={() => setIsFormModalOpen(false)} className="text-gray-500 hover:text-white">
                        <X size={20} />
                    </button>
                </div>
                
                <form onSubmit={handleSaveForm} className="p-6 space-y-4">
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Fornecedor</label>
                        <input 
                            type="text" 
                            value={formData.supplier}
                            onChange={(e) => setFormData({...formData, supplier: e.target.value})}
                            placeholder="Nome da empresa ou prestador"
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        />
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Descrição</label>
                        <input 
                            type="text" 
                            value={formData.description}
                            onChange={(e) => setFormData({...formData, description: e.target.value})}
                            placeholder="Referência do pagamento"
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        />
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Vencimento</label>
                            <input 
                                type="date" 
                                value={formData.dueDate}
                                onChange={(e) => setFormData({...formData, dueDate: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                required
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Valor (R$)</label>
                            <input 
                                type="number" 
                                step="0.01"
                                value={formData.amount}
                                onChange={(e) => setFormData({...formData, amount: e.target.value})}
                                placeholder="0,00"
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                required
                            />
                        </div>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1 flex items-center gap-1">
                            <Repeat size={12} /> Recorrência
                        </label>
                        <select 
                            value={formData.recurrence}
                            onChange={(e) => setFormData({...formData, recurrence: e.target.value as any})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                        >
                            <option value="none">Nenhuma (Pagamento Único)</option>
                            <option value="monthly">Mensal</option>
                            <option value="yearly">Anual</option>
                        </select>
                        <p className="text-xs text-gray-500 mt-1">
                            Selecione para gerar lembretes automáticos para os próximos períodos.
                        </p>
                    </div> 

                    <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
                        <button 
                            type="button" 
                            onClick={() => setIsFormModalOpen(false)}
                            className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                        >
                            Cancelar
                        </button>
                        <button 
                            type="submit" 
                            className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg"
                        >
                            Salvar
                        </button>
                    </div>
                </form>
            </div>
        </div>
      )}

    </div>
  );
};

export default FinancialView;
