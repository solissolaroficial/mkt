import { apiClient } from '@/infrastructure/api/client';
import type { Task, Subtask, Comment } from '../types';

export const taskService = {
  // Get all tasks
  list: async (): Promise<Task[]> => {
    const response = await apiClient.get('/api/tasks');
    return response.data;
  },

  // Get task by ID
  getById: async (id: string): Promise<Task> => {
    const response = await apiClient.get(`/api/tasks/${id}`);
    return response.data;
  },

  // Create new task
  create: async (data: Partial<Task>): Promise<Task> => {
    const response = await apiClient.post('/api/tasks', data);
    return response.data;
  },

  // Update task
  update: async (id: string, data: Partial<Task>): Promise<Task> => {
    const response = await apiClient.put(`/api/tasks/${id}`, data);
    return response.data;
  },

  // Delete task
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/tasks/${id}`);
  },

  // Add subtask
  addSubtask: async (taskId: string, subtask: Partial<Subtask>): Promise<Task> => {
    const response = await apiClient.post(`/api/tasks/${taskId}/subtasks`, subtask);
    return response.data;
  },

  // Update subtask
  updateSubtask: async (taskId: string, subtaskId: string, data: Partial<Subtask>): Promise<Task> => {
    const response = await apiClient.put(`/api/tasks/${taskId}/subtasks/${subtaskId}`, data);
    return response.data;
  },

  // Delete subtask
  deleteSubtask: async (taskId: string, subtaskId: string): Promise<Task> => {
    const response = await apiClient.delete(`/api/tasks/${taskId}/subtasks/${subtaskId}`);
    return response.data;
  },

  // Add comment
  addComment: async (taskId: string, comment: Partial<Comment>): Promise<Task> => {
    const response = await apiClient.post(`/api/tasks/${taskId}/comments`, comment);
    return response.data;
  },

  // Reorder tasks
  reorder: async (tasks: Task[]): Promise<Task[]> => {
    const response = await apiClient.post('/api/tasks/reorder', { tasks });
    return response.data;
  },
};
