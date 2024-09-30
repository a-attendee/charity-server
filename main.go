package main

import (
	"log/slog"
	"os"

	"github.com/a-attendee/charity-server/app"
	"github.com/a-attendee/charity-server/database"
	"github.com/a-attendee/charity-server/internal/config"
	"github.com/a-attendee/charity-server/router"
)


func main() {

    // Init config //
    cfg := config.MustLoad()

    // Init logger //
    logger :=  slog.New(slog.NewTextHandler(os.Stdout, nil))

    // Init db //
    database.InitDB()

    // Init app instance //
    app := app.New(
        app.WithPort(3000),
        app.WithConfig(cfg),
        app.WithLogger(logger),
    )

    // Router init //
    router.Api(app.App)
   
    // Start server //
    if err := app.Listen(); err != nil {
        logger.Error("Server has been crashed: %v" + err.Error())
    }

    return
}
