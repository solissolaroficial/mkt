import type { Task, Subtask, Comment, AllowedUser } from '@/shared/types';

export type { Task, Subtask, Comment, AllowedUser };

export interface TaskViewProps {
  initialTasks: Task[];
}

export interface TaskWidgetProps {
  tasks: Task[];
  onViewAll?: () => void;
  onTaskClick?: (taskId: string) => void;
}

export interface TaskModalProps {
  task: Task;
  isOpen: boolean;
  onClose: () => void;
  onUpdate: (updatedTask: Task) => void;
  onDelete: (taskId: string) => void;
  onMention?: (taskId: string, taskTitle: string) => void;
}

export interface KanbanBoardProps {
  tasks: Task[];
  onAddTask: (task: Task) => void;
  onUpdateTask: (task: Task) => void;
  onDeleteTask: (taskId: string) => void;
  onReorderTasks: (tasks: Task[]) => void;
  onTaskClick: (taskId: string) => void;
}
