import React, { useState, useMemo } from 'react';
import { Store, Megaphone, MapPin, Calendar, CheckCircle, XCircle, Clock, Target, Globe, LayoutGrid, List, ArrowLeft, ChevronRight, User, Package, Handshake, Award, MonitorPlay, Tent, MessageSquareText, X, Filter, Plus, Save, Trash2, Edit, History } from 'lucide-react';
import { MOCK_SHOWROOM_ITEMS, REPS_MARKETING_DATA, FULL_MONTH_MAP, OFFLINE_ACTIONS_DATA, MONTHS, REP_NAMES } from '../../../shared/utils/legacy.constants';
import { RepTable } from '../../representatives';
import { PdvTrackingView } from '../../pdv';
import type { OfflineAction, ShowroomItem } from '../../../shared/types';
import type { CooperativeActionsViewProps, RepMarketingAction, SubTabType, OfflineCategory } from '../types';

const OFFLINE_CATEGORIES: OfflineCategory[] = [
  "PARCERIA", 
  "AÇÃO COOPERADA", 
  "ENTREGA DE BRINDES EXCLUSIVOS", 
  "MINIATURAS", 
  "BRINDES - FEIRA", 
  "FEIRA - EXPOSIÇÃO"
];

const CooperativeActionsView: React.FC<CooperativeActionsViewProps> = ({ selectedMonth = '---', onRepClick }) => {
  const [activeSubTab, setActiveSubTab] = useState<SubTabType>('menu');
  
  // Offline Actions State
  const [offlineActions, setOfflineActions] = useState<OfflineAction[]>(OFFLINE_ACTIONS_DATA);
  const [offlineFilterCategory, setOfflineFilterCategory] = useState<OfflineCategory>('TODAS');
  const [offlineFilterRep, setOfflineFilterRep] = useState('TODOS');
  const [observationText, setObservationText] = useState<string | null>(null);

  // Showroom State
  const [showroomItems, setShowroomItems] = useState<ShowroomItem[]>(MOCK_SHOWROOM_ITEMS);
  const [isAddingShowroom, setIsAddingShowroom] = useState(false);
  const [newShowroomForm, setNewShowroomForm] = useState({
      pdv: '',
      city: '',
      contact: '',
      repName: '',
      deliveryForecast: '',
      workshopDate: ''
  });

  // Add/Edit Modal State (Offline)
  const [isAddingOffline, setIsAddingOffline] = useState(false);
  const [editingOfflineId, setEditingOfflineId] = useState<number | null>(null); // Track which index is being edited
  const [newActionForm, setNewActionForm] = useState({
      val: '',
      date: new Date().toISOString().split('T')[0],
      pdv: '',
      rep: '',
      city: '',
      uf: '',
      desc: '', // Obs
      category: 'AÇÃO COOPERADA' as OfflineCategory,
      pedido: '',
      saida: '',
      previsao: '',
      pontuado: 'AINDA NÃO'
  });

  // --- REPS MARKETING ACTIONS STATE ---
  const [marketingHistory, setMarketingHistory] = useState<RepMarketingAction[]>(() => {
      // Initialize history from static REPS_MARKETING_DATA
      const actions: RepMarketingAction[] = [];
      const monthToNum: Record<string, string> = {
          'JANEIRO': '01', 'FEVEREIRO': '02', 'MARÇO': '03', 'ABRIL': '04', 'MAIO': '05', 'JUNHO': '06',
          'JULHO': '07', 'AGOSTO': '08', 'SETEMBRO': '09', 'OUTUBRO': '10', 'NOVEMBRO': '11', 'DEZEMBRO': '12'
      };

      REPS_MARKETING_DATA.rows.forEach((row) => {
          Object.entries(row.values).forEach(([rep, count]) => {
              const c = count as number;
              if (c > 0) {
                  // Create "count" number of entries
                  for(let i=0; i<c; i++) {
                      // Construct a hypothetical date
                      const monthNum = monthToNum[row.month] || '01';
                      const dateStr = `2025-${monthNum}-01`; 
                      
                      actions.push({
                          id: `init-${row.month}-${rep}-${i}`,
                          repName: rep,
                          date: dateStr,
                          description: 'Ação registrada (Histórico)',
                          month: row.month
                      });
                  }
              }
          });
      });
      return actions;
  });

  const [isAddingMarketing, setIsAddingMarketing] = useState(false);
  const [isHistoryMarketingOpen, setIsHistoryMarketingOpen] = useState(false);
  const [editingMarketingId, setEditingMarketingId] = useState<string | null>(null);
  
  const [marketingForm, setMarketingForm] = useState({
      repName: '',
      date: new Date().toISOString().split('T')[0],
      description: ''
  });

  // Filter Marketing Data for Table based on History State
  const computedMarketingTableData = useMemo(() => {
      // Clone structure from constants to keep Targets and Months order
      const dataRows = REPS_MARKETING_DATA.rows.map(row => ({
          ...row,
          values: { ...row.values } // Clone values object
      }));

      // Re-calculate values based on history
      dataRows.forEach(row => {
          REP_NAMES.forEach(rep => {
              // Count actions for this rep in this month
              const count = marketingHistory.filter(h => h.month === row.month && h.repName === rep).length;
              row.values[rep] = count;
          });
      });

      // Filter by selectedMonth if needed
      let filteredRows = dataRows;
      if (selectedMonth && selectedMonth !== '---') {
          const fullMonth = FULL_MONTH_MAP[selectedMonth];
          filteredRows = dataRows.filter(r => r.month === fullMonth);
      }

      return {
          ...REPS_MARKETING_DATA,
          rows: filteredRows
      };
  }, [marketingHistory, selectedMonth]);

  const menuOptions = [
      { 
          id: 'marketing' as SubTabType, 
          title: 'Ações Cooperadas pelos Representantes Solis', 
          description: 'Gestão de iniciativas conjuntas e metas por representante.',
          icon: Target,
          color: 'text-purple-400',
          bg: 'bg-purple-500/10',
          border: 'border-purple-500/20',
          hoverBorder: 'hover:border-purple-500/50'
      },
      { 
          id: 'online_marketing' as SubTabType, 
          title: 'Ações de Marketing Online', 
          description: 'Rastreamento de postagens, presença digital e parcerias.',
          icon: Globe,
          color: 'text-blue-400',
          bg: 'bg-blue-500/10',
          border: 'border-blue-500/20',
          hoverBorder: 'hover:border-blue-500/50'
      },
      { 
          id: 'showroom' as SubTabType, 
          title: 'Showroom Funcional', 
          description: 'Acompanhamento de entrega, instalação e workshops.',
          icon: Store,
          color: 'text-emerald-400',
          bg: 'bg-emerald-500/10',
          border: 'border-emerald-500/20',
          hoverBorder: 'hover:border-emerald-500/50'
      },
      { 
          id: 'offline_marketing' as SubTabType, 
          title: 'Ações de Marketing Offline', 
          description: 'Controle de solicitações, aprovações e entregas.',
          icon: Megaphone,
          color: 'text-amber-400',
          bg: 'bg-amber-500/10',
          border: 'border-amber-500/20',
          hoverBorder: 'hover:border-amber-500/50'
      },
  ];

  const handleBack = () => setActiveSubTab('menu');

  // Helper to categorize offline actions (Simulation based on text content)
  const getOfflineCategory = (action: OfflineAction): OfflineCategory => {
      const text = (action.status + action.pdv).toUpperCase();
      if (text.includes('FEIRA') || text.includes('EXPO')) return 'FEIRA - EXPOSIÇÃO';
      if (text.includes('AMOSTRA') || text.includes('MINIATURA')) return 'MINIATURAS';
      if (text.includes('BRINDE')) return 'ENTREGA DE BRINDES EXCLUSIVOS';
      if (text.includes('PARCERIA')) return 'PARCERIA';
      return 'AÇÃO COOPERADA'; 
  };

  // Extract unique reps for filters
  const uniqueReps = React.useMemo(() => {
      return Array.from(new Set(OFFLINE_ACTIONS_DATA.map(a => a.responsavel || 'Desconhecido'))).sort();
  }, []);

  const filteredOfflineActions = offlineActions.filter(action => {
      const matchesCategory = offlineFilterCategory === 'TODAS' || getOfflineCategory(action) === offlineFilterCategory;
      const matchesRep = offlineFilterRep === 'TODOS' || action.responsavel === offlineFilterRep;
      
      let matchesMonth = true;
      if (selectedMonth !== '---') {
          const parts = action.data.split('/');
          if (parts.length === 3) {
              const monthIndex = parseInt(parts[1], 10) - 1;
              matchesMonth = MONTHS[monthIndex] === selectedMonth;
          } else {
              matchesMonth = false;
          }
      }

      return matchesCategory && matchesRep && matchesMonth;
  });

  const handleAddOfflineAction = (e: React.FormEvent) => {
      e.preventDefault();
      
      const formattedDate = newActionForm.date.split('-').reverse().join('/');
      const formattedValue = parseFloat(newActionForm.val).toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
      
      // Formatting Dates for display (Saída/Previsão)
      const formattedSaida = newActionForm.saida ? newActionForm.saida.split('-').reverse().join('/') : '';
      const formattedPrevisao = newActionForm.previsao ? newActionForm.previsao.split('-').reverse().join('/') : '';

      const actionData: OfflineAction = {
          solicitado: formattedValue,
          data: formattedDate,
          aprovado: '', // Pendente/Vazio
          pedido: newActionForm.pedido,
          saida: formattedSaida,
          previsao: formattedPrevisao,
          entrega: '', // Inicialmente vazio
          cidade: newActionForm.city,
          uf: newActionForm.uf,
          pontuado: newActionForm.pontuado,
          status: newActionForm.desc, // Obs/Status
          pdv: newActionForm.pdv,
          responsavel: newActionForm.rep
      };

      if (editingOfflineId !== null) {
          // Edit existing
          const updated = [...offlineActions];
          const original = updated[editingOfflineId];
          updated[editingOfflineId] = { 
              ...actionData, 
              aprovado: original.aprovado, 
              entrega: original.entrega 
          };
          setOfflineActions(updated);
      } else {
          // Add new
          setOfflineActions([actionData, ...offlineActions]);
      }

      setIsAddingOffline(false);
      setEditingOfflineId(null);
      resetOfflineForm();
  };

  const handleEditOffline = (action: OfflineAction, index: number) => {
      // Helper to convert DD/MM/YYYY to YYYY-MM-DD
      const toIso = (dateStr: string) => {
          if(!dateStr) return '';
          const parts = dateStr.split('/');
          if(parts.length === 3) return `${parts[2]}-${parts[1]}-${parts[0]}`;
          return '';
      };

      // Clean currency string to number string
      const cleanVal = (val: string) => {
          return val.replace('R$', '').replace(/\./g, '').replace(',', '.').trim();
      };

      setNewActionForm({
          val: cleanVal(action.solicitado),
          date: toIso(action.data),
          pdv: action.pdv || '',
          rep: action.responsavel || '',
          city: action.cidade,
          uf: action.uf,
          desc: action.status,
          category: getOfflineCategory(action), // Approximate category
          pedido: action.pedido,
          saida: toIso(action.saida),
          previsao: toIso(action.previsao),
          pontuado: action.pontuado || 'AINDA NÃO'
      });
      setEditingOfflineId(index);
      setIsAddingOffline(true);
  };

  const handleDeleteOffline = (index: number) => {
      if(window.confirm('Tem certeza que deseja excluir este registro?')) {
          const updated = [...offlineActions];
          updated.splice(index, 1);
          setOfflineActions(updated);
      }
  };

  const resetOfflineForm = () => {
      setNewActionForm({
          val: '',
          date: new Date().toISOString().split('T')[0],
          pdv: '',
          rep: '',
          city: '',
          uf: '',
          desc: '',
          category: 'AÇÃO COOPERADA',
          pedido: '',
          saida: '',
          previsao: '',
          pontuado: 'AINDA NÃO'
      });
  };

  // --- MARKETING ACTIONS HANDLERS ---
  const handleSaveMarketing = (e: React.FormEvent) => {
      e.preventDefault();
      
      // Determine Full Month Name from Date
      const [y, m, d] = marketingForm.date.split('-');
      const monthIndex = parseInt(m) - 1;
      const shortMonth = MONTHS[monthIndex];
      const fullMonth = FULL_MONTH_MAP[shortMonth];

      const newAction: RepMarketingAction = {
          id: editingMarketingId || Date.now().toString(),
          repName: marketingForm.repName,
          date: marketingForm.date,
          description: marketingForm.description,
          month: fullMonth
      };

      if (editingMarketingId) {
          setMarketingHistory(prev => prev.map(a => a.id === editingMarketingId ? newAction : a));
      } else {
          setMarketingHistory(prev => [newAction, ...prev]);
      }

      setIsAddingMarketing(false);
      setEditingMarketingId(null);
      setMarketingForm({ repName: '', date: new Date().toISOString().split('T')[0], description: '' });
  };

  const handleEditMarketing = (action: RepMarketingAction) => {
      setMarketingForm({
          repName: action.repName,
          date: action.date,
          description: action.description
      });
      setEditingMarketingId(action.id);
      setIsAddingMarketing(true);
      // If opened from history modal, close history modal or keep it open? 
      // Usually better to close history modal or stack them. Let's close history modal for clarity.
      setIsHistoryMarketingOpen(false); 
  };

  const handleDeleteMarketing = (id: string) => {
      if(window.confirm('Tem certeza que deseja excluir esta ação?')) {
          setMarketingHistory(prev => prev.filter(a => a.id !== id));
      }
  };

  // --- SHOWROOM HANDLERS ---
  const handleAddShowroom = (e: React.FormEvent) => {
      e.preventDefault();
      
      const newItem: ShowroomItem = {
          id: Date.now().toString(),
          pdv: newShowroomForm.pdv,
          city: newShowroomForm.city,
          contact: newShowroomForm.contact,
          repName: newShowroomForm.repName,
          notDelivered: true,
          delivered: false,
          deliveryForecast: newShowroomForm.deliveryForecast ? newShowroomForm.deliveryForecast.split('-').reverse().join('/') : '',
          workshopDate: newShowroomForm.workshopDate ? newShowroomForm.workshopDate.split('-').reverse().join('/') : ''
      };

      setShowroomItems([newItem, ...showroomItems]);
      setIsAddingShowroom(false);
      
      // Reset
      setNewShowroomForm({
          pdv: '',
          city: '',
          contact: '',
          repName: '',
          deliveryForecast: '',
          workshopDate: ''
      });
  };

  const handleDeleteShowroom = (id: string) => {
      if(window.confirm('Tem certeza que deseja excluir este registro?')) {
          setShowroomItems(prev => prev.filter(item => item.id !== id));
      }
  };

  const activeOption = menuOptions.find(opt => opt.id === activeSubTab);
  const ActiveIcon = activeOption ? activeOption.icon : LayoutGrid;

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      
      {/* Header */}
      <div className="mb-6 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-3">
                {activeSubTab !== 'menu' && (
                    <button 
                        onClick={handleBack}
                        className="p-2 -ml-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
                    >
                        <ArrowLeft size={24} />
                    </button>
                )}
                {activeSubTab === 'menu' ? (
                    <>
                        <LayoutGrid className="text-[#1e5144]" />
                        Ações de Marketing
                    </>
                ) : (
                    <>
                        {activeOption && <ActiveIcon className={activeOption.color} />}
                        {activeOption?.title}
                    </>
                )}
            </h1>
            <p className="text-gray-500 mt-1 ml-1">
                {activeSubTab === 'menu' 
                    ? 'Gestão centralizada de iniciativas conjuntas com PDVs'
                    : activeOption?.description}
            </p>
          </div>
      </div>

      {/* MENU VIEW */}
      {activeSubTab === 'menu' && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 animate-in fade-in zoom-in-95 duration-300">
              {menuOptions.map((option) => {
                  const Icon = option.icon;
                  return (
                      <button
                          key={option.id}
                          onClick={() => setActiveSubTab(option.id)}
                          className={`
                              relative overflow-hidden group p-8 rounded-2xl border bg-[#1a1d24] text-left transition-all duration-300 hover:-translate-y-1 hover:shadow-xl
                              ${option.border} ${option.hoverBorder}
                          `}
                      >
                          <div className={`absolute top-0 right-0 p-32 rounded-full blur-3xl opacity-0 group-hover:opacity-10 transition-opacity duration-500 ${option.bg}`} />
                          
                          <div className="relative z-10 flex flex-col h-full">
                              <div className={`w-14 h-14 rounded-2xl flex items-center justify-center mb-6 ${option.bg} ${option.color} border ${option.border}`}>
                                  <Icon size={32} />
                              </div>
                              
                              <h3 className="text-xl font-bold text-gray-100 mb-2 group-hover:text-white transition-colors">
                                  {option.title}
                              </h3>
                              
                              <p className="text-gray-400 text-sm leading-relaxed mb-6">
                                  {option.description}
                              </p>
                              
                              <div className="mt-auto flex items-center gap-2 text-sm font-medium text-gray-500 group-hover:text-white transition-colors">
                                  Acessar Módulo <ChevronRight size={16} className="group-hover:translate-x-1 transition-transform" />
                              </div>
                          </div>
                      </button>
                  );
              })}
          </div>
      )}

      {/* DETAIL VIEWS */}
      <div className="flex-grow flex flex-col min-h-0 relative">
        
        {/* TAB: Online Marketing (Embedded View) */}
        {activeSubTab === 'online_marketing' && (
             <div className="h-full flex flex-col animate-in fade-in slide-in-from-right-8 duration-300">
                <PdvTrackingView />
             </div>
        )}

        {/* TAB: Reps Marketing */}
        {activeSubTab === 'marketing' && (
            <div className="h-full bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden animate-in fade-in slide-in-from-right-8 duration-300 flex flex-col">
                <div className="p-5 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <div>
                        <h3 className="font-bold text-gray-100">Visão Geral</h3>
                        <p className="text-xs text-gray-500">Ações consolidadas</p>
                    </div>
                    <div className="flex gap-2">
                        <button 
                            onClick={() => setIsHistoryMarketingOpen(true)}
                            className="px-4 py-2 bg-[#1a1d24] border border-gray-700 hover:border-gray-600 text-gray-300 rounded-lg text-sm font-medium transition-colors flex items-center gap-2"
                        >
                            <History size={16} /> Histórico
                        </button>
                        <button 
                            onClick={() => { setEditingMarketingId(null); setMarketingForm({ repName: '', date: new Date().toISOString().split('T')[0], description: '' }); setIsAddingMarketing(true); }}
                            className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/20"
                        >
                            <Plus size={16} /> Novo Lançamento
                        </button>
                    </div>
                </div>

                {computedMarketingTableData.rows.length > 0 ? (
                    <RepTable 
                        data={computedMarketingTableData} 
                        onRepClick={onRepClick}
                    />
                ) : (
                    <div className="h-full flex flex-col items-center justify-center text-gray-500 p-8">
                        <List size={48} className="mb-4 opacity-20" />
                        <p>Sem dados de ações para {FULL_MONTH_MAP[selectedMonth]}.</p>
                    </div>
                )}
            </div>
        )}

        {/* TAB: Showroom */}
        {activeSubTab === 'showroom' && (
            <div className="h-full bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden flex flex-col animate-in fade-in slide-in-from-right-8 duration-300">
                <div className="p-5 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <div className="flex gap-2 text-xs">
                        <div className="flex items-center gap-1.5 px-3 py-1.5 bg-[#0f1115] rounded-lg border border-gray-800 text-gray-400">
                            <div className="w-2 h-2 rounded-full bg-emerald-500"></div> Entregue
                        </div>
                        <div className="flex items-center gap-1.5 px-3 py-1.5 bg-[#0f1115] rounded-lg border border-gray-800 text-gray-400">
                            <div className="w-2 h-2 rounded-full bg-rose-500"></div> Pendente
                        </div>
                    </div>
                    <button 
                        onClick={() => setIsAddingShowroom(true)}
                        className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/20"
                    >
                        <Plus size={16} /> Novo Registro
                    </button>
                </div>
                
                <div className="overflow-x-auto custom-scrollbar flex-grow p-4">
                    <div className="min-w-[1000px] border border-gray-800 rounded-xl overflow-hidden">
                        <table className="w-full text-sm text-left">
                            <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800">
                                <tr>
                                    <th className="px-6 py-4">PDV / Parceiro</th>
                                    <th className="px-6 py-4">Localização</th>
                                    <th className="px-6 py-4">Representante</th>
                                    <th className="px-6 py-4 text-center">Status Entrega</th>
                                    <th className="px-6 py-4 text-center">Previsão</th>
                                    <th className="px-6 py-4 text-center">Workshop</th>
                                    <th className="px-6 py-4 text-center">Ações</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-800 bg-[#1a1d24]">
                                {showroomItems.map((item) => (
                                    <tr key={item.id} className="hover:bg-[#20232b] transition-colors group">
                                        <td className="px-6 py-4">
                                            <div className="font-medium text-gray-200">{item.pdv}</div>
                                            <div className="text-xs text-gray-500 mt-0.5">Contato: {item.contact}</div>
                                        </td>
                                        <td className="px-6 py-4 text-gray-400">
                                            <span className="flex items-center gap-1.5 bg-[#0f1115] px-2 py-1 rounded w-fit text-xs border border-gray-800">
                                                <MapPin size={10} className="text-gray-500" />
                                                {item.city}
                                            </span>
                                        </td>
                                        <td className="px-6 py-4">
                                            <button 
                                                onClick={() => onRepClick && onRepClick(item.repName)}
                                                className="font-medium text-gray-200 hover:text-emerald-400 hover:underline transition-colors text-left"
                                            >
                                                {item.repName}
                                            </button>
                                        </td>
                                        
                                        <td className="px-6 py-4 text-center">
                                            {item.delivered ? (
                                                <span className="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 text-xs font-medium">
                                                    <CheckCircle size={12} /> Entregue
                                                </span>
                                            ) : (
                                                <span className="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-rose-500/10 text-rose-400 border border-rose-500/20 text-xs font-medium">
                                                    <Clock size={12} /> Pendente
                                                </span>
                                            )}
                                        </td>
                                        
                                        <td className="px-6 py-4 text-center">
                                            <span className={`text-xs ${item.deliveryForecast ? 'text-gray-300' : 'text-gray-600 italic'}`}>
                                                {item.deliveryForecast || 'Sem previsão'}
                                            </span>
                                        </td>
                                        
                                        <td className="px-6 py-4 text-center">
                                            {item.workshopDate ? (
                                                <span className="inline-flex items-center gap-1.5 px-2.5 py-1 bg-blue-500/10 text-blue-400 rounded-lg text-xs font-medium border border-blue-500/20">
                                                    <Calendar size={12} />
                                                    {item.workshopDate}
                                                </span>
                                            ) : (
                                                <span className="text-gray-600 text-[10px] uppercase font-bold tracking-wider">A Agendar</span>
                                            )}
                                        </td>

                                        <td className="px-6 py-4 text-center">
                                            <button 
                                                onClick={() => handleDeleteShowroom(item.id)}
                                                className="p-1.5 text-gray-500 hover:text-rose-400 hover:bg-rose-500/10 rounded-lg transition-colors"
                                                title="Excluir Registro"
                                            >
                                                <Trash2 size={16} />
                                            </button>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        )}

        {/* TAB: Offline Marketing */}
        {activeSubTab === 'offline_marketing' && (
            <div className="h-full bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden flex flex-col animate-in fade-in slide-in-from-right-8 duration-300">
                <div className="p-5 border-b border-gray-800 bg-[#20232b] flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
                    <div className="flex flex-wrap gap-4">
                        {/* Category Filter */}
                        <div className="relative">
                            <select
                                value={offlineFilterCategory}
                                onChange={(e) => setOfflineFilterCategory(e.target.value as OfflineCategory)}
                                className="appearance-none bg-[#15171c] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-xs font-medium focus:outline-none focus:ring-1 focus:ring-[#1e5144] hover:border-gray-600 transition-colors"
                            >
                                <option value="TODAS">Todas Categorias</option>
                                {OFFLINE_CATEGORIES.map(cat => <option key={cat} value={cat}>{cat}</option>)}
                            </select>
                            <Filter size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                        </div>

                        {/* Rep Filter */}
                        <div className="relative">
                            <select
                                value={offlineFilterRep}
                                onChange={(e) => setOfflineFilterRep(e.target.value)}
                                className="appearance-none bg-[#15171c] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-xs font-medium focus:outline-none focus:ring-1 focus:ring-[#1e5144] hover:border-gray-600 transition-colors"
                            >
                                <option value="TODOS">Todos Representantes</option>
                                {uniqueReps.map(rep => <option key={rep} value={rep}>{rep}</option>)}
                            </select>
                            <User size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                        </div>
                    </div>

                    <button 
                        onClick={() => { resetOfflineForm(); setEditingOfflineId(null); setIsAddingOffline(true); }}
                        className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/20"
                    >
                        <Plus size={16} /> Novo Registro
                    </button>
                </div>
                
                <div className="overflow-x-auto custom-scrollbar flex-grow p-4">
                    <div className="border border-gray-800 rounded-xl overflow-hidden min-w-[1200px]">
                        <table className="w-full text-xs text-left whitespace-nowrap">
                            <thead className="bg-[#15171c] text-gray-400 uppercase font-semibold border-b border-gray-800">
                                <tr>
                                    <th className="px-4 py-3 bg-[#15171c] sticky left-0 z-10 border-r border-gray-800 shadow-lg">Solicitação</th>
                                    <th className="px-4 py-3 text-center">Categoria</th>
                                    <th className="px-4 py-3 text-center">Status</th>
                                    <th className="px-4 py-3">Aprovado</th>
                                    <th className="px-4 py-3">PDV</th>
                                    <th className="px-4 py-3">Representante</th>
                                    <th className="px-4 py-3">Pedido</th>
                                    <th className="px-4 py-3">Saída / Prev.</th>
                                    <th className="px-4 py-3 text-center">Pontuado</th>
                                    <th className="px-4 py-3 text-center">Obs.</th>
                                    <th className="px-4 py-3 text-center">Conclusão</th>
                                    <th className="px-4 py-3 text-center">Ações</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-800 bg-[#1a1d24]">
                                {filteredOfflineActions.length > 0 ? (
                                    filteredOfflineActions.map((action, index) => (
                                        <tr key={index} className="hover:bg-[#20232b] transition-colors group">
                                            <td className="px-4 py-3 text-white font-medium bg-[#1a1d24] group-hover:bg-[#20232b] sticky left-0 z-10 border-r border-gray-800 shadow-sm">
                                                {action.solicitado}
                                                <span className="block text-[10px] text-gray-500 font-normal mt-0.5">{action.data}</span>
                                            </td>
                                            
                                            <td className="px-4 py-3 text-center">
                                                <span className="inline-block text-[10px] font-bold text-gray-400 bg-gray-800/50 px-2 py-1 rounded border border-gray-700/50 whitespace-nowrap">
                                                    {getOfflineCategory(action)}
                                                </span>
                                            </td>

                                            <td className="px-4 py-3 text-center">
                                                {action.aprovado ? (
                                                    <div title="Aprovado">
                                                        <CheckCircle size={16} className="text-emerald-500 mx-auto" />
                                                    </div>
                                                ) : (
                                                    <div title="Em Análise">
                                                        <Clock size={16} className="text-amber-500 mx-auto" />
                                                    </div>
                                                )}
                                            </td>

                                            <td className="px-4 py-3 text-emerald-400 font-medium font-mono">
                                                {action.aprovado || 'Análise'}
                                            </td>

                                            <td className="px-4 py-3 text-gray-200 font-medium">
                                                <div className="flex flex-col gap-1">
                                                    <div className="block">
                                                        {action.pdv || '-'}
                                                    </div>
                                                    {action.cidade && (
                                                        <span className="flex items-center gap-1.5 bg-[#0f1115] px-2 py-0.5 rounded w-fit text-[10px] border border-gray-800 text-gray-400">
                                                            <MapPin size={10} className="text-gray-600" />
                                                            {action.cidade} {action.uf && <span className="text-gray-600 ml-0.5">({action.uf})</span>}
                                                        </span>
                                                    )}
                                                </div>
                                            </td>

                                            <td className="px-4 py-3">
                                                {action.responsavel ? (
                                                    <button 
                                                        onClick={() => onRepClick && onRepClick(action.responsavel!)}
                                                        className="flex items-center gap-2 text-gray-300 hover:text-emerald-400 hover:underline transition-colors"
                                                    >
                                                        <User size={14} className="text-gray-500" />
                                                        {action.responsavel}
                                                    </button>
                                                ) : (
                                                    <span className="text-gray-600">-</span>
                                                )}
                                            </td>
                                            
                                            <td className="px-4 py-3 text-gray-300">
                                                {action.pedido ? (
                                                    <span className="font-mono bg-black/30 px-1.5 py-0.5 rounded border border-gray-700 text-[11px]">#{action.pedido}</span>
                                                ) : '-'}
                                            </td>
                                            
                                            <td className="px-4 py-3">
                                                <div className="flex flex-col gap-1 text-[10px]">
                                                    {action.saida ? <span className="text-gray-300">Saída: {action.saida}</span> : <span className="text-gray-600">-</span>}
                                                    {action.previsao && <span className="text-blue-400">Prev: {action.previsao}</span>}
                                                </div>
                                            </td>
                                            
                                            <td className="px-4 py-3 text-center">
                                                {action.pontuado === 'SIM' ? (
                                                    <span className="text-emerald-500 text-[10px] font-bold">SIM</span>
                                                ) : action.pontuado === 'AINDA NÃO' ? (
                                                    <span className="text-amber-500 text-[10px] font-bold">NÃO</span>
                                                ) : (
                                                    <span className="text-gray-700">-</span>
                                                )}
                                            </td>
                                            
                                            <td className="px-4 py-3 text-center">
                                                {action.status ? (
                                                    <button
                                                        onClick={() => setObservationText(action.status)}
                                                        className="text-gray-400 hover:text-blue-400 transition-colors p-1.5 hover:bg-blue-500/10 rounded-lg"
                                                        title="Ler Observação"
                                                    >
                                                        <MessageSquareText size={16} />
                                                    </button>
                                                ) : (
                                                    <span className="text-gray-700">-</span>
                                                )}
                                            </td>

                                            <td className="px-4 py-3 text-center">
                                                {action.entrega && /\d{2}\/\d{2}\/\d{4}/.test(action.entrega) ? (
                                                    <span className="text-emerald-400 font-mono text-[11px] font-medium bg-emerald-500/10 px-2 py-0.5 rounded border border-emerald-500/20">
                                                        {action.entrega.match(/\d{2}\/\d{2}\/\d{4}/)![0]}
                                                    </span>
                                                ) : (
                                                    <span className="text-gray-700">-</span>
                                                )}
                                            </td>

                                            <td className="px-4 py-3 text-center">
                                                <div className="flex items-center gap-1 justify-center">
                                                    <button 
                                                        onClick={() => handleEditOffline(action, index)}
                                                        className="p-1.5 text-gray-500 hover:text-emerald-400 hover:bg-emerald-500/10 rounded-lg transition-colors"
                                                        title="Editar"
                                                    >
                                                        <Edit size={14} />
                                                    </button>
                                                    <button 
                                                        onClick={() => handleDeleteOffline(index)}
                                                        className="p-1.5 text-gray-500 hover:text-rose-400 hover:bg-rose-500/10 rounded-lg transition-colors"
                                                        title="Excluir"
                                                    >
                                                        <Trash2 size={14} />
                                                    </button>
                                                </div>
                                            </td>
                                        </tr>
                                    ))
                                ) : (
                                    <tr>
                                        <td colSpan={12} className="px-6 py-12 text-center text-gray-500">
                                            <div className="flex flex-col items-center justify-center gap-2">
                                                <Target size={32} className="opacity-20" />
                                                <p>Nenhuma ação encontrada com os filtros selecionados.</p>
                                            </div>
                                        </td>
                                    </tr>
                                )}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        )}

      </div>

      {/* Observation Modal */}
      {observationText && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm" onClick={() => setObservationText(null)}>
            <div className="bg-[#1a1d24] w-full max-w-lg rounded-2xl border border-gray-700 shadow-2xl overflow-hidden animate-in zoom-in-95" onClick={e => e.stopPropagation()}>
                <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <MessageSquareText size={18} className="text-blue-400" />
                        Observações
                    </h3>
                    <button onClick={() => setObservationText(null)} className="text-gray-500 hover:text-white transition-colors">
                        <X size={20} />
                    </button>
                </div>
                <div className="p-6 text-gray-300 text-sm leading-relaxed whitespace-pre-wrap">
                    {observationText}
                </div>
            </div>
        </div>
      )}

      {/* ADD/EDIT MARKETING ACTION MODAL */}
      {isAddingMarketing && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm" onClick={() => setIsAddingMarketing(false)}>
            <div className="bg-[#1a1d24] w-full max-w-md rounded-2xl border border-gray-700 shadow-2xl overflow-hidden animate-in zoom-in-95" onClick={e => e.stopPropagation()}>
                <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <Target size={18} className="text-emerald-400" />
                        {editingMarketingId ? 'Editar Lançamento' : 'Novo Lançamento'}
                    </h3>
                    <button onClick={() => setIsAddingMarketing(false)} className="text-gray-500 hover:text-white transition-colors">
                        <X size={20} />
                    </button>
                </div>
                
                <form onSubmit={handleSaveMarketing} className="p-6 space-y-4">
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Representante</label>
                        <select 
                            value={marketingForm.repName}
                            onChange={(e) => setMarketingForm({...marketingForm, repName: e.target.value})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        >
                            <option value="" disabled>Selecione...</option>
                            {REP_NAMES.map(r => <option key={r} value={r}>{r}</option>)}
                        </select>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Data da Ação</label>
                        <input 
                            type="date" 
                            value={marketingForm.date}
                            onChange={(e) => setMarketingForm({...marketingForm, date: e.target.value})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        />
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Descrição</label>
                        <textarea 
                            value={marketingForm.description}
                            onChange={(e) => setMarketingForm({...marketingForm, description: e.target.value})}
                            placeholder="Detalhes da ação realizada..."
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144] h-24 resize-none"
                            required
                        />
                    </div>

                    <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
                        <button 
                            type="button" 
                            onClick={() => setIsAddingMarketing(false)}
                            className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                        >
                            Cancelar
                        </button>
                        <button 
                            type="submit" 
                            className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2"
                        >
                            <Save size={16} /> Salvar
                        </button>
                    </div>
                </form>
            </div>
        </div>
      )}

      {/* MARKETING HISTORY MODAL */}
      {isHistoryMarketingOpen && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm" onClick={() => setIsHistoryMarketingOpen(false)}>
            <div className="bg-[#1a1d24] w-full max-w-4xl rounded-2xl border border-gray-700 shadow-2xl overflow-hidden animate-in zoom-in-95 flex flex-col max-h-[80vh]" onClick={e => e.stopPropagation()}>
                <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <History size={18} className="text-gray-400" />
                        Histórico de Lançamentos
                    </h3>
                    <button onClick={() => setIsHistoryMarketingOpen(false)} className="text-gray-500 hover:text-white transition-colors">
                        <X size={20} />
                    </button>
                </div>
                
                <div className="flex-grow overflow-y-auto custom-scrollbar p-6">
                    <table className="w-full text-sm text-left">
                        <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
                            <tr>
                                <th className="px-4 py-3">Data</th>
                                <th className="px-4 py-3">Mês (Ref)</th>
                                <th className="px-4 py-3">Representante</th>
                                <th className="px-4 py-3 w-1/3">Descrição</th>
                                <th className="px-4 py-3 text-center">Ações</th>
                            </tr>
                        </thead>
                        <tbody className="divide-y divide-gray-800">
                            {marketingHistory.length > 0 ? (
                                marketingHistory
                                    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
                                    .map((action) => (
                                    <tr key={action.id} className="hover:bg-[#20232b] transition-colors">
                                        <td className="px-4 py-3 text-gray-300 whitespace-nowrap">
                                            {action.date.split('-').reverse().join('/')}
                                        </td>
                                        <td className="px-4 py-3 text-gray-400 text-xs">
                                            {action.month}
                                        </td>
                                        <td className="px-4 py-3 text-gray-200 font-medium">
                                            {action.repName}
                                        </td>
                                        <td className="px-4 py-3 text-gray-400 text-xs">
                                            {action.description}
                                        </td>
                                        <td className="px-4 py-3 text-center">
                                            <div className="flex items-center justify-center gap-2">
                                                <button 
                                                    onClick={() => handleEditMarketing(action)}
                                                    className="p-1.5 text-gray-500 hover:text-emerald-400 hover:bg-emerald-500/10 rounded-lg transition-colors"
                                                    title="Editar"
                                                >
                                                    <Edit size={14} />
                                                </button>
                                                <button 
                                                    onClick={() => handleDeleteMarketing(action.id)}
                                                    className="p-1.5 text-gray-500 hover:text-rose-400 hover:bg-rose-500/10 rounded-lg transition-colors"
                                                    title="Excluir"
                                                >
                                                    <Trash2 size={14} />
                                                </button>
                                            </div>
                                        </td>
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan={5} className="px-6 py-8 text-center text-gray-500">
                                        Nenhum registro encontrado.
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
      )}

      {/* ADD/EDIT OFFLINE ACTION MODAL */}
      {isAddingOffline && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm" onClick={() => setIsAddingOffline(false)}>
            <div className="bg-[#1a1d24] w-full max-w-lg rounded-2xl border border-gray-700 shadow-2xl overflow-hidden animate-in zoom-in-95" onClick={e => e.stopPropagation()}>
                <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <Megaphone size={18} className="text-amber-400" />
                        {editingOfflineId !== null ? 'Editar Ação Offline' : 'Nova Ação Offline'}
                    </h3>
                    <button onClick={() => setIsAddingOffline(false)} className="text-gray-500 hover:text-white transition-colors">
                        <X size={20} />
                    </button>
                </div>
                
                <form onSubmit={handleAddOfflineAction} className="p-6 space-y-4 max-h-[70vh] overflow-y-auto custom-scrollbar">
                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Data da Solicitação</label>
                            <input 
                                type="date" 
                                value={newActionForm.date}
                                onChange={(e) => setNewActionForm({...newActionForm, date: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                required
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Valor Solicitado (R$)</label>
                            <input 
                                type="number" 
                                step="0.01"
                                placeholder="0,00"
                                value={newActionForm.val}
                                onChange={(e) => setNewActionForm({...newActionForm, val: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                required
                            />
                        </div>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Representante Responsável</label>
                        <select 
                            value={newActionForm.rep}
                            onChange={(e) => setNewActionForm({...newActionForm, rep: e.target.value})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        >
                            <option value="" disabled>Selecione...</option>
                            {REP_NAMES.map(r => <option key={r} value={r}>{r}</option>)}
                        </select>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">PDV / Cliente</label>
                        <input 
                            type="text" 
                            placeholder="Nome do estabelecimento"
                            value={newActionForm.pdv}
                            onChange={(e) => setNewActionForm({...newActionForm, pdv: e.target.value})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        />
                    </div>

                    <div className="grid grid-cols-3 gap-4">
                        <div className="col-span-2">
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Cidade</label>
                            <input 
                                type="text" 
                                placeholder="Ex: Campinas"
                                value={newActionForm.city}
                                onChange={(e) => setNewActionForm({...newActionForm, city: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">UF</label>
                            <input 
                                type="text" 
                                placeholder="SP"
                                maxLength={2}
                                value={newActionForm.uf}
                                onChange={(e) => setNewActionForm({...newActionForm, uf: e.target.value.toUpperCase()})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Categoria da Ação</label>
                        <select 
                            value={newActionForm.category}
                            onChange={(e) => setNewActionForm({...newActionForm, category: e.target.value as OfflineCategory})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                        >
                            {OFFLINE_CATEGORIES.map(cat => <option key={cat} value={cat}>{cat}</option>)}
                        </select>
                    </div>

                    {/* New Fields Requested */}
                    <div className="grid grid-cols-2 gap-4 border-t border-gray-800 pt-4">
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Pedido (Nº)</label>
                            <input 
                                type="text" 
                                placeholder="Ex: 46372"
                                value={newActionForm.pedido}
                                onChange={(e) => setNewActionForm({...newActionForm, pedido: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Pontuado</label>
                            <select 
                                value={newActionForm.pontuado}
                                onChange={(e) => setNewActionForm({...newActionForm, pontuado: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            >
                                <option value="AINDA NÃO">AINDA NÃO</option>
                                <option value="SIM">SIM</option>
                                <option value="NÃO">NÃO</option>
                            </select>
                        </div>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Data Saída</label>
                            <input 
                                type="date" 
                                value={newActionForm.saida}
                                onChange={(e) => setNewActionForm({...newActionForm, saida: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Previsão Entrega</label>
                            <input 
                                type="date" 
                                value={newActionForm.previsao}
                                onChange={(e) => setNewActionForm({...newActionForm, previsao: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Observações (Status)</label>
                        <textarea 
                            value={newActionForm.desc}
                            onChange={(e) => setNewActionForm({...newActionForm, desc: e.target.value})}
                            placeholder="Detalhes sobre a ação, contrapartidas, etc."
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144] h-20 resize-none"
                        />
                    </div>

                    <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
                        <button 
                            type="button" 
                            onClick={() => setIsAddingOffline(false)}
                            className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                        >
                            Cancelar
                        </button>
                        <button 
                            type="submit" 
                            className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2"
                        >
                            <Save size={16} /> Salvar Solicitação
                        </button>
                    </div>
                </form>
            </div>
        </div>
      )}

      {/* ADD SHOWROOM MODAL */}
      {isAddingShowroom && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm" onClick={() => setIsAddingShowroom(false)}>
            <div className="bg-[#1a1d24] w-full max-w-lg rounded-2xl border border-gray-700 shadow-2xl overflow-hidden animate-in zoom-in-95" onClick={e => e.stopPropagation()}>
                <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <Store size={18} className="text-emerald-400" />
                        Novo Showroom
                    </h3>
                    <button onClick={() => setIsAddingShowroom(false)} className="text-gray-500 hover:text-white transition-colors">
                        <X size={20} />
                    </button>
                </div>
                
                <form onSubmit={handleAddShowroom} className="p-6 space-y-4">
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">PDV / Parceiro</label>
                        <input 
                            type="text" 
                            placeholder="Nome da loja"
                            value={newShowroomForm.pdv}
                            onChange={(e) => setNewShowroomForm({...newShowroomForm, pdv: e.target.value})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        />
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Responsável</label>
                        <select 
                            value={newShowroomForm.repName}
                            onChange={(e) => setNewShowroomForm({...newShowroomForm, repName: e.target.value})}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            required
                        >
                            <option value="" disabled>Selecione...</option>
                            {REP_NAMES.map(r => <option key={r} value={r}>{r}</option>)}
                        </select>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Cidade</label>
                            <input 
                                type="text" 
                                placeholder="Cidade - UF"
                                value={newShowroomForm.city}
                                onChange={(e) => setNewShowroomForm({...newShowroomForm, city: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                required
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Contato</label>
                            <input 
                                type="text" 
                                placeholder="Nome do contato"
                                value={newShowroomForm.contact}
                                onChange={(e) => setNewShowroomForm({...newShowroomForm, contact: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Previsão Entrega</label>
                            <input 
                                type="date" 
                                value={newShowroomForm.deliveryForecast}
                                onChange={(e) => setNewShowroomForm({...newShowroomForm, deliveryForecast: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Data Workshop</label>
                            <input 
                                type="date" 
                                value={newShowroomForm.workshopDate}
                                onChange={(e) => setNewShowroomForm({...newShowroomForm, workshopDate: e.target.value})}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            />
                        </div>
                    </div>

                    <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
                        <button 
                            type="button" 
                            onClick={() => setIsAddingShowroom(false)}
                            className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                        >
                            Cancelar
                        </button>
                        <button 
                            type="submit" 
                            className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2"
                        >
                            <Save size={16} /> Salvar Showroom
                        </button>
                    </div>
                </form>
            </div>
        </div>
      )}

    </div>
  );
};

export default CooperativeActionsView;
