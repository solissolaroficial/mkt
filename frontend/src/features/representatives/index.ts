export { default as RepTable } from './ui/RepTable';
export { default as RepProfileModal } from './ui/RepProfileModal';

export { useRepTrainingData, useRepMarketingData, useRepProfile, useAllRepProfiles } from './hooks/useRepresentatives';

export type { 
  RepTableData, 
  RepresentativeProfile,
  RepTableProps,
  RepProfileModalProps,
  COMPANY_MAP
} from './types';
