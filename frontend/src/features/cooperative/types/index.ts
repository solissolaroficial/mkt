import type { OfflineAction, ShowroomItem } from '@/shared/types';

export type { OfflineAction, ShowroomItem };

export interface CooperativeActionsViewProps {
  selectedMonth?: string;
  onRepClick?: (rep: string) => void;
}

export interface RepMarketingAction {
  id: string;
  repName: string;
  date: string; // YYYY-MM-DD
  description: string;
  month: string; // Full month name e.g. "NOVEMBRO"
}

export type SubTabType = 'menu' | 'marketing' | 'online_marketing' | 'showroom' | 'offline_marketing';

export type OfflineCategory = 
  | 'TODAS' 
  | 'PARCERIA' 
  | 'AÇÃO COOPERADA' 
  | 'ENTREGA DE BRINDES EXCLUSIVOS' 
  | 'MINIATURAS' 
  | 'BRINDES - FEIRA' 
  | 'FEIRA - EXPOSIÇÃO';

export interface OfflineActionForm {
  val: string;
  date: string;
  pdv: string;
  rep: string;
  city: string;
  uf: string;
  desc: string;
  category: OfflineCategory;
  pedido: string;
  saida: string;
  previsao: string;
  pontuado: string;
}

export interface ShowroomForm {
  pdv: string;
  city: string;
  contact: string;
  repName: string;
  deliveryForecast: string;
  workshopDate: string;
}

export interface MarketingActionForm {
  repName: string;
  date: string;
  description: string;
}
