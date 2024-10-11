package main

// func main() {
// 	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/stack?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
// 	if err != nil {
// 		log.Error().Err(err)
// 		return
// 	}
// 	pip := pipline.NewCreate(db)
//
// 	sp := spiders.NewStarkSpider(pip)
// 	// sp := spiders.NewNameCode(pip)
// 	//
// 	// sp.ListSh()
// 	sp.Start()
// }

func main() {
	cmd := app.NewExportServerCommand()
	if err := cmd.Execute(); err != nil {
		// zap.S().Fatal(err)
		logging.L.Fatal().Err(err).Send()
	}
}
