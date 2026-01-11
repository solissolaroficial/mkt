export { default as SocialBenchmarkingView } from './ui/SocialBenchmarkingView';
export { default as SocialRankingModal } from './ui/SocialRankingModal';
export { default as SocialPostsView } from './ui/SocialPostsView';
export { default as SocialDailyAggregationsView } from './ui/SocialDailyAggregationsView';
export { default as SocialPage } from './pages/SocialPage';

export * from './services/socialService';
export {
  useSocialBenchmarkings,
  useSocialBenchmarking,
  useSocialPosts,
  useSocialPost,
  useSocialDailyAggregations,
  useSocialDailyAggregation,
} from './hooks/useSocial';
export { useSocialMutations } from './hooks/useSocialMutations';

export type {
  SocialBenchmarking,
  SocialBenchmarkingViewProps,
  SocialRankingModalProps
} from './types';

// Re-export types from legacy.types.ts for convenience
export type {
  SocialPost,
  SocialPostListResponse,
  SocialPostFilters,
  CreateSocialPostRequest,
  UpdateSocialPostRequest,
  SocialDailyAggregation,
  SocialDailyAggregationListResponse,
  SocialDailyAggregationFilters,
} from '@/shared/types/legacy.types';
