import { apiClient } from '@/infrastructure/api/client';
import type { Task, Subtask, Comment, PaginatedTasks } from '../types';

export const taskService = {
  // Get all tasks with pagination
  list: async (page = 1, limit = 100): Promise<PaginatedTasks> => {
    const response = await apiClient.get('/api/tasks', {
      params: { page, limit }
    });
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

  // Reorder tasks
  reorder: async (taskIds: string[]): Promise<void> => {
    await apiClient.post('/api/tasks/reorder', { task_ids: taskIds });
  },
};

export const subtaskService = {
  // Get subtasks by task ID
  list: async (taskId: string): Promise<Subtask[]> => {
    const response = await apiClient.get(`/api/subtasks?task_id=${taskId}`);
    return response.data;
  },

  // Get subtask by ID
  getById: async (id: string): Promise<Subtask> => {
    const response = await apiClient.get(`/api/subtasks/${id}`);
    return response.data;
  },

  // Create new subtask
  create: async (data: Partial<Subtask>): Promise<Subtask> => {
    const response = await apiClient.post('/api/subtasks', data);
    return response.data;
  },

  // Update subtask
  update: async (id: string, data: Partial<Subtask>): Promise<Subtask> => {
    const response = await apiClient.put(`/api/subtasks/${id}`, data);
    return response.data;
  },

  // Delete subtask
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/subtasks/${id}`);
  },
};

export const commentService = {
  // Get comments by task ID
  list: async (taskId: string): Promise<Comment[]> => {
    const response = await apiClient.get(`/api/comments?task_id=${taskId}`);
    return response.data;
  },

  // Get comment by ID
  getById: async (id: string): Promise<Comment> => {
    const response = await apiClient.get(`/api/comments/${id}`);
    return response.data;
  },

  // Create new comment
  create: async (data: Partial<Comment>): Promise<Comment> => {
    const response = await apiClient.post('/api/comments', data);
    return response.data;
  },

  // Update comment
  update: async (id: string, data: Partial<Comment>): Promise<Comment> => {
    const response = await apiClient.put(`/api/comments/${id}`, data);
    return response.data;
  },

  // Delete comment
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/comments/${id}`);
  },
};
