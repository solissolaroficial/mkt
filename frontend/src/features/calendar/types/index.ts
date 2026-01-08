import type { CalendarPost, CalendarPostsListResponse, PostCategory, PostType, PostStatus, PostHistoryEvent } from '@/shared/types/legacy.types';

// Calendar component props
export interface PostCalendarViewProps {
  posts: CalendarPost[];
  pagination?: CalendarPostsListResponse;
  mutations?: any; // TODO: Definir tipo específico para mutations
  selectedCategory?: PostCategory | 'all';
  setSelectedCategory?: (category: PostCategory | 'all') => void;
  selectedType?: PostType | 'all';
  setSelectedType?: (type: PostType | 'all') => void;
  selectedAssignee?: string;
  setSelectedAssignee?: (assignee: string) => void;
}

export interface UpcomingPostsWidgetProps {
  posts: CalendarPost[];
  onViewCalendar: () => void;
}

// Re-export types from legacy.types
export type {
  CalendarPost,
  CalendarPostsListResponse,
  PostCategory,
  PostType,
  PostStatus,
  PostHistoryEvent
} from '@/shared/types/legacy.types';
