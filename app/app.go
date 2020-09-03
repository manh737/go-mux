package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/manh737/go-mux/app/handler"
	"github.com/manh737/go-mux/app/mongodb"
	"github.com/manh737/go-mux/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// App has the mongo database and router instances
type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

// ConfigAndRunApp will create and initialize App structure. App factory function.
func ConfigAndRunApp(config *config.Config) {
	app := new(App)
	app.Initialize(config)
	app.Run(config.ServerHost)
}

// Initialize initialize the app with
func (app *App) Initialize(config *config.Config) {
	app.DB = mongodb.InitialConnection("id-card-ocr", config.MongoURI())

	app.Router = mux.NewRouter()
	app.UseMiddleware(handler.JSONContentTypeMiddleware)
	app.setRouters()
}

// UseMiddleware will add global middleware in router
func (app *App) UseMiddleware(middleware mux.MiddlewareFunc) {
	app.Router.Use(middleware)
}

// SetupRouters will register routes in router
func (app *App) setRouters() {
	app.Post("/", app.handleRequest(handler.UploadImage))
}

// Post will register Post method for an endpoint
func (app *App) Post(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("POST").Queries(queries...)
}

// RequestHandlerFunction is a custome type that help us to pass db arg to all endpoints
type RequestHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}

// Run will start the http server on host that you pass in. host:<ip:port>
func (app *App) Run(host string) {
	// use signals for shutdown server gracefully.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		log.Fatal(http.ListenAndServe(host, app.Router))
	}()
	log.Printf("Server is listning on http://%s\n", host)
	sig := <-sigs
	log.Println("Signal: ", sig)

	log.Println("Stoping MongoDB Connection...")
	app.DB.Client().Disconnect(context.Background())
}