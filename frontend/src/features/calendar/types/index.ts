// Re-export calendar types from shared types
export type { CalendarPost, PostCategory, PostStatus, PostType, PostHistoryEvent } from '@/shared/types/legacy.types';

import type { CalendarPost } from '@/shared/types/legacy.types';

// Calendar component props
export interface PostCalendarViewProps {
  posts: CalendarPost[];
  onUpdatePosts: (posts: CalendarPost[]) => void;
}

export interface UpcomingPostsWidgetProps {
  posts: CalendarPost[];
  onViewCalendar: () => void;
}
