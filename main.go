package main

import (
	"flow-editor-mock/datasource"
	"flow-editor-mock/repositories"
	"flow-editor-mock/repositories/nifi"
	"flow-editor-mock/services"
	"flow-editor-mock/web/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/logger"
	"gopkg.in/resty.v1"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		//Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)

	app.RegisterView(iris.HTML("./web/views", ".html"))

	client := resty.New()
	{
		client.SetDebug(true)
	}

	//repoTypeGroup := repositories.NewTypeGroupRepository(datasource.TypeGroups)
	repoTypeGroup := nifi.NewTypeGroupRepository(client)
	typeGroupService := services.NewTypeGroupService(repoTypeGroup)
	hero.Register(typeGroupService)

	repoProcessGroup := repositories.NewProcessGroupRepository(datasource.ProcessGroups)
	processGroupService := services.NewProcessGroupService(repoProcessGroup, repoTypeGroup)
	hero.Register(processGroupService)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "HEAD", "PUT", "DELETE"},
	})

	flow := app.Party("/flow", crs).AllowMethods(iris.MethodOptions)
	{
		flow.PartyFunc("/dataflow", registerDataFlowRoutes)
		flow.PartyFunc("/process-groups", registerProcessGroupRoutes)
	}

	return app
}

func main() {

	app := newApp()

	_ = app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func registerDataFlowRoutes(party iris.Party) {
	party.Get("/processor-types", hero.Handler(routes.TypeGroups))
	party.Get("/process-groups/{id}", hero.Handler(routes.ProcessGroups))
}

func registerProcessGroupRoutes(party iris.Party) {
	party.Post("/{gid:string}/processors", hero.Handler(routes.CreateProcessor))
	party.Put("/{gid:string}/processors", hero.Handler(routes.UpdateProcessors))
	party.Post("/{gid:string}/connections", hero.Handler(routes.CreateConnection))
	party.Post("/{gid:string}/process-groups", hero.Handler(routes.CreateProcessGroup))
	party.Post("/{gid:string}/snippet", hero.Handler(routes.CloneSnippet))
	party.Delete("/{gid:string}/snippet", hero.Handler(routes.DeleteSnippet))
	party.Put("/{gid:string}/snippet", hero.Handler(routes.UpdateSnippet))
	party.Put("/{gid:string}/ungroup", hero.Handler(routes.UngroupProcessGroup))
}
