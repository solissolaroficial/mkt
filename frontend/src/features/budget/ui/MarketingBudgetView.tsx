 
import React, { useState, useMemo } from 'react';
import { ArrowLeft, Wallet, TrendingUp, AlertCircle, Plus, FileText, Check, X, Building, Phone, Mail, FileSpreadsheet, DollarSign, User, Calendar } from 'lucide-react';
import { MONTHS, FULL_MONTH_MAP } from '@/shared/utils/constants';
import type { OfflineAction } from '@/shared/types';
import type { MarketingBudgetViewProps, BudgetSummaryItem } from '../types';

const parseCurrency = (valStr: string) => {
    if (!valStr) return 0;
    // Remove R$, spaces, dots, replace comma with dot
    const clean = valStr.toString().replace(/[R$\s.]/g, '').replace(',', '.');
    return parseFloat(clean) || 0;
  }

// Helper to check if date string matches selected month
const isDateInMonth = (dateStr: string, selectedMonth: string) => {
    if (selectedMonth === '---') return true;
    
    // dateStr format: DD/MM/YYYY
    const parts = dateStr.split('/');
    if (parts.length !== 3) return false;
    
    const monthIndex = parseInt(parts[1], 10) - 1; // 0-based
    const monthName = MONTHS[monthIndex];
    
    return monthName === selectedMonth;
  }

// Helper to normalize category matching between Budget Items and Requests
const normalizeCategory = (cat: string) => {
    if (!cat) return '';
    const upper = cat.toUpperCase().trim();
    // Manual mapping for known discrepancies
    if (upper === 'PROPAGANDAS E PUBLICIDADES') return 'DESPESAS C/ PROPAGANDA/PUBLIC';
    if (upper === 'DESPESAS C/ PROPAGANDA/PUBLIC') return 'DESPESAS C/ PROPAGANDA/PUBLIC';
    
    return upper;
  }

const MarketingBudgetView: React.FC<MarketingBudgetViewProps> = ({ onBack, selectedMonth }) => {
  // --- STATE ---
  const [requests, setRequests] = useState<OfflineAction[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  
  // Form State
  const [formData, setFormData] = useState({
      rep: '',
      client: '',
      contact: '',
      phone: '',
      email: '',
      proposal: '',
      volume12m: '',
      requestedVal: '',
      approvedVal: '',
      category: '',
  });

  // --- DATA CALCULATION ---
  // Filter for Marketing Actions (Group 13, Subgroup 41)
  const budgetItems = useMemo(() => {
      return [];
  }, []);

  const budgetSummary = useMemo(() => {
      const monthIndex = selectedMonth === '---' ? -1 : MONTHS.indexOf(selectedMonth);

      return budgetItems.map(item => {
          let totalBudget = 0;
          let totalRealizedBase = 0;

          // Calculate Base Budget & Realized from ERP Data (RAW_BUDGET_ITEMS)
          if (monthIndex === -1) {
              totalBudget = item.vals.reduce((acc, curr) => acc + curr, 0);
              totalRealizedBase = item.realizedVals.reduce((acc, curr) => acc + curr, 0);
          } else {
              totalBudget = item.vals[monthIndex] || 0;
              totalRealizedBase = item.realizedVals[monthIndex] || 0;
          }
          
          let totalRealized = totalRealizedBase;
          let totalRequested = 0;

          // Normalize item description for matching
          const itemDescNormalized = normalizeCategory(item.desc);

          // Process Requests (Add to Realized or Requested based on status)
          requests.forEach(req => {
              const reqCatNormalized = normalizeCategory(req.category || '');
              
              // Match category AND selected month
              if (reqCatNormalized === itemDescNormalized && isDateInMonth(req.data, selectedMonth)) { 
                  const approvedVal = parseCurrency(req.aprovado);
                  const requestedVal = parseCurrency(req.solicitado);

                  if (approvedVal > 0) {
                      // If approved, it adds to realized (Request is fulfilled)
                      totalRealized += approvedVal;
                  } else {
                      // If not approved yet (and requested > 0), it is pending request
                      totalRequested += requestedVal;
                  }
              }
          });

          const available = totalBudget - totalRealized;
          // Calculate percentage used based on Realized vs Budget
          const percentUsed = totalBudget > 0 ? (totalRealized / totalBudget) * 100 : 0;

          return {
              id: item.id,
              name: item.desc,
              budget: totalBudget,
              realized: totalRealized,
              requested: totalRequested,
              available,
              percentUsed
          };
      });
  }, [budgetItems, requests, selectedMonth]);

  // Filter Requests List for Display
  const filteredRequests = useMemo(() => {
      return requests.filter(req => isDateInMonth(req.data, selectedMonth));
  }, [requests, selectedMonth]);

  // --- HANDLERS ---
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
      const { name, value } = e.target;
      setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
      e.preventDefault();
      
      const newAction: OfflineAction = {
          solicitado: `R$ ${formData.requestedVal}`,
          data: new Date().toLocaleDateString('pt-BR'),
          aprovado: formData.approvedVal ? `R$ ${formData.approvedVal}` : '',
          pedido: '',
          saida: '',
          previsao: '',
          entrega: '',
          cidade: '',
          uf: '',
          pontuado: 'AINDA NÃO',
          status: 'Nova solicitação registrada',
          pdv: formData.client,
          representative_uuid: formData.rep,
          contact: formData.contact,
          phone: formData.phone,
          email: formData.email,
          proposal: formData.proposal,
          purchaseVolume: formData.volume12m,
          category: formData.category
      };

      setRequests([newAction, ...requests]);
      setIsModalOpen(false);
      resetForm();
  };

  const resetForm = () => {
      setFormData({
          rep: '',
          client: '',
          contact: '',
          phone: '',
          email: '',
          proposal: '',
          volume12m: '',
          requestedVal: '',
          approvedVal: '',
          category: ''
      });
  };

  // Helper for formatting currency
  const fmt = (val: number) => val.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

  // Helper for Status Badge Color
  const getStatusBadge = (status: string, pontuado: string) => {
      const lower = status.toLowerCase();
      if (lower.includes('entregue')) return 'bg-blue-900/30 text-blue-400 border-blue-500/30';
      if (lower.includes('rota')) return 'bg-blue-900/30 text-blue-400 border-blue-500/30';
      if (lower.includes('pago')) return 'bg-blue-900/30 text-blue-400 border-blue-500/30';
      if (pontuado === 'SIM') return 'bg-emerald-900/30 text-emerald-400 border-emerald-500/30';
      return 'bg-gray-800 text-gray-400 border-gray-700';
  }

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header */}
      <div className="flex items-center gap-4 mb-8">
        <button onClick={onBack} className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white">
              <ArrowLeft size={24} />
        </button>
        <div>
            <h1 className="text-2xl font-bold text-gray-100">Orçamento de Marketing</h1>
            <p className="text-gray-500 flex items-center gap-2">
                Gestão de Verba e Solicitações de Parceria
            </p>
        </div>
        <div className="ml-auto flex items-center gap-2 px-3 py-1 bg-[#1a1d24] border border-gray-700 rounded-lg text-sm text-gray-300">
            <Calendar size={16} className="text-emerald-500" />
            <span>{selectedMonth === '---' ? 'Ano Completo' : FULL_MONTH_MAP[selectedMonth]}</span>
        </div>
      </div>

      {/* Budget Summary Cards */}
      <div className="mb-8">
          <h2 className="text-lg font-semibold text-gray-200 mb-4 flex items-center gap-2">
              <Wallet size={20} className="text-emerald-500" />
              Saldos por Categoria
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              {budgetSummary.map(item => (
                  <div key={item.id} className="bg-[#1a1d24] border border-gray-800 p-5 rounded-xl shadow-lg relative overflow-hidden group hover:border-emerald-500/30 transition-all flex flex-col justify-between">
                      <div>
                          <div className="flex justify-between items-start mb-4 relative z-10">
                              <h3 className="text-sm font-bold text-gray-300 uppercase tracking-wide truncate max-w-[200px]" title={item.name}>{item.name}</h3>
                              <div className="p-2 bg-gray-800 rounded-lg text-emerald-400 border border-gray-700">
                                  <DollarSign size={16} />
                              </div>
                          </div>
                          
                          <div className="space-y-1 relative z-10 mb-4">
                              <p className="text-xs text-gray-500">Saldo Disponível</p>
                              <p className={`text-3xl font-bold ${item.available < 0 ? 'text-rose-500' : 'text-white'}`}>
                                  {fmt(item.available)}
                              </p>
                          </div>
                      </div>

                      <div className="relative z-10">
                          {/* Metrics Grid */}
                          <div className="grid grid-cols-3 gap-2 py-3 border-t border-gray-800/50">
                              <div className="flex flex-col">
                                  <span className="text-[10px] text-gray-500 uppercase font-semibold">Orçado</span>
                                  <span className="text-xs text-gray-300 font-medium">{fmt(item.budget)}</span>
                              </div>
                              <div className="flex flex-col text-center border-l border-r border-gray-800/50 px-1 relative">
                                  <span className="text-[10px] text-amber-500 uppercase font-bold tracking-tight mb-0.5" title="Valor Solicitado (Pendente)">Solicitado</span>
                                  <span className="text-xs text-amber-400 font-bold">{fmt(item.requested)}</span>
                                  {item.requested > 0 && (
                                      <div className="absolute top-0 left-1/2 -translate-x-1/2 w-8 h-0.5 bg-amber-500/50 mt-3.5"></div>
                                  )}
                              </div>
                              <div className="flex flex-col text-right">
                                  <span className="text-[10px] text-blue-500 uppercase font-bold tracking-tight mb-0.5">Realizado</span>
                                  <span className="text-xs text-blue-400 font-bold">{fmt(item.realized)}</span>
                              </div>
                          </div>
                          {/* Visual Indicator Lines for metrics */}
                          <div className="flex w-full h-1 mt-1 rounded-full overflow-hidden bg-gray-800">
                              {/* Budget line (gray base) */}
                              <div className="h-full bg-gray-600" style={{ width: '0%' }}></div> 
                              {/* Realized line (blue) */}
                              <div 
                                  className="h-full bg-blue-500" 
                                  style={{ width: `${Math.min(item.percentUsed, 100)}%` }}
                              ></div>
                          </div>
                      </div>
                  </div>
              ))}
          </div>
      </div>

      {/* Requests List */}
      <div className="flex-grow flex flex-col bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden">
          <div className="p-5 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
              <h2 className="font-bold text-gray-100 flex items-center gap-2">
                  <FileSpreadsheet size={18} className="text-blue-400" />
                  Solicitações de Parceria
              </h2>
              <button 
                  onClick={() => setIsModalOpen(true)}
                  className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/20"
              >
                  <Plus size={16} /> Nova Solicitação
              </button>
          </div>

          <div className="overflow-auto custom-scrollbar flex-grow p-4">
              <table className="w-full text-sm text-left border-separate border-spacing-y-1">
                  <thead className="text-xs text-gray-400 uppercase font-semibold sticky top-0 bg-[#1a1d24] z-10">
                      <tr>
                          <th className="px-4 py-3 border-b border-gray-800">Data</th>
                          <th className="px-4 py-3 border-b border-gray-800">Cliente / PDV</th>
                          <th className="px-4 py-3 border-b border-gray-800">Representante</th>
                          <th className="px-4 py-3 border-b border-gray-800">Categoria</th>
                          <th className="px-4 py-3 border-b border-gray-800">Contato</th>
                          <th className="px-4 py-3 border-b border-gray-800 text-right text-white font-bold tracking-wide border-r border-gray-700 bg-[#20232b]">SOLICITADO</th>
                          <th className="px-4 py-3 border-b border-gray-800 text-right">Aprovado</th>
                          <th className="px-4 py-3 border-b border-gray-800 text-center">Status</th>
                      </tr>
                  </thead>
                  <tbody className="text-gray-300">
                      {filteredRequests.length === 0 ? (
                          <tr>
                              <td colSpan={8} className="px-4 py-8 text-center text-gray-500 italic">
                                  Nenhuma solicitação encontrada para o período selecionado.
                              </td>
                          </tr>
                      ) : (
                          filteredRequests.map((req, idx) => {
                              const approvedVal = parseCurrency(req.aprovado);
                              return (
                                <tr key={idx} className="hover:bg-[#20232b] transition-colors group text-xs">
                                        <td className="px-4 py-3 text-gray-400">{req.data}</td>
                                        <td className="px-4 py-3 font-bold text-white">{req.pdv}</td>
                                        <td className="px-4 py-3">{req.representative_uuid}</td>
                                        <td className="px-4 py-3 text-gray-400 truncate max-w-[150px]" title={req.category}>
                                            {req.category || '-'}
                                        </td>
                                      <td className="px-4 py-3 text-gray-500">
                                          {req.contact || '-'}
                                      </td>
                                      <td className="px-4 py-3 text-right font-mono font-bold text-white border-r border-gray-800 bg-[#1e2128] group-hover:bg-[#252830]">
                                          {req.solicitado}
                                      </td>
                                      <td className={`px-4 py-3 text-right font-mono font-bold ${approvedVal > 0 ? 'text-emerald-400' : 'text-gray-600'}`}>
                                          {req.aprovado || '-'}
                                      </td>
                                      <td className="px-4 py-3 text-center">
                                          {req.status ? (
                                              <span className={`inline-flex px-2 py-1 rounded border text-[10px] uppercase font-bold tracking-wider ${getStatusBadge(req.status, req.pontuado)}`}>
                                                  {req.status.length > 20 ? 'Ver Obs' : req.status}
                                              </span>
                                          ) : '-'}
                                      </td>
                                </tr>
                              );
                          })
                      )}
                  </tbody>
              </table>
          </div>
      </div>

      {/* --- NEW REQUEST MODAL --- */}
      {isModalOpen && (
          <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsModalOpen(false)}>
              <div 
                  className="bg-[#1a1d24] w-full max-w-4xl max-h-[90vh] rounded-2xl shadow-2xl border border-gray-700 flex flex-col overflow-hidden animate-in zoom-in-95"
                  onClick={e => e.stopPropagation()}
              >
                  {/* Modal Header */}
                  <div className="p-6 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                      <h3 className="text-xl font-bold text-gray-100 flex items-center gap-2">
                          <FileText size={24} className="text-emerald-400" />
                          Ação Cooperada - Solicitação
                      </h3>
                      <button onClick={() => setIsModalOpen(false)} className="text-gray-500 hover:text-white transition-colors">
                          <X size={24} />
                      </button>
                  </div>

                  {/* Form Content */}
                  <div className="flex-grow overflow-y-auto custom-scrollbar p-8">
                      <form id="marketing-request-form" onSubmit={handleSubmit} className="space-y-6">
                          
                          {/* Section 1: Basic Info */}
                          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                              <div>
                                  <label className="block text-xs font-bold text-gray-400 uppercase mb-2 flex items-center gap-2">
                                      <User size={14} /> Representante
                                  </label>
                                  <select 
                                      name="rep"
                                      value={formData.rep}
                                      onChange={handleInputChange}
                                      className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent transition-all"
                                      required
                                  >
                                      <option value="" disabled>Selecione o Representante...</option>
                                  </select>
                              </div>
                              <div>
                                  <label className="block text-xs font-bold text-gray-400 uppercase mb-2 flex items-center gap-2">
                                      <Building size={14} /> Cliente (Razão Social)
                                  </label>
                                  <input 
                                      type="text" 
                                      name="client"
                                      value={formData.client}
                                      onChange={handleInputChange}
                                      className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent transition-all"
                                      placeholder="Nome do Cliente ou PDV"
                                      required
                                  />
                              </div>
                          </div>

                          {/* Category Selection */}
                          <div>
                              <label className="block text-xs font-bold text-gray-400 uppercase mb-2 flex items-center gap-2">
                                  <TrendingUp size={14} /> Categoria da Ação
                              </label>
                              <select 
                                  name="category"
                                  value={formData.category}
                                  onChange={handleInputChange}
                                  className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent transition-all"
                                  required
                              >
                                  <option value="" disabled>Selecione a Categoria...</option>
                                  {budgetSummary.map(item => (
                                      <option key={item.id} value={item.name}>{item.name}</option>
                                  ))}
                              </select>
                              <p className="text-xs text-gray-500 mt-1">O valor aprovado impactará o realizado desta categoria.</p>
                          </div>

                          {/* Section 2: Contact Info */}
                          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 bg-[#15171c] p-4 rounded-xl border border-gray-800">
                              <div>
                                  <label className="block text-xs font-bold text-gray-500 uppercase mb-2">Contato</label>
                                  <input 
                                      type="text" 
                                      name="contact"
                                      value={formData.contact}
                                      onChange={handleInputChange}
                                      className="w-full bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:ring-1 focus:ring-[#1e5144] outline-none"
                                      placeholder="Nome do contato"
                                  />
                              </div>
                              <div>
                                  <label className="block text-xs font-bold text-gray-500 uppercase mb-2 flex items-center gap-1"><Phone size={12}/> Telefone</label>
                                  <input 
                                      type="text" 
                                      name="phone"
                                      value={formData.phone}
                                      onChange={handleInputChange}
                                      className="w-full bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:ring-1 focus:ring-[#1e5144] outline-none"
                                      placeholder="(00) 00000-0000"
                                  />
                              </div>
                              <div>
                                  <label className="block text-xs font-bold text-gray-500 uppercase mb-2 flex items-center gap-1"><Mail size={12}/> E-mail</label>
                                  <input 
                                      type="email" 
                                      name="email"
                                      value={formData.email}
                                      onChange={handleInputChange}
                                      className="w-full bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:ring-1 focus:ring-[#1e5144] outline-none"
                                      placeholder="contato@cliente.com.br"
                                  />
                              </div>
                          </div>

                          {/* Section 3: Proposal */}
                          <div>
                              <label className="block text-xs font-bold text-gray-400 uppercase mb-2">Proposta da Parceria</label>
                              <textarea 
                                  name="proposal"
                                  value={formData.proposal}
                                  onChange={handleInputChange}
                                  className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent transition-all h-32 resize-none"
                                  placeholder="Descrever com clareza as informações solicitadas sobre a demanda..."
                                  required
                              />
                          </div>

                          {/* Section 4: Values */}
                          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                              <div>
                                  <label className="block text-xs font-bold text-gray-400 uppercase mb-2">Volume de compras (12 meses)</label>
                                  <input 
                                      type="text" 
                                      name="volume12m"
                                      value={formData.volume12m}
                                      onChange={handleInputChange}
                                      className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                      placeholder="R$ 0,00"
                                  />
                              </div>
                              <div>
                                  <label className="block text-xs font-bold text-gray-400 uppercase mb-2">Valor Solicitado</label>
                                  <div className="relative">
                                      <span className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 font-bold">R$</span>
                                      <input 
                                          type="number" 
                                          name="requestedVal"
                                          value={formData.requestedVal}
                                          onChange={handleInputChange}
                                          className="w-full bg-[#0f1115] border border-gray-700 rounded-lg pl-10 pr-4 py-3 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                          placeholder="0,00"
                                          required
                                      />
                                  </div>
                              </div>
                              <div>
                                  <label className="block text-xs font-bold text-gray-400 uppercase mb-2">Valor Aprovado</label>
                                  <div className="relative">
                                      <span className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 font-bold">R$</span>
                                      <input 
                                          type="number" 
                                          name="approvedVal"
                                          value={formData.approvedVal}
                                          onChange={handleInputChange}
                                          className="w-full bg-[#0f1115] border border-gray-700 rounded-lg pl-10 pr-4 py-3 text-emerald-400 font-bold focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                          placeholder="0,00"
                                      />
                                  </div>
                              </div>
                          </div>

                          {/* Instructions / Observations */}
                          <div className="bg-amber-900/10 border border-amber-500/20 rounded-xl p-5">
                              <h4 className="text-amber-500 font-bold text-sm mb-3 flex items-center gap-2">
                                  <AlertCircle size={16} /> OBSERVAÇÕES GERAIS
                              </h4>
                              <ul className="text-xs text-amber-200/70 space-y-1.5 list-disc pl-4">
                                  <li>Preencher corretamente o formulário;</li>
                                  <li>Será analisado o histórico de compra dos últimos 12 meses do PDV;</li>
                                  <li>O pagamento da parceria deverá ser, preferencialmente, em produtos produzidos pela SOLIS através de nota remessa de bonificação emitida pela mesma;</li>
                                  <li>A parceria terá andamento apenas com o formulário preenchido corretamente.</li>
                              </ul>
                          </div>

                      </form>
                  </div>

                  {/* Modal Footer */}
                  <div className="p-6 border-t border-gray-800 bg-[#20232b] flex justify-between items-center">
                      <div className="text-gray-500 text-xs italic font-medium">
                          "Tudo o que se faz junto é melhor!"
                      </div>
                      <div className="flex gap-3">
                          <button 
                              onClick={() => setIsModalOpen(false)}
                              className="px-6 py-2.5 text-gray-400 hover:text-white hover:bg-gray-800 rounded-xl transition-colors font-medium text-sm"
                          >
                              Cancelar
                          </button>
                          <button 
                              form="marketing-request-form"
                              type="submit"
                              className="px-6 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-xl shadow-lg shadow-[#1e5144]/20 transition-all font-medium text-sm flex items-center gap-2"
                          >
                              <Check size={18} /> Enviar Solicitação
                          </button>
                      </div>
                  </div>
              </div>
          </div>
      )}

    </div>
  );
};

export default MarketingBudgetView;
