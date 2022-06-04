package main

import (
	"context"

	"github.com/KaranAhlawat/ddgf/internal/app/service"
	repo "github.com/KaranAhlawat/ddgf/internal/repo/postgresql"
	"github.com/KaranAhlawat/ddgf/internal/server"
)

func main() {
	ctx := context.Background()

	apiConf := server.NewAPIConfig()

	dbConf, err := repo.InitPostgresConn()
	if err != nil {
		println("Unable to init postgres connection")
		println(err.Error())
		return
	}
	defer dbConf.Close()

	pageRepo := repo.NewPage(dbConf)
	adviceRepo := repo.NewAdvice(dbConf)
	tagRepo := repo.NewTag(dbConf)

	pageService := service.NewPage(ctx, pageRepo)
	adviceService := service.NewAdvice(ctx, adviceRepo)
	tagService := service.NewTag(ctx, tagRepo)

	app := server.NewServer(apiConf)
	app.Setup(pageService, adviceService, tagService)
	app.Run()
}
