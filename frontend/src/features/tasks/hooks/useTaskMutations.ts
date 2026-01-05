import { useMutation, useQueryClient } from '@tanstack/react-query';
import { taskService, subtaskService, commentService } from '../services/taskService';
import type { Task, Subtask, Comment } from '../types';

export const useTaskMutations = () => {
  const queryClient = useQueryClient();

  const createTask = useMutation({
    mutationFn: (data: Partial<Task>) => taskService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
    onError: (error) => {
      console.error('Error creating task:', error);
    },
  });

  const updateTask = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Task> }) =>
      taskService.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
    onError: (error) => {
      console.error('Error updating task:', error);
    },
  });

  const deleteTask = useMutation({
    mutationFn: (id: string) => taskService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
    onError: (error) => {
      console.error('Error deleting task:', error);
    },
  });

  const createSubtask = useMutation({
    mutationFn: (data: Partial<Subtask>) => subtaskService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
      queryClient.invalidateQueries({ queryKey: ['subtasks'] });
    },
    onError: (error) => {
      console.error('Error creating subtask:', error);
    },
  });

  const updateSubtask = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Subtask> }) =>
      subtaskService.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
      queryClient.invalidateQueries({ queryKey: ['subtasks'] });
    },
    onError: (error) => {
      console.error('Error updating subtask:', error);
    },
  });

  const deleteSubtask = useMutation({
    mutationFn: (id: string) => subtaskService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
      queryClient.invalidateQueries({ queryKey: ['subtasks'] });
    },
    onError: (error) => {
      console.error('Error deleting subtask:', error);
    },
  });

  const createComment = useMutation({
    mutationFn: (data: Partial<Comment>) => commentService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
      queryClient.invalidateQueries({ queryKey: ['comments'] });
    },
    onError: (error) => {
      console.error('Error creating comment:', error);
    },
  });

  const updateComment = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Comment> }) =>
      commentService.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
      queryClient.invalidateQueries({ queryKey: ['comments'] });
    },
    onError: (error) => {
      console.error('Error updating comment:', error);
    },
  });

  const deleteComment = useMutation({
    mutationFn: (id: string) => commentService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
      queryClient.invalidateQueries({ queryKey: ['comments'] });
    },
    onError: (error) => {
      console.error('Error deleting comment:', error);
    },
  });

  const reorderTasks = useMutation({
    mutationFn: (taskIds: string[]) => taskService.reorder(taskIds),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
    onError: (error) => {
      console.error('Error reordering tasks:', error);
    },
  });

  return {
    createTask,
    updateTask,
    deleteTask,
    createSubtask,
    updateSubtask,
    deleteSubtask,
    createComment,
    updateComment,
    deleteComment,
    reorderTasks,
  };
};
