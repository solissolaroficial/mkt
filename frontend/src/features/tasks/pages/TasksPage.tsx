import React from 'react';
import TaskView from '../ui/TaskView';
import { MOCK_TASKS } from '@/shared/utils/legacy.constants';

const TasksPage: React.FC = () => {
  return <TaskView initialTasks={MOCK_TASKS} />;
};

export default TasksPage;
