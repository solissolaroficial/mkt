import { useMemo } from 'react';
import type { KpiCategory, MonthlyData, BreakdownItem } from '../types';

/**
 * Hook para cálculos derivados de KPIs
 * Mantém a lógica de cálculo do CPL e outros KPIs derivados
 */
export const useKpiCalculations = (kpis: KpiCategory[]) => {
  return useMemo(() => {
    return kpis.map((kpi) => {
      // Cálculo automático de CPL baseado em Ad Spend / Opportunities
      if (kpi.id === 'cpl') {
        const adSpendKpi = kpis.find((k) => k.id === 'ad_spend');
        const oppKpi = kpis.find((k) => k.id === 'opportunities');

        const newData = kpi.data.map((d) => {
          const adData = adSpendKpi?.data.find((ad) => ad.month === d.month);
          const oppData = oppKpi?.data.find((opp) => opp.month === d.month);

          const spend = adData?.realized || 0;
          const leads = oppData?.realized || 0;

          // Evitar divisão por zero
          const calculatedCpl = leads > 0 ? spend / leads : 0;

          // Calcular breakdown do CPL por canal
          let cplBreakdown: BreakdownItem[] | undefined = undefined;
          if (adData?.breakdown && oppData?.breakdown) {
            cplBreakdown = adData.breakdown.map((adChannel, idx) => {
              const oppChannel = oppData.breakdown?.[idx];
              const channelLeads = oppChannel?.value || 0;
              const channelCpl = channelLeads > 0 ? adChannel.value / channelLeads : 0;
              return {
                label: adChannel.label,
                value: channelCpl,
              };
            });
          }

          return {
            ...d,
            realized: calculatedCpl,
            breakdown: cplBreakdown,
          };
        });

        return { ...kpi, data: newData };
      }

      // Retorna KPI sem modificações
      return kpi;
    });
  }, [kpis]);
};