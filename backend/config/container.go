package config

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/seu-usuario/solis-backend/application/usecase/auth"
	"github.com/seu-usuario/solis-backend/application/usecase/brand"
	"github.com/seu-usuario/solis-backend/application/usecase/budget"
	"github.com/seu-usuario/solis-backend/application/usecase/calendar"
	offlineactionusecase "github.com/seu-usuario/solis-backend/application/usecase/cooperative/offlineaction"
	repmarketingactionusecase "github.com/seu-usuario/solis-backend/application/usecase/cooperative/repmarketingaction"
	showroomitemusecase "github.com/seu-usuario/solis-backend/application/usecase/cooperative/showroomitem"
	accountpayableusecase "github.com/seu-usuario/solis-backend/application/usecase/financial/accountpayable"
	"github.com/seu-usuario/solis-backend/application/usecase/gifts"
	"github.com/seu-usuario/solis-backend/application/usecase/kpis"
	usecasepdv "github.com/seu-usuario/solis-backend/application/usecase/pdv"
	representativemonthlygoal "github.com/seu-usuario/solis-backend/application/usecase/representativemonthlygoal"
	"github.com/seu-usuario/solis-backend/application/usecase/representatives"
	"github.com/seu-usuario/solis-backend/application/usecase/social"
	"github.com/seu-usuario/solis-backend/application/usecase/tasks"
	"github.com/seu-usuario/solis-backend/application/usecase/users"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/service"
	dbgateway "github.com/seu-usuario/solis-backend/dataprovider/database/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/infrastructure"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/routes"
	"github.com/seu-usuario/solis-backend/seeders"
)

// Container struct que agrupa TODAS as dependências da aplicação
type Container struct {
	// Database
	DB *gorm.DB

	// Infrastructure Services
	HasherService service.HasherService
	JwtService    service.JWTService

	// Gateways
	UserGateway                      gateway.UserGateway
	KpiGateway                       gateway.KpiGateway
	MonthlyDataGateway               gateway.MonthlyDataGateway
	TaskGateway                      gateway.TaskGateway
	SubtaskGateway                   gateway.SubtaskGateway
	CommentGateway                   gateway.CommentGateway
	NotificationGateway              gateway.NotificationGateway
	CalendarPostGateway              gateway.CalendarPostGateway
	PdvPostGateway                   gateway.PdvPostGateway
	RecurrentPdvGateway              gateway.RecurrentPdvGateway
	SocialBenchmarkingGateway        gateway.SocialBenchmarkingGateway
	SocialPostGateway                gateway.SocialPostGateway
	SocialDailyAggregationGateway    gateway.SocialDailyAggregationGateway
	OfflineActionGateway             gateway.OfflineActionGateway
	ShowroomItemGateway              gateway.ShowroomItemGateway
	RepMarketingActionGateway        gateway.RepMarketingActionGateway
	GiftItemGateway                  gateway.GiftItemGateway
	GiftTransactionGateway           gateway.GiftTransactionGateway
	AccountPayableGateway            gateway.AccountPayableGateway
	BudgetGateway                    gateway.BudgetGateway
	RepresentativeGateway            gateway.RepresentativeGateway
	RepresentativeStatsGateway       gateway.RepresentativeStatsGateway
	RepresentativeMonthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway
	BrandGateway                     gateway.BrandGateway

	// Seeders
	UserSeeder                      *seeders.UserSeeder
	KpiSeeder                       *seeders.KpiSeeder
	SocialBenchmarkingSeeder        *seeders.SocialBenchmarkingSeeder
	CooperativeSeeder               *seeders.CooperativeSeeder
	GiftSeeder                      *seeders.GiftSeeder
	BudgetSeeder                    *seeders.BudgetSeeder
	RepresentativeMonthlyGoalSeeder *seeders.RepresentativeMonthlyGoalSeeder

	// Use Cases - Auth
	LoginUseCase *auth.LoginUseCase

	// Use Cases - Users
	ListUsersUseCase *users.ListUsersUseCase

	// Use Cases - Brands
	CreateBrandUseCase *brand.CreateBrandUseCase
	ListBrandsUseCase  *brand.ListBrandsUseCase
	DeleteBrandUseCase *brand.DeleteBrandUseCase

	// Use Cases - Social Benchmarking
	CreateSocialBenchmarkingUseCase *social.CreateSocialBenchmarkingUseCase
	ListSocialBenchmarkingsUseCase  *social.ListSocialBenchmarkingsUseCase
	GetSocialBenchmarkingUseCase    *social.GetSocialBenchmarkingUseCase
	UpdateSocialBenchmarkingUseCase *social.UpdateSocialBenchmarkingUseCase
	DeleteSocialBenchmarkingUseCase *social.DeleteSocialBenchmarkingUseCase

	// Use Cases - Social Posts
	CreateSocialPostUseCase *social.CreateSocialPostUseCase
	GetSocialPostUseCase    *social.GetSocialPostUseCase
	ListSocialPostsUseCase  *social.ListSocialPostsUseCase
	UpdateSocialPostUseCase *social.UpdateSocialPostUseCase
	DeleteSocialPostUseCase *social.DeleteSocialPostUseCase

	// Use Cases - Social Daily Aggregations
	RecalculateDailyAggregationsUseCase *social.RecalculateDailyAggregationsUseCase
	ListSocialDailyAggregationsUseCase  *social.ListSocialDailyAggregationsUseCase
	GetSocialDailyAggregationUseCase    *social.GetSocialDailyAggregationUseCase

	// Use Cases - Offline Actions
	CreateOfflineActionUseCase *offlineactionusecase.CreateOfflineActionUseCase
	ListOfflineActionsUseCase  *offlineactionusecase.ListOfflineActionsUseCase
	GetOfflineActionUseCase    *offlineactionusecase.GetOfflineActionUseCase
	UpdateOfflineActionUseCase *offlineactionusecase.UpdateOfflineActionUseCase
	DeleteOfflineActionUseCase *offlineactionusecase.DeleteOfflineActionUseCase

	// Use Cases - Showroom Items
	CreateShowroomItemUseCase *showroomitemusecase.CreateShowroomItemUseCase
	ListShowroomItemsUseCase  *showroomitemusecase.ListShowroomItemsUseCase
	GetShowroomItemUseCase    *showroomitemusecase.GetShowroomItemUseCase
	UpdateShowroomItemUseCase *showroomitemusecase.UpdateShowroomItemUseCase
	DeleteShowroomItemUseCase *showroomitemusecase.DeleteShowroomItemUseCase

	// Use Cases - Rep Marketing Actions
	CreateRepMarketingActionUseCase *repmarketingactionusecase.CreateRepMarketingActionUseCase
	ListRepMarketingActionsUseCase  *repmarketingactionusecase.ListRepMarketingActionsUseCase
	GetRepMarketingActionUseCase    *repmarketingactionusecase.GetRepMarketingActionUseCase
	UpdateRepMarketingActionUseCase *repmarketingactionusecase.UpdateRepMarketingActionUseCase
	DeleteRepMarketingActionUseCase *repmarketingactionusecase.DeleteRepMarketingActionUseCase

	// Use Cases - Gift Items
	CreateGiftItemUseCase *gifts.CreateGiftItemUseCase
	GetGiftItemUseCase    *gifts.GetGiftItemUseCase
	ListGiftItemsUseCase  *gifts.ListGiftItemsUseCase
	UpdateGiftItemUseCase *gifts.UpdateGiftItemUseCase
	DeleteGiftItemUseCase *gifts.DeleteGiftItemUseCase

	// Use Cases - Gift Transactions
	CreateGiftTransactionUseCase *gifts.CreateGiftTransactionUseCase
	GetGiftTransactionUseCase    *gifts.GetGiftTransactionUseCase
	ListGiftTransactionsUseCase  *gifts.ListGiftTransactionsUseCase
	UpdateGiftTransactionUseCase *gifts.UpdateGiftTransactionUseCase
	DeleteGiftTransactionUseCase *gifts.DeleteGiftTransactionUseCase

	// Use Cases - Account Payable
	CreateAccountPayableUseCase *accountpayableusecase.CreateAccountPayableUseCase
	ListAccountsPayableUseCase  *accountpayableusecase.ListAccountsPayableUseCase
	GetAccountPayableUseCase    *accountpayableusecase.GetAccountPayableUseCase
	UpdateAccountPayableUseCase *accountpayableusecase.UpdateAccountPayableUseCase
	DeleteAccountPayableUseCase *accountpayableusecase.DeleteAccountPayableUseCase
	ToggleNFUseCase             *accountpayableusecase.ToggleNFUseCase
	ToggleBoletoUseCase         *accountpayableusecase.ToggleBoletoUseCase
	SendToFinanceUseCase        *accountpayableusecase.SendToFinanceUseCase

	// Use Cases - Budget
	CreateBudgetItemUseCase       *budget.CreateBudgetItemUseCase
	ListBudgetItemsUseCase        *budget.ListBudgetItemsUseCase
	GetBudgetItemUseCase          *budget.GetBudgetItemUseCase
	UpdateBudgetItemUseCase       *budget.UpdateBudgetItemUseCase
	DeleteBudgetItemUseCase       *budget.DeleteBudgetItemUseCase
	BatchCreateBudgetItemsUseCase *budget.BatchCreateBudgetItemsUseCase
	GetBudgetSummaryUseCase       *budget.GetBudgetSummaryUseCase

	// Use Cases - Representatives
	CreateRepresentativeUseCase         *representatives.CreateRepresentativeUseCase
	GetRepresentativeUseCase            *representatives.GetRepresentativeUseCase
	UpdateRepresentativeUseCase         *representatives.UpdateRepresentativeUseCase
	DeleteRepresentativeUseCase         *representatives.DeleteRepresentativeUseCase
	ListRepresentativesUseCase          *representatives.ListRepresentativesUseCase
	GetRepresentativeStatsUseCase       *representatives.GetRepresentativeStatsUseCase
	GetRepresentativeProfileUseCase     *representatives.GetRepresentativeProfileUseCase
	GetAllRepresentativeProfilesUseCase *representatives.GetAllRepresentativeProfilesUseCase

	// Use Cases - Representative Monthly Goals
	CreateRepresentativeMonthlyGoalUseCase *representativemonthlygoal.CreateRepresentativeMonthlyGoalUseCase
	GetRepresentativeMonthlyGoalUseCase    *representativemonthlygoal.GetRepresentativeMonthlyGoalUseCase
	UpdateRepresentativeMonthlyGoalUseCase *representativemonthlygoal.UpdateRepresentativeMonthlyGoalUseCase
	DeleteRepresentativeMonthlyGoalUseCase *representativemonthlygoal.DeleteRepresentativeMonthlyGoalUseCase
	ListRepresentativeMonthlyGoalsUseCase  *representativemonthlygoal.ListRepresentativeMonthlyGoalsUseCase
	GetRepresentativeGoalsTableDataUseCase *representativemonthlygoal.GetRepresentativeGoalsTableDataUseCase

	// Use Cases - Calendar
	CreateCalendarPostUseCase            *calendar.CreateCalendarPostUseCase
	GetCalendarPostUseCase               *calendar.GetCalendarPostUseCase
	UpdateCalendarPostUseCase            *calendar.UpdateCalendarPostUseCase
	UpdateCalendarPostStatusUseCase      *calendar.UpdateCalendarPostStatusUseCase
	ConfirmCalendarPostPublishingUseCase *calendar.ConfirmCalendarPostPublishingUseCase
	DeleteCalendarPostUseCase            *calendar.DeleteCalendarPostUseCase
	ListCalendarPostsUseCase             *calendar.ListCalendarPostsUseCase

	// Use Cases - KPIs
	CreateKpiUseCase         *kpis.CreateKpiUseCase
	GetKpiUseCase            *kpis.GetKpiUseCase
	ListKpisUseCase          *kpis.ListKpisUseCase
	GetKpisBySlugsUseCase    *kpis.GetKpisBySlugsUseCase
	UpdateKpiUseCase         *kpis.UpdateKpiUseCase
	DeleteKpiUseCase         *kpis.DeleteKpiUseCase
	UpdateMonthlyDataUseCase *kpis.UpdateMonthlyDataUseCase

	// Use Cases - Tasks
	CreateTaskUseCase                  *tasks.CreateTaskUseCase
	UpdateTaskUseCase                  *tasks.UpdateTaskUseCase
	DeleteTaskUseCase                  *tasks.DeleteTaskUseCase
	GetTaskUseCase                     *tasks.GetTaskUseCase
	ListTasksUseCase                   *tasks.ListTasksUseCase
	ReorderTasksUseCase                *tasks.ReorderTasksUseCase
	CreateSubtaskUseCase               *tasks.CreateSubtaskUseCase
	UpdateSubtaskUseCase               *tasks.UpdateSubtaskUseCase
	DeleteSubtaskUseCase               *tasks.DeleteSubtaskUseCase
	GetSubtaskUseCase                  *tasks.GetSubtaskUseCase
	ListSubtasksUseCase                *tasks.ListSubtasksUseCase
	CreateCommentUseCase               *tasks.CreateCommentUseCase
	UpdateCommentUseCase               *tasks.UpdateCommentUseCase
	DeleteCommentUseCase               *tasks.DeleteCommentUseCase
	GetCommentUseCase                  *tasks.GetCommentUseCase
	ListCommentsUseCase                *tasks.ListCommentsUseCase
	CreateNotificationUseCase          *tasks.CreateNotificationUseCase
	UpdateNotificationUseCase          *tasks.UpdateNotificationUseCase
	DeleteNotificationUseCase          *tasks.DeleteNotificationUseCase
	GetNotificationUseCase             *tasks.GetNotificationUseCase
	ListNotificationsUseCase           *tasks.ListNotificationsUseCase
	MarkAsReadNotificationUseCase      *tasks.MarkAsReadNotificationUseCase
	MarkAllAsReadNotificationsUseCase  *tasks.MarkAllAsReadNotificationsUseCase
	DeleteNotificationsByTaskIDUseCase *tasks.DeleteNotificationsByTaskIDUseCase

	// Controllers
	AuthController                      *controller.AuthController
	KpiController                       *controller.KpiController
	TaskController                      *controller.TaskController
	SubtaskController                   *controller.SubtaskController
	CommentController                   *controller.CommentController
	NotificationController              *controller.NotificationController
	UserController                      *controller.UserController
	CalendarPostController              *controller.CalendarPostController
	PdvController                       *controller.PdvController
	SocialController                    *controller.SocialController
	SocialPostController                *controller.SocialPostController
	OfflineActionController             *controller.OfflineActionController
	ShowroomItemController              *controller.ShowroomItemController
	RepMarketingActionController        *controller.RepMarketingActionController
	GiftItemController                  *controller.GiftItemController
	GiftTransactionController           *controller.GiftTransactionController
	AccountPayableController            *controller.AccountPayableController
	BudgetController                    *controller.BudgetController
	RepresentativeController            *controller.RepresentativeController
	RepresentativeMonthlyGoalController *controller.RepresentativeMonthlyGoalController
	BrandController                     *controller.BrandController

	// Middlewares
	AuthMiddleware *middleware.AuthMiddleware
	CorsMiddleware fiber.Handler
}

// NewContainer cria e inicializa todas as dependências da aplicação
func NewContainer(cfg *Config) (*Container, error) {
	// 1. Database Connection
	db, err := initDatabase(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	log.Println("✅ Database connected successfully")

	// 2. Infrastructure Services (sem dependências)
	hasherService := infrastructure.NewHasherService()
	jwtService := infrastructure.NewJwtService(
		cfg.JWT.Secret,
		time.Duration(cfg.JWT.AccessTokenExpiryHours)*time.Hour,
		time.Duration(cfg.JWT.RefreshTokenExpiryHours)*time.Hour,
	)
	log.Println("✅ Services initialized")

	// 3. Gateways (dependem do DB)
	userGateway := dbgateway.NewUserGateway(db)
	kpiGateway := dbgateway.NewKpiGateway(db)
	monthlyDataGateway := dbgateway.NewMonthlyDataGateway(db)
	taskGateway := dbgateway.NewTaskGateway(db)
	subtaskGateway := dbgateway.NewSubtaskGateway(db)
	commentGateway := dbgateway.NewCommentGateway(db)
	notificationGateway := dbgateway.NewNotificationGateway(db)
	calendarPostGateway := dbgateway.NewCalendarPostGateway(db)
	pdvPostGateway := dbgateway.NewPdvPostGateway(db)
	recurrentPdvGateway := dbgateway.NewRecurrentPdvGateway(db)
	socialBenchmarkingGateway := dbgateway.NewSocialBenchmarkingGateway(db)
	socialPostGateway := dbgateway.NewSocialPostGateway(db)
	socialDailyAggregationGateway := dbgateway.NewSocialDailyAggregationGateway(db)
	offlineActionGateway := dbgateway.NewOfflineActionGateway(db)
	showroomItemGateway := dbgateway.NewShowroomItemGateway(db)
	repMarketingActionGateway := dbgateway.NewRepMarketingActionGateway(db)
	giftItemGateway := dbgateway.NewGiftItemGateway(db)
	giftTransactionGateway := dbgateway.NewGiftTransactionGateway(db)
	accountPayableGateway := dbgateway.NewAccountPayableGateway(db)
	budgetGateway := dbgateway.NewBudgetGateway(db)
	representativeGateway := dbgateway.NewRepresentativeGateway(db)
	representativeStatsGateway := dbgateway.NewRepresentativeStatsGateway(db)
	representativeMonthlyGoalGateway := dbgateway.NewRepresentativeMonthlyGoalGateway(db)
	brandGateway := dbgateway.NewBrandGateway(db)
	log.Println("✅ Gateways initialized")

	// 3.1 Seeders (dependem de gateways e services)
	userSeeder := seeders.NewUserSeeder(userGateway, hasherService)
	kpiSeeder := seeders.NewKpiSeeder(kpiGateway, monthlyDataGateway)
	socialBenchmarkingSeeder := seeders.NewSocialBenchmarkingSeeder(socialBenchmarkingGateway)
	cooperativeSeeder := seeders.NewCooperativeSeeder(offlineActionGateway, showroomItemGateway, repMarketingActionGateway, representativeGateway)
	giftSeeder := seeders.NewGiftSeeder(giftItemGateway, giftTransactionGateway)
	budgetSeeder := seeders.NewBudgetSeeder(budgetGateway)
	representativeMonthlyGoalSeeder := seeders.NewRepresentativeMonthlyGoalSeeder(representativeMonthlyGoalGateway, representativeGateway)
	log.Println("✅ Seeders initialized")

	// 4. Use Cases (dependem de gateways e services)
	// Auth Use Cases
	loginUseCase := auth.NewLoginUseCase(userGateway, hasherService, jwtService)

	// KPI Use Cases
	createKpiUseCase := kpis.NewCreateKpiUseCase(kpiGateway)
	getKpiUseCase := kpis.NewGetKpiUseCase(kpiGateway)
	listKpisUseCase := kpis.NewListKpisUseCase(kpiGateway)
	getKpisBySlugsUseCase := kpis.NewGetKpisBySlugsUseCase(kpiGateway)
	updateKpiUseCase := kpis.NewUpdateKpiUseCase(kpiGateway)
	deleteKpiUseCase := kpis.NewDeleteKpiUseCase(kpiGateway)
	updateMonthlyDataUseCase := kpis.NewUpdateMonthlyDataUseCase(
		monthlyDataGateway,
		kpiGateway,
	)

	// Task Use Cases
	createTaskUseCase := tasks.NewCreateTaskUseCase(taskGateway, subtaskGateway)
	updateTaskUseCase := tasks.NewUpdateTaskUseCase(taskGateway)
	deleteTaskUseCase := tasks.NewDeleteTaskUseCase(taskGateway)
	getTaskUseCase := tasks.NewGetTaskUseCase(taskGateway)
	listTasksUseCase := tasks.NewListTasksUseCase(taskGateway)
	reorderTasksUseCase := tasks.NewReorderTasksUseCase(taskGateway)
	createSubtaskUseCase := tasks.NewCreateSubtaskUseCase(subtaskGateway)
	updateSubtaskUseCase := tasks.NewUpdateSubtaskUseCase(subtaskGateway)
	deleteSubtaskUseCase := tasks.NewDeleteSubtaskUseCase(subtaskGateway)
	getSubtaskUseCase := tasks.NewGetSubtaskUseCase(subtaskGateway)
	listSubtasksUseCase := tasks.NewListSubtasksUseCase(subtaskGateway)
	createCommentUseCase := tasks.NewCreateCommentUseCase(
		commentGateway,
		notificationGateway,
		userGateway,
	)
	updateCommentUseCase := tasks.NewUpdateCommentUseCase(commentGateway)
	deleteCommentUseCase := tasks.NewDeleteCommentUseCase(commentGateway)
	getCommentUseCase := tasks.NewGetCommentUseCase(commentGateway)
	listCommentsUseCase := tasks.NewListCommentsUseCase(commentGateway)
	createNotificationUseCase := tasks.NewCreateNotificationUseCase(notificationGateway)
	updateNotificationUseCase := tasks.NewUpdateNotificationUseCase(notificationGateway)
	deleteNotificationUseCase := tasks.NewDeleteNotificationUseCase(notificationGateway)
	getNotificationUseCase := tasks.NewGetNotificationUseCase(notificationGateway)
	listNotificationsUseCase := tasks.NewListNotificationsUseCase(notificationGateway)
	markAsReadNotificationUseCase := tasks.NewMarkAsReadNotificationUseCase(notificationGateway)
	markAllAsReadNotificationsUseCase := tasks.NewMarkAllAsReadNotificationsUseCase(notificationGateway)
	deleteNotificationsByTaskIDUseCase := tasks.NewDeleteNotificationsByTaskIDUseCase(notificationGateway)

	// User Use Cases
	listUsersUseCase := users.NewListUsersUseCase(userGateway)

	// Brand Use Cases
	createBrandUseCase := brand.NewCreateBrandUseCase(brandGateway)
	listBrandsUseCase := brand.NewListBrandsUseCase(brandGateway)
	deleteBrandUseCase := brand.NewDeleteBrandUseCase(brandGateway)

	// Calendar Use Cases
	createCalendarPostUseCase := calendar.NewCreateCalendarPost(calendarPostGateway)
	getCalendarPostUseCase := calendar.NewGetCalendarPost(calendarPostGateway)
	updateCalendarPostUseCase := calendar.NewUpdateCalendarPost(calendarPostGateway)
	updateCalendarPostStatusUseCase := calendar.NewUpdateCalendarPostStatus(calendarPostGateway)
	confirmCalendarPostPublishingUseCase := calendar.NewConfirmCalendarPostPublishing(calendarPostGateway)
	deleteCalendarPostUseCase := calendar.NewDeleteCalendarPost(calendarPostGateway)
	listCalendarPostsUseCase := calendar.NewListCalendarPosts(calendarPostGateway)

	// PDV Use Cases - PdvPost
	createPdvPostUseCase := usecasepdv.NewCreatePdvPost(pdvPostGateway)
	listPdvPostsUseCase := usecasepdv.NewListPdvPosts(pdvPostGateway)
	getPdvPostUseCase := usecasepdv.NewGetPdvPost(pdvPostGateway)
	updatePdvPostUseCase := usecasepdv.NewUpdatePdvPost(pdvPostGateway)
	updatePdvPostStatusUseCase := usecasepdv.NewUpdatePdvPostStatus(pdvPostGateway)
	deletePdvPostUseCase := usecasepdv.NewDeletePdvPost(pdvPostGateway)

	// PDV Use Cases - RecurrentPdv
	createRecurrentPdvUseCase := usecasepdv.NewCreateRecurrentPdv(recurrentPdvGateway)
	getRecurrentPdvUseCase := usecasepdv.NewGetRecurrentPdv(recurrentPdvGateway)
	updateRecurrentPdvUseCase := usecasepdv.NewUpdateRecurrentPdv(recurrentPdvGateway)
	listRecurrentPdvsUseCase := usecasepdv.NewListRecurrentPdvs(recurrentPdvGateway)
	deleteRecurrentPdvUseCase := usecasepdv.NewDeleteRecurrentPdv(recurrentPdvGateway)

	// Social Benchmarking Use Cases
	createSocialBenchmarkingUseCase := social.NewCreateSocialBenchmarking(socialBenchmarkingGateway)
	listSocialBenchmarkingsUseCase := social.NewListSocialBenchmarkings(socialBenchmarkingGateway)
	getSocialBenchmarkingUseCase := social.NewGetSocialBenchmarking(socialBenchmarkingGateway)
	updateSocialBenchmarkingUseCase := social.NewUpdateSocialBenchmarking(socialBenchmarkingGateway)
	deleteSocialBenchmarkingUseCase := social.NewDeleteSocialBenchmarking(socialBenchmarkingGateway)

	// Social Daily Aggregation Use Cases (must be before Social Post Use Cases)
	recalculateDailyAggregationsUseCase := social.NewRecalculateDailyAggregationsUseCase(socialPostGateway, socialDailyAggregationGateway, socialBenchmarkingGateway)
	listSocialDailyAggregationsUseCase := social.NewListSocialDailyAggregationsUseCase(socialDailyAggregationGateway)
	getSocialDailyAggregationUseCase := social.NewGetSocialDailyAggregationUseCase(socialDailyAggregationGateway)

	// Social Post Use Cases
	createSocialPostUseCase := social.NewCreateSocialPostUseCase(socialPostGateway, brandGateway, recalculateDailyAggregationsUseCase)
	getSocialPostUseCase := social.NewGetSocialPostUseCase(socialPostGateway)
	listSocialPostsUseCase := social.NewListSocialPostsUseCase(socialPostGateway)
	updateSocialPostUseCase := social.NewUpdateSocialPostUseCase(socialPostGateway, brandGateway, recalculateDailyAggregationsUseCase)
	deleteSocialPostUseCase := social.NewDeleteSocialPostUseCase(socialPostGateway, recalculateDailyAggregationsUseCase)

	// Cooperative Use Cases - Offline Actions
	createOfflineActionUseCase := offlineactionusecase.NewCreateOfflineAction(offlineActionGateway)
	listOfflineActionsUseCase := offlineactionusecase.NewListOfflineActions(offlineActionGateway)
	getOfflineActionUseCase := offlineactionusecase.NewGetOfflineAction(offlineActionGateway)
	updateOfflineActionUseCase := offlineactionusecase.NewUpdateOfflineAction(offlineActionGateway)
	deleteOfflineActionUseCase := offlineactionusecase.NewDeleteOfflineAction(offlineActionGateway)

	// Cooperative Use Cases - Showroom Items
	createShowroomItemUseCase := showroomitemusecase.NewCreateShowroomItem(showroomItemGateway)
	listShowroomItemsUseCase := showroomitemusecase.NewListShowroomItems(showroomItemGateway)
	getShowroomItemUseCase := showroomitemusecase.NewGetShowroomItem(showroomItemGateway)
	updateShowroomItemUseCase := showroomitemusecase.NewUpdateShowroomItem(showroomItemGateway)
	deleteShowroomItemUseCase := showroomitemusecase.NewDeleteShowroomItem(showroomItemGateway)

	// Cooperative Use Cases - Rep Marketing Actions
	createRepMarketingActionUseCase := repmarketingactionusecase.NewCreateRepMarketingAction(repMarketingActionGateway)
	listRepMarketingActionsUseCase := repmarketingactionusecase.NewListRepMarketingActions(repMarketingActionGateway)
	getRepMarketingActionUseCase := repmarketingactionusecase.NewGetRepMarketingAction(repMarketingActionGateway)
	updateRepMarketingActionUseCase := repmarketingactionusecase.NewUpdateRepMarketingAction(repMarketingActionGateway)
	deleteRepMarketingActionUseCase := repmarketingactionusecase.NewDeleteRepMarketingAction(repMarketingActionGateway)

	// Gift Items Use Cases
	createGiftItemUseCase := gifts.NewCreateGiftItemUseCase(giftItemGateway)
	getGiftItemUseCase := gifts.NewGetGiftItemUseCase(giftItemGateway)
	listGiftItemsUseCase := gifts.NewListGiftItemsUseCase(giftItemGateway)
	updateGiftItemUseCase := gifts.NewUpdateGiftItemUseCase(giftItemGateway)
	deleteGiftItemUseCase := gifts.NewDeleteGiftItemUseCase(giftItemGateway, giftTransactionGateway)

	// Gift Transactions Use Cases
	createGiftTransactionUseCase := gifts.NewCreateGiftTransactionUseCase(giftTransactionGateway, giftItemGateway)
	getGiftTransactionUseCase := gifts.NewGetGiftTransactionUseCase(giftTransactionGateway)
	listGiftTransactionsUseCase := gifts.NewListGiftTransactionsUseCase(giftTransactionGateway)
	updateGiftTransactionUseCase := gifts.NewUpdateGiftTransactionUseCase(giftTransactionGateway)
	deleteGiftTransactionUseCase := gifts.NewDeleteGiftTransactionUseCase(giftTransactionGateway)

	// Account Payable Use Cases
	createAccountPayableUseCase := accountpayableusecase.NewCreateAccountPayable(accountPayableGateway)
	listAccountsPayableUseCase := accountpayableusecase.NewListAccountsPayable(accountPayableGateway)
	getAccountPayableUseCase := accountpayableusecase.NewGetAccountPayable(accountPayableGateway)
	updateAccountPayableUseCase := accountpayableusecase.NewUpdateAccountPayable(accountPayableGateway)
	deleteAccountPayableUseCase := accountpayableusecase.NewDeleteAccountPayable(accountPayableGateway)
	toggleNFUseCase := accountpayableusecase.NewToggleNFUseCase(accountPayableGateway)
	toggleBoletoUseCase := accountpayableusecase.NewToggleBoletoUseCase(accountPayableGateway)
	sendToFinanceUseCase := accountpayableusecase.NewSendToFinanceUseCase(accountPayableGateway)

	// Budget Use Cases
	createBudgetItemUseCase := budget.NewCreateBudgetItemUseCase(budgetGateway)
	listBudgetItemsUseCase := budget.NewListBudgetItemsUseCase(budgetGateway)
	getBudgetItemUseCase := budget.NewGetBudgetItemUseCase(budgetGateway)
	updateBudgetItemUseCase := budget.NewUpdateBudgetItemUseCase(budgetGateway)
	deleteBudgetItemUseCase := budget.NewDeleteBudgetItemUseCase(budgetGateway)
	batchCreateBudgetItemsUseCase := budget.NewBatchCreateBudgetItemsUseCase(budgetGateway)
	getBudgetSummaryUseCase := budget.NewGetBudgetSummaryUseCase(budgetGateway)

	// Representatives Use Cases
	createRepresentativeUseCase := representatives.NewCreateRepresentativeUseCase(representativeGateway)
	getRepresentativeUseCase := representatives.NewGetRepresentativeUseCase(representativeGateway)
	updateRepresentativeUseCase := representatives.NewUpdateRepresentativeUseCase(representativeGateway)
	deleteRepresentativeUseCase := representatives.NewDeleteRepresentativeUseCase(representativeGateway)
	listRepresentativesUseCase := representatives.NewListRepresentativesUseCase(representativeGateway)
	getRepresentativeStatsUseCase := representatives.NewGetRepresentativeStatsUseCase(representativeGateway, representativeStatsGateway)
	getRepresentativeProfileUseCase := representatives.NewGetRepresentativeProfileUseCase(representativeGateway, representativeStatsGateway)
	getAllRepresentativeProfilesUseCase := representatives.NewGetAllRepresentativeProfilesUseCase(representativeGateway, representativeStatsGateway)

	// Representative Monthly Goals Use Cases
	createRepresentativeMonthlyGoalUseCase := representativemonthlygoal.NewCreateRepresentativeMonthlyGoalUseCase(representativeMonthlyGoalGateway, representativeGateway)
	getRepresentativeMonthlyGoalUseCase := representativemonthlygoal.NewGetRepresentativeMonthlyGoalUseCase(representativeMonthlyGoalGateway)
	updateRepresentativeMonthlyGoalUseCase := representativemonthlygoal.NewUpdateRepresentativeMonthlyGoalUseCase(representativeMonthlyGoalGateway)
	deleteRepresentativeMonthlyGoalUseCase := representativemonthlygoal.NewDeleteRepresentativeMonthlyGoalUseCase(representativeMonthlyGoalGateway)
	listRepresentativeMonthlyGoalsUseCase := representativemonthlygoal.NewListRepresentativeMonthlyGoalsUseCase(representativeMonthlyGoalGateway, representativeGateway)
	getRepresentativeGoalsTableDataUseCase := representativemonthlygoal.NewGetRepresentativeGoalsTableDataUseCase(representativeMonthlyGoalGateway)

	log.Println("✅ Use cases initialized")

	// 5. Controllers (dependem de use cases)
	authController := controller.NewAuthController(loginUseCase)
	kpiController := controller.NewKpiController(
		createKpiUseCase,
		getKpiUseCase,
		listKpisUseCase,
		getKpisBySlugsUseCase,
		updateKpiUseCase,
		deleteKpiUseCase,
		updateMonthlyDataUseCase,
	)
	taskController := controller.NewTaskController(
		createTaskUseCase,
		updateTaskUseCase,
		deleteTaskUseCase,
		getTaskUseCase,
		listTasksUseCase,
		reorderTasksUseCase,
	)
	subtaskController := controller.NewSubtaskController(
		createSubtaskUseCase,
		updateSubtaskUseCase,
		deleteSubtaskUseCase,
		getSubtaskUseCase,
		listSubtasksUseCase,
	)
	commentController := controller.NewCommentController(
		createCommentUseCase,
		updateCommentUseCase,
		deleteCommentUseCase,
		getCommentUseCase,
		listCommentsUseCase,
	)
	notificationController := controller.NewNotificationController(
		createNotificationUseCase,
		updateNotificationUseCase,
		deleteNotificationUseCase,
		getNotificationUseCase,
		listNotificationsUseCase,
		markAsReadNotificationUseCase,
		markAllAsReadNotificationsUseCase,
		deleteNotificationsByTaskIDUseCase,
	)

	userController := controller.NewUserController(listUsersUseCase)

	calendarPostController := controller.NewCalendarPostController(
		createCalendarPostUseCase,
		getCalendarPostUseCase,
		updateCalendarPostUseCase,
		updateCalendarPostStatusUseCase,
		confirmCalendarPostPublishingUseCase,
		deleteCalendarPostUseCase,
		listCalendarPostsUseCase,
	)

	pdvController := controller.NewPdvController(
		createPdvPostUseCase,
		listPdvPostsUseCase,
		getPdvPostUseCase,
		updatePdvPostUseCase,
		updatePdvPostStatusUseCase,
		deletePdvPostUseCase,
		createRecurrentPdvUseCase,
		getRecurrentPdvUseCase,
		updateRecurrentPdvUseCase,
		listRecurrentPdvsUseCase,
		deleteRecurrentPdvUseCase,
	)

	// Social Benchmarking Controller
	socialController := controller.NewSocialController(
		createSocialBenchmarkingUseCase,
		listSocialBenchmarkingsUseCase,
		getSocialBenchmarkingUseCase,
		updateSocialBenchmarkingUseCase,
		deleteSocialBenchmarkingUseCase,
	)

	// Social Post Controller
	socialPostController := controller.NewSocialPostController(
		createSocialPostUseCase,
		updateSocialPostUseCase,
		deleteSocialPostUseCase,
		listSocialPostsUseCase,
		getSocialPostUseCase,
		recalculateDailyAggregationsUseCase,
		listSocialDailyAggregationsUseCase,
		getSocialDailyAggregationUseCase,
	)

	// Cooperative Controllers
	offlineActionController := controller.NewOfflineActionController(
		createOfflineActionUseCase,
		listOfflineActionsUseCase,
		getOfflineActionUseCase,
		updateOfflineActionUseCase,
		deleteOfflineActionUseCase,
	)
	showroomItemController := controller.NewShowroomItemController(
		createShowroomItemUseCase,
		listShowroomItemsUseCase,
		getShowroomItemUseCase,
		updateShowroomItemUseCase,
		deleteShowroomItemUseCase,
	)
	repMarketingActionController := controller.NewRepMarketingActionController(
		createRepMarketingActionUseCase,
		listRepMarketingActionsUseCase,
		getRepMarketingActionUseCase,
		updateRepMarketingActionUseCase,
		deleteRepMarketingActionUseCase,
	)

	budgetController := controller.NewBudgetController(
		createBudgetItemUseCase,
		listBudgetItemsUseCase,
		getBudgetItemUseCase,
		updateBudgetItemUseCase,
		deleteBudgetItemUseCase,
		batchCreateBudgetItemsUseCase,
		getBudgetSummaryUseCase,
		budgetGateway,
	)

	representativeController := controller.NewRepresentativeController(
		createRepresentativeUseCase,
		getRepresentativeUseCase,
		updateRepresentativeUseCase,
		deleteRepresentativeUseCase,
		listRepresentativesUseCase,
		getRepresentativeStatsUseCase,
		getRepresentativeProfileUseCase,
		getAllRepresentativeProfilesUseCase,
		representativeGateway,
	)

	representativeMonthlyGoalController := controller.NewRepresentativeMonthlyGoalController(
		createRepresentativeMonthlyGoalUseCase,
		getRepresentativeMonthlyGoalUseCase,
		updateRepresentativeMonthlyGoalUseCase,
		deleteRepresentativeMonthlyGoalUseCase,
		listRepresentativeMonthlyGoalsUseCase,
		getRepresentativeGoalsTableDataUseCase,
		representativeMonthlyGoalGateway,
	)

	// Brand Controller
	brandController := controller.NewBrandController(
		createBrandUseCase,
		listBrandsUseCase,
		deleteBrandUseCase,
	)

	// Gift Controllers
	giftItemController := controller.NewGiftItemController(
		createGiftItemUseCase,
		listGiftItemsUseCase,
		getGiftItemUseCase,
		updateGiftItemUseCase,
		deleteGiftItemUseCase,
	)

	giftTransactionController := controller.NewGiftTransactionController(
		createGiftTransactionUseCase,
		listGiftTransactionsUseCase,
		getGiftTransactionUseCase,
		updateGiftTransactionUseCase,
		deleteGiftTransactionUseCase,
	)

	log.Println("✅ Gift controllers initialized")
	log.Println("✅ Controllers initialized")

	// 6. Middlewares (dependem de services)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	corsMiddleware := middleware.NewCorsMiddleware()
	log.Println("✅ Middlewares initialized")

	// 7. Retornar Container
	return &Container{
		DB:                                     db,
		HasherService:                          hasherService,
		JwtService:                             jwtService,
		UserGateway:                            userGateway,
		KpiGateway:                             kpiGateway,
		MonthlyDataGateway:                     monthlyDataGateway,
		TaskGateway:                            taskGateway,
		SubtaskGateway:                         subtaskGateway,
		CommentGateway:                         commentGateway,
		NotificationGateway:                    notificationGateway,
		CalendarPostGateway:                    calendarPostGateway,
		PdvPostGateway:                         pdvPostGateway,
		RecurrentPdvGateway:                    recurrentPdvGateway,
		UserSeeder:                             userSeeder,
		KpiSeeder:                              kpiSeeder,
		SocialBenchmarkingSeeder:               socialBenchmarkingSeeder,
		CooperativeSeeder:                      cooperativeSeeder,
		SocialBenchmarkingGateway:              socialBenchmarkingGateway,
		SocialPostGateway:                      socialPostGateway,
		SocialDailyAggregationGateway:          socialDailyAggregationGateway,
		OfflineActionGateway:                   offlineActionGateway,
		ShowroomItemGateway:                    showroomItemGateway,
		RepMarketingActionGateway:              repMarketingActionGateway,
		GiftItemGateway:                        giftItemGateway,
		GiftTransactionGateway:                 giftTransactionGateway,
		AccountPayableGateway:                  accountPayableGateway,
		BudgetGateway:                          budgetGateway,
		RepresentativeGateway:                  representativeGateway,
		RepresentativeStatsGateway:             representativeStatsGateway,
		RepresentativeMonthlyGoalGateway:       representativeMonthlyGoalGateway,
		BrandGateway:                           brandGateway,
		GiftSeeder:                             giftSeeder,
		BudgetSeeder:                           budgetSeeder,
		RepresentativeMonthlyGoalSeeder:        representativeMonthlyGoalSeeder,
		LoginUseCase:                           loginUseCase,
		CreateKpiUseCase:                       createKpiUseCase,
		GetKpiUseCase:                          getKpiUseCase,
		ListKpisUseCase:                        listKpisUseCase,
		GetKpisBySlugsUseCase:                  getKpisBySlugsUseCase,
		UpdateKpiUseCase:                       updateKpiUseCase,
		DeleteKpiUseCase:                       deleteKpiUseCase,
		UpdateMonthlyDataUseCase:               updateMonthlyDataUseCase,
		CreateTaskUseCase:                      createTaskUseCase,
		UpdateTaskUseCase:                      updateTaskUseCase,
		DeleteTaskUseCase:                      deleteTaskUseCase,
		GetTaskUseCase:                         getTaskUseCase,
		ListTasksUseCase:                       listTasksUseCase,
		CreateSubtaskUseCase:                   createSubtaskUseCase,
		UpdateSubtaskUseCase:                   updateSubtaskUseCase,
		DeleteSubtaskUseCase:                   deleteSubtaskUseCase,
		GetSubtaskUseCase:                      getSubtaskUseCase,
		ListSubtasksUseCase:                    listSubtasksUseCase,
		CreateCommentUseCase:                   createCommentUseCase,
		UpdateCommentUseCase:                   updateCommentUseCase,
		DeleteCommentUseCase:                   deleteCommentUseCase,
		GetCommentUseCase:                      getCommentUseCase,
		ListCommentsUseCase:                    listCommentsUseCase,
		CreateNotificationUseCase:              createNotificationUseCase,
		UpdateNotificationUseCase:              updateNotificationUseCase,
		DeleteNotificationUseCase:              deleteNotificationUseCase,
		GetNotificationUseCase:                 getNotificationUseCase,
		ListNotificationsUseCase:               listNotificationsUseCase,
		MarkAsReadNotificationUseCase:          markAsReadNotificationUseCase,
		MarkAllAsReadNotificationsUseCase:      markAllAsReadNotificationsUseCase,
		DeleteNotificationsByTaskIDUseCase:     deleteNotificationsByTaskIDUseCase,
		ListUsersUseCase:                       listUsersUseCase,
		CreateCalendarPostUseCase:              createCalendarPostUseCase,
		GetCalendarPostUseCase:                 getCalendarPostUseCase,
		UpdateCalendarPostUseCase:              updateCalendarPostUseCase,
		UpdateCalendarPostStatusUseCase:        updateCalendarPostStatusUseCase,
		ConfirmCalendarPostPublishingUseCase:   confirmCalendarPostPublishingUseCase,
		DeleteCalendarPostUseCase:              deleteCalendarPostUseCase,
		ListCalendarPostsUseCase:               listCalendarPostsUseCase,
		AuthController:                         authController,
		KpiController:                          kpiController,
		TaskController:                         taskController,
		SubtaskController:                      subtaskController,
		CommentController:                      commentController,
		NotificationController:                 notificationController,
		UserController:                         userController,
		CalendarPostController:                 calendarPostController,
		PdvController:                          pdvController,
		SocialController:                       socialController,
		SocialPostController:                   socialPostController,
		OfflineActionController:                offlineActionController,
		ShowroomItemController:                 showroomItemController,
		RepMarketingActionController:           repMarketingActionController,
		GiftItemController:                     giftItemController,
		GiftTransactionController:              giftTransactionController,
		CreateSocialBenchmarkingUseCase:        createSocialBenchmarkingUseCase,
		ListSocialBenchmarkingsUseCase:         listSocialBenchmarkingsUseCase,
		GetSocialBenchmarkingUseCase:           getSocialBenchmarkingUseCase,
		UpdateSocialBenchmarkingUseCase:        updateSocialBenchmarkingUseCase,
		DeleteSocialBenchmarkingUseCase:        deleteSocialBenchmarkingUseCase,
		CreateSocialPostUseCase:                createSocialPostUseCase,
		GetSocialPostUseCase:                   getSocialPostUseCase,
		ListSocialPostsUseCase:                 listSocialPostsUseCase,
		UpdateSocialPostUseCase:                updateSocialPostUseCase,
		DeleteSocialPostUseCase:                deleteSocialPostUseCase,
		RecalculateDailyAggregationsUseCase:    recalculateDailyAggregationsUseCase,
		ListSocialDailyAggregationsUseCase:     listSocialDailyAggregationsUseCase,
		GetSocialDailyAggregationUseCase:       getSocialDailyAggregationUseCase,
		CreateOfflineActionUseCase:             createOfflineActionUseCase,
		ListOfflineActionsUseCase:              listOfflineActionsUseCase,
		GetOfflineActionUseCase:                getOfflineActionUseCase,
		UpdateOfflineActionUseCase:             updateOfflineActionUseCase,
		DeleteOfflineActionUseCase:             deleteOfflineActionUseCase,
		CreateShowroomItemUseCase:              createShowroomItemUseCase,
		ListShowroomItemsUseCase:               listShowroomItemsUseCase,
		GetShowroomItemUseCase:                 getShowroomItemUseCase,
		UpdateShowroomItemUseCase:              updateShowroomItemUseCase,
		DeleteShowroomItemUseCase:              deleteShowroomItemUseCase,
		CreateRepMarketingActionUseCase:        createRepMarketingActionUseCase,
		ListRepMarketingActionsUseCase:         listRepMarketingActionsUseCase,
		GetRepMarketingActionUseCase:           getRepMarketingActionUseCase,
		UpdateRepMarketingActionUseCase:        updateRepMarketingActionUseCase,
		DeleteRepMarketingActionUseCase:        deleteRepMarketingActionUseCase,
		CreateGiftItemUseCase:                  createGiftItemUseCase,
		GetGiftItemUseCase:                     getGiftItemUseCase,
		ListGiftItemsUseCase:                   listGiftItemsUseCase,
		UpdateGiftItemUseCase:                  updateGiftItemUseCase,
		DeleteGiftItemUseCase:                  deleteGiftItemUseCase,
		CreateGiftTransactionUseCase:           createGiftTransactionUseCase,
		GetGiftTransactionUseCase:              getGiftTransactionUseCase,
		ListGiftTransactionsUseCase:            listGiftTransactionsUseCase,
		UpdateGiftTransactionUseCase:           updateGiftTransactionUseCase,
		DeleteGiftTransactionUseCase:           deleteGiftTransactionUseCase,
		CreateAccountPayableUseCase:            createAccountPayableUseCase,
		ListAccountsPayableUseCase:             listAccountsPayableUseCase,
		GetAccountPayableUseCase:               getAccountPayableUseCase,
		UpdateAccountPayableUseCase:            updateAccountPayableUseCase,
		DeleteAccountPayableUseCase:            deleteAccountPayableUseCase,
		ToggleNFUseCase:                        toggleNFUseCase,
		ToggleBoletoUseCase:                    toggleBoletoUseCase,
		SendToFinanceUseCase:                   sendToFinanceUseCase,
		CreateBudgetItemUseCase:                createBudgetItemUseCase,
		ListBudgetItemsUseCase:                 listBudgetItemsUseCase,
		GetBudgetItemUseCase:                   getBudgetItemUseCase,
		UpdateBudgetItemUseCase:                updateBudgetItemUseCase,
		DeleteBudgetItemUseCase:                deleteBudgetItemUseCase,
		BatchCreateBudgetItemsUseCase:          batchCreateBudgetItemsUseCase,
		GetBudgetSummaryUseCase:                getBudgetSummaryUseCase,
		BudgetController:                       budgetController,
		CreateRepresentativeUseCase:            createRepresentativeUseCase,
		GetRepresentativeUseCase:               getRepresentativeUseCase,
		UpdateRepresentativeUseCase:            updateRepresentativeUseCase,
		DeleteRepresentativeUseCase:            deleteRepresentativeUseCase,
		ListRepresentativesUseCase:             listRepresentativesUseCase,
		GetRepresentativeStatsUseCase:          getRepresentativeStatsUseCase,
		GetRepresentativeProfileUseCase:        getRepresentativeProfileUseCase,
		GetAllRepresentativeProfilesUseCase:    getAllRepresentativeProfilesUseCase,
		RepresentativeController:               representativeController,
		RepresentativeMonthlyGoalController:    representativeMonthlyGoalController,
		CreateRepresentativeMonthlyGoalUseCase: createRepresentativeMonthlyGoalUseCase,
		GetRepresentativeMonthlyGoalUseCase:    getRepresentativeMonthlyGoalUseCase,
		UpdateRepresentativeMonthlyGoalUseCase: updateRepresentativeMonthlyGoalUseCase,
		DeleteRepresentativeMonthlyGoalUseCase: deleteRepresentativeMonthlyGoalUseCase,
		ListRepresentativeMonthlyGoalsUseCase:  listRepresentativeMonthlyGoalsUseCase,
		GetRepresentativeGoalsTableDataUseCase: getRepresentativeGoalsTableDataUseCase,
		CreateBrandUseCase:                     createBrandUseCase,
		ListBrandsUseCase:                      listBrandsUseCase,
		DeleteBrandUseCase:                     deleteBrandUseCase,
		BrandController:                        brandController,
		AuthMiddleware:                         authMiddleware,
		CorsMiddleware:                         corsMiddleware,
	}, nil
}

// initDatabase conecta ao PostgreSQL usando GORM e configura pool de conexões
func initDatabase(cfg *Config) (*gorm.DB, error) {
	// 1. Construir DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	// 2. Conectar usando driver PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 3. Configurar pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.Database.ConnMaxIdleTime) * time.Second)

	// 4. Testar conexão
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// Close fecha conexões e libera recursos quando a aplicação encerrar
func (c *Container) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// GetControllers retorna struct para usar nas rotas
func (c *Container) GetControllers() *routes.Controllers {
	return &routes.Controllers{
		AuthController:                      c.AuthController,
		KpiController:                       c.KpiController,
		TaskController:                      c.TaskController,
		SubtaskController:                   c.SubtaskController,
		CommentController:                   c.CommentController,
		NotificationController:              c.NotificationController,
		UserController:                      c.UserController,
		CalendarPostController:              c.CalendarPostController,
		PdvController:                       c.PdvController,
		SocialController:                    c.SocialController,
		SocialPostController:                c.SocialPostController,
		OfflineActionController:             c.OfflineActionController,
		ShowroomItemController:              c.ShowroomItemController,
		RepMarketingActionController:        c.RepMarketingActionController,
		GiftItemController:                  c.GiftItemController,
		GiftTransactionController:           c.GiftTransactionController,
		AccountPayableController:            c.AccountPayableController,
		BudgetController:                    c.BudgetController,
		RepresentativeController:            c.RepresentativeController,
		RepresentativeMonthlyGoalController: c.RepresentativeMonthlyGoalController,
		BrandController:                     c.BrandController,
	}
}

// GetMiddlewares retorna struct para usar nas rotas
func (c *Container) GetMiddlewares() *routes.Middlewares {
	return &routes.Middlewares{
		AuthMiddleware: c.AuthMiddleware,
		CorsMiddleware: c.CorsMiddleware,
	}
}
