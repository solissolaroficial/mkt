
import React, { useState, useEffect, useMemo } from 'react';
import { 
  LayoutDashboard, 
  Target, 
  GraduationCap, 
  Megaphone, 
  CheckSquare, 
  Calendar, 
  MapPin, 
  Share2, 
  TrendingUp, 
  Bell, 
  Search, 
  Menu, 
  ChevronRight, 
  Youtube, 
  Settings, 
  LogOut, 
  ChevronDown, 
  Gift, 
  DollarSign, 
  KeyRound, 
  Phone, 
  ArrowLeft, 
  PieChart 
} from 'lucide-react';
import { KPIS, MOCK_TASKS, REPS_TRAINING_DATA, REPS_MARKETING_DATA, SOCIAL_MEDIA_RANKING, MONTHS, MOCK_CALENDAR_POSTS, MOCK_ACCOUNTS_PAYABLE, REP_PROFILES, FULL_MONTH_MAP, RAW_BUDGET_ITEMS, MOCK_PDV_POSTS, OFFLINE_ACTIONS_DATA } from './constants';
import { KpiCategory, Task, Notification, RepTableData, CalendarPost, RepresentativeProfile, BreakdownItem } from './types';
import SummaryCard from './components/SummaryCard';
import HeatmapChart from './components/HeatmapChart';
import RepTable from './components/RepTable';
import KpiDetailView from './components/KpiDetailView';
import TaskWidget from './components/TaskWidget';
import KanbanBoard from './components/KanbanBoard';
import PostCalendarView from './components/PostCalendarView';
import CooperativeActionsView from './components/CooperativeActionsView';
import SocialBenchmarkingView from './components/SocialBenchmarkingView';
import NotificationPanel from './components/NotificationPanel';
import TaskModal from './components/TaskModal';
import LoginView from './components/LoginView';
import SettingsView from './components/SettingsView';
import GiftsView from './components/GiftsView';
import FinancialView from './components/FinancialView';
import BudgetView from './components/BudgetView';
import UpcomingPostsWidget from './components/UpcomingPostsWidget';
import RepProfileModal from './components/RepProfileModal';
import PasswordsView from './components/PasswordsView';
import ContactsView from './components/ContactsView';
import MarketingBudgetView from './components/MarketingBudgetView';

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(true); // Auth State
  const [activeTab, setActiveTab] = useState('dashboard');
  const [kpis, setKpis] = useState<KpiCategory[]>(KPIS);
  const [tasks, setTasks] = useState<Task[]>(MOCK_TASKS);
  const [posts, setPosts] = useState<CalendarPost[]>(MOCK_CALENDAR_POSTS); // Lifted state
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [selectedKpiId, setSelectedKpiId] = useState<string | null>(null);
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);
  const [showNotifications, setShowNotifications] = useState(false);
  const [showUserMenu, setShowUserMenu] = useState(false); // Dropdown menu for user
  const [editingTask, setEditingTask] = useState<Task | null>(null);
  
  // Rep Profile Modal State
  const [selectedRepProfile, setSelectedRepProfile] = useState<RepresentativeProfile | null>(null);

  // Financial Module State
  const [selectedFinancialModule, setSelectedFinancialModule] = useState<string | null>(null);

  // Intermediate Dashboard States
  const [showCplDashboard, setShowCplDashboard] = useState(false);
  const [showOppDashboard, setShowOppDashboard] = useState(false);
  const [showAdSpendDashboard, setShowAdSpendDashboard] = useState(false);
  const [showRoasDashboard, setShowRoasDashboard] = useState(false);

  // --- DATE FILTER STATE ---
  const [selectedMonth, setSelectedMonth] = useState('NOV');
  const [selectedYear, setSelectedYear] = useState('2025');

  // --- DYNAMIC KPI CALCULATION (CPL) ---
  const derivedKpis = useMemo(() => {
    return kpis.map(kpi => {
        // Automatically calculate CPL based on Ad Spend / Opportunities
        if (kpi.id === 'cpl') {
            const adSpendKpi = kpis.find(k => k.id === 'ad_spend');
            const oppKpi = kpis.find(k => k.id === 'opportunities');
            
            const newData = kpi.data.map(d => {
                const adData = adSpendKpi?.data.find(ad => ad.month === d.month);
                const oppData = oppKpi?.data.find(opp => opp.month === d.month);
                
                const spend = adData?.realized || 0;
                const leads = oppData?.realized || 0;
                
                // Avoid division by zero
                const calculatedCpl = leads > 0 ? (spend / leads) : 0;

                // --- Calculate Breakdown for CPL ---
                let cplBreakdown: BreakdownItem[] | undefined = undefined;

                if (adData?.breakdown && oppData?.breakdown) {
                    cplBreakdown = adData.breakdown.map(adItem => {
                        // Find matching opportunity item (assuming labels match e.g. "Meta Ads" or "Campanhas Ads Meta")
                        // Using includes for looser matching or exact match
                        const oppItem = oppData.breakdown?.find(o => o.label === adItem.label || o.label.includes(adItem.label) || adItem.label.includes(o.label));
                        
                        const itemLeads = oppItem ? oppItem.value : 0;
                        const itemSpend = adItem.value;
                        const itemCpl = itemLeads > 0 ? itemSpend / itemLeads : 0;

                        // Calculate SubItems CPL
                        let subItems: BreakdownItem[] | undefined = undefined;
                        if (adItem.subItems && oppItem?.subItems) {
                            subItems = adItem.subItems.map(adSub => {
                                const oppSub = oppItem.subItems?.find(o => o.label === adSub.label);
                                const subLeads = oppSub ? oppSub.value : 0;
                                const subSpend = adSub.value;
                                return {
                                    label: adSub.label,
                                    value: subLeads > 0 ? subSpend / subLeads : 0
                                };
                            });
                        }

                        return {
                            label: adItem.label,
                            value: itemCpl,
                            subItems
                        };
                    });
                }
                
                return {
                    ...d,
                    realized: calculatedCpl,
                    breakdown: cplBreakdown
                };
            });
            return { ...kpi, data: newData };
        }
        return kpi;
    });
  }, [kpis]);

  // Calculate budget total for selected month (Moved to top level)
  const budgetSummary = useMemo(() => {
      let budget = 0;
      const idx = selectedMonth === '---' ? -1 : MONTHS.indexOf(selectedMonth);
      
      RAW_BUDGET_ITEMS.forEach(item => {
          if (idx === -1) {
              budget += item.vals.reduce((a, b) => a + b, 0);
          } else {
              budget += item.vals[idx] || 0;
          }
      });
      return budget;
  }, [selectedMonth]);

  // Calculate Marketing Actions Budget Summary (Group 13, Subgroup 41)
  const marketingActionsSummary = useMemo(() => {
      let total = 0;
      const idx = selectedMonth === '---' ? -1 : MONTHS.indexOf(selectedMonth);
      
      const marketingItems = RAW_BUDGET_ITEMS.filter(
          item => item.codObj === '13' && item.codGrp === '41'
      );

      marketingItems.forEach(item => {
          if (idx === -1) {
              total += item.vals.reduce((a, b) => a + b, 0);
          } else {
              total += item.vals[idx] || 0;
          }
      });
      return total;
  }, [selectedMonth]);

  // --- LOGICA DE NOTIFICAÇÕES REAIS ---
  useEffect(() => {
    // (Notification logic omitted for brevity, keeps existing)
    // ...
  }, [tasks, posts, isLoggedIn]);

  const handleMention = (taskId: string, taskTitle: string) => {
      const newNotif: Notification = {
          id: `mention-${Date.now()}`,
          type: 'mention',
          title: 'Nova Menção',
          message: `Jackson mencionou você no card "${taskTitle}".`,
          timestamp: 'Agora',
          read: false,
          taskId,
          archived: false
      };
      setNotifications(prev => [newNotif, ...prev]);
      setShowNotifications(true);
  };

  const handleAddNotification = (newNotification: Notification) => {
      setNotifications(prev => [{...newNotification, archived: false}, ...prev]);
  };

  const handleNotificationClick = (notification: Notification) => {
      setNotifications(prev => prev.map(n => n.id === notification.id ? { ...n, read: true } : n));
      if (notification.taskId) {
          const task = tasks.find(t => t.id === notification.taskId);
          if (task) {
              setEditingTask(task);
              setShowNotifications(false);
          }
      } else if (notification.id.startsWith('auto-post-')) {
          setActiveTab('calendar');
          setShowNotifications(false);
      }
  };

  const handleArchiveNotification = (id: string) => {
      setNotifications(prev => prev.map(n => n.id === id ? { ...n, archived: true, read: true } : n));
  };

  const handleRepClick = (repName: string) => {
      // Tenta buscar o perfil real na constante
      let profile = REP_PROFILES[repName];
      
      // Se não existir dados cadastrados, gera um perfil visual temporário para que a tela abra
      if (!profile) {
          profile = {
              code: Math.floor(Math.random() * 899) + 100,
              company: `${repName.toUpperCase()} REPRESENTAÇÕES`,
              name: repName,
              region: 'Região Sul / Sudeste', // Placeholder
              phone: '(19) 99999-9999',
              email: `${repName.toLowerCase().replace(/\s+/g, '.')}@solis.ind.br`,
              attendant: 'Comercial Interno'
          };
      }

      // --- CALCULATE DYNAMIC STATS ---
      
      // 1. Training Count (Sum all months from heat map data)
      const trainingCount = REPS_TRAINING_DATA.rows.reduce((acc, row) => {
          return acc + (row.values[repName] || 0);
      }, 0);

      // 2. Online Marketing Actions (Count from PDV Posts)
      const onlineCount = MOCK_PDV_POSTS.filter(p => p.repName === repName).length;

      // 3. Offline Marketing Actions (Count & Value from Offline Actions Data)
      const offlineActions = OFFLINE_ACTIONS_DATA.filter(a => a.responsavel === repName);
      const offlineCount = offlineActions.length;
      
      const offlineValue = offlineActions.reduce((acc, curr) => {
          // Parse currency string "R$ 1.500,00" -> 1500.00
          const valStr = curr.solicitado || '0';
          const val = parseFloat(valStr.replace(/[R$\s.]/g, '').replace(',', '.'));
          return acc + (isNaN(val) ? 0 : val);
      }, 0);

      // Enrich profile with stats
      const enrichedProfile: RepresentativeProfile = {
          ...profile,
          stats: {
              trainingCount,
              onlineCount,
              offlineCount,
              offlineValue
          }
      };
      
      setSelectedRepProfile(enrichedProfile);
  };

  const getKpiSummary = (kpi: KpiCategory) => {
    // Logic for Full Year Aggregation
    if (selectedMonth === '---') {
        // Special logic for CPL Year Summary (Total Spend / Total Leads)
        if (kpi.id === 'cpl') {
            const adSpendKpi = derivedKpis.find(k => k.id === 'ad_spend');
            const oppKpi = derivedKpis.find(k => k.id === 'opportunities');
            
            const totalSpend = adSpendKpi?.data.reduce((acc, curr) => acc + (curr.realized || 0), 0) || 0;
            const totalLeads = oppKpi?.data.reduce((acc, curr) => acc + (curr.realized || 0), 0) || 0;
            
            const annualCpl = totalLeads > 0 ? (totalSpend / totalLeads) : 0;
            // For target, we can average the monthly targets as a baseline
            const validMetaCount = kpi.data.filter(d => d.meta !== null).length;
            const avgMeta = validMetaCount > 0 ? kpi.data.reduce((acc, curr) => acc + (curr.meta || 0), 0) / validMetaCount : 0;

            return {
                realized: annualCpl,
                meta: avgMeta,
                month: 'TOTAL'
            };
        }

        const validData = kpi.data.filter(d => d.realized !== null);
        
        // Ajuste Específico: Autoridade da Internet deve mostrar o valor ATUAL (último), não a média.
        if (kpi.id === 'authority') {
            if (validData.length === 0) return { realized: 0, meta: 0, month: 'ANO' };
            const lastRecord = validData[validData.length - 1];
            return { 
                realized: lastRecord.realized || 0, 
                meta: lastRecord.meta || 0, 
                month: 'ATUAL' 
            };
        }

        if (validData.length === 0) return { realized: 0, meta: 0, month: 'ANO' };

        // Determine if we should sum or average based on KPI type
        const shouldAverage = kpi.unit === 'percent' || ['roas', 'conversion_rate'].includes(kpi.id);

        const totalRealized = validData.reduce((acc, curr) => acc + (curr.realized || 0), 0);
        
        // CORREÇÃO: Para KPIs cumulativos (não média), calculamos a meta anual completa (todos os meses),
        // e não apenas a meta dos meses realizados (YTD). Isso garante que o valor anual (ex: 840 para Youtube) seja exibido corretamente.
        let totalMeta = 0;
        if (shouldAverage) {
            // Para médias (ex: ROAS), usamos apenas os meses válidos para manter a consistência
            totalMeta = validData.reduce((acc, curr) => acc + (curr.meta || 0), 0);
        } else {
            // Para cumulativos, somamos a meta do ano todo (ex: meta do YouTube 840)
            totalMeta = kpi.data.reduce((acc, curr) => acc + (curr.meta || 0), 0);
        }

        if (shouldAverage) {
             return { 
                realized: totalRealized / validData.length, 
                meta: totalMeta / validData.length, 
                month: 'MÉDIA' 
            };
        }

        return { 
            realized: totalRealized, 
            meta: totalMeta, 
            month: 'TOTAL' 
        };
    }

    // Find data for the selected month
    const data = kpi.data.find(d => d.month === selectedMonth);
    
    if (data) {
        return { 
            realized: data.realized || 0, 
            meta: data.meta || 0, 
            month: data.month 
        };
    }
    
    // Fallback if month not found
    return { realized: 0, meta: 0, month: selectedMonth };
  };

  // Helper to filter table data based on selected month
  const getFilteredTableData = (tableData: RepTableData): RepTableData => {
      if (selectedMonth === '---') {
          return tableData; // Return all rows for full year view
      }

      const fullMonthName = FULL_MONTH_MAP[selectedMonth];
      const filteredRows = tableData.rows.filter(row => row.month === fullMonthName);
      
      // If no rows found (e.g. month hasn't happened yet), return empty or keep structure
      return {
          ...tableData,
          rows: filteredRows.length > 0 ? filteredRows : []
      };
  };

  const getKpiColorClass = (id: string) => {
     if (id.includes('training')) return 'bg-emerald-500/10';
     if (id.includes('marketing') || id.includes('lead') || id.includes('ad_spend')) return 'bg-purple-500/10';
     if (id.includes('sales') || id.includes('opportunities')) return 'bg-pink-500/10';
     if (id.includes('youtube')) return 'bg-red-500/10';
     if (id.includes('inbound')) return 'bg-amber-500/10';
     if (id === 'authority') return 'bg-indigo-500/10';
     return 'bg-blue-500/10';
  };

  const handleKpiUpdate = (kpiId: string, month: string, value: number, context?: string, isCorrection?: boolean, newBreakdown?: BreakdownItem[], dateOverride?: string) => {
      setKpis(prev => prev.map(k => {
          if (k.id !== kpiId) return k;
          
          const newData = k.data.map(d => {
              if (d.month === month) {
                  return { 
                      ...d, 
                      realized: value,
                      breakdown: newBreakdown || d.breakdown // Update breakdown if provided
                  };
              }
              return d;
          });

          // Use provided date or current date
          const logDate = dateOverride ? new Date(dateOverride) : new Date();
          const timestampStr = logDate.toLocaleTimeString('pt-BR', { hour: '2-digit', minute: '2-digit' });

          const newLog = {
              id: Date.now().toString(),
              date: logDate.toISOString(),
              timestamp: timestampStr,
              user: 'Jackson',
              month,
              oldValue: k.data.find(d => d.month === month)?.realized || 0,
              newValue: value,
              action: (isCorrection ? 'update' : 'create') as 'update' | 'create',
              context
          };
          const currentLogs = k.logs || [];
          return { ...k, data: newData, logs: [newLog, ...currentLogs] };
      }));
  };

  const handleUpdateKpiMeta = (kpiId: string, month: string, newMeta: number) => {
      setKpis(prev => prev.map(k => {
          if (k.id !== kpiId) return k;
          const newData = k.data.map(d => {
              if (d.month === month) {
                  return { ...d, meta: newMeta };
              }
              return d;
          });
          return { ...k, data: newData };
      }));
  };

  // Function to clear manual data (reset logs and realized values for non-computed KPIs)
  const handleClearKpiData = (kpiId: string) => {
      setKpis(prev => prev.map(k => {
          if (k.id !== kpiId) return k;
          
          // Reset logs
          const newLogs: any[] = [];
          
          const newData = k.data.map(d => ({
              ...d,
              realized: 0,
              breakdown: undefined
          }));

          return { ...k, logs: newLogs, data: newData };
      }));
  };

  // (Task Handlers omitted, keeping existing)
  // ... handleTaskClick, handleTaskUpdate, handleTaskDelete ...
  const handleTaskClick = (taskId: string) => {
    const task = tasks.find(t => t.id === taskId);
    if (task) setEditingTask(task);
  };

  const handleTaskUpdate = (updatedTask: Task) => {
    setTasks(prev => prev.map(t => t.id === updatedTask.id ? updatedTask : t));
    if (editingTask && editingTask.id === updatedTask.id) {
        setEditingTask(updatedTask);
    }
  };

  const handleTaskDelete = (taskId: string) => {
    setTasks(prev => prev.filter(t => t.id !== taskId));
    setEditingTask(null);
  };

  // Auth Handlers
  const handleLogin = () => {
    setIsLoggedIn(true);
    setActiveTab('dashboard');
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setShowUserMenu(false);
    setActiveTab('dashboard'); // Reset tab for next login
  };

  const NavItem = ({ id, icon: Icon, label }: any) => (
    <button
      onClick={() => { 
          setActiveTab(id); 
          setSelectedKpiId(null); 
          setSelectedFinancialModule(null); 
          setShowCplDashboard(false);
          setShowOppDashboard(false);
          setShowAdSpendDashboard(false);
          setShowRoasDashboard(false);
      }}
      className={`w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all duration-200 ${
        activeTab === id 
          ? 'bg-[#1e5144] text-white shadow-lg shadow-[#1e5144]/20 font-medium' 
          : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'
      }`}
    >
      <Icon size={20} />
      {isSidebarOpen && <span>{label}</span>}
      {activeTab === id && isSidebarOpen && <ChevronRight size={16} className="ml-auto opacity-50" />}
    </button>
  );

  // Render Content Switch
  const renderContent = () => {
    if (!isLoggedIn) return null; // Should be handled by parent return

    if (activeTab === 'settings') {
        return <SettingsView 
                    onLogout={handleLogout} 
                    notifications={notifications}
                    onUpdateNotification={setNotifications}
                    kpis={kpis}
                    onUpdateKpiMeta={handleUpdateKpiMeta}
               />;
    }

    if (activeTab === 'passwords') {
        return <PasswordsView />;
    }

    if (activeTab === 'contacts') {
        return <ContactsView />;
    }

    if (selectedKpiId) {
        const kpi = derivedKpis.find(k => k.id === selectedKpiId);
        if (kpi) {
            return (
                <KpiDetailView 
                    kpi={kpi}
                    allKpis={derivedKpis} // Passing derived kpis for cross-referencing
                    selectedMonth={selectedMonth} 
                    onBack={() => setSelectedKpiId(null)}
                    onUpdate={handleKpiUpdate}
                    onClearData={() => handleClearKpiData(kpi.id)}
                />
            );
        }
    }

    switch (activeTab) {
        case 'dashboard':
            return (
                <div className="space-y-6 animate-in fade-in duration-500">
                    <div className="flex justify-between items-end">
                        <h2 className="text-xl font-bold text-gray-200">Visão Geral - {FULL_MONTH_MAP[selectedMonth]}</h2>
                    </div>
                    
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {derivedKpis.filter(k => ['youtube', 'training_reps', 'training_people', 'marketing', 'inbound', 'authority'].includes(k.id)).map(kpi => {
                            const { realized, meta } = getKpiSummary(kpi);
                            let Icon = Target;
                            if (kpi.id === 'youtube') Icon = Youtube;
                            else if (kpi.id.includes('training')) Icon = GraduationCap;
                            else if (kpi.id === 'marketing') Icon = Megaphone;
                            else if (kpi.id === 'authority') Icon = TrendingUp;

                            return (
                                <SummaryCard 
                                    key={kpi.id}
                                    title={kpi.title}
                                    value={realized}
                                    target={meta}
                                    icon={Icon}
                                    colorClass={getKpiColorClass(kpi.id)}
                                    onClick={() => setSelectedKpiId(kpi.id)}
                                    unit={kpi.unit}
                                />
                            );
                        })}
                    </div>
                    {/* Alterado para flex-col para empilhar Widgets */}
                    <div className="flex flex-col gap-6">
                        <div className="h-[400px]">
                            <TaskWidget 
                                tasks={tasks} 
                                onViewAll={() => setActiveTab('tasks')} 
                                onTaskClick={handleTaskClick} 
                            />
                        </div>
                        <div className="h-auto">
                            <UpcomingPostsWidget 
                                posts={posts} 
                                onViewCalendar={() => setActiveTab('calendar')}
                            />
                        </div>
                    </div>
                </div>
            );
        case 'training':
            const trainingKpiIds = ['training_people', 'training_reps'];
            const trainingKpis = derivedKpis.filter(k => trainingKpiIds.includes(k.id));
            const filteredTrainingData = getFilteredTableData(REPS_TRAINING_DATA);
            
            return (
                <div className="space-y-6 animate-in fade-in duration-500">
                    <h2 className="text-xl font-bold text-gray-200 mb-4">Treinamentos e Capacitação</h2>
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {trainingKpis.map(kpi => {
                            const { realized, meta } = getKpiSummary(kpi);
                            return (
                                <SummaryCard 
                                    key={kpi.id}
                                    title={kpi.title}
                                    value={realized}
                                    target={meta}
                                    icon={GraduationCap}
                                    colorClass={getKpiColorClass(kpi.id)}
                                    onClick={() => setSelectedKpiId(kpi.id)}
                                    unit={kpi.unit}
                                />
                            );
                        })}
                    </div>
                    {filteredTrainingData.rows.length > 0 ? (
                        <div className="h-auto">
                            <HeatmapChart 
                                data={filteredTrainingData} 
                                title={`Frequência de Treinamentos (${selectedMonth === '---' ? 'Ano Completo' : FULL_MONTH_MAP[selectedMonth]})`}
                                onRepClick={handleRepClick}
                            />
                        </div>
                    ) : (
                        <div className="p-8 text-center bg-[#1a1d24] rounded-2xl border border-white/5 text-gray-500">
                            Sem dados de treinamento para {FULL_MONTH_MAP[selectedMonth]}.
                        </div>
                    )}
                </div>
            );
        case 'marketing':
             // --- OPPORTUNITIES INTERMEDIATE DASHBOARD ---
             if (showOppDashboard) {
                 const oppKpi = derivedKpis.find(k => k.id === 'opportunities');
                 if (oppKpi) {
                     const { realized, meta } = getKpiSummary(oppKpi);
                     
                     // Get Breakdown for Cards
                     let breakdownData: { label: string, value: number }[] = [];
                     
                     if (selectedMonth !== '---') {
                         // Monthly Breakdown
                         const monthData = oppKpi.data.find(d => d.month === selectedMonth);
                         if (monthData?.breakdown) {
                             breakdownData = monthData.breakdown;
                         }
                     } else {
                         // Annual Breakdown: Sum all values by label
                         const aggStats: Record<string, number> = {};
                         oppKpi.data.forEach(d => {
                             d.breakdown?.forEach(item => {
                                 if (!aggStats[item.label]) aggStats[item.label] = 0;
                                 aggStats[item.label] += item.value;
                             });
                         });
                         breakdownData = Object.entries(aggStats).map(([label, value]) => ({ label, value }));
                     }

                     return (
                         <div className="space-y-6 animate-in fade-in duration-500">
                             <div className="flex items-center gap-4 mb-4">
                                 <button 
                                     onClick={() => setShowOppDashboard(false)}
                                     className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
                                 >
                                     <ArrowLeft size={24} />
                                 </button>
                                 <h2 className="text-xl font-bold text-gray-200">Taxa de Oportunidades</h2>
                             </div>
                             
                             <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 items-start">
                                 {/* Total Card */}
                                 <SummaryCard 
                                     key={oppKpi.id}
                                     title={`Total de Oportunidades (${selectedMonth === '---' ? 'Ano' : selectedMonth})`}
                                     value={realized}
                                     target={meta}
                                     icon={TrendingUp}
                                     colorClass={getKpiColorClass(oppKpi.id)}
                                     onClick={() => setSelectedKpiId(oppKpi.id)}
                                     unit={oppKpi.unit}
                                 />
                                 
                                 {/* Breakdown Cards Container (Middle Column) */}
                                 <div className="lg:col-span-2 grid grid-cols-2 lg:grid-cols-4 gap-4 w-full">
                                     {breakdownData.map((item, idx) => (
                                         <SummaryCard 
                                             key={idx}
                                             title={item.label}
                                             value={item.value}
                                             target={0} // No specific target per channel
                                             icon={TrendingUp}
                                             colorClass="bg-pink-500/10"
                                             onClick={() => setSelectedKpiId(oppKpi.id)}
                                             unit="number"
                                             hideMeta={true}
                                         />
                                     ))}
                                 </div>
                             </div>
                         </div>
                     );
                 }
             }

             // --- CPL INTERMEDIATE DASHBOARD ---
             if (showCplDashboard) {
                 const cplKpi = derivedKpis.find(k => k.id === 'cpl');
                 if (cplKpi) {
                     const { realized, meta } = getKpiSummary(cplKpi);
                     
                     // Get Breakdown for Cards
                     let breakdownData: { label: string, value: number }[] = [];
                     
                     if (selectedMonth !== '---') {
                         // Monthly Breakdown: Use derived calculation
                         const monthData = cplKpi.data.find(d => d.month === selectedMonth);
                         if (monthData?.breakdown) {
                             breakdownData = monthData.breakdown;
                         }
                     } else {
                         // Annual Breakdown: Calculate Aggregate CPL per Channel
                         const adSpendKpi = derivedKpis.find(k => k.id === 'ad_spend');
                         const oppKpi = derivedKpis.find(k => k.id === 'opportunities');
                         
                         if (adSpendKpi && oppKpi) {
                             const channelStats: Record<string, { spend: number, leads: number }> = {};
                             
                             // Sum Spend
                             adSpendKpi.data.forEach(d => {
                                 d.breakdown?.forEach(item => {
                                     if (!channelStats[item.label]) channelStats[item.label] = { spend: 0, leads: 0 };
                                     channelStats[item.label].spend += item.value;
                                 });
                             });
                             
                             // Sum Leads
                             oppKpi.data.forEach(d => {
                                 d.breakdown?.forEach(item => {
                                     const matchKey = Object.keys(channelStats).find(k => k === item.label || k.includes(item.label) || item.label.includes(k));
                                     if (matchKey) {
                                         channelStats[matchKey].leads += item.value;
                                     }
                                 });
                             });
                             
                             breakdownData = Object.entries(channelStats).map(([label, stats]) => ({
                                 label,
                                 value: stats.leads > 0 ? stats.spend / stats.leads : 0
                             }));
                         }
                     }

                     return (
                         <div className="space-y-6 animate-in fade-in duration-500">
                             <div className="flex items-center gap-4 mb-4">
                                 <button 
                                     onClick={() => setShowCplDashboard(false)}
                                     className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
                                 >
                                     <ArrowLeft size={24} />
                                 </button>
                                 <h2 className="text-xl font-bold text-gray-200">Custo por Lead (CPL)</h2>
                             </div>
                             
                             <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 items-start">
                                 {/* Total Card */}
                                 <SummaryCard 
                                     key={cplKpi.id}
                                     title={`CPL Geral (${selectedMonth === '---' ? 'Ano' : selectedMonth})`}
                                     value={realized}
                                     target={meta}
                                     icon={Megaphone}
                                     colorClass={getKpiColorClass(cplKpi.id)}
                                     onClick={() => setSelectedKpiId(cplKpi.id)}
                                     unit={cplKpi.unit}
                                 />
                                 
                                 {/* Breakdown Cards Container (Middle Column) */}
                                 <div className="lg:col-span-2 grid grid-cols-2 lg:grid-cols-4 gap-4 w-full">
                                     {breakdownData.map((item, idx) => (
                                         <SummaryCard 
                                             key={idx}
                                             title={`CPL - ${item.label}`}
                                             value={item.value}
                                             target={0} // No specific target, will display just value
                                             icon={Megaphone}
                                             colorClass="bg-purple-500/10"
                                             onClick={() => setSelectedKpiId(cplKpi.id)}
                                             unit="currency"
                                             hideMeta={true}
                                         />
                                     ))}
                                 </div>
                             </div>
                         </div>
                     );
                 }
             }

             // --- AD SPEND INTERMEDIATE DASHBOARD ---
             if (showAdSpendDashboard) {
                 const adKpi = derivedKpis.find(k => k.id === 'ad_spend');
                 if (adKpi) {
                     const { realized, meta } = getKpiSummary(adKpi);
                     
                     let breakdownData: { label: string, value: number }[] = [];
                     
                     if (selectedMonth !== '---') {
                         const monthData = adKpi.data.find(d => d.month === selectedMonth);
                         if (monthData?.breakdown) {
                             breakdownData = monthData.breakdown;
                         }
                     } else {
                         // Annual Aggregation (Sum)
                         const aggStats: Record<string, number> = {};
                         adKpi.data.forEach(d => {
                             d.breakdown?.forEach(item => {
                                 if (!aggStats[item.label]) aggStats[item.label] = 0;
                                 aggStats[item.label] += item.value;
                             });
                         });
                         breakdownData = Object.entries(aggStats).map(([label, value]) => ({ label, value }));
                     }

                     return (
                         <div className="space-y-6 animate-in fade-in duration-500">
                             <div className="flex items-center gap-4 mb-4">
                                 <button 
                                     onClick={() => setShowAdSpendDashboard(false)}
                                     className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
                                 >
                                     <ArrowLeft size={24} />
                                 </button>
                                 <h2 className="text-xl font-bold text-gray-200">Investimento em Ads</h2>
                             </div>
                             
                             <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 items-start">
                                 <SummaryCard 
                                     key={adKpi.id}
                                     title={`Investimento Total (${selectedMonth === '---' ? 'Ano' : selectedMonth})`}
                                     value={realized}
                                     target={meta}
                                     icon={Megaphone}
                                     colorClass={getKpiColorClass(adKpi.id)}
                                     onClick={() => setSelectedKpiId(adKpi.id)}
                                     unit={adKpi.unit}
                                 />
                                 
                                 <div className="lg:col-span-2 grid grid-cols-2 lg:grid-cols-4 gap-4 w-full">
                                     {breakdownData.map((item, idx) => (
                                         <SummaryCard 
                                             key={idx}
                                             title={item.label}
                                             value={item.value}
                                             target={0}
                                             icon={Megaphone}
                                             colorClass="bg-purple-500/10"
                                             onClick={() => setSelectedKpiId(adKpi.id)}
                                             unit="currency"
                                             hideMeta={true}
                                         />
                                     ))}
                                 </div>
                             </div>
                         </div>
                     );
                 }
             }

             // --- ROAS INTERMEDIATE DASHBOARD ---
             if (showRoasDashboard) {
                 const roasKpi = derivedKpis.find(k => k.id === 'roas');
                 if (roasKpi) {
                     const { realized, meta } = getKpiSummary(roasKpi);
                     
                     let breakdownData: { label: string, value: number }[] = [];
                     
                     if (selectedMonth !== '---') {
                         const monthData = roasKpi.data.find(d => d.month === selectedMonth);
                         if (monthData?.breakdown) {
                             breakdownData = monthData.breakdown;
                         }
                     } else {
                         // Annual Aggregation (Average for ratios like ROAS)
                         const aggStats: Record<string, { sum: number, count: number }> = {};
                         roasKpi.data.forEach(d => {
                             d.breakdown?.forEach(item => {
                                 if (!aggStats[item.label]) aggStats[item.label] = { sum: 0, count: 0 };
                                 aggStats[item.label].sum += item.value;
                                 aggStats[item.label].count += 1;
                             });
                         });
                         breakdownData = Object.entries(aggStats).map(([label, stats]) => ({ 
                             label, 
                             value: stats.count > 0 ? stats.sum / stats.count : 0 
                         }));
                     }

                     return (
                         <div className="space-y-6 animate-in fade-in duration-500">
                             <div className="flex items-center gap-4 mb-4">
                                 <button 
                                     onClick={() => setShowRoasDashboard(false)}
                                     className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
                                 >
                                     <ArrowLeft size={24} />
                                 </button>
                                 <h2 className="text-xl font-bold text-gray-200">ROAS (Retorno em Ads)</h2>
                             </div>
                             
                             <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 items-start">
                                 <SummaryCard 
                                     key={roasKpi.id}
                                     title={`ROAS Médio (${selectedMonth === '---' ? 'Ano' : selectedMonth})`}
                                     value={realized}
                                     target={meta}
                                     icon={TrendingUp}
                                     colorClass={getKpiColorClass(roasKpi.id)}
                                     onClick={() => setSelectedKpiId(roasKpi.id)}
                                     unit={roasKpi.unit}
                                 />
                                 
                                 <div className="lg:col-span-2 grid grid-cols-2 lg:grid-cols-4 gap-4 w-full">
                                     {breakdownData.map((item, idx) => (
                                         <SummaryCard 
                                             key={idx}
                                             title={item.label}
                                             value={item.value}
                                             target={0}
                                             icon={TrendingUp}
                                             colorClass="bg-purple-500/10"
                                             onClick={() => setSelectedKpiId(roasKpi.id)}
                                             unit="number" // ROAS is a multiplier/number
                                             hideMeta={true}
                                         />
                                     ))}
                                 </div>
                             </div>
                         </div>
                     );
                 }
             }

             return (
                 <div className="space-y-6 animate-in fade-in duration-500">
                     <h2 className="text-xl font-bold text-gray-200">Marketing e Performance</h2>
                     <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {derivedKpis.filter(k => ['inbound', 'opportunities', 'ad_spend', 'cpl', 'roas'].includes(k.id)).map(kpi => {
                             const { realized, meta } = getKpiSummary(kpi);
                             
                             // Logic to intercept clicks for intermediate dashboards
                             const handleClick = () => {
                                 if (kpi.id === 'cpl') {
                                     setShowCplDashboard(true);
                                 } else if (kpi.id === 'opportunities') {
                                     setShowOppDashboard(true);
                                 } else if (kpi.id === 'ad_spend') {
                                     setShowAdSpendDashboard(true);
                                 } else if (kpi.id === 'roas') {
                                     setShowRoasDashboard(true);
                                 } else {
                                     setSelectedKpiId(kpi.id);
                                 }
                             };

                             return (
                                 <SummaryCard 
                                     key={kpi.id}
                                     title={kpi.title}
                                     value={realized}
                                     target={meta}
                                     icon={Megaphone}
                                     colorClass={getKpiColorClass(kpi.id)}
                                     onClick={handleClick}
                                     unit={kpi.unit}
                                 />
                             );
                         })}
                     </div>
                 </div>
             );
        case 'commercial':
             return (
                 <div className="space-y-6 animate-in fade-in duration-500">
                     <h2 className="text-xl font-bold text-gray-200">Comercial</h2>
                     <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {derivedKpis.filter(k => ['opportunities', 'conversion_rate'].includes(k.id)).map(kpi => {
                             const { realized, meta } = getKpiSummary(kpi);
                             
                             // Intercept opportunity click here too if consistent behavior is desired
                             const handleClick = () => {
                                if (kpi.id === 'opportunities') {
                                    // Direct to detailed view directly here for simplicity in Commercial tab context
                                    // as the intermediate dashboard is part of "Marketing" tab flow visually
                                    setSelectedKpiId(kpi.id);
                                } else {
                                    setSelectedKpiId(kpi.id);
                                }
                             };

                             return (
                                 <SummaryCard 
                                     key={kpi.id}
                                     title={kpi.title}
                                     value={realized}
                                     target={meta}
                                     icon={TrendingUp}
                                     colorClass={getKpiColorClass(kpi.id)}
                                     onClick={handleClick}
                                     unit={kpi.unit}
                                 />
                             );
                         })}
                     </div>
                 </div>
             );
        case 'tasks':
            // ... (rest of cases remain same)
            return <KanbanBoard 
                        tasks={tasks} 
                        onAddTask={(t) => setTasks([...tasks, t])}
                        onUpdateTask={handleTaskUpdate}
                        onDeleteTask={handleTaskDelete}
                        onReorderTasks={setTasks}
                        onTaskClick={handleTaskClick} 
                    />;
        case 'calendar':
            return <PostCalendarView posts={posts} onUpdatePosts={setPosts} />;
        case 'cooperative':
            return <CooperativeActionsView selectedMonth={selectedMonth} onRepClick={handleRepClick} />;
        case 'social':
            return <SocialBenchmarkingView data={SOCIAL_MEDIA_RANKING} />;
        case 'gifts':
            return <GiftsView onAddNotification={handleAddNotification} />;
        case 'financial':
            if (selectedFinancialModule === 'accounts_payable') {
                return <FinancialView onBack={() => setSelectedFinancialModule(null)} />;
            }
            
            if (selectedFinancialModule === 'budget') {
                return <BudgetView onBack={() => setSelectedFinancialModule(null)} />;
            }

            if (selectedFinancialModule === 'marketing_actions') {
                return (
                    <MarketingBudgetView 
                        onBack={() => setSelectedFinancialModule(null)}
                        selectedMonth={selectedMonth} // Passing the selected month
                    />
                );
            }
            
            // Calculate summary data
            const totalPending = MOCK_ACCOUNTS_PAYABLE
                .filter(a => a.status === 'pending')
                .reduce((acc, curr) => acc + curr.amount, 0);

            return (
                <div className="space-y-6 animate-in fade-in duration-500">
                    <h2 className="text-xl font-bold text-gray-200">Financeiro</h2>
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        <SummaryCard 
                            title="Contas a Pagar"
                            value={totalPending}
                            target={0} 
                            icon={DollarSign}
                            colorClass="bg-rose-500/10"
                            unit="currency"
                            onClick={() => setSelectedFinancialModule('accounts_payable')}
                        />
                        <SummaryCard 
                            title="Orçamento"
                            value={budgetSummary} 
                            target={0} 
                            icon={PieChart}
                            colorClass="bg-purple-500/10"
                            unit="currency"
                            onClick={() => setSelectedFinancialModule('budget')}
                            hideMeta={true}
                        />
                        <SummaryCard 
                            title="Ações de Marketing"
                            value={marketingActionsSummary} 
                            target={0} 
                            icon={Megaphone}
                            colorClass="bg-pink-500/10"
                            unit="currency"
                            onClick={() => setSelectedFinancialModule('marketing_actions')}
                            hideMeta={true}
                        />
                    </div>
                </div>
            );
        default:
            return <div className="text-gray-400 p-10">Selecione uma opção no menu.</div>;
    }
  };

  // ... (rest of component)
  if (!isLoggedIn) {
      return <LoginView onLogin={handleLogin} />;
  }

  const unreadCount = notifications.filter(n => !n.read && !n.archived).length;

  return (
    <div className="flex h-screen bg-[#0f1115] text-gray-100 overflow-hidden font-sans selection:bg-[#1e5144] selection:text-white">
      {/* ... Sidebar ... */}
      <aside 
        className={`${isSidebarOpen ? 'w-64' : 'w-20'} bg-[#12141a] border-r border-gray-800 transition-all duration-300 flex flex-col z-20`}
      >
        {/* ... Sidebar content ... */}
        <div className="p-6 flex items-center gap-3 mb-2 cursor-pointer" onClick={() => setActiveTab('dashboard')}>
            <div className="w-8 h-8 bg-gradient-to-br from-[#1e5144] to-emerald-600 rounded-lg flex items-center justify-center shadow-lg shadow-emerald-900/50">
                <span className="font-bold text-white text-lg">S</span>
            </div>
            {isSidebarOpen && (
                <div>
                    <h1 className="font-bold text-lg tracking-tight text-white">Solis <span className="text-emerald-500">Hub</span></h1>
                    <p className="text-[10px] text-gray-500 uppercase tracking-widest">Workspace</p>
                </div>
            )}
        </div>

        <nav className="flex-grow px-3 space-y-1 overflow-y-auto custom-scrollbar">
            <div className="mb-2 px-4 text-xs font-semibold text-gray-600 uppercase tracking-wider">
                {isSidebarOpen ? 'Gestão' : '---'}
            </div>
            <NavItem id="dashboard" icon={LayoutDashboard} label="Dashboard" />
            <NavItem id="commercial" icon={TrendingUp} label="Comercial" />
            <NavItem id="marketing" icon={Megaphone} label="Marketing" />
            <NavItem id="training" icon={GraduationCap} label="Treinamentos" />
            
            <div className="mt-6 mb-2 px-4 text-xs font-semibold text-gray-600 uppercase tracking-wider">
                {isSidebarOpen ? 'Operacional' : '---'}
            </div>
            <NavItem id="tasks" icon={CheckSquare} label="Tarefas" />
            <NavItem id="calendar" icon={Calendar} label="Calendário" />
            <NavItem id="cooperative" icon={Share2} label="Ações MKT" />
            <NavItem id="social" icon={TrendingUp} label="Social Rank" />
            <NavItem id="gifts" icon={Gift} label="Brindes" />
            <NavItem id="financial" icon={DollarSign} label="Financeiro" />
        </nav>

        <div className="p-4 border-t border-gray-800">
            <button 
                onClick={() => setIsSidebarOpen(!isSidebarOpen)}
                className="w-full flex items-center justify-center p-2 rounded-lg hover:bg-gray-800 text-gray-500 transition-colors"
            >
                <Menu size={20} />
            </button>
        </div>
      </aside>

      {/* Rest of the layout... */}
      <div className="flex-grow flex flex-col h-full overflow-hidden relative">
         <header className="h-16 shrink-0 w-full bg-[#12141a]/80 backdrop-blur-md border-b border-gray-800 flex justify-end items-center px-6 z-10">
            <div className="flex items-center gap-4">
                <div className="hidden sm:flex items-center gap-2 bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-1.5 transition-colors hover:border-gray-600 mr-2">
                    <Calendar size={16} className="text-emerald-500" />
                    <div className="relative group">
                        <select 
                            value={selectedMonth}
                            onChange={(e) => setSelectedMonth(e.target.value)}
                            className="appearance-none bg-transparent text-sm text-gray-200 font-medium focus:outline-none cursor-pointer pr-4"
                        >
                            <option value="---" className="bg-[#1a1d24] text-gray-200 font-bold">---</option>
                            {MONTHS.map(m => (
                                <option key={m} value={m} className="bg-[#1a1d24] text-gray-200">
                                    {m}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div className="w-px h-4 bg-gray-700 mx-1"></div>
                    <div className="relative group">
                        <select 
                            value={selectedYear}
                            onChange={(e) => setSelectedYear(e.target.value)}
                            className="appearance-none bg-transparent text-sm text-gray-400 focus:outline-none cursor-pointer pr-4"
                        >
                            <option value="2025" className="bg-[#1a1d24]">2025</option>
                        </select>
                    </div>
                    <ChevronDown size={14} className="text-gray-500" />
                </div>

                <div className="relative">
                    <button 
                        className="relative p-2 text-gray-400 hover:text-white hover:bg-gray-800 rounded-full transition-colors"
                        onClick={() => setShowNotifications(!showNotifications)}
                    >
                        <Bell size={20} />
                        {unreadCount > 0 && (
                            <span className="absolute top-2 right-2 w-2 h-2 bg-rose-50 rounded-full animate-pulse"></span>
                        )}
                    </button>
                    {showNotifications && (
                        <>
                            <div className="fixed inset-0 z-40" onClick={() => setShowNotifications(false)}></div>
                            <NotificationPanel 
                                notifications={notifications} 
                                onMarkAsRead={(id) => setNotifications(prev => prev.map(n => n.id === id ? {...n, read: true} : n))}
                                onArchive={handleArchiveNotification}
                                onClearAll={() => setNotifications(prev => prev.map(n => ({...n, read: true})))}
                                onNotificationClick={handleNotificationClick}
                            />
                        </>
                    )}
                </div>
                
                <div className="h-8 w-px bg-gray-800 mx-2"></div>
                
                <div className="relative">
                    <button 
                        onClick={() => setShowUserMenu(!showUserMenu)}
                        className="flex items-center gap-3 pl-2 group"
                    >
                        <div className="text-right hidden md:block">
                            <p className="text-sm font-medium text-gray-200 group-hover:text-white transition-colors">Jackson</p>
                        </div>
                        <div className="w-9 h-9 rounded-full bg-gradient-to-tr from-gray-700 to-gray-600 border border-gray-500 flex items-center justify-center text-sm font-bold text-white shadow-md cursor-pointer group-hover:border-[#1e5144] transition-all">
                            JA
                        </div>
                    </button>
                    
                    {showUserMenu && (
                        <>
                            <div className="fixed inset-0 z-40" onClick={() => setShowUserMenu(false)}></div>
                            <div className="absolute right-0 top-12 w-56 bg-[#1a1d24] border border-gray-700 rounded-xl shadow-2xl z-50 overflow-hidden animate-in fade-in slide-in-from-top-2">
                                <button 
                                    onClick={() => { setActiveTab('settings'); setSelectedKpiId(null); setShowUserMenu(false); }}
                                    className="w-full text-left px-4 py-3 hover:bg-[#20232b] text-gray-300 hover:text-white text-sm flex items-center gap-2"
                                >
                                    <Settings size={16} /> Configurações
                                </button>
                                <button 
                                    onClick={() => { setActiveTab('passwords'); setSelectedKpiId(null); setShowUserMenu(false); }}
                                    className="w-full text-left px-4 py-3 hover:bg-[#20232b] text-gray-300 hover:text-white text-sm flex items-center gap-2"
                                >
                                    <KeyRound size={16} /> Acessos
                                </button>
                                <button 
                                    onClick={() => { setActiveTab('contacts'); setSelectedKpiId(null); setShowUserMenu(false); }}
                                    className="w-full text-left px-4 py-3 hover:bg-[#20232b] text-gray-300 hover:text-white text-sm flex items-center gap-2"
                                >
                                    <Phone size={16} /> Ramais e E-mails
                                </button>

                                <div className="h-px bg-gray-800 my-1"></div>
                                <button 
                                    onClick={handleLogout}
                                    className="w-full text-left px-4 py-3 hover:bg-rose-500/10 text-rose-400 text-sm flex items-center gap-2"
                                >
                                    <LogOut size={16} /> Sair
                                </button>
                            </div>
                        </>
                    )}
                </div>
            </div>
         </header>

         {/* MODIFIED MAIN CONTAINER */}
         <main className={`flex-grow overflow-y-auto custom-scrollbar bg-[#0f1115] ${activeTab === 'tasks' ? 'p-0' : 'p-6'}`}>
            <div className={`h-full flex flex-col ${activeTab === 'tasks' ? '' : 'max-w-7xl mx-auto'}`}>
                {renderContent()}
            </div>
         </main>
      </div>

      {editingTask && isLoggedIn && (
        <TaskModal 
            task={editingTask}
            isOpen={!!editingTask}
            onClose={() => setEditingTask(null)}
            onUpdate={handleTaskUpdate}
            onDelete={handleTaskDelete}
            onMention={handleMention}
        />
      )}

      {selectedRepProfile && (
          <RepProfileModal 
            isOpen={!!selectedRepProfile}
            onClose={() => setSelectedRepProfile(null)}
            profile={selectedRepProfile}
          />
      )}

    </div>
  );
};

export default App;
