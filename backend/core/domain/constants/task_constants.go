package constants

// TaskStatus define os possíveis status de uma tarefa
type TaskStatus string

const (
	// TaskStatusPending tarefa pendente
	TaskStatusPending TaskStatus = "pending"
	// TaskStatusInProgress tarefa em andamento
	TaskStatusInProgress TaskStatus = "in_progress"
	// TaskStatusCompleted tarefa concluída
	TaskStatusCompleted TaskStatus = "completed"
	// TaskStatusCancelled tarefa cancelada
	TaskStatusCancelled TaskStatus = "cancelled"
)

// ValidTaskStatuses lista de status válidos
var ValidTaskStatuses = []TaskStatus{
	TaskStatusPending,
	TaskStatusInProgress,
	TaskStatusCompleted,
	TaskStatusCancelled,
}

// IsValidTaskStatus verifica se o status é válido
func IsValidTaskStatus(status TaskStatus) bool {
	for _, s := range ValidTaskStatuses {
		if s == status {
			return true
		}
	}
	return false
}

// TaskPriority define as possíveis prioridades de uma tarefa
type TaskPriority string

const (
	// TaskPriorityLow prioridade baixa
	TaskPriorityLow TaskPriority = "low"
	// TaskPriorityMedium prioridade média
	TaskPriorityMedium TaskPriority = "medium"
	// TaskPriorityHigh prioridade alta
	TaskPriorityHigh TaskPriority = "high"
	// TaskPriorityUrgent prioridade urgente
	TaskPriorityUrgent TaskPriority = "urgent"
)

// ValidTaskPriorities lista de prioridades válidas
var ValidTaskPriorities = []TaskPriority{
	TaskPriorityLow,
	TaskPriorityMedium,
	TaskPriorityHigh,
	TaskPriorityUrgent,
}

// IsValidTaskPriority verifica se a prioridade é válida
func IsValidTaskPriority(priority TaskPriority) bool {
	for _, p := range ValidTaskPriorities {
		if p == priority {
			return true
		}
	}
	return false
}

// TaskCategory define as possíveis categorias de uma tarefa
type TaskCategory string

const (
	// TaskCategoryCommercial categoria comercial
	TaskCategoryCommercial TaskCategory = "commercial"
	// TaskCategoryMarketing categoria marketing
	TaskCategoryMarketing TaskCategory = "marketing"
	// TaskCategoryTraining categoria treinamento
	TaskCategoryTraining TaskCategory = "training"
	// TaskCategoryAdministrative categoria administrativa
	TaskCategoryAdministrative TaskCategory = "administrative"
	// TaskCategorySupport categoria suporte
	TaskCategorySupport TaskCategory = "support"
)

// ValidTaskCategories lista de categorias válidas
var ValidTaskCategories = []TaskCategory{
	TaskCategoryCommercial,
	TaskCategoryMarketing,
	TaskCategoryTraining,
	TaskCategoryAdministrative,
	TaskCategorySupport,
}

// IsValidTaskCategory verifica se a categoria é válida
func IsValidTaskCategory(category TaskCategory) bool {
	for _, c := range ValidTaskCategories {
		if c == category {
			return true
		}
	}
	return false
}

// TaskFlow define os possíveis fluxos de uma tarefa
type TaskFlow string

const (
	// TaskFlowBacklog fluxo backlog
	TaskFlowBacklog TaskFlow = "backlog"
	// TaskFlowToDo fluxo to do
	TaskFlowToDo TaskFlow = "to_do"
	// TaskFlowInProgress fluxo in progress
	TaskFlowInProgress TaskFlow = "in_progress"
	// TaskFlowReview fluxo review
	TaskFlowReview TaskFlow = "review"
	// TaskFlowDone fluxo done
	TaskFlowDone TaskFlow = "done"
)

// ValidTaskFlows lista de fluxos válidos
var ValidTaskFlows = []TaskFlow{
	TaskFlowBacklog,
	TaskFlowToDo,
	TaskFlowInProgress,
	TaskFlowReview,
	TaskFlowDone,
}

// IsValidTaskFlow verifica se o fluxo é válido
func IsValidTaskFlow(flow TaskFlow) bool {
	for _, f := range ValidTaskFlows {
		if f == flow {
			return true
		}
	}
	return false
}

// NotificationType define os possíveis tipos de notificação
type NotificationType string

const (
	// NotificationTypeTaskAssigned notificação de tarefa atribuída
	NotificationTypeTaskAssigned NotificationType = "task_assigned"
	// NotificationTypeTaskUpdated notificação de tarefa atualizada
	NotificationTypeTaskUpdated NotificationType = "task_updated"
	// NotificationTypeTaskCompleted notificação de tarefa concluída
	NotificationTypeTaskCompleted NotificationType = "task_completed"
	// NotificationTypeTaskOverdue notificação de tarefa atrasada
	NotificationTypeTaskOverdue NotificationType = "task_overdue"
	// NotificationTypeCommentAdded notificação de comentário adicionado
	NotificationTypeCommentAdded NotificationType = "comment_added"
	// NotificationTypeMention notificação de menção
	NotificationTypeMention NotificationType = "mention"
)

// ValidNotificationTypes lista de tipos de notificação válidos
var ValidNotificationTypes = []NotificationType{
	NotificationTypeTaskAssigned,
	NotificationTypeTaskUpdated,
	NotificationTypeTaskCompleted,
	NotificationTypeTaskOverdue,
	NotificationTypeCommentAdded,
	NotificationTypeMention,
}

// IsValidNotificationType verifica se o tipo de notificação é válido
func IsValidNotificationType(notificationType NotificationType) bool {
	for _, t := range ValidNotificationTypes {
		if t == notificationType {
			return true
		}
	}
	return false
}
