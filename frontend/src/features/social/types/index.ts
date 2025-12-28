import type { SocialBenchmarking } from '@/shared/types';

export type { SocialBenchmarking };

export interface SocialBenchmarkingViewProps {
  data: SocialBenchmarking[];
  onBack?: () => void;
}

export interface SocialRankingModalProps {
  isOpen: boolean;
  onClose: () => void;
  data: SocialBenchmarking[];
}
