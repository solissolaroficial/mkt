// UI Components
export { default as TaskView } from './ui/TaskView';
export { default as TaskWidget } from './ui/TaskWidget';
export { default as TaskModal } from './ui/TaskModal';
export { default as KanbanBoard } from './ui/KanbanBoard';

// Hooks
export { useTasks, useTask } from './hooks/useTasks';
export { useTaskMutations } from './hooks/useTaskMutations';

// Services
export { taskService } from './services/taskService';

// Types
export type { 
  TaskViewProps, 
  TaskWidgetProps, 
  TaskModalProps, 
  KanbanBoardProps 
} from './types';
