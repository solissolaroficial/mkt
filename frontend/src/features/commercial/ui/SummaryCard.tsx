import React from 'react';
import { ArrowUpRight, ArrowDownRight, MousePointerClick } from 'lucide-react';

interface SummaryCardProps {
  title: string;
  value: number;
  target: number;
  icon: React.ElementType;
  colorClass: string;
  onClick?: () => void;
  unit?: string;
  hideMeta?: boolean;
  className?: string;
}

const SummaryCard: React.FC<SummaryCardProps> = ({ title, value, target, icon: Icon, colorClass, onClick, unit, hideMeta = false, className = '' }) => {
  // Determine format based on unit prop OR title keywords (heuristic fallback)
  const isCurrency = unit === 'currency' || (!unit && (title.includes('R$') || title.includes('Custo') || title.includes('Investimento')));
  const isPercent = unit === 'percent' || (!unit && (title.includes('%') || title.includes('Taxa') || title.includes('ROAS')));

  const percent = target > 0 ? (value / target) * 100 : 0;
  // Logic inverse for Costs (lower is better)
  const isCostMetric = title.includes('Custo');
  const isPositive = isCostMetric ? percent <= 100 : percent >= 100;

  const formatValue = (val: number) => {
    if (isCurrency) return val.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL', maximumFractionDigits: 0 });
    if (isPercent && title.includes('ROAS')) return `${val.toFixed(1)}x`;
    if (isPercent) return `${val.toFixed(1)}%`;
    // If unit is explicitly 'number' or unknown, format as simple number
    return val.toLocaleString();
  };

  return (
    <div
      onClick={onClick}
      className={`bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 transition-all duration-300 group relative ${onClick ? 'cursor-pointer hover:border-[#1e5144]/30 hover:bg-[#20232b]' : ''} ${hideMeta ? 'p-4' : 'p-5'} ${className}`}
    >
      {onClick && (
        <div className="absolute top-3 right-3 opacity-0 group-hover:opacity-100 transition-opacity text-gray-500">
          <MousePointerClick size={16} />
        </div>
      )}

      {hideMeta ? (
        <div className="flex flex-col items-start h-full">
            <div className={`p-2 rounded-xl mb-3 ${colorClass}`}>
                <Icon className={`w-5 h-5 ${colorClass.replace('bg-', 'text-').replace('/10', '')}`} />
            </div>
            <h3 className="text-gray-400 text-xs font-medium uppercase tracking-wide mb-2 leading-tight">{title}</h3>
            <span className="text-2xl font-bold text-white tracking-tight mt-auto">{formatValue(value)}</span>
        </div>
      ) : (
        <>
            <div className="flex justify-between items-start mb-4">
                <div className={`p-2 rounded-xl ${colorClass}`}>
                <Icon className={`w-5 h-5 ${colorClass.replace('bg-', 'text-').replace('/10', '')}`} />
                </div>
                <div className={`flex items-center gap-1 text-xs font-medium px-2 py-1 rounded-full border ${
                    isPositive
                        ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20'
                        : 'bg-rose-500/10 text-rose-400 border-rose-500/20'
                }`}>
                    {isPositive ? <ArrowUpRight className="w-3 h-3" /> : <ArrowDownRight className="w-3 h-3" />}
                    {isCostMetric && percent > 100 ? '+' : ''}{Math.round(percent - 100)}% vs Meta
                </div>
            </div>

            <h3 className="text-gray-400 text-sm font-medium mb-1">{title}</h3>
            <div className="flex items-baseline gap-2 flex-wrap">
                <span className="text-2xl font-bold text-white tracking-tight">{formatValue(value)}</span>
                <span className="text-xs text-gray-500 whitespace-nowrap">/ meta {formatValue(target)}</span>
            </div>

            <div className="w-full bg-gray-800 rounded-full h-1.5 mt-4 overflow-hidden">
                <div
                    className={`h-1.5 rounded-full shadow-[0_0_10px_rgba(0,0,0,0.5)] ${isPositive ? 'bg-emerald-500' : 'bg-rose-500'}`}
                    style={{ width: `${Math.min(isCostMetric ? (10000/percent) : percent, 100)}%` }}
                ></div>
            </div>
        </>
      )}
    </div>
  );
};

export default SummaryCard;
