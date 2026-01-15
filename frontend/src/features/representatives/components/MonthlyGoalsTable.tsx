import React from 'react';
import { useMonthlyGoalsTableData } from '../hooks/useMonthlyGoals';
import type { GetRepresentativeGoalsTableDataResponse } from '../types/monthlyGoal';

interface MonthlyGoalsTableProps {
  year?: number;
  month?: number;
  onRepClick?: (repName: string) => void;
}

/**
 * MonthlyGoalsTable component displays representatives' monthly goals in a transposed table format
 * (representatives as rows, months as columns)
 */
export function MonthlyGoalsTable({ year, month, onRepClick }: MonthlyGoalsTableProps) {
  const { data, isLoading, error } = useMonthlyGoalsTableData({ year, month });

  if (isLoading) {
    return <div className="flex justify-center p-8">Carregando...</div>;
  }

  if (error) {
    return <div className="text-red-500 p-4">Erro ao carregar dados: {error.message}</div>;
  }

  if (!data) {
    return <div className="p-4">Nenhum dado disponível</div>;
  }

  return (
    <div className="bg-[#1a1d24] rounded-2xl shadow-lg border border-white/5 overflow-hidden flex flex-col">
      <div className="overflow-x-auto flex-grow">
        <table className="w-full text-sm text-left">
          <thead className="text-xs text-gray-400 uppercase bg-[#20232b] sticky top-0 z-20">
            <tr>
              <th className="px-6 py-4 font-semibold border-b border-gray-800 min-w-[200px] z-20 sticky left-0 bg-[#20232b] border-r border-gray-800">
                Representante
              </th>
              <th className="px-4 py-4 font-medium border-b border-gray-800 whitespace-nowrap">
                Empresa
              </th>
              <th className="px-4 py-4 font-medium border-b border-gray-800 whitespace-nowrap">
                Região
              </th>
              <th className="px-4 py-4 font-medium border-b border-gray-800 whitespace-nowrap">
                Cidade
              </th>
              {data.months.map((monthData) => (
                <th
                  key={monthData.month}
                  className="px-4 py-4 font-medium border-b border-gray-800 whitespace-nowrap text-center min-w-[80px]"
                >
                  {monthData.monthName}
                </th>
              ))}
              <th className="px-4 py-4 font-bold text-gray-300 border-b border-gray-800 text-center min-w-[80px] border-l border-gray-800 bg-[#20232b]">
                TOTAL
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-800">
            {data.rows.map((row) => (
              <tr key={row.representativeId} className="hover:bg-[#20232b]">
                <td className="px-6 py-4 font-medium text-gray-200 bg-[#1a1d24] border-r border-gray-800 shadow-sm">
                   <div
                      className={`font-medium text-gray-200 ${onRepClick ? 'cursor-pointer hover:text-emerald-400 hover:underline' : ''}`}
                      onClick={() => onRepClick && onRepClick(row.representativeName)}
                   >
                      {row.representativeName}
                   </div>
                </td>
                <td className="px-4 py-4 whitespace-nowrap text-sm text-gray-400 border-b">
                  {row.company}
                </td>
                <td className="px-4 py-4 whitespace-nowrap text-sm text-gray-400 border-b">
                  {row.region}
                </td>
                <td className="px-4 py-4 whitespace-nowrap text-sm text-gray-400 border-b">
                  {row.city}
                </td>
                {row.monthValues.map((monthValue) => (
                  <td
                    key={monthValue.month}
                    className="px-4 py-4 text-center text-sm border-b"
                  >
                    <div className="flex flex-col items-center space-y-1">
                      <div className="flex items-center space-x-2">
                        <span className="text-gray-500">{monthValue.target.toLocaleString()}</span>
                        <span className="text-gray-400">/</span>
                        <span
                          className={`font-medium ${
                            monthValue.isMet ? 'text-emerald-400' : 'text-rose-400'
                          }`}
                        >
                          {monthValue.realized.toLocaleString()}
                        </span>
                      </div>
                      <span className="text-xs text-gray-500">
                        {monthValue.percentage.toFixed(1)}%
                      </span>
                    </div>
                  </td>
                ))}
                <td className="px-4 py-4 text-center text-sm border-l border-gray-800 bg-[#1a1d24]/50">
                  <div className="flex flex-col items-center space-y-1">
                    <div className="flex items-center space-x-2">
                      <span className="text-gray-500">{row.totalTarget.toLocaleString()}</span>
                      <span className="text-gray-400">/</span>
                      <span className="font-medium text-gray-200">
                        {row.totalRealized.toLocaleString()}
                      </span>
                    </div>
                    <span className="text-xs text-gray-500">
                      {row.totalPercentage.toFixed(1)}%
                    </span>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
          <tfoot>
            <tr className="bg-[#20232b] font-medium">
              <td colSpan={4} className="px-4 py-4 text-sm text-gray-200 border-t">
                Resumo
              </td>
              {data.months.map((monthData) => (
                <td
                  key={monthData.month}
                  className="px-4 py-4 text-center text-sm text-gray-200 border-t"
                >
                  <div className="flex flex-col items-center space-y-1">
                    <div className="flex items-center space-x-2">
                      <span className="text-gray-500">{monthData.targetSum.toLocaleString()}</span>
                      <span className="text-gray-400">/</span>
                      <span className="font-medium">{monthData.realizedSum.toLocaleString()}</span>
                    </div>
                  </div>
                </td>
              ))}
              <td className="px-4 py-4 text-center text-sm text-gray-200 border-t">
                <div className="flex flex-col items-center space-y-1">
                  <div className="flex items-center space-x-2">
                    <span className="text-gray-500">{data.summary.totalTarget.toLocaleString()}</span>
                    <span className="text-gray-400">/</span>
                    <span className="font-medium">{data.summary.totalRealized.toLocaleString()}</span>
                  </div>
                  <span className="text-xs text-gray-500">
                    {data.summary.overallPercentage.toFixed(1)}%
                  </span>
                </div>
              </td>
            </tr>
          </tfoot>
        </table>
      </div>

      {/* Summary Statistics */}
      <div className="mt-4 grid grid-cols-2 md:grid-cols-4 gap-4">
        <div className="bg-[#20232b] p-4 rounded-lg shadow border border-gray-800">
          <div className="text-sm text-gray-500">Total de Representantes</div>
          <div className="text-2xl font-bold text-gray-200">
            {data.summary.totalRepresentatives}
          </div>
        </div>
        <div className="bg-[#20232b] p-4 rounded-lg shadow border border-gray-800">
          <div className="text-sm text-gray-500">Meta Total</div>
          <div className="text-2xl font-bold text-gray-200">
            {data.summary.totalTarget.toLocaleString()}
          </div>
        </div>
        <div className="bg-[#20232b] p-4 rounded-lg shadow border border-gray-800">
          <div className="text-sm text-gray-500">Realizado Total</div>
          <div className="text-2xl font-bold text-gray-200">
            {data.summary.totalRealized.toLocaleString()}
          </div>
        </div>
        <div className="bg-[#20232b] p-4 rounded-lg shadow border border-gray-800">
          <div className="text-sm text-gray-500">Percentual Geral</div>
          <div className="text-2xl font-bold text-gray-200">
            {data.summary.overallPercentage.toFixed(1)}%
          </div>
        </div>
      </div>

      <div className="mt-4 flex space-x-4">
        <div className="flex items-center space-x-2">
          <div className="w-3 h-3 bg-emerald-500/20 border border-emerald-500/50 rounded"></div>
          <span className="text-sm text-gray-400">Meta Atingida: {data.summary.goalsMetCount}</span>
        </div>
        <div className="flex items-center space-x-2">
          <div className="w-3 h-3 bg-rose-500/20 border border-rose-500/50 rounded"></div>
          <span className="text-sm text-gray-400">Meta Não Atingida: {data.summary.goalsNotMetCount}</span>
        </div>
      </div>
    </div>
  );
}
