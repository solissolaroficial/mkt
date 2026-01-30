import React, { useState, useEffect } from 'react';
import { 
  ResponsiveContainer, 
  AreaChart, 
  Area, 
  XAxis, 
  YAxis, 
  CartesianGrid, 
  Tooltip,
  BarChart,
  Bar,
  Cell,
  ComposedChart,
  Line
} from 'recharts';
import { KpiCategory, KpiLog, BreakdownItem, MonthlyData } from '@/shared/types/legacy.types';
import { ArrowLeft, Plus, History, X, AlertCircle, ChevronDown, ChevronRight, Sparkles, Calendar, Calculator, Trash2, Pencil } from 'lucide-react';
import { MONTHS } from '@/shared/utils/legacy.constants';
import { useUpdateMonthlyData } from '../hooks/useKpiMutations';
import { useDeleteMonthlyData } from '../hooks/useKpis';
import { useRepresentatives } from '@/features/pdv/hooks';
interface KpiDetailViewProps {
  kpi: KpiCategory;
  allKpis?: KpiCategory[]; // Access to other KPIs for cross-referencing
  selectedMonth: string;
  monthNames?: Record<string, string>;
  onBack: () => void;
}

const KpiDetailView: React.FC<KpiDetailViewProps> = ({ kpi, allKpis, selectedMonth, monthNames, onBack }) => {
  const [isHistoryOpen, setIsHistoryOpen] = useState(false);
  const [isLaunchOpen, setIsLaunchOpen] = useState(false);
  const [isEnrichOpen, setIsEnrichOpen] = useState(false); // New state for Enrichment Modal
  const { data: representatives = [], isLoading: isLoadingReps } = useRepresentatives();
  
  // Launch State
  const [launchDate, setLaunchDate] = useState(''); // Specific Date YYYY-MM-DD
  const [launchMonth, setLaunchMonth] = useState('NOV'); // Extracted Month for internal logic
  const [launchValue, setLaunchValue] = useState('');
  const [launchContext, setLaunchContext] = useState('');

// Opportunity Launch Specifics
const [oppChannel, setOppChannel] = useState('');
const [oppSubItem, setOppSubItem] = useState('');
const [isNewChannel, setIsNewChannel] = useState(false);

// Enrichment Specifics (Ad Spend)
// Store values as Record<label_key, value>
const [enrichValues, setEnrichValues] = useState<Record<string, string>>({});

// Expanded breakdown rows state
const [expandedRows, setExpandedRows] = useState<Record<string, boolean>>({});

// ADICIONAR: Hook de mutation para salvar dados
const { mutate: updateMonthlyData, isPending } = useUpdateMonthlyData();
const { deleteMonthlyData, isPending: isDeleting } = useDeleteMonthlyData();

 // Delete confirmation state
 const [deleteConfirm, setDeleteConfirm] = useState<{ id: string; month: string } | null>(null);

 // Edit modal state
 const [isEditOpen, setIsEditOpen] = useState(false);
 const [editData, setEditData] = useState<MonthlyData | null>(null);
 const [editRealized, setEditRealized] = useState('');
 const [editMeta, setEditMeta] = useState('');
 const [editContext, setEditContext] = useState('');

 const currentYear = new Date().getFullYear();

  useEffect(() => {
      // Set default date to today or start of selected month
      const now = new Date();
      if (selectedMonth !== '---') {
          const monthIdx = MONTHS.indexOf(selectedMonth);
          if (monthIdx !== -1) {
              // Construct a date in the selected month
              // If selected month is current month, use today, else 1st of month
              if (monthIdx === now.getMonth()) {
                  setLaunchDate(now.toISOString().split('T')[0]);
              } else {
                  const dateStr = `${currentYear}-${String(monthIdx + 1).padStart(2, '0')}-01`;
                  setLaunchDate(dateStr);
              }
              setLaunchMonth(selectedMonth);
          }
      } else {
          setLaunchDate(now.toISOString().split('T')[0]);
          setLaunchMonth(MONTHS[now.getMonth()]);
      }
  }, [selectedMonth]);

  const handleDateChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
      const dateVal = e.target.value;
      setLaunchDate(dateVal);
      
      // Extract Month to update launchMonth
      if (dateVal) {
          const d = new Date(dateVal);
          // Month is 0-indexed
          const mIdx = d.getMonth(); // Note: input date assumes local, but we just need month index
          // Adjust for timezone potentially shifting date?
          // To be safe, split string
          const [y, m, dstr] = dateVal.split('-').map(Number);
          if (MONTHS[m - 1]) {
              setLaunchMonth(MONTHS[m - 1]);
          }
      }
  };

  const toggleRow = (month: string) => {
      setExpandedRows(prev => ({ ...prev, [month]: !prev[month] }));
  };

  // Determine input configuration based on KPI slug
  const isRepKpi = ['treinamentos_realizados_reps', 'acoes_de_marketing_reps'].includes(kpi.slug);
  const isChannelKpi = ['roas_retorno_em_ads'].includes(kpi.slug);
  const isOppKpi = kpi.slug === 'taxa_de_oportunidades';
  const isAdSpend = kpi.slug === 'investimento_em_ads';
  const isCplKpi = kpi.slug === 'custo_por_lead_cpl';
  
  const handleLaunch = (e: React.FormEvent) => {
      e.preventDefault();
      if (!launchValue) return;

      const numValue = Number(launchValue);
      
      // Special logic for Opportunities Breakdown
      if (isOppKpi) {
          const currentData = kpi.data.find(d => d.month === launchMonth);
          const currentRealized = currentData?.realized || 0;
          const currentBreakdown = currentData?.breakdown ? JSON.parse(JSON.stringify(currentData.breakdown)) : []; // Deep clone

          const channelLabel = isNewChannel ? oppChannel : oppChannel;
          
          let channelIndex = currentBreakdown.findIndex((b: BreakdownItem) => b.label === channelLabel);
          
          if (channelIndex === -1) {
              // Create new channel
              currentBreakdown.push({
                  label: channelLabel,
                  value: 0,
                  subItems: []
              });
              channelIndex = currentBreakdown.length - 1;
          }

          // Update Channel Value
          currentBreakdown[channelIndex].value += numValue;

          // Handle Sub-item if present (e.g. Region for Meta Ads)
          if (oppSubItem) {
              if (!currentBreakdown[channelIndex].subItems) {
                  currentBreakdown[channelIndex].subItems = [];
              }
              const subIndex = currentBreakdown[channelIndex].subItems.findIndex((s: BreakdownItem) => s.label === oppSubItem);
              
              if (subIndex === -1) {
                  currentBreakdown[channelIndex].subItems.push({
                      label: oppSubItem,
                      value: numValue
                  });
              } else {
                  currentBreakdown[channelIndex].subItems[subIndex].value += numValue;
              }
          }

          const newTotal = currentRealized + numValue;
          const contextMsg = `${channelLabel}${oppSubItem ? ` - ${oppSubItem}` : ''}`;

          // SUBSTITUIR onUpdate por mutation
          const monthData = currentData;
          if (monthData) {
            updateMonthlyData({
              kpiId: kpi.id,
              data: {
                year: currentYear,
                realized: newTotal,
                breakdown: currentBreakdown,
                month: launchMonth,
              }
            }, {
              onSuccess: () => {
                setIsLaunchOpen(false);
                // Limpar form
                setLaunchValue('');
                setLaunchContext('');
                setOppChannel('');
                setOppSubItem('');
                setIsNewChannel(false);
              }
            });
          }

      } else {
          // Standard Update
          const monthData = kpi.data.find(d => d.month === launchMonth);
          if (monthData) {
            updateMonthlyData({
              kpiId: kpi.id,
              data: {
                year: currentYear,
                realized: Number(launchValue),
                month: launchMonth,
              }
            }, {
              onSuccess: () => {
                setIsLaunchOpen(false);
                // Limpar form
                setLaunchValue('');
                setLaunchContext('');
              }
            });
          }
      }
  };

  // Helper recursive function to build Ad Spend breakdown
  const buildAdSpendBreakdown = (oppBreakdown: BreakdownItem[], values: Record<string, string>): { breakdown: BreakdownItem[], total: number } => {
      let total = 0;
      const breakdown: BreakdownItem[] = [];

      oppBreakdown.forEach(item => {
          const itemKey = item.label;
          let itemValue = 0;
          let subItems: BreakdownItem[] = [];

          if (item.subItems && item.subItems.length > 0) {
              // Process subItems
              item.subItems.forEach(sub => {
                  const subKey = `${item.label}-${sub.label}`;
                  const val = Number(values[subKey] || 0);
                  itemValue += val;
                  if (val > 0) {
                      subItems.push({ label: sub.label, value: val });
                  }
              });
          } else {
              // Process main item directly
              itemValue = Number(values[itemKey] || 0);
          }

          total += itemValue;
          if (itemValue > 0) {
              breakdown.push({
                  label: item.label,
                  value: itemValue,
                  subItems: subItems.length > 0 ? subItems : undefined
              });
          }
      });

      return { breakdown, total };
  };

  const handleEnrichSave = (e: React.FormEvent) => {
      e.preventDefault();
      if (!allKpis) return;
      
      // Find Opportunity Data for reference
      const oppKpi = allKpis.find(k => k.slug === 'taxa_de_oportunidades');
      const oppData = oppKpi?.data.find(d => d.month === launchMonth);
      
      if (!oppData || !oppData.breakdown) {
          // Fallback if no opp data, simple manual entry? Or error.
          // For now, assume Opp data exists as per requirement
          return; 
      }

      const { breakdown, total } = buildAdSpendBreakdown(oppData.breakdown, enrichValues);

      const monthData = kpi.data.find(d => d.month === launchMonth);
      if (monthData) {
        updateMonthlyData({
          kpiId: kpi.id,
          data: {
            year: currentYear,
            realized: total,
            breakdown: breakdown,
            month: launchMonth,
          }
        }, {
          onSuccess: () => {
            setIsEnrichOpen(false);
            setEnrichValues({});
          }
        });
      }
  };

  const handleEditSave = (e: React.FormEvent) => {
    e.preventDefault();
    if (!editData) return;

    // Determinar valores finais
    const finalRealized = editRealized !== '' ? Number(editRealized) : undefined;
    const finalMeta = editMeta !== '' ? Number(editMeta) : undefined;

    // Validar que pelo menos um valor foi alterado
    if (finalRealized === undefined && finalMeta === undefined) {
      return;
    }

    updateMonthlyData({
      kpiId: kpi.id,
      data: {
        year: currentYear,
        realized: finalRealized,
        meta: finalMeta,
        month: editData.month,
        context: editContext || 'Correção manual via interface',
      }
    }, {
      onSuccess: () => {
        setIsEditOpen(false);
        setEditData(null);
        setEditRealized('');
        setEditMeta('');
        setEditContext('');
      }
    });
  };

  // --- CHART LOGIC ---
  const isMonthlyView = selectedMonth !== '---';
  
  // Get all logs from all MonthlyData (logs are stored at MonthlyData level)
  const allLogs = React.useMemo(() => {
    return kpi.data.flatMap(monthData => monthData.logs || []);
  }, [kpi.data]);
  
  const getDailyChartData = (currentRealized: number, currentMeta: number | null, monthStr: string) => {
      const monthIdx = MONTHS.indexOf(monthStr);
      const year = currentYear;
      const daysInMonth = new Date(year, monthIdx + 1, 0).getDate();
      
      // Filter logs for this month
      const relevantLogs = allLogs.filter(l => {
          if (!l.date) return false;
          const d = new Date(l.date);
          // Check if log belongs to this month/year
          return l.month === monthStr;
      });

      const dailyData = [];
      let accumulatedRealized = 0;
      const hasLogs = relevantLogs.length > 0;
      
      // Calculate linear step for fallback (no logs)
      const linearStep = currentRealized / daysInMonth;

      for (let i = 1; i <= daysInMonth; i++) {
          let dayValue = 0;
          
          if (hasLogs) {
              const dayLogs = relevantLogs.filter(l => new Date(l.date).getDate() === i);
              dayLogs.forEach(l => {
                  // Only count positive increments for daily bar chart to show "production"
                  // Use diff if available
                  const diff = l.newValue - (l.oldValue || 0);
                  // We allow negative diffs in accumulation, but logs typically store total snapshot or incremental.
                  // Current handleKpiUpdate stores total in newValue.
                  // But logs list contains multiple entries per day. 
                  // We need to be careful: logs track updates to the MONTH total in the app state, 
                  // but here we are trying to reconstruct the daily flow.
                  // For simplicity: We sum the *net changes* recorded on this day.
                  dayValue += diff; 
              });
              accumulatedRealized += dayValue;
          } else {
              // Fallback: Linear accumulation if no detailed logs exist but we have a total
              accumulatedRealized = linearStep * i;
          }
          
          // Cumulative Target: Proportional to the day of the month
          // Formula: (Total Meta / Total Days) * Current Day
          const cumulativeTarget = currentMeta ? (currentMeta / daysInMonth) * i : 0;
          
          dailyData.push({
              day: i,
              value: accumulatedRealized,
              target: cumulativeTarget
          });
      }

      return dailyData;
  };

  const currentMonthData = isMonthlyView ? kpi.data.find(d => d.month === selectedMonth) : null;
  const chartData = isMonthlyView 
      ? getDailyChartData(currentMonthData?.realized || 0, currentMonthData?.meta || 0, selectedMonth) 
      : kpi.data;

  // Get existing channels for dropdown
  const existingChannels = Array.from(new Set(
      kpi.data.flatMap(d => d.breakdown?.map(b => b.label) || [])
  ));

  // --- PREPARE DATA FOR ENRICHMENT FORM ---
  const getOpportunitiesForMonth = () => {
      if (!allKpis) return null;
      const oppKpi = allKpis.find(k => k.slug === 'taxa_de_oportunidades');
      return oppKpi?.data.find(d => d.month === launchMonth);
  };

  // Logic to find which days have Opportunity logs to show in the date picker
  const availableOppLogs = React.useMemo(() => {
      if (!allKpis || !isAdSpend) return [];
      const oppKpi = allKpis.find(k => k.slug === 'taxa_de_oportunidades');
      if (!oppKpi) return [];

      // Get all logs from all MonthlyData of Opportunities KPI
      const oppLogs = oppKpi.data.flatMap(monthData => monthData.logs || []);

      // Filter logs for the currently selected month in the Enrich Modal
      return oppLogs
          .filter(l => l.month === launchMonth)
          .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
  }, [allKpis, isAdSpend, launchMonth]);

  const oppDataForEnrich = isAdSpend ? getOpportunitiesForMonth() : null;

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div className="flex items-center gap-4">
            <button 
            onClick={onBack}
            className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
            >
            <ArrowLeft size={24} />
            </button>
            <div>
            <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-2">
                {kpi.title}
                {isCplKpi && (
                    <span className="text-xs font-normal bg-amber-500/10 text-amber-400 border border-amber-500/20 px-2 py-0.5 rounded-full flex items-center gap-1">
                        <Calculator size={12} /> Automático
                    </span>
                )}
            </h1>
            <p className="text-gray-500">
                {isMonthlyView ? `Resultado Acumulado - ${selectedMonth}` : 'Visão Anual Consolidada'}
            </p>
            </div>
        </div>
        
        {/* ACTION BUTTONS */}
        <div className="flex gap-2">
            {/* Regular History Button - Hide for CPL */}
            {!isCplKpi && (
                <button 
                    onClick={() => setIsHistoryOpen(true)}
                    className="flex items-center gap-2 px-4 py-2 bg-[#1a1d24] border border-gray-700 hover:border-gray-600 text-gray-300 rounded-lg text-sm font-medium transition-colors"
                >
                    <History size={16} /> Histórico
                </button>
            )}

            {/* Conditional Buttons based on KPI Type */}
            {isAdSpend ? (
                <button
                    onClick={() => setIsEnrichOpen(true)}
                    className="flex items-center gap-2 px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-sm font-medium transition-colors shadow-lg shadow-[#1e5144]/20"
                >
                    <Sparkles size={16} /> Enriquecer Dados
                </button>
            ) : !isCplKpi && (
                // Default Launch Button - Hidden for CPL
                <button
                    onClick={() => setIsLaunchOpen(true)}
                    className="flex items-center gap-2 px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-sm font-medium transition-colors shadow-lg shadow-[#1e5144]/20"
                >
                    <Plus size={16} /> Lançar Dados
                </button>
            )}
        </div>
      </div>

      <div className="bg-[#1a1d24] p-6 rounded-2xl border border-white/5 shadow-lg flex-grow flex flex-col">
          <div className="h-[300px] w-full">
            <ResponsiveContainer width="100%" height="100%">
                {!isMonthlyView ? (
                    <AreaChart data={chartData as any[]} margin={{ top: 20, right: 30, left: 10, bottom: 20 }}>
                        <defs>
                            <linearGradient id="colorRealized" x1="0" y1="0" x2="0" y2="1">
                                <stop offset="5%" stopColor={kpi.color} stopOpacity={0.3}/>
                                <stop offset="95%" stopColor={kpi.color} stopOpacity={0}/>
                            </linearGradient>
                        </defs>
                        <XAxis dataKey="month" axisLine={false} tickLine={false} tick={{fill: '#6b7280'}} dy={10} />
                        <YAxis axisLine={false} tickLine={false} tick={{fill: '#6b7280'}} dx={-10} />
                        <CartesianGrid vertical={false} stroke="#374151" strokeDasharray="3 3" />
                        <Tooltip 
                            contentStyle={{ backgroundColor: '#1f2937', borderColor: '#374151', color: '#f3f4f6' }}
                            itemStyle={{ color: '#e5e7eb' }}
                        />
                        <Area 
                            type="monotone" 
                            dataKey="realized" 
                            stroke={kpi.color} 
                            fillOpacity={1} 
                            fill="url(#colorRealized)" 
                        />
                        <Area 
                            type="monotone" 
                            dataKey="meta" 
                            stroke="#6b7280" 
                            strokeDasharray="5 5" 
                            fill="none" 
                        />
                    </AreaChart>
                ) : (
                    <ComposedChart data={chartData as any[]} margin={{ top: 20, right: 30, left: 10, bottom: 20 }}>
                        <defs>
                            <linearGradient id="colorGradientDaily" x1="0" y1="0" x2="0" y2="1">
                                <stop offset="5%" stopColor={kpi.color} stopOpacity={0.3}/>
                                <stop offset="95%" stopColor={kpi.color} stopOpacity={0}/>
                            </linearGradient>
                        </defs>
                        <XAxis dataKey="day" axisLine={false} tickLine={false} tick={{fill: '#6b7280'}} dy={10} interval={2} />
                        <YAxis axisLine={false} tickLine={false} tick={{fill: '#6b7280'}} dx={-10} />
                        <CartesianGrid vertical={false} stroke="#374151" strokeDasharray="3 3" />
                        <Tooltip 
                            cursor={{fill: '#20232b'}}
                            content={({ active, payload, label }) => {
                                if (active && payload && payload.length) {
                                return (
                                    <div className="bg-[#1f2937] border border-gray-700 p-3 rounded shadow-lg text-xs">
                                        <p className="text-gray-400 mb-1">Dia {label}</p>
                                        <p className="font-bold text-white" style={{ color: kpi.color }}>
                                            Realizado (Acum.): {Number(payload[0].value).toLocaleString('pt-BR', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                        </p>
                                        {payload[1] && (
                                            <p className="font-bold text-amber-400 mt-1">
                                                Meta (Projeção): {Number(payload[1].value).toLocaleString('pt-BR', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                            </p>
                                        )}
                                    </div>
                                );
                                }
                                return null;
                            }}
                        />
                        <Area 
                            type="monotone" 
                            dataKey="value" 
                            name="Realizado (Acumulado)"
                            stroke={kpi.color}
                            fill={`url(#colorGradientDaily)`}
                            fillOpacity={1}
                        />
                        <Line 
                            type="monotone" 
                            dataKey="target" 
                            stroke="#fbbf24" 
                            strokeWidth={2} 
                            strokeDasharray="5 5" 
                            dot={false}
                            name="Meta (Acumulada)"
                        />
                    </ComposedChart>
                )}
            </ResponsiveContainer>
        </div>
        
        <div className="mt-8">
            <h3 className="text-lg font-semibold text-gray-200 mb-4">Registros Mensais</h3>
            <div className="overflow-x-auto">
                <table className="w-full text-left text-sm text-gray-400">
                    <thead className="bg-[#20232b] text-gray-200 uppercase font-semibold text-xs">
                        <tr>
                            <th className="px-4 py-3 w-10"></th> {/* Expand icon col */}
                            <th className="px-4 py-3">Mês</th>
                            <th className="px-4 py-3">Realizado</th>
                            <th className="px-4 py-3">Meta</th>
                            <th className="px-4 py-3">Status</th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-800">
                        {kpi.data.map((row, idx) => (
                            <React.Fragment key={idx}>
                                <tr 
                                    className={`hover:bg-[#20232b]/50 transition-colors ${row.breakdown ? 'cursor-pointer' : ''}`}
                                    onClick={() => row.breakdown && toggleRow(row.month)}
                                >
                                    <td className="px-4 py-3 text-center">
                                        <div className="flex items-center justify-center gap-2">
                                            {row.breakdown && (
                                                <span className="text-gray-500 cursor-pointer" onClick={(e) => { e.stopPropagation(); toggleRow(row.month); }}>
                                                    {expandedRows[row.month] ? <ChevronDown size={16} /> : <ChevronRight size={16} />}
                                                </span>
                                            )}
                                            {/* Edit Button */}
                                            <button
                                                onClick={(e) => {
                                                    e.stopPropagation();
                                                    setEditData(row);
                                                    setEditRealized(row.realized?.toString() || '');
                                                    setEditMeta(row.meta?.toString() || '');
                                                    setEditContext('');
                                                    setIsEditOpen(true);
                                                }}
                                                className="text-gray-500 hover:text-emerald-400 transition-colors"
                                                title="Editar dados mensais"
                                            >
                                                <Pencil size={14} />
                                            </button>
                                            {/* Delete Button - Only show if user has permission */}
                                            <button
                                                onClick={(e) => {
                                                    e.stopPropagation();
                                                    setDeleteConfirm({ id: row.id || '', month: row.month });
                                                }}
                                                className="text-gray-500 hover:text-rose-400 transition-colors"
                                                title="Remover dados mensais"
                                            >
                                                <Trash2 size={16} />
                                            </button>
                                        </div>
                                    </td>
                                    <td className="px-4 py-3 font-medium text-white">{row.month}</td>
                                    <td className="px-4 py-3">{row.realized?.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' }) || '-'}</td>
                                    <td className="px-4 py-3">{row.meta?.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' }) || '-'}</td>
                                    <td className="px-4 py-3">
                                        {row.realized !== null ? (
                                            // Inverse logic for Costs (CPL) - Lower is better
                                            (isCplKpi ? (row.realized <= (row.meta || 0)) : (row.realized >= (row.meta || 0)))
                                            ? <span className="text-emerald-400">Meta OK</span> 
                                            : <span className="text-rose-400">{isCplKpi ? 'Acima' : 'Abaixo'}</span>
                                        ) : '-'}
                                    </td>
                                </tr>
                                
                                {/* Detailed Breakdown Row */}
                                {row.breakdown && expandedRows[row.month] && (
                                    <tr className="bg-[#15171c] animate-in slide-in-from-top-2">
                                        <td colSpan={5} className="px-8 py-4 border-l-2 border-[#1e5144]">
                                            <div className="space-y-4">
                                                <h4 className="text-xs font-bold text-gray-500 uppercase tracking-wider mb-2">Detalhamento do Mês</h4>
                                                {row.breakdown.map((item, bIdx) => (
                                                    <div key={bIdx} className="space-y-1">
                                                        <div className="flex justify-between items-center text-sm">
                                                            <span className="text-gray-300 font-medium">{item.label}</span>
                                                            <span className="text-white font-bold">{item.value.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' })}</span>
                                                        </div>
                                                        {/* Sub-items (Level 2) */}
                                                        {item.subItems && (
                                                            <div className="pl-4 space-y-1 mt-1 border-l border-gray-700 ml-1">
                                                                {item.subItems.map((sub, sIdx) => (
                                                                    <div key={sIdx} className="flex justify-between items-center text-xs text-gray-500">
                                                                        <span>{sub.label}</span>
                                                                        <span>{sub.value.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' })}</span>
                                                                    </div>
                                                                ))}
                                                                <div className="flex justify-between items-center text-xs text-gray-400 pt-1 border-t border-gray-800 mt-1">
                                                                    <span className="italic">Subtotal</span>
                                                                    <span className="font-semibold">{item.value.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' })}</span>
                                                                </div>
                                                            </div>
                                                        )}
                                                    </div>
                                                ))}
                                                <div className="pt-3 border-t border-gray-700 flex justify-between items-center">
                                                    <span className="text-sm font-bold text-emerald-400">TOTAL</span>
                                                    <span className="text-sm font-bold text-white">{row.realized?.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' })}</span>
                                                </div>
                                            </div>
                                        </td>
                                    </tr>
                                )}
                            </React.Fragment>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
      </div>

      {/* --- LAUNCH MODAL (GENERIC) --- */}
      {isLaunchOpen && !isAdSpend && !isCplKpi && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsLaunchOpen(false)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-700 overflow-hidden"
                onClick={e => e.stopPropagation()}
            >
                <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
                    <h3 className="text-lg font-bold text-gray-100">Lançar Novos Dados</h3>
                    <button onClick={() => setIsLaunchOpen(false)} className="text-gray-500 hover:text-white">
                        <X size={20} />
                    </button>
                </div>
                <form onSubmit={handleLaunch} className="p-6 space-y-4">
                    
                    <div>
                        <label className="block text-sm font-medium text-gray-400 mb-1">Data do Registro</label>
                        <input 
                            type="date"
                            value={launchDate}
                            onChange={handleDateChange}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none"
                            required
                        />
                        <div className="mt-1 flex justify-between text-xs text-gray-500">
                            <span>Mês detectado: <span className="font-bold text-emerald-400">{launchMonth}</span></span>
                        </div>
                    </div>

                    {/* OPPORTUNITIES SPECIFIC FIELDS */}
                    {isOppKpi ? (
                        <>
                            <div>
                                <label className="block text-sm font-medium text-gray-400 mb-1">Canal / Origem</label>
                                {!isNewChannel ? (
                                    <select 
                                        value={oppChannel}
                                        onChange={(e) => {
                                            const val = e.target.value;
                                            if (val === '__new__') {
                                                setIsNewChannel(true);
                                                setOppChannel('');
                                            } else {
                                                setOppChannel(val);
                                            }
                                        }}
                                        className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none mb-2"
                                        required
                                    >
                                        <option value="" disabled>Selecione o canal...</option>
                                        {existingChannels.map(c => <option key={c} value={c}>{c}</option>)}
                                        <option value="__new__" className="text-emerald-400 font-bold">+ Adicionar Novo Canal</option>
                                    </select>
                                ) : (
                                    <div className="flex gap-2 items-center mb-2">
                                        <input 
                                            type="text" 
                                            value={oppChannel}
                                            onChange={(e) => setOppChannel(e.target.value)}
                                            placeholder="Nome do novo canal..."
                                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144]"
                                            autoFocus
                                            required
                                        />
                                        <button 
                                            type="button" 
                                            onClick={() => setIsNewChannel(false)}
                                            className="text-gray-500 hover:text-white"
                                            title="Cancelar"
                                        >
                                            <X size={20} />
                                        </button>
                                    </div>
                                )}
                            </div>

                            {(oppChannel === 'Campanhas Ads Meta (Instagram/Facebook)' || isNewChannel) && (
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">
                                        Sub-item <span className="text-gray-500 font-normal">(Opcional: ex: Região, Campanha)</span>
                                    </label>
                                    <input 
                                        type="text" 
                                        value={oppSubItem}
                                        onChange={(e) => setOppSubItem(e.target.value)}
                                        placeholder="Ex: Brasil, RJ, Campanha Verão..."
                                        className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none"
                                    />
                                </div>
                            )}
                        </>
                    ) : (
                        /* GENERIC CONTEXT FIELDS */
                        <>
                            {isRepKpi && (
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">Representante</label>
                                    <select 
                                        value={launchContext}
                                        onChange={(e) => setLaunchContext(e.target.value)}
                                        className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none"
                                        required
                                    >
                                        <option value="" disabled>Selecione...</option>
                                        {representatives.map(rep => <option key={rep.uuid} value={rep.name}>{rep.name}</option>)}
                                    </select>
                                </div>
                            )}

                            {isChannelKpi && (
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">Canal de Mídia</label>
                                    <select 
                                        value={launchContext}
                                        onChange={(e) => setLaunchContext(e.target.value)}
                                        className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none"
                                        required
                                    >
                                        <option value="" disabled>Selecione...</option>
                                        <option value="Google Ads">Google Ads</option>
                                        <option value="Meta Ads">Meta Ads</option>
                                        <option value="Orgânico">Orgânico</option>
                                        <option value="E-mail">E-mail Marketing</option>
                                    </select>
                                </div>
                            )}
                        </>
                    )}

                    <div>
                        <label className="block text-sm font-medium text-gray-400 mb-1">
                            {kpi.unit === 'currency' ? 'Valor (R$)' : kpi.unit === 'percent' ? 'Percentual (%)' : 'Quantidade'}
                        </label>
                        <input 
                            type="number" 
                            step={kpi.unit === 'percent' ? "0.01" : "1"}
                            value={launchValue}
                            onChange={(e) => setLaunchValue(e.target.value)}
                            placeholder="0"
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none"
                            required
                        />
                        {isOppKpi && (
                            <p className="text-xs text-emerald-500 mt-1">
                                * Este valor será somado ao total atual do canal/mês.
                            </p>
                        )}
                    </div>

                    <div className="pt-4 flex justify-end gap-3">
                        <button 
                            type="button" 
                            onClick={() => setIsLaunchOpen(false)}
                            className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium"
                        >
                            Cancelar
                        </button>
                        <button 
                            type="submit" 
                            disabled={isPending}
                            className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium shadow-lg disabled:opacity-50"
                        >
                            {isPending ? 'Salvando...' : 'Salvar Lançamento'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
      )}

      {/* --- ENRICHMENT MODAL (AD SPEND) --- */}
      {isEnrichOpen && isAdSpend && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsEnrichOpen(false)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-lg rounded-2xl shadow-2xl border border-gray-700 overflow-hidden flex flex-col max-h-[90vh]"
                onClick={e => e.stopPropagation()}
            >
                <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
                    <div>
                        <h3 className="text-lg font-bold text-gray-100 flex items-center gap-2">
                            <Sparkles size={18} className="text-amber-400" />
                            Enriquecer Dados
                        </h3>
                        <p className="text-sm text-gray-500">Cruzar investimento com oportunidades</p>
                    </div>
                    <button onClick={() => setIsEnrichOpen(false)} className="text-gray-500 hover:text-white">
                        <X size={20} />
                    </button>
                </div>
                
                <form onSubmit={handleEnrichSave} className="flex-grow overflow-y-auto custom-scrollbar p-6 space-y-6">
                    
                    {/* Step 1: Select Month */}
                    <div>
                        <label className="block text-sm font-medium text-gray-400 mb-1">Mês de Referência</label>
                        <select
                            value={launchMonth}
                            onChange={(e) => {
                                setLaunchMonth(e.target.value);
                                setLaunchDate(''); // Reset specific day/batch when month changes
                            }}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none appearance-none cursor-pointer"
                        >
                            {MONTHS.map(m => (
                                <option key={m} value={m}>
                                    {monthNames ? monthNames[m] : m}
                                </option>
                            ))}
                        </select>
                    </div>

                    {/* Step 2: Select Batch/Day */}
                    <div>
                        <label className="block text-sm font-medium text-gray-400 mb-1">Lote de Oportunidades</label>
                        <div className="relative">
                            <select
                                value={launchDate}
                                onChange={(e) => setLaunchDate(e.target.value)}
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none appearance-none cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
                                required
                                disabled={!launchMonth}
                            >
                                <option value="" disabled>Selecione um lote no mês de {launchMonth}...</option>
                                {availableOppLogs.length > 0 ? (
                                    availableOppLogs.map((log) => {
                                        const displayDate = log.date.split('T')[0].split('-').reverse().join('/');
                                        return (
                                            <option key={log.id} value={log.date.split('T')[0]}>
                                                Dia {displayDate.split('/')[0]} - {log.context} ({log.newValue > (log.oldValue || 0) ? `+${log.newValue - (log.oldValue || 0)}` : log.newValue} leads)
                                            </option>
                                        );
                                    })
                                ) : (
                                    <option value="" disabled>Nenhum lançamento de oportunidade encontrado em {launchMonth}</option>
                                )}
                            </select>
                            <Calendar size={16} className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                        </div>
                        <p className="text-xs text-gray-500 mt-1">Selecione o lote de leads para atribuir o custo corretamente.</p>
                    </div>

                    {/* Step 3: Inputs */}
                    <div className="space-y-4">
                        {oppDataForEnrich?.breakdown ? (
                            oppDataForEnrich.breakdown.map((item, idx) => {
                                // Recursive render for nested items
                                const renderInput = (label: string, key: string, volume: number) => (
                                    <div key={key} className="flex items-center gap-4">
                                        <div className="flex-grow">
                                            <label className="block text-sm font-medium text-gray-300">{label}</label>
                                            <p className="text-xs text-gray-500">Volume Total Mês: <span className="text-emerald-400 font-bold">{volume} leads</span></p>
                                        </div>
                                        <div className="w-40 relative">
                                            <span className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 text-xs">R$</span>
                                            <input 
                                                type="number" 
                                                step="0.01"
                                                placeholder="0.00"
                                                value={enrichValues[key] || ''}
                                                onChange={(e) => setEnrichValues(prev => ({ ...prev, [key]: e.target.value }))}
                                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg pl-8 pr-3 py-2 text-right text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent text-sm"
                                            />
                                        </div>
                                    </div>
                                );

                                if (item.subItems && item.subItems.length > 0) {
                                    return (
                                        <div key={idx} className="border border-gray-800 rounded-xl p-4 bg-[#0f1115]">
                                            <h4 className="font-bold text-gray-200 mb-3 border-b border-gray-800 pb-2">{item.label}</h4>
                                            <div className="space-y-3 pl-2">
                                                {item.subItems.map((sub, sIdx) => 
                                                    renderInput(sub.label, `${item.label}-${sub.label}`, sub.value)
                                                )}
                                            </div>
                                        </div>
                                    );
                                } else {
                                    return (
                                        <div key={idx} className="border border-gray-800 rounded-xl p-4 bg-[#0f1115]">
                                            {renderInput(item.label, item.label, item.value)}
                                        </div>
                                    );
                                }
                            })
                        ) : (
                            <div className="text-center py-8 text-gray-500">
                                <AlertCircle className="mx-auto mb-2 opacity-50" />
                                <p>Não há dados de oportunidades lançados para {launchMonth} para cruzar informações.</p>
                            </div>
                        )}
                    </div>
                </form>

                <div className="p-4 border-t border-gray-800 bg-[#20232b] flex justify-end gap-3">
                    <button 
                        type="button" 
                        onClick={() => setIsEnrichOpen(false)}
                        className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium"
                    >
                        Cancelar
                    </button>
                    <button 
                        onClick={handleEnrichSave}
                        disabled={isPending || !oppDataForEnrich?.breakdown}
                        className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium shadow-lg disabled:opacity-50"
                    >
                        {isPending ? 'Salvando...' : 'Salvar Enriquecimento'}
                    </button>
                </div>
            </div>
        </div>
      )}

      {/* --- HISTORY MODAL --- */}
      {isHistoryOpen && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsHistoryOpen(false)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-2xl rounded-2xl shadow-2xl border border-gray-700 overflow-hidden flex flex-col max-h-[80vh]"
                onClick={e => e.stopPropagation()}
            >
                <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
                    <div>
                        <h3 className="text-lg font-bold text-gray-100">Histórico de Alterações</h3>
                        <p className="text-sm text-gray-500">Log de auditoria e correções</p>
                    </div>
                    <button onClick={() => setIsHistoryOpen(false)} className="text-gray-500 hover:text-white">
                        <X size={20} />
                    </button>
                </div>
                
                <div className="flex-grow overflow-y-auto custom-scrollbar p-6">
                    {allLogs.length === 0 ? (
                        <div className="text-center text-gray-500 py-8">
                            Nenhum histórico disponível.
                        </div>
                    ) : (
                        <div className="space-y-4">
                            {allLogs.map((log) => (
                                <div key={log.id} className="flex items-start gap-4 p-4 bg-[#0f1115] rounded-xl border border-gray-800 hover:border-gray-700 transition-colors group">
                                    <div className="mt-1 p-2 bg-[#1e5144]/10 text-emerald-400 rounded-lg">
                                        {log.action === 'create' ? <Plus size={16} /> : <X size={16} />}
                                    </div>
                                    <div className="flex-grow">
                                        <div className="flex justify-between items-start">
                                            <div>
                                                <p className="text-sm font-medium text-gray-200">
                                                    {log.action === 'create' ? 'Novo Lançamento' : 'Correção Manual'}
                                                    <span className="text-gray-500 font-normal ml-2">por {log.user}</span>
                                                </p>
                                                <p className="text-xs text-gray-500 mt-0.5">{log.date.split('T')[0].split('-').reverse().join('/')} às {log.timestamp}</p>
                                            </div>
                                            
                                            <div className="text-right">
                                                <p className="text-sm font-bold text-white">
                                                    {log.newValue.toLocaleString('pt-BR', { style: kpi.unit === 'currency' ? 'currency' : 'decimal', currency: 'BRL' })}
                                                </p>
                                                {log.context && (
                                                    <span className="text-xs text-[#1e5144] bg-[#1e5144]/10 px-2 py-0.5 rounded border border-[#1e5144]/20">
                                                        {log.context}
                                                    </span>
                                                )}
                                            </div>
                                        </div>
                                        {log.month && (
                                            <p className="text-xs text-gray-600 mt-2">
                                                Referência: <span className="font-medium text-gray-400">{log.month}</span>
                                            </p>
                                        )}
                                    </div>
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
      )}
      
      {/* --- DELETE CONFIRMATION MODAL --- */}
      {deleteConfirm && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm">
          <div className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-700 overflow-hidden">
            <div className="p-6 border-b border-gray-800">
              <h3 className="text-lg font-bold text-gray-100">Confirmar Remoção</h3>
              <p className="text-sm text-gray-400 mt-2">
                Tem certeza que deseja remover os dados mensais de <span className="font-bold text-white">{deleteConfirm.month}</span>?
              </p>
            </div>
            <div className="p-6 flex justify-end gap-3 bg-[#20232b]">
              <button
                onClick={() => setDeleteConfirm(null)}
                className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium"
              >
                Cancelar
              </button>
              <button
                onClick={() => {
                  if (deleteConfirm.id) {
                    deleteMonthlyData({ kpiId: kpi.id, monthlyDataId: deleteConfirm.id });
                    setDeleteConfirm(null);
                  }
                }}
                disabled={isDeleting}
                className="px-4 py-2 bg-rose-600 text-white hover:bg-rose-700 rounded-lg text-sm font-medium shadow-lg disabled:opacity-50"
              >
                {isDeleting ? 'Removendo...' : 'Remover'}
              </button>
            </div>
          </div>
        </div>
      )}

      {/* --- EDIT MODAL --- */}
      {isEditOpen && editData && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsEditOpen(false)}>
          <div
            className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-700 overflow-hidden"
            onClick={e => e.stopPropagation()}
          >
            {/* Header */}
            <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
              <div>
                <h3 className="text-lg font-bold text-gray-100 flex items-center gap-2">
                  <Pencil size={18} className="text-emerald-400" />
                  Editar Dados Mensais
                </h3>
                <p className="text-sm text-gray-500">
                  Mês: <span className="font-bold text-white">{editData.month}</span>
                </p>
              </div>
              <button onClick={() => setIsEditOpen(false)} className="text-gray-500 hover:text-white">
                <X size={20} />
              </button>
            </div>

            {/* Form */}
            <form onSubmit={handleEditSave} className="p-6 space-y-4">

              {/* Campo: Realized */}
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Valor Realizado
                </label>
                <div className="relative">
                  {kpi.unit === 'currency' && (
                    <span className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 text-sm">R$</span>
                  )}
                  <input
                    type="number"
                    step={kpi.unit === 'percent' ? "0.01" : "1"}
                    value={editRealized}
                    onChange={(e) => setEditRealized(e.target.value)}
                    placeholder={editData.realized?.toString() || '0'}
                    className={`w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none ${kpi.unit === 'currency' ? 'pl-10' : ''}`}
                  />
                </div>
                <p className="text-xs text-gray-500 mt-1">
                  Valor atual: {editData.realized?.toLocaleString('pt-BR', {
                    style: kpi.unit === 'currency' ? 'currency' : 'decimal',
                    currency: 'BRL'
                  })}
                </p>
              </div>

              {/* Campo: Meta */}
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Meta do Mês
                </label>
                <div className="relative">
                  {kpi.unit === 'currency' && (
                    <span className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 text-sm">R$</span>
                  )}
                  <input
                    type="number"
                    step={kpi.unit === 'percent' ? "0.01" : "1"}
                    value={editMeta}
                    onChange={(e) => setEditMeta(e.target.value)}
                    placeholder={editData.meta?.toString() || '0'}
                    className={`w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none ${kpi.unit === 'currency' ? 'pl-10' : ''}`}
                  />
                </div>
                <p className="text-xs text-gray-500 mt-1">
                  Meta atual: {editData.meta?.toLocaleString('pt-BR', {
                    style: kpi.unit === 'currency' ? 'currency' : 'decimal',
                    currency: 'BRL'
                  })}
                </p>
              </div>

              {/* Campo: Contexto (opcional) */}
              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">
                  Contexto da Edição <span className="text-gray-600 font-normal">(opcional)</span>
                </label>
                <input
                  type="text"
                  value={editContext}
                  onChange={(e) => setEditContext(e.target.value)}
                  placeholder="Ex: Correção de valor lançado errado..."
                  className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:ring-1 focus:ring-[#1e5144] focus:border-transparent outline-none"
                />
              </div>

              {/* Aviso de Auditoria */}
              <div className="bg-amber-500/10 border border-amber-500/20 rounded-lg p-3">
                <p className="text-xs text-amber-400 flex items-start gap-2">
                  <AlertCircle size={14} className="mt-0.5 flex-shrink-0" />
                  <span>
                    Esta alteração será registrada no log de auditoria com seu nome e data/hora.
                  </span>
                </p>
              </div>

              {/* Botões */}
              <div className="pt-4 flex justify-end gap-3">
                <button
                  type="button"
                  onClick={() => setIsEditOpen(false)}
                  className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium"
                >
                  Cancelar
                </button>
                <button
                  type="submit"
                  disabled={isPending || (!editRealized && !editMeta)}
                  className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium shadow-lg disabled:opacity-50"
                >
                  {isPending ? 'Salvando...' : 'Salvar Alterações'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};

export default KpiDetailView;