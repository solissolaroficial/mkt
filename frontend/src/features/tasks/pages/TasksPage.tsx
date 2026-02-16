import React, { useState } from 'react';
import { useTasks, useTaskMutations } from '@/features/tasks/hooks';
import KanbanBoard from '../ui/KanbanBoard';
import TaskModal from '../ui/TaskModal';
import { ErrorDisplay } from '@/shared/components/ErrorDisplay';
import { ToastContainer } from '@/shared/components/ToastContainer';
import { useToast } from '@/shared/hooks/useToast';
import ConfirmationModal from '@/shared/components/ConfirmationModal';
import type { Task } from '@/shared/types';
import type { PaginatedTasks } from '../types';
import type { Subtask, Comment } from '../types';

const TasksPage: React.FC = () => {
  const { data: paginatedTasks, isLoading, error, refetch } = useTasks(1, 100);
  const tasks = paginatedTasks?.data || [];
  const { toasts, removeToast, success, error: toastError, info } = useToast();
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState<string | null>(null);
  const [showDeleteCommentConfirm, setShowDeleteCommentConfirm] = useState<string | null>(null);
  const {
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
  } = useTaskMutations();

  const handleAddTask = (task: Task) => {
    createTask.mutate(task, {
      onSuccess: () => {
        success('Tarefa criada com sucesso!');
      },
      onError: (err) => {
        toastError('Erro ao criar tarefa. Tente novamente.');
      },
    });
  };

  const handleUpdateTask = (task: Task) => {
    const { id, created_at, updated_at, subtasks, comments, notifications, ...data } = task;
    
    updateTask.mutate({
      id,
      data: {
        ...data,
        due_date: data.due_date || undefined,
        assignee_id: data.assignee_id || undefined,
        description: data.description || undefined,
      }
    }, {
      onSuccess: (updatedTaskData) => {
        // Atualizar o selectedTask com os dados retornados pela API
        setSelectedTask(prev => prev ? { ...prev, ...updatedTaskData } : null);
        success('Tarefa atualizada com sucesso!');
      },
      onError: (err) => {
        toastError('Erro ao atualizar tarefa. Tente novamente.');
      },
    });
  };

  const handleDeleteTask = (taskId: string) => {
    if (!confirm('Tem certeza que deseja excluir esta tarefa?')) {
      return;
    }
    deleteTask.mutate(taskId, {
      onSuccess: () => {
        success('Tarefa excluída com sucesso!');
        handleCloseModal();
      },
      onError: (err) => {
        toastError('Erro ao excluir tarefa. Tente novamente.');
      },
    });
  };

  const handleReorderTasks = (taskIds: string[]) => {
    reorderTasks.mutate(taskIds, {
      onError: (err) => {
        toastError('Erro ao reordenar tarefas. Tente novamente.');
      },
    });
  };

  const handleTaskClick = (taskId: string) => {
    const task = tasks.find(t => t.id === taskId);
    if (task) {
      setSelectedTask(task);
      setIsModalOpen(true);
    }
  };

  const handleCloseModal = () => {
    setSelectedTask(null);
    setIsModalOpen(false);
  };

  const handleMention = (taskId: string, taskTitle: string) => {
    // Mostrar um toast informando que o usuário foi mencionado
    info(`Você mencionou alguém na tarefa "${taskTitle}"`);
    
    // Log para debug
    console.log(`Usuário mencionado na tarefa ${taskId}: ${taskTitle}`);
    
    // TODO: Implementar notificação real no backend
    // Exemplo: criar uma notificação para o usuário mencionado
  };

  const handleAddSubtask = (subtask: Subtask) => {
    createSubtask.mutate(subtask, {
      onSuccess: () => {
        success('Subtarefa criada com sucesso!');
      },
      onError: (err) => {
        toastError('Erro ao criar subtarefa. Tente novamente.');
      },
    });
  };

  const handleUpdateSubtask = (id: string, data: Partial<Subtask>) => {
    updateSubtask.mutate({ id, data }, {
      onSuccess: (updatedSubtaskData) => {
        // Atualizar o selectedTask com os dados retornados pela API
        if (selectedTask) {
          setSelectedTask(prev => {
            // Encontrar e atualizar a subtarefa
            const updatedSubtasks = (prev.subtasks || []).map(st =>
              st.id === id ? { ...st, ...updatedSubtaskData } : st
            );
            return { ...prev, subtasks: updatedSubtasks };
          });
        }
        success('Subtarefa atualizada com sucesso!');
      },
      onError: (err) => {
        toastError('Erro ao atualizar subtarefa. Tente novamente.');
      },
    });
  };

  const handleDeleteSubtask = (id: string) => {
    setShowDeleteConfirm(id);
  };

  const handleConfirmDelete = (id: string) => {
    deleteSubtask.mutate(id, {
      onSuccess: () => {
        success('Subtarefa excluída com sucesso!');
        setShowDeleteConfirm(null);
      },
      onError: (err) => {
        toastError('Erro ao excluir subtarefa. Tente novamente.');
        setShowDeleteConfirm(null);
      },
    });
  };

  const handleAddComment = (comment: Comment) => {
    createComment.mutate(comment, {
      onSuccess: () => {
        success('Comentário adicionado com sucesso!');
      },
      onError: (err) => {
        toastError('Erro ao adicionar comentário. Tente novamente.');
      },
    });
  };

  const handleUpdateComment = (id: string, data: Partial<Comment>) => {
    updateComment.mutate({ id, data }, {
      onSuccess: () => {
        success('Comentário atualizado com sucesso!');
      },
      onError: (err) => {
        toastError('Erro ao atualizar comentário. Tente novamente.');
      },
    });
  };

  const handleDeleteComment = (id: string) => {
    setShowDeleteCommentConfirm(id);
  };

  const handleConfirmDeleteComment = (id: string) => {
    deleteComment.mutate(id, {
      onSuccess: () => {
        success('Comentário excluído com sucesso!');
        setShowDeleteCommentConfirm(null);
      },
      onError: (err) => {
        toastError('Erro ao excluir comentário. Tente novamente.');
        setShowDeleteCommentConfirm(null);
      },
    });
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="flex items-center gap-3 text-gray-500">
          <div className="w-5 h-5 border-2 border-gray-300 border-t-transparent rounded-full animate-spin"></div>
          <span>Carregando tarefas...</span>
        </div>
      </div>
    );
  }
  
  if (error) {
    return <ErrorDisplay error={error} onRetry={() => refetch()} />;
  }

  return (
    <>
      <ToastContainer toasts={toasts} onRemove={removeToast} />
      <KanbanBoard
        tasks={tasks || []}
        onAddTask={handleAddTask}
        onUpdateTask={handleUpdateTask}
        onDeleteTask={handleDeleteTask}
        onReorderTasks={handleReorderTasks}
        onTaskClick={handleTaskClick}
      />
      {selectedTask && (
        <TaskModal
          task={selectedTask}
          isOpen={isModalOpen}
          onClose={handleCloseModal}
          onUpdate={handleUpdateTask}
          onDelete={handleDeleteTask}
          onMention={handleMention}
          onAddSubtask={handleAddSubtask}
          onUpdateSubtask={handleUpdateSubtask}
          onDeleteSubtask={handleDeleteSubtask}
          onAddComment={handleAddComment}
          onUpdateComment={handleUpdateComment}
          onDeleteComment={handleDeleteComment}
        />
      )}
      <ConfirmationModal
        isOpen={!!showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(null)}
        onConfirm={() => handleConfirmDelete(showDeleteConfirm!)}
        title="Confirmar Exclusão"
        message="Tem certeza que deseja excluir esta subtarefa? Esta ação não pode ser desfeita."
        isPending={deleteSubtask.isPending}
        type="danger"
      />
      <ConfirmationModal
        isOpen={!!showDeleteCommentConfirm}
        onClose={() => setShowDeleteCommentConfirm(null)}
        onConfirm={() => handleConfirmDeleteComment(showDeleteCommentConfirm!)}
        title="Confirmar Exclusão"
        message="Tem certeza que deseja excluir este comentário? Esta ação não pode ser desfeita."
        isPending={deleteComment.isPending}
        type="danger"
      />
    </>
  );
};

export default TasksPage;
