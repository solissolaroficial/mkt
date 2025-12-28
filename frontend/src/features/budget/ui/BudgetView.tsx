 
import React, { useState, useMemo } from 'react';
import { ArrowLeft, Download, Search, Edit2, Calculator, ChevronRight, ChevronDown, Folder, FolderOpen, FileText } from 'lucide-react';
import { MONTHS } from '@/shared/utils/constants';
import type { BudgetItem, BudgetViewProps, EditingCell } from '../types';

const MONTHS_LIST = [
  "Jan/25", "Fev/26", "Mar/26", "Abr/26", "Mai/26", "Jun/26",
  "Jul/26", "Ago/26", "Set/26", "Out/26", "Nov/26", "Dez/26"
];

// Helper to parse "1.000" or "-" to number
const parseVal = (v: string): number => {
  if (!v || v === '-') return 0;
  return parseFloat(v.replace(/\./g, '').replace(',', '.'));
};

const BudgetView: React.FC<BudgetViewProps> = ({ onBack, filterFn, customTitle, customSubtitle }) => {
  const [dataItems, setDataItems] = useState<BudgetItem[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  
  // Edit State: Now includes type (budget or realized)
  const [editingCell, setEditingCell] = useState<EditingCell | null>(null);
  const [editValue, setEditValue] = useState('');
  
  // Collapse State
  const [expandedGroups, setExpandedGroups] = useState<Record<string, boolean>>({
      'EVENTOS & REUNIÕES': true, // Open by default
      'BRINDES': false,
      'PROPAGANDAS & PUBLICIDADE': true,
      'MÍDIA DIGITAL': false,
      'MÍDIA TRADICIONAL': false,
      'APOIO MARKETING': true,
      'PROMOÇÕES LOCAIS': false
  });
  
  const [expandedSubgroups, setExpandedSubgroups] = useState<Record<string, boolean>>({});

  const toggleGroup = (grp: string) => {
      setExpandedGroups(prev => ({ ...prev, [grp]: !prev[grp] }));
  };

  const toggleSubgroup = (key: string) => {
      setExpandedSubgroups(prev => ({ ...prev, [key]: !prev[key] }));
  };

  // --- HIERARCHY BUILDING & SUMMATION ---
  const hierarchy = useMemo(() => {
      const groups: Record<string, any> = {};

      dataItems.forEach(item => {
          // Filter by search
          if (searchTerm) {
              const term = searchTerm.toLowerCase();
              const match = 
                  item.obj.toLowerCase().includes(term) ||
                  item.grp.toLowerCase().includes(term) ||
                  item.desc.toLowerCase().includes(term) ||
                  item.codObj.includes(term) ||
                  item.codGrp.includes(term);
              if (!match) return;
          }

          if (!groups[item.obj]) {
              groups[item.obj] = {
                  id: item.obj,
                  codObj: item.codObj,
                  name: item.obj,
                  items: [],
                  subgroups: {},
                  totals: Array(12).fill(0),
                  realizedTotals: Array(12).fill(0)
              };
          }

          const group = groups[item.obj];
          const subKey = item.grp || 'Outros';

          if (!group.subgroups[subKey]) {
              group.subgroups[subKey] = {
                  id: `${item.obj}-${subKey}`,
                  codGrp: item.codGrp,
                  name: subKey,
                  items: [],
                  totals: Array(12).fill(0),
                  realizedTotals: Array(12).fill(0)
              };
          }

          const subgroup = group.subgroups[subKey];
          subgroup.items.push(item);

          // Add to totals
          item.vals.forEach((v, i) => {
              subgroup.totals[i] += v;
              group.totals[i] += v;
          });
          // Add to realized totals
          item.realizedVals.forEach((v, i) => {
              subgroup.realizedTotals[i] += v;
              group.realizedTotals[i] += v;
          });
      });

      // Automatically expand groups if we are in a filtered view (like Marketing Actions)
      if (filterFn) {
          const keys = Object.values(groups).map((g: any) => g.name);
          const newExpanded = { ...expandedGroups };
          keys.forEach(k => { newExpanded[k] = true; });
      }

      return Object.values(groups);
  }, [dataItems, searchTerm, filterFn]);

  // Calculate Grand Totals for Footer (Budget vs Realized)
  const monthlyGrandTotals = useMemo(() => {
      const budgetTotals = Array(12).fill(0);
      const realizedTotals = Array(12).fill(0);

      hierarchy.forEach((group: any) => {
          group.totals.forEach((val: number, idx: number) => {
              budgetTotals[idx] += val;
          });
          group.realizedTotals.forEach((val: number, idx: number) => {
              realizedTotals[idx] += val;
          });
      });

      return { budgetTotals, realizedTotals };
  }, [hierarchy]);

  // Expand groups automatically when searching
  useMemo(() => {
      if (searchTerm) {
          const newExpandedGroups = { ...expandedGroups };
          const newExpandedSub = { ...expandedSubgroups };
          hierarchy.forEach((g: any) => {
              newExpandedGroups[g.name] = true;
              Object.values(g.subgroups).forEach((s: any) => {
                  newExpandedSub[s.id] = true;
              });
          });
          setExpandedGroups(newExpandedGroups);
          setExpandedSubgroups(newExpandedSub);
      }
  }, [searchTerm, hierarchy.length]);

  // Formatting
  const formatCurrency = (val: number) => {
    if (val === 0) return '-';
    // Format: 1.000 or 1.000,50
    return val.toLocaleString('pt-BR', { minimumFractionDigits: 0, maximumFractionDigits: 0 });
  };

  // Edit Logic
  const handleCellClick = (item: BudgetItem, colIndex: number, type: 'budget' | 'realized') => {
    setEditingCell({ id: item.id, colIndex, type });
    const val = type === 'budget' ? item.vals[colIndex] : item.realizedVals[colIndex];
    setEditValue(val === 0 ? '' : val.toLocaleString('pt-BR'));
  };

  const handleSaveEdit = () => {
    if (!editingCell) return;

    const numVal = parseVal(editValue);
    
    setDataItems(prev => prev.map(item => {
        if (item.id === editingCell.id) {
            if (editingCell.type === 'budget') {
                const newVals = [...item.vals];
                newVals[editingCell.colIndex] = numVal;
                return { ...item, vals: newVals };
            } else {
                const newRealizedVals = [...item.realizedVals];
                newRealizedVals[editingCell.colIndex] = numVal;
                return { ...item, realizedVals: newRealizedVals };
            }
        }
        return item;
    }));

    setEditingCell(null);
    setEditValue('');
  };

  const handleExport = () => {
    let csvContent = "data:text/csv;charset=utf-8,";
    // Header for 3-column structure in CSV
    let headerRow = "Nivel;Codigo;Descricao;";
    MONTHS_LIST.forEach(m => {
        headerRow += `${m} - Orçado;${m} - Realizado;${m} - Diferença;`;
    });
    headerRow += "Total Orçado;Total Realizado;Total Diferença\n";
    csvContent += headerRow;

    const appendRow = (level: string, code: string, name: string, budgets: number[], realizeds: number[]) => {
        let rowStr = `${level};${code};${name};`;
        let totalB = 0, totalR = 0;
        
        budgets.forEach((b, i) => {
            const r = realizeds[i];
            const diff = b - r;
            totalB += b;
            totalR += r;
            rowStr += `${formatCurrency(b)};${formatCurrency(r)};${formatCurrency(diff)};`;
        });
        
        const totalDiff = totalB - totalR;
        rowStr += `${formatCurrency(totalB)};${formatCurrency(totalR)};${formatCurrency(totalDiff)}\n`;
        return rowStr;
    };

    hierarchy.forEach((g: any) => {
        csvContent += appendRow("Grupo", g.codObj, g.name, g.totals, g.realizedTotals);

        Object.values(g.subgroups).forEach((s: any) => {
            csvContent += appendRow("Subgrupo", s.codGrp, s.name, s.totals, s.realizedTotals);

            s.items.forEach((item: BudgetItem) => {
                csvContent += appendRow("Item", item.cod, item.desc, item.vals, item.realizedVals);
            });
        });
    });

    const encodedUri = encodeURI(csvContent);
    const link = document.createElement("a");
    link.setAttribute("href", encodedUri);
    link.setAttribute("download", "orcamento_detalhado.csv");
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
        {/* Header */}
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
            <div className="flex items-center gap-4">
                <button onClick={onBack} className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white">
                    <ArrowLeft size={24} />
                </button>
                <div>
                    <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-2">
                        {customTitle || "Orçamento 2025"} <span className="text-emerald-500 text-lg font-normal">(Simulação 2026)</span>
                    </h1>
                    <p className="text-gray-500 flex items-center gap-2">
                        <Calculator size={14} /> {customSubtitle || "Detalhamento com Orçado x Realizado"}
                    </p>
                </div>
            </div>
            <div className="flex gap-3 w-full md:w-auto">
                <div className="relative flex-grow md:flex-grow-0">
                    <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" size={16} />
                    <input 
                        type="text" 
                        placeholder="Buscar item ou grupo..." 
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="w-full md:w-64 pl-10 pr-4 py-2 bg-[#1a1d24] border border-gray-700 rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm placeholder-gray-600"
                    />
                </div>
                <button onClick={handleExport} className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm flex items-center gap-2 transition-colors shadow-lg">
                    <Download size={16} /> Exportar
                </button>
            </div>
        </div>

        {/* Tree Table */}
        <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden flex-grow flex flex-col relative">
            <div className="overflow-auto custom-scrollbar flex-grow">
                <table className="w-full text-xs text-left whitespace-nowrap border-collapse">
                    <thead className="bg-[#15171c] text-gray-400 uppercase font-semibold sticky top-0 z-20 shadow-md">
                        {/* Row 1: Months */}
                        <tr>
                            <th className="px-4 py-3 border-b border-gray-800 min-w-[300px] sticky left-0 z-30 bg-[#15171c] border-r shadow-[2px_0_5px_-2px_rgba(0,0,0,0.3)]" rowSpan={2}>
                                Estrutura de Custos
                            </th>
                            {MONTHS_LIST.map((month, i) => (
                                <th key={i} colSpan={3} className="px-2 py-2 border-b border-gray-800 text-center border-r border-gray-800/50">
                                    {month}
                                </th>
                            ))}
                            <th colSpan={3} className="px-2 py-2 border-b border-gray-800 text-center bg-[#15171c] sticky right-0 z-30 border-l shadow-[-2px_0_5px_-2px_rgba(0,0,0,0.3)]">
                                Total 2026
                            </th>
                        </tr>
                        {/* Row 2: Sub-columns */}
                        <tr>
                            {MONTHS_LIST.map((_, i) => (
                                <React.Fragment key={i}>
                                    <th className="px-2 py-1 border-b border-gray-800 text-right min-w-[70px] text-[10px] text-gray-500">Orç</th>
                                    <th className="px-2 py-1 border-b border-gray-800 text-right min-w-[70px] text-[10px] text-blue-400">Real</th>
                                    <th className="px-2 py-1 border-b border-gray-800 text-right min-w-[70px] text-[10px] font-bold border-r border-gray-800">Dif</th>
                                </React.Fragment>
                            ))}
                            {/* Total Columns */}
                            <th className="px-2 py-1 border-b border-gray-800 text-right min-w-[80px] text-[10px] text-gray-500 bg-[#15171c] sticky z-30 border-l" style={{right: 160}}>Orç</th>
                            <th className="px-2 py-1 border-b border-gray-800 text-right min-w-[80px] text-[10px] text-blue-400 bg-[#15171c] sticky z-30" style={{right: 80}}>Real</th>
                            <th className="px-2 py-1 border-b border-gray-800 text-right min-w-[80px] text-[10px] font-bold bg-[#15171c] sticky right-0 z-30">Dif</th>
                        </tr>
                    </thead>
                    <tbody className="bg-[#1a1d24]">
                        {hierarchy.map((group: any) => {
                            const isGroupExpanded = expandedGroups[group.name];
                            
                            const groupBudgetTotal = group.totals.reduce((a:number,b:number)=>a+b,0);
                            const groupRealizedTotal = group.realizedTotals.reduce((a:number,b:number)=>a+b,0);
                            const groupDiffTotal = groupBudgetTotal - groupRealizedTotal;

                            return (
                                <React.Fragment key={group.id}>
                                    {/* Level 1: Group Header */}
                                    <tr 
                                        className="bg-[#152e29] hover:bg-[#1b3a34] cursor-pointer transition-colors border-b border-gray-800/50"
                                        onClick={() => toggleGroup(group.name)}
                                    >
                                        <td className="px-4 py-3 sticky left-0 z-10 bg-[#152e29] border-r border-gray-800/30 flex items-center gap-2">
                                            {isGroupExpanded ? <ChevronDown size={14} className="text-emerald-400" /> : <ChevronRight size={14} className="text-gray-400" />}
                                            <span className="font-bold text-emerald-400">{group.codObj} - {group.name}</span>
                                        </td>
                                        {group.totals.map((budget: number, i: number) => {
                                            const realized = group.realizedTotals[i];
                                            const diff = budget - realized;
                                            return (
                                                <React.Fragment key={i}>
                                                    <td className="px-2 py-3 text-right font-medium text-emerald-100/70 text-[11px]">{formatCurrency(budget)}</td>
                                                    <td className="px-2 py-3 text-right font-medium text-blue-300 text-[11px]">{formatCurrency(realized)}</td>
                                                    <td className={`px-2 py-3 text-right font-bold text-[11px] border-r border-gray-800/30 ${diff >= 0 ? 'text-emerald-400' : 'text-rose-400'}`}>{formatCurrency(diff)}</td>
                                                </React.Fragment>
                                            );
                                        })}
                                        {/* Sticky Total Columns */}
                                        <td className="px-2 py-3 text-right font-bold text-emerald-100 bg-[#152e29] sticky z-10 border-l border-gray-800/30" style={{right: 160}}>{formatCurrency(groupBudgetTotal)}</td>
                                        <td className="px-2 py-3 text-right font-bold text-blue-300 bg-[#152e29] sticky z-10" style={{right: 80}}>{formatCurrency(groupRealizedTotal)}</td>
                                        <td className={`px-2 py-3 text-right font-extrabold bg-[#152e29] sticky right-0 z-10 ${groupDiffTotal >= 0 ? 'text-emerald-400' : 'text-rose-400'}`}>{formatCurrency(groupDiffTotal)}</td>
                                    </tr>

                                    {/* Level 2: Subgroups */}
                                    {isGroupExpanded && Object.values(group.subgroups).map((subgroup: any) => {
                                        const isSubExpanded = expandedSubgroups[subgroup.id];
                                        
                                        const subBudgetTotal = subgroup.totals.reduce((a:number,b:number)=>a+b,0);
                                        const subRealizedTotal = subgroup.realizedTotals.reduce((a:number,b:number)=>a+b,0);
                                        const subDiffTotal = subBudgetTotal - subRealizedTotal;

                                        return (
                                            <React.Fragment key={subgroup.id}>
                                                <tr 
                                                    className="bg-[#20232b] hover:bg-[#252a33] cursor-pointer transition-colors border-b border-gray-800/30"
                                                    onClick={() => toggleSubgroup(subgroup.id)}
                                                >
                                                    <td className="px-4 py-2 sticky left-0 z-10 bg-[#20232b] border-r border-gray-800/30 pl-8 flex items-center gap-2">
                                                        {isSubExpanded ? <FolderOpen size={14} className="text-blue-400" /> : <Folder size={14} className="text-gray-500" />}
                                                        <span className="font-semibold text-gray-300">{subgroup.codGrp} - {subgroup.name}</span>
                                                    </td>
                                                    {subgroup.totals.map((budget: number, i: number) => {
                                                        const realized = subgroup.realizedTotals[i];
                                                        const diff = budget - realized;
                                                        return (
                                                            <React.Fragment key={i}>
                                                                <td className="px-2 py-2 text-right font-medium text-gray-400 text-[11px]">{formatCurrency(budget)}</td>
                                                                <td className="px-2 py-2 text-right font-medium text-blue-400/70 text-[11px]">{formatCurrency(realized)}</td>
                                                                <td className={`px-2 py-2 text-right font-bold text-[11px] border-r border-gray-800/30 ${diff >= 0 ? 'text-emerald-500/70' : 'text-rose-500/70'}`}>{formatCurrency(diff)}</td>
                                                            </React.Fragment>
                                                        );
                                                    })}
                                                    <td className="px-2 py-2 text-right font-medium text-gray-300 bg-[#20232b] sticky z-10 border-l border-gray-800/30" style={{right: 160}}>{formatCurrency(subBudgetTotal)}</td>
                                                    <td className="px-2 py-2 text-right font-medium text-blue-300 bg-[#20232b] sticky z-10" style={{right: 80}}>{formatCurrency(subRealizedTotal)}</td>
                                                    <td className={`px-2 py-2 text-right font-bold bg-[#20232b] sticky right-0 z-10 ${subDiffTotal >= 0 ? 'text-emerald-500' : 'text-rose-500'}`}>{formatCurrency(subDiffTotal)}</td>
                                                </tr>

                                                {/* Level 3: Items */}
                                                {isSubExpanded && subgroup.items.map((item: BudgetItem) => {
                                                    const itemBudgetTotal = item.vals.reduce((a,b)=>a+b,0);
                                                    const itemRealizedTotal = item.realizedVals.reduce((a,b)=>a+b,0);
                                                    const itemDiffTotal = itemBudgetTotal - itemRealizedTotal;

                                                    return (
                                                        <tr key={item.id} className="hover:bg-white/5 transition-colors border-b border-gray-800/20 group">
                                                            <td className="px-4 py-2 sticky left-0 z-10 bg-[#1a1d24] group-hover:bg-[#20232b] border-r border-gray-800 pl-12 flex items-center gap-2">
                                                                <FileText size={12} className="text-gray-600" />
                                                                <span className="text-gray-400 truncate max-w-xs" title={item.desc}>
                                                                    {item.cod} - {item.desc}
                                                                </span>
                                                            </td>
                                                             
                                                            {item.vals.map((budget, i) => {
                                                                const realized = item.realizedVals[i];
                                                                const diff = budget - realized;
                                                                const isEditingBudget = editingCell?.id === item.id && editingCell?.colIndex === i && editingCell?.type === 'budget';
                                                                const isEditingRealized = editingCell?.id === item.id && editingCell?.colIndex === i && editingCell?.type === 'realized';

                                                                return (
                                                                    <React.Fragment key={i}>
                                                                        {/* Budget Cell */}
                                                                        <td 
                                                                            onDoubleClick={() => handleCellClick(item, i, 'budget')}
                                                                            className={`px-2 py-2 text-right relative cursor-pointer text-[11px] ${isEditingBudget ? 'p-0' : 'text-gray-500 hover:text-white hover:bg-white/5'}`}
                                                                        >
                                                                            {isEditingBudget ? (
                                                                                <input 
                                                                                    autoFocus
                                                                                    type="text" 
                                                                                    value={editValue}
                                                                                    onChange={(e) => setEditValue(e.target.value)}
                                                                                    onBlur={handleSaveEdit}
                                                                                    onKeyDown={(e) => e.key === 'Enter' && handleSaveEdit()}
                                                                                    className="w-full h-full bg-[#0f1115] text-white text-right px-1 focus:outline-none focus:ring-1 focus:ring-[#1e5144] absolute inset-0 text-[11px]"
                                                                                />
                                                                            ) : formatCurrency(budget)}
                                                                        </td>
                                                                        
                                                                        {/* Realized Cell */}
                                                                        <td 
                                                                            onDoubleClick={() => handleCellClick(item, i, 'realized')}
                                                                            className={`px-2 py-2 text-right relative cursor-pointer text-[11px] ${isEditingRealized ? 'p-0' : 'text-blue-500/70 hover:text-blue-300 hover:bg-white/5'}`}
                                                                        >
                                                                            {isEditingRealized ? (
                                                                                <input 
                                                                                    autoFocus
                                                                                    type="text" 
                                                                                    value={editValue}
                                                                                    onChange={(e) => setEditValue(e.target.value)}
                                                                                    onBlur={handleSaveEdit}
                                                                                    onKeyDown={(e) => e.key === 'Enter' && handleSaveEdit()}
                                                                                    className="w-full h-full bg-[#0f1115] text-blue-300 text-right px-1 focus:outline-none focus:ring-1 focus:ring-blue-500 absolute inset-0 text-[11px]"
                                                                                />
                                                                            ) : formatCurrency(realized)}
                                                                        </td>

                                                                        {/* Diff Cell */}
                                                                        <td className={`px-2 py-2 text-right font-medium text-[11px] border-r border-gray-800 ${diff >= 0 ? 'text-emerald-600' : 'text-rose-600'}`}>
                                                                            {formatCurrency(diff)}
                                                                        </td>
                                                                    </React.Fragment>
                                                                );
                                                            })}

                                                            {/* Total Item Columns */}
                                                            <td className="px-2 py-2 text-right text-gray-400 bg-[#1a1d24] group-hover:bg-[#20232b] sticky z-10 border-l border-gray-800" style={{right: 160}}>{formatCurrency(itemBudgetTotal)}</td>
                                                            <td className="px-2 py-2 text-right text-blue-400 bg-[#1a1d24] group-hover:bg-[#20232b] sticky z-10" style={{right: 80}}>{formatCurrency(itemRealizedTotal)}</td>
                                                            <td className={`px-2 py-2 text-right font-bold bg-[#1a1d24] group-hover:bg-[#20232b] sticky right-0 z-10 ${itemDiffTotal >= 0 ? 'text-emerald-500' : 'text-rose-500'}`}>{formatCurrency(itemDiffTotal)}</td>
                                                        </tr>
                                                    );
                                                })}
                                            </React.Fragment>
                                        );
                                    })}
                                </React.Fragment>
                            );
                        })}
                    </tbody>
                    <tfoot className="sticky bottom-0 z-40 bg-[#0f1115] shadow-2xl border-t-2 border-[#1e5144]">
                        <tr>
                            <td className="px-4 py-3 sticky left-0 z-50 bg-[#0f1115] border-r border-gray-800 font-bold text-gray-200 uppercase tracking-wider text-sm">
                                Totais Consolidados
                            </td>
                            {monthlyGrandTotals.budgetTotals.map((budgetVal, i) => {
                                const realizedVal = monthlyGrandTotals.realizedTotals[i];
                                const diff = budgetVal - realizedVal;
                                const isPositive = diff >= 0;

                                return (
                                    <React.Fragment key={i}>
                                        <td className="px-2 py-3 border-b border-gray-800 text-right bg-[#0f1115] text-gray-300 font-mono font-bold text-xs">
                                            {formatCurrency(budgetVal)}
                                        </td>
                                        <td className="px-2 py-3 border-b border-gray-800 text-right bg-[#0f1115] text-blue-400 font-mono font-bold text-xs">
                                            {formatCurrency(realizedVal)}
                                        </td>
                                        <td className={`px-2 py-3 border-b border-gray-800 text-right bg-[#0f1115] font-mono font-extrabold text-xs border-r border-gray-800 ${isPositive ? 'text-emerald-500' : 'text-rose-500'}`}>
                                            {formatCurrency(diff)}
                                        </td>
                                    </React.Fragment>
                                );
                            })}
                            
                            {/* Grand Totals Footer */}
                            {(() => {
                                const totalBudget = monthlyGrandTotals.budgetTotals.reduce((a, b) => a + b, 0);
                                const totalRealized = monthlyGrandTotals.realizedTotals.reduce((a, b) => a + b, 0);
                                const totalDiff = totalBudget - totalRealized;
                                const isPos = totalDiff >= 0;

                                return (
                                    <>
                                        <td className="px-2 py-3 border-l-2 border-[#1e5144] text-right bg-[#0f1115] text-white font-mono font-bold text-xs sticky z-50" style={{right: 160}}>
                                            {formatCurrency(totalBudget)}
                                        </td>
                                        <td className="px-2 py-3 text-right bg-[#0f1115] text-blue-300 font-mono font-bold text-xs sticky z-50" style={{right: 80}}>
                                            {formatCurrency(totalRealized)}
                                        </td>
                                        <td className={`px-2 py-3 text-right bg-[#0f1115] font-mono font-extrabold text-xs sticky right-0 z-50 ${isPos ? 'text-emerald-400' : 'text-rose-400'}`}>
                                            {formatCurrency(totalDiff)}
                                        </td>
                                    </>
                                );
                            })()}
                        </tr>
                    </tfoot>
                </table>
            </div>
            
            <div className="bg-[#15171c] p-2 border-t border-gray-800 flex justify-end gap-6 text-[10px] text-gray-500 px-6">
                <div className="flex items-center gap-1.5"><div className="w-2 h-2 bg-[#152e29] border border-emerald-500 rounded"></div> Grupo</div>
                <div className="flex items-center gap-1.5"><div className="w-2 h-2 bg-[#20232b] border border-gray-600 rounded"></div> Subgrupo</div>
                <div className="flex items-center gap-1.5"><Edit2 size={10} /> Clique duplo para editar valores (Orçado ou Realizado)</div>
            </div>
        </div>
    </div>
  );
};

export default BudgetView;
