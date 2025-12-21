
import React from 'react';
import { RepTableData } from '../types';

interface HeatmapChartProps {
  data: RepTableData;
  title: string;
  onRepClick?: (repName: string) => void;
}

const COMPANY_MAP: Record<string, string> = {
  'André': 'MELO & DOMINGUES REP',
  'César': 'TORQUATO REPRESENTACOES',
  'Cleber': 'SANTOS & BRAMBILA REP.',
  'Cristiano e Ranoika': 'RN REPRESENTAÇÕES COMERCIAIS',
  'Fausto': 'HIKARI REPRESENTAÇÃO',
  'Gonçalves': 'SOLAR PRÁTICO',
  'Márcio Henrique': '4F REPRESENTAÇÕES',
  'Marcos': 'MARCOS JUNQUEIRA VILELA',
  'Nilton': 'QUALITYENG SERVICE REPRESENTACOES',
  'Jorge': 'JK GUIMARAES REPRES',
  'Otávio': 'MONICA C. MENDES ME',
  'Rafael Betoni': 'RAFAEL NONATO BETONI TOMAZ',
  'Rafael Lazzarotto': 'LAZZAROTTO VENDAS E REP. LTDA',
  'Wilson': 'SOLAR FLUX REPRESENTACOES'
};

const HeatmapChart: React.FC<HeatmapChartProps> = ({ data, title, onRepClick }) => {
  // Extract unique columns (Reps) and Rows (Months)
  // RepTableData has rows as months, and values keys as Reps.
  // We want Matrix: X = Months, Y = Reps
  
  const reps = data.columns;
  // Reverse rows to show Jan -> Dec left to right, or maintain recent first. 
  // Standard heatmap usually time is X axis (Left->Right: Jan->Dec).
  // The data provided in constants is Nov -> Jan (Reverse chronological).
  // Let's reverse it to have Jan on left.
  const timeSortedRows = [...data.rows].reverse(); 

  // Function to determine color intensity
  const getColor = (value: number, target: number) => {
    if (value === 0) return 'bg-gray-800/50 text-gray-600'; // No activity
    
    const percentage = value / target;
    
    if (percentage < 0.5) return 'bg-rose-900/50 text-rose-300';
    if (percentage < 0.8) return 'bg-amber-900/50 text-amber-300';
    if (percentage < 1.0) return 'bg-blue-900/40 text-blue-300';
    if (percentage < 1.5) return 'bg-blue-600 text-white';
    return 'bg-blue-500 text-white shadow-lg shadow-blue-500/30'; // High performance
  };

  return (
    <div className="bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 overflow-hidden flex flex-col">
      <div className="p-6 border-b border-gray-800">
        <h3 className="text-lg font-semibold text-gray-100">{title}</h3>
        <p className="text-sm text-gray-500">Mapa de Calor de Intensidade</p>
      </div>

      <div className="p-6 overflow-x-auto custom-scrollbar">
        <div className="min-w-[800px]">
          {/* Header Row (Months) */}
          <div className="flex mb-2">
            <div className="w-56 flex-shrink-0"></div> {/* Spacer for Rep Names */}
            {timeSortedRows.map((row, idx) => (
              <div key={idx} className="flex-1 text-center text-xs font-semibold text-gray-500 uppercase">
                {row.month.substring(0, 3)}
              </div>
            ))}
            {/* Total Header */}
            <div className="flex-1 text-center text-xs font-bold text-gray-400 uppercase border-l border-gray-800 ml-2 pl-2">
              TOTAL
            </div>
          </div>

          {/* Data Rows (Reps) */}
          <div className="space-y-2">
            {reps.map((rep, repIdx) => {
              // Calculate Total for this rep across all timeSortedRows
              const totalValue = timeSortedRows.reduce((acc, row) => acc + (row.values[rep] || 0), 0);
              const company = COMPANY_MAP[rep];

              return (
                <div key={repIdx} className="flex items-center hover:bg-[#20232b] rounded-lg transition-colors p-1 group">
                  {/* Rep Name Label */}
                  <div className="w-56 flex-shrink-0 flex flex-col justify-center pr-4">
                    <span 
                        className={`text-sm font-bold text-gray-200 group-hover:text-white transition-colors ${onRepClick ? 'cursor-pointer hover:underline hover:text-emerald-400' : ''}`}
                        onClick={() => onRepClick && onRepClick(rep)}
                    >
                        {rep}
                    </span>
                    {company && (
                        <span className="text-[10px] text-gray-500 uppercase tracking-tight truncate" title={company}>
                            {company}
                        </span>
                    )}
                  </div>
                  
                  {/* Month Cells */}
                  {timeSortedRows.map((row, monthIdx) => {
                    const value = row.values[rep] || 0;
                    const colorClass = getColor(value, row.target);
                    
                    return (
                      <div key={monthIdx} className="flex-1 px-1">
                        <div 
                          className={`h-10 w-full rounded flex items-center justify-center text-xs font-bold transition-all hover:scale-105 cursor-default ${colorClass}`}
                          title={`${rep} em ${row.month}: ${value} (Meta: ${row.target})`}
                        >
                          {value > 0 ? value : ''}
                        </div>
                      </div>
                    );
                  })}

                  {/* Total Cell */}
                  <div className="flex-1 px-1 border-l border-gray-800 ml-2 pl-2">
                    <div 
                      className="h-10 w-full rounded flex items-center justify-center text-xs font-bold bg-gray-800 text-gray-300 border border-gray-700"
                      title={`Total de ${rep}: ${totalValue}`}
                    >
                      {totalValue}
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </div>

      {/* Legend */}
      <div className="bg-[#15171c] p-3 border-t border-gray-800 flex flex-wrap items-center gap-4 text-xs text-gray-500 justify-center">
        <div className="flex items-center gap-1"><div className="w-3 h-3 bg-gray-800 border border-gray-700 rounded"></div> 0</div>
        <div className="flex items-center gap-1"><div className="w-3 h-3 bg-rose-900/50 rounded"></div> Baixo</div>
        <div className="flex items-center gap-1"><div className="w-3 h-3 bg-amber-900/50 rounded"></div> Médio</div>
        <div className="flex items-center gap-1"><div className="w-3 h-3 bg-blue-900/40 rounded"></div> Meta</div>
        <div className="flex items-center gap-1"><div className="w-3 h-3 bg-blue-500 rounded"></div> Superação</div>
      </div>
    </div>
  );
};

export default HeatmapChart;