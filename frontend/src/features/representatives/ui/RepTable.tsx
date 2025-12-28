import React from 'react';
import type { RepTableData } from '@/shared/types';
import { COMPANY_MAP } from '../types';
import type { RepTableProps } from '../types';

const RepTable: React.FC<RepTableProps> = ({ data, onRepClick }) => {
  // Transpose logic:
  // Rows become Reps (from data.columns)
  // Columns become Months (from data.rows)
  
  const months = data.rows; 
  const reps = data.columns;

  return (
    <div className="bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 overflow-hidden flex flex-col">
      <div className="p-6 border-b border-gray-800 flex justify-between items-center">
        <h3 className="text-lg font-semibold text-gray-100">{data.title}</h3>
        <span className="text-xs font-medium bg-[#1e5144]/10 text-emerald-400 px-3 py-1 rounded-full border border-[#1e5144]/20">
          Visão por Representante
        </span>
      </div>
      
      <div className="overflow-x-auto custom-scrollbar flex-grow">
        <table className="w-full text-sm text-left">
          <thead className="text-xs text-gray-400 uppercase bg-[#20232b] sticky top-0 z-20">
            <tr>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 min-w-[200px] z-20 sticky left-0 bg-[#20232b] border-r border-gray-800">
                Representante
              </th>
              {months.map((row, idx) => (
                <th key={idx} scope="col" className="px-4 py-4 font-medium border-b border-gray-800 whitespace-nowrap text-center min-w-[80px]">
                  <div className="flex flex-col items-center">
                     <span>{row.month.substring(0, 3)}</span>
                     <span className="text-[10px] text-gray-600 font-normal mt-0.5">Meta: {row.target}</span>
                  </div>
                </th>
              ))}
              <th scope="col" className="px-4 py-4 font-bold text-gray-300 border-b border-gray-800 text-center min-w-[80px] border-l border-gray-800 bg-[#20232b]">
                TOTAL
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-800">
            {reps.map((rep, repIndex) => {
              const company = COMPANY_MAP[rep];
              // Calculate Total for this rep across all visible months
              const total = months.reduce((acc, row) => acc + (row.values[rep] || 0), 0);

              return (
                <tr key={repIndex} className="hover:bg-[#20232b] transition-colors group">
                  <td className="px-6 py-4 font-medium text-gray-200 bg-[#1a1d24] sticky left-0 border-r border-gray-800 shadow-sm z-10 group-hover:bg-[#20232b]">
                     <div 
                        className={`font-medium text-gray-200 ${onRepClick ? 'cursor-pointer hover:text-emerald-400 hover:underline' : ''}`}
                        onClick={() => onRepClick && onRepClick(rep)}
                     >
                        {rep}
                     </div>
                     {company && (
                       <div className="text-[10px] text-gray-500 uppercase tracking-tight truncate max-w-[180px] mt-0.5" title={company}>
                          {company}
                       </div>
                     )}
                  </td>
                  {months.map((row, monthIndex) => {
                    const value = row.values[rep] || 0;
                    const isMet = value >= row.target;
                    const isZero = value === 0;
                    
                    return (
                      <td key={monthIndex} className="px-4 py-4 text-center">
                        <div className={`
                          inline-flex items-center justify-center w-8 h-8 rounded-full font-bold text-xs border
                          ${isMet 
                            ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' 
                            : isZero 
                              ? 'bg-gray-800 text-gray-500 border-gray-700' 
                              : 'bg-rose-500/10 text-rose-400 border-rose-500/20'}
                        `}>
                          {value}
                        </div>
                      </td>
                    );
                  })}
                  <td className="px-4 py-4 text-center border-l border-gray-800 bg-[#1a1d24]/50">
                    <div className="inline-flex items-center justify-center w-10 h-8 rounded font-bold text-xs bg-gray-800 text-gray-300 border border-gray-700">
                       {total}
                    </div>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
      <div className="bg-[#15171c] p-3 border-t border-gray-800 flex items-center gap-4 text-xs text-gray-500 justify-end">
        <div className="flex items-center gap-1">
          <div className="w-3 h-3 rounded-full bg-emerald-500/20 border border-emerald-500/50"></div>
          <span>Meta Atingida</span>
        </div>
        <div className="flex items-center gap-1">
          <div className="w-3 h-3 rounded-full bg-rose-500/20 border border-rose-500/50"></div>
          <span>Abaixo da Meta</span>
        </div>
        <div className="flex items-center gap-1">
          <div className="w-3 h-3 rounded-full bg-gray-700 border border-gray-600"></div>
          <span>Sem Atividade</span>
        </div>
      </div>
    </div>
  );
};

export default RepTable;
