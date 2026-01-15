import React from 'react';
import type { RepMarketingAction } from '../../cooperative/types';

interface RepTableProps {
  data: {
    headers: string[];
    rows: string[];
    values: Record<string, Record<string, number>>;
  };
  onRepClick: (repName: string) => void;
}

const RepTable: React.FC<RepTableProps> = ({ data, onRepClick }) => {
  const MONTHS = ['JANEIRO', 'FEVEREIRO', 'MARÇO', 'ABRIL', 'MAIO', 'JUNHO', 'JULHO', 'AGOSTO', 'SETEMBRO', 'OUTUBRO', 'NOVEMBRO', 'DEZEMBRO'];

  return (
    <div className="overflow-x-auto custom-scrollbar">
      <table className="w-full text-sm text-left border-collapse">
        <thead className="bg-[#15171c] text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0 z-20">
          <tr>
            <th className="px-4 py-3 min-w-[150px] sticky left-0 z-30 bg-[#15171c] border-r">
              Representante
            </th>
            {data.headers.map((header, index) => (
              <th key={index} colSpan={data.rows.length} className="px-2 py-3 text-center min-w-[80px]">
                {header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="bg-[#1a1d24]">
          {data.rows.map((repName, rowIndex) => (
            <tr key={rowIndex} className="hover:bg-white/5 transition-colors border-b border-gray-800/20">
              <td className="px-4 py-3 sticky left-0 z-10 bg-[#1a1d24] border-r border-gray-800/30">
                <button
                  onClick={() => onRepClick(repName)}
                  className="font-medium text-emerald-400 hover:text-emerald-300 hover:underline transition-colors text-left"
                >
                  {repName}
                </button>
              </td>
              {data.rows.map((month, colIndex) => {
                const value = data.values[month]?.[repName] || 0;
                return (
                  <td key={`${rowIndex}-${colIndex}`} className="px-2 py-3 text-center">
                    {value}
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RepTable;
