package Routes

import (
	"EmployeeService/Controller"
	_ "EmployeeService/Library/Swagger/docs"
	"EmployeeService/Routes/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	Controller Controller.ControllerInterface
	Middleware middleware.MiddlewareInterface
	Chi        *chi.Mux
}

func (app *Routes) CollectRoutes() *chi.Mux {
	logger := httplog.NewLogger("EmployeeService", httplog.Options{
		JSON: true,
	})

	appRoute := app.Chi
	appRoute.Use(chiMiddleware.RequestID)
	appRoute.Use(chiMiddleware.RealIP)
	appRoute.Use(chiMiddleware.RedirectSlashes)
	appRoute.Use(chiMiddleware.Recoverer)
	appRoute.Use(middleware.Logger(&logger))

	appRoute.Route("/employee", func(appRoute chi.Router) {
		appRoute.Group(func(appRoute chi.Router) {
			appRoute.Use(app.Middleware.ApiKey())
			appRoute.Get("/check-mobile-phone-number/{mobilePhoneNumber}", app.Controller.CheckExistMobilePhoneNumber)
			appRoute.Get("/check-email/{email}", app.Controller.CheckExistEmail)
			appRoute.Get("/check-employee-id/{id}", app.Controller.CheckExistEmployeeID)
			//Leave Balance
			appRoute.Get("/reset-leave-balance", app.Controller.LeaveBalanceReset)
			appRoute.Get("/count-employees", app.Controller.CountEmployees)
		})
		appRoute.Get("/{id}/read-photo", app.Controller.EmployeeReadPhoto)
		appRoute.Group(func(appRoute chi.Router) {
			appRoute.Use(app.Middleware.UserAuth())
			appRoute.Get("/employee", app.Controller.EmployeeSearch)
			appRoute.Route("/employees", func(appRoute chi.Router) {
				appRoute.Get("/", app.Controller.EmployeeBrowse)
				appRoute.Post("/", app.Controller.EmployeeAdded)
				appRoute.Route("/{id}", func(appRoute chi.Router) {
					appRoute.Get("/head", app.Controller.EmployeeGetHead)
					appRoute.Get("/", app.Controller.EmployeeBrowseDetail)
					appRoute.Put("/", app.Controller.EmployeePersonalInfoUpdated)
					appRoute.Delete("/", app.Controller.EmployeeDeleted)
				})
				appRoute.Get("/{id}/subordinates", app.Controller.SubordinatesData)
				appRoute.Get("/{id_parent}/subordinates/{id}", app.Controller.DetailSubordinatesData)
				appRoute.Put("/{id}/workstatus", app.Controller.EmployeeWorkStatusUpdated)
				appRoute.Put("/{id}/fingerprint", app.Controller.EmployeeFingerUpdated)
				appRoute.Put("/{id}/machineid", app.Controller.EmployeeMachineIdUpdated)
				appRoute.Put("/{id}/faceid", app.Controller.EmployeeFaceIdUpdated)
				appRoute.Put("/{id}/resign", app.Controller.EmployeeResigned)
				appRoute.Post("/{id}/upload-photo", app.Controller.EmployeeUploadPhoto)
				appRoute.Get("/{id}/read-photo", app.Controller.EmployeeReadPhoto)
			})

			appRoute.Route("/leave-balance", func(appRoute chi.Router) {
				appRoute.Post("/", app.Controller.LeaveBalanceAdjusmentCreate)
				appRoute.Get("/", app.Controller.LeaveBalanceData)
				appRoute.Route("/{employeeId}", func(appRoute chi.Router) {
					appRoute.Get("/", app.Controller.LeaveBalanceDetail)
				})
				appRoute.Get("/{employeeId}/{year}/history", app.Controller.LeaveBalanceAdjustmentDetail)
			})

			appRoute.Route("/transfers", func(appRoute chi.Router) {
				appRoute.Post("/", app.Controller.TransferCreate)
				appRoute.Get("/", app.Controller.TransferData)
				appRoute.Route("/{id}", func(appRoute chi.Router) {
					appRoute.Get("/", app.Controller.TransferDetail)
				})
			})

			appRoute.Route("/dashboards", func(appRoute chi.Router) {
				appRoute.Route("/employee", func(appRoute chi.Router) {
					appRoute.Get("/new", app.Controller.NewEmployeeData)
					appRoute.Get("/religion/summary", app.Controller.TotalReligionSummary)
					appRoute.Get("/work-status/summary", app.Controller.TotalWorkStatusSummary)
					appRoute.Get("/work-status/will-expire/summary", app.Controller.TotalWillExpireEmployeeContract)
					appRoute.Get("/length-of-work/{numberOfYear}", app.Controller.TotalEmployeeByLengthOfWork)
				})
			})

			appRoute.Route("/leave-balance/policy", func(appRoute chi.Router) {
				appRoute.Get("/", app.Controller.SelectLeaveBalancePolicy)
				appRoute.Post("/", app.Controller.SaveLeaveBalancePolicy)
				appRoute.Put("/{id}", app.Controller.UpdateLeaveBalancePolicy)
			})
		})
		//sample controller function
		appRoute.Mount("/swagger", httpSwagger.WrapHandler)
		appRoute.Get("/ping", app.Controller.Ping)
		appRoute.Post("/dapr-state/{stateStore}", app.Controller.DaprState)
		appRoute.Get("/dapr-state/{stateStore}/{key}", app.Controller.DaprStateGet)
		appRoute.Delete("/dapr-state/{stateStore}/{key}", app.Controller.DaprStateDelete)
	})

	return appRoute
}
