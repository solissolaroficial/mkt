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
	"github.com/seu-usuario/solis-backend/application/usecase/kpis"
	"github.com/seu-usuario/solis-backend/application/usecase/tasks"
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
	UserGateway         gateway.UserGateway
	KpiGateway          gateway.KpiGateway
	MonthlyDataGateway  gateway.MonthlyDataGateway
	TaskGateway         gateway.TaskGateway
	SubtaskGateway      gateway.SubtaskGateway
	CommentGateway      gateway.CommentGateway
	NotificationGateway gateway.NotificationGateway

	// Seeders
	UserSeeder *seeders.UserSeeder
	KpiSeeder  *seeders.KpiSeeder

	// Use Cases - Auth
	LoginUseCase *auth.LoginUseCase

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
	AuthController         *controller.AuthController
	KpiController          *controller.KpiController
	TaskController         *controller.TaskController
	SubtaskController      *controller.SubtaskController
	CommentController      *controller.CommentController
	NotificationController *controller.NotificationController

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
	log.Println("✅ Gateways initialized")

	// 3.1 Seeders (dependem de gateways e services)
	userSeeder := seeders.NewUserSeeder(userGateway, hasherService)
	kpiSeeder := seeders.NewKpiSeeder(kpiGateway, monthlyDataGateway)
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
	createSubtaskUseCase := tasks.NewCreateSubtaskUseCase(subtaskGateway)
	updateSubtaskUseCase := tasks.NewUpdateSubtaskUseCase(subtaskGateway)
	deleteSubtaskUseCase := tasks.NewDeleteSubtaskUseCase(subtaskGateway)
	getSubtaskUseCase := tasks.NewGetSubtaskUseCase(subtaskGateway)
	listSubtasksUseCase := tasks.NewListSubtasksUseCase(subtaskGateway)
	createCommentUseCase := tasks.NewCreateCommentUseCase(commentGateway)
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

	log.Println("✅ Use cases initialized")

	//5. Controllers (dependem de use cases)
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
	log.Println("✅ Controllers initialized")

	// 6. Middlewares (dependem de services)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	corsMiddleware := middleware.NewCorsMiddleware()
	log.Println("✅ Middlewares initialized")

	// 7. Retornar Container
	return &Container{
		DB:                                 db,
		HasherService:                      hasherService,
		JwtService:                         jwtService,
		UserGateway:                        userGateway,
		KpiGateway:                         kpiGateway,
		MonthlyDataGateway:                 monthlyDataGateway,
		TaskGateway:                        taskGateway,
		SubtaskGateway:                     subtaskGateway,
		CommentGateway:                     commentGateway,
		NotificationGateway:                notificationGateway,
		UserSeeder:                         userSeeder,
		KpiSeeder:                          kpiSeeder,
		LoginUseCase:                       loginUseCase,
		CreateKpiUseCase:                   createKpiUseCase,
		GetKpiUseCase:                      getKpiUseCase,
		ListKpisUseCase:                    listKpisUseCase,
		GetKpisBySlugsUseCase:              getKpisBySlugsUseCase,
		UpdateKpiUseCase:                   updateKpiUseCase,
		DeleteKpiUseCase:                   deleteKpiUseCase,
		UpdateMonthlyDataUseCase:           updateMonthlyDataUseCase,
		CreateTaskUseCase:                  createTaskUseCase,
		UpdateTaskUseCase:                  updateTaskUseCase,
		DeleteTaskUseCase:                  deleteTaskUseCase,
		GetTaskUseCase:                     getTaskUseCase,
		ListTasksUseCase:                   listTasksUseCase,
		CreateSubtaskUseCase:               createSubtaskUseCase,
		UpdateSubtaskUseCase:               updateSubtaskUseCase,
		DeleteSubtaskUseCase:               deleteSubtaskUseCase,
		GetSubtaskUseCase:                  getSubtaskUseCase,
		ListSubtasksUseCase:                listSubtasksUseCase,
		CreateCommentUseCase:               createCommentUseCase,
		UpdateCommentUseCase:               updateCommentUseCase,
		DeleteCommentUseCase:               deleteCommentUseCase,
		GetCommentUseCase:                  getCommentUseCase,
		ListCommentsUseCase:                listCommentsUseCase,
		CreateNotificationUseCase:          createNotificationUseCase,
		UpdateNotificationUseCase:          updateNotificationUseCase,
		DeleteNotificationUseCase:          deleteNotificationUseCase,
		GetNotificationUseCase:             getNotificationUseCase,
		ListNotificationsUseCase:           listNotificationsUseCase,
		MarkAsReadNotificationUseCase:      markAsReadNotificationUseCase,
		MarkAllAsReadNotificationsUseCase:  markAllAsReadNotificationsUseCase,
		DeleteNotificationsByTaskIDUseCase: deleteNotificationsByTaskIDUseCase,
		AuthController:                     authController,
		KpiController:                      kpiController,
		TaskController:                     taskController,
		SubtaskController:                  subtaskController,
		CommentController:                  commentController,
		NotificationController:             notificationController,
		AuthMiddleware:                     authMiddleware,
		CorsMiddleware:                     corsMiddleware,
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
		AuthController:         c.AuthController,
		KpiController:          c.KpiController,
		TaskController:         c.TaskController,
		SubtaskController:      c.SubtaskController,
		CommentController:      c.CommentController,
		NotificationController: c.NotificationController,
	}
}

// GetMiddlewares retorna struct para usar nas rotas
func (c *Container) GetMiddlewares() *routes.Middlewares {
	return &routes.Middlewares{
		AuthMiddleware: c.AuthMiddleware,
		CorsMiddleware: c.CorsMiddleware,
	}
}
