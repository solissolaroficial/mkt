import type { Task, Subtask, Comment, AllowedUser } from '@/shared/types';

export type { Task, Subtask, Comment, AllowedUser };

export interface PaginatedTasks {
  data: Task[];
  total: number;
  page: number;
  limit: number;
  totalPages: number;
}

export interface TaskViewProps {
  initialTasks: Task[];
  onAddTask?: (task: Task) => void;
  onUpdateTask?: (task: Task) => void;
  onDeleteTask?: (taskId: string) => void;
  onAddSubtask?: (subtask: Subtask) => void;
  onUpdateSubtask?: (subtaskId: string, data: Partial<Subtask>) => void;
  onDeleteSubtask?: (subtaskId: string) => void;
  onAddComment?: (comment: Comment) => void;
  onUpdateComment?: (commentId: string, data: Partial<Comment>) => void;
  onDeleteComment?: (commentId: string) => void;
  onReorderTasks?: (taskIds: string[]) => void;
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
  onAddSubtask?: (subtask: Subtask) => void;
  onUpdateSubtask?: (subtaskId: string, data: Partial<Subtask>) => void;
  onDeleteSubtask?: (subtaskId: string) => void;
  onAddComment?: (comment: Comment) => void;
  onUpdateComment?: (commentId: string, data: Partial<Comment>) => void;
  onDeleteComment?: (commentId: string) => void;
}

export interface KanbanBoardProps {
  tasks: Task[];
  onAddTask: (task: Task) => void;
  onUpdateTask: (task: Task) => void;
  onDeleteTask: (taskId: string) => void;
  onReorderTasks: (taskIds: string[]) => void;
  onTaskClick: (taskId: string) => void;
}
