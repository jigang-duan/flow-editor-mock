package main

import (
	"flow-editor-mock/datasource"
	"flow-editor-mock/repositories"
	"flow-editor-mock/services"
	"flow-editor-mock/web/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./web/views", ".html"))

	repo := repositories.NewTypeGroupRepository(datasource.TypeGroups)
	typeGroupService := services.NewTypeGroupService(repo)
	hero.Register(typeGroupService)

	crs := cors.New(cors.Options{
		AllowedOrigins:[]string{"*"},
		AllowCredentials:true,
	})

	flow := app.Party("/flow", crs).AllowMethods(iris.MethodOptions)
	{
		flow.PartyFunc("/dataflow", registerDataFlowRoutes)
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

func registerDataFlowRoutes(router iris.Party)  {
	router.Get("/processor-types", hero.Handler(routes.TypeGroups))
}