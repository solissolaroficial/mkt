import { apiClient } from '@/infrastructure/api/client';
import { ENDPOINTS } from '@/infrastructure/api/endpoints';
import type { Notification } from '@/shared/types/legacy.types';

// Interface para resposta paginada de notificações
export interface NotificationsListResponse {
  data: Notification[];
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

// Buscar todas as notificações de um usuário
export const fetchNotifications = async (
  userId: string,
  page: number = 1,
  limit: number = 20
): Promise<NotificationsListResponse> => {
  const response = await apiClient.get(ENDPOINTS.NOTIFICATIONS.LIST, {
    params: { user_id: userId, page, limit },
  });
  return response.data;
};

// Buscar uma notificação específica
export const fetchNotification = async (
  id: string
): Promise<Notification> => {
  const response = await apiClient.get(ENDPOINTS.NOTIFICATIONS.GET(id));
  return response.data;
};

// Criar uma nova notificação
export const createNotification = async (
  data: {
    user_id: string;
    task_id?: string;
    notification_type: 'mention' | 'deadline' | 'system';
    title: string;
    message: string;
  }
): Promise<Notification> => {
  const response = await apiClient.post(ENDPOINTS.NOTIFICATIONS.CREATE, data);
  return response.data;
};

// Atualizar uma notificação
export const updateNotification = async (
  id: string,
  data: { title?: string; message?: string }
): Promise<Notification> => {
  const response = await apiClient.put(ENDPOINTS.NOTIFICATIONS.UPDATE(id), data);
  return response.data;
};

// Deletar uma notificação
export const deleteNotification = async (id: string): Promise<void> => {
  await apiClient.delete(ENDPOINTS.NOTIFICATIONS.DELETE(id));
};

// Marcar uma notificação como lida
export const markAsRead = async (id: string): Promise<Notification> => {
  const response = await apiClient.post(ENDPOINTS.NOTIFICATIONS.MARK_AS_READ(id));
  return response.data;
};

// Marcar todas as notificações como lidas
export const markAllAsRead = async (
  userId: string
): Promise<Notification[]> => {
  const response = await apiClient.post(ENDPOINTS.NOTIFICATIONS.MARK_ALL_AS_READ, null, {
    params: { user_id: userId },
  });
  return response.data;
};

// Deletar notificações por task_id
export const deleteNotificationsByTask = async (
  taskId: string
): Promise<void> => {
  await apiClient.delete(ENDPOINTS.NOTIFICATIONS.DELETE_BY_TASK(taskId));
};
