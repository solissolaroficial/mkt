import React from 'react';
import { MarketingChannelData, ChannelPerformance } from '../types';
import { MARKETING_CHANNELS_DATA, ANNUAL_CHANNEL_DATA } from '../constants';

interface ChannelTableProps {
  selectedMonth: string;
}

const ChannelTable: React.FC<ChannelTableProps> = ({ selectedMonth }) => {
  // Select data based on filter
  const data: ChannelPerformance[] = selectedMonth === 'ALL' 
    ? ANNUAL_CHANNEL_DATA 
    : MARKETING_CHANNELS_DATA[selectedMonth] || [];

  return (
    <div className="bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 overflow-hidden flex flex-col h-full">
      <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
        <div>
          <h3 className="text-lg font-semibold text-gray-100">Performance por Canal de Aquisição</h3>
          <p className="text-sm text-gray-500 mt-1">
            {selectedMonth === 'ALL' ? 'Acumulado Anual' : `Dados de ${selectedMonth}`}
          </p>
        </div>
      </div>
      
      <div className="overflow-x-auto custom-scrollbar flex-grow">
        <table className="w-full text-sm text-left">
          <thead className="text-xs text-gray-400 uppercase bg-[#15171c] sticky top-0 z-10">
            <tr>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800">Canal</th>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 text-right">Visitas</th>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 text-right">Leads</th>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 text-right">Conv. (%)</th>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 text-right">Investimento</th>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 text-right">CPL</th>
              <th scope="col" className="px-6 py-4 font-semibold border-b border-gray-800 text-right">ROAS</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-800">
            {data.map((row, index) => (
              <tr key={index} className="hover:bg-[#20232b] transition-colors">
                <td className="px-6 py-4 font-medium text-gray-200">{row.channel}</td>
                <td className="px-6 py-4 text-right text-gray-400">{row.visits.toLocaleString()}</td>
                <td className="px-6 py-4 text-right font-medium text-blue-400">{row.leads.toLocaleString()}</td>
                <td className="px-6 py-4 text-right">
                  <span className={`px-2 py-1 rounded-full text-xs font-medium border ${row.conversion > 5 ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' : 'bg-gray-800 text-gray-400 border-gray-700'}`}>
                    {row.conversion.toFixed(1)}%
                  </span>
                </td>
                <td className="px-6 py-4 text-right text-gray-400">
                    {row.investment > 0 ? `R$ ${row.investment.toLocaleString()}` : '-'}
                </td>
                <td className="px-6 py-4 text-right font-medium text-gray-300">
                   {row.cpl > 0 ? `R$ ${row.cpl.toLocaleString('pt-BR', { minimumFractionDigits: 2 })}` : '-'}
                </td>
                <td className="px-6 py-4 text-right">
                   {row.roas > 0 ? (
                     <span className={`font-bold ${row.roas >= 4 ? 'text-emerald-400' : 'text-amber-400'}`}>
                        {row.roas.toFixed(1)}x
                     </span>
                   ) : '-'}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ChannelTable;