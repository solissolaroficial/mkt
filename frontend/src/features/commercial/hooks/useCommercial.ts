import { useMemo } from 'react';
import { useKpis } from '@/features/kpis/hooks/useKpis';

export const useCommercial = () => {
  const { data } = useKpis();
  const allKpis = data?.data || [];

  // Filtrar apenas os KPIs comerciais
  const commercialKpis = useMemo(() => {
    return allKpis.filter(kpi =>
      kpi.title === 'Taxa de Oportunidades' ||
      kpi.title === 'Taxa de Conversão Global (%)'
    );
  }, [allKpis]);

  return {
    commercialKpis,
    isLoading: false,
    error: null
  };
};
