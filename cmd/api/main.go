package main

import "github.com/Abdulqudri/fintech/internal/app"

func main() {
	app := app.NewApp()

	srv := app.Mount()
	srv.Run()
}