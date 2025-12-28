import type { RepTableData, RepresentativeProfile } from '@/shared/types';

export type { RepTableData, RepresentativeProfile };

export interface RepTableProps {
  data: RepTableData;
  onRepClick?: (repName: string) => void;
}

export interface RepProfileModalProps {
  isOpen: boolean;
  onClose: () => void;
  profile: RepresentativeProfile;
}

export const COMPANY_MAP: Record<string, string> = {
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
