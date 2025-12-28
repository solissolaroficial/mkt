import React, { useState } from 'react';
import PostCalendarView from '../ui/PostCalendarView';
import { MOCK_CALENDAR_POSTS } from '@/shared/utils/legacy.constants';

const CalendarPage: React.FC = () => {
  const [posts, setPosts] = useState(MOCK_CALENDAR_POSTS);

  const handleUpdatePosts = (updatedPosts: any[]) => {
    setPosts(updatedPosts);
  };

  return <PostCalendarView posts={posts} onUpdatePosts={handleUpdatePosts} />;
};

export default CalendarPage;
