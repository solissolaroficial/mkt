import { useEffect, useCallback } from 'react';
import { useAuth } from '@/features/auth';
import { useNotificationStore } from '@/shared/store/notificationStore';
import {
  fetchNotifications,
  markAsRead as markAsReadAPI,
  markAllAsRead as markAllAsReadAPI,
  deleteNotification as deleteNotificationAPI,
} from '../services/notificationService';

export const useNotifications = () => {
  const { user } = useAuth();
  const {
    notifications,
    setNotifications,
    updateNotification,
    removeNotification,
    clearNotifications,
  } = useNotificationStore();

  // Buscar notificações do backend
  const fetchFromBackend = useCallback(async () => {
    if (!user?.id) return;

    try {
      const response = await fetchNotifications(user.id, 1, 50);
      setNotifications(response.data); // Importante: usar response.data
    } catch (error) {
      console.error('Erro ao buscar notificações:', error);
    }
  }, [user?.id, setNotifications]);

  // Marcar uma notificação como lida
  const markAsRead = useCallback(
    async (id: string) => {
      try {
        await markAsReadAPI(id);
        updateNotification(id, { read: true });
      } catch (error) {
        console.error('Erro ao marcar notificação como lida:', error);
      }
    },
    [updateNotification]
  );

  // Marcar todas as notificações como lidas
  const markAllAsRead = useCallback(async () => {
    if (!user?.id) return;

    try {
      await markAllAsReadAPI(user.id);
      notifications.forEach((n) => updateNotification(n.id, { read: true }));
    } catch (error) {
      console.error('Erro ao marcar todas como lidas:', error);
    }
  }, [user?.id, notifications, updateNotification]);

  // Deletar uma notificação
  const deleteNotification = useCallback(
    async (id: string) => {
      try {
        await deleteNotificationAPI(id);
        removeNotification(id);
      } catch (error) {
        console.error('Erro ao deletar notificação:', error);
      }
    },
    [removeNotification]
  );

  // Arquivar uma notificação (marcar como archived)
  const archiveNotification = useCallback(
    async (id: string) => {
      try {
        updateNotification(id, { archived: true });
      } catch (error) {
        console.error('Erro ao arquivar notificação:', error);
      }
    },
    [updateNotification]
  );

  // Desarquivar uma notificação
  const unarchiveNotification = useCallback(
    async (id: string) => {
      try {
        updateNotification(id, { archived: false, read: false });
      } catch (error) {
        console.error('Erro ao desarquivar notificação:', error);
      }
    },
    [updateNotification]
  );

  // Carregar notificações ao montar o componente
  useEffect(() => {
    fetchFromBackend();
  }, [fetchFromBackend]);

  // Calcular contadores
  const unreadCount = notifications.filter((n) => !n.read && !n.archived).length;
  const activeNotifications = notifications.filter((n) => !n.archived);

  return {
    notifications,
    activeNotifications,
    unreadCount,
    markAsRead,
    markAllAsRead,
    deleteNotification,
    archiveNotification,
    unarchiveNotification,
    clearNotifications,
    refetch: fetchFromBackend,
  };
};
