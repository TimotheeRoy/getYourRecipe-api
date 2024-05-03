package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//db connection
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Erreur lors de la connection à la db", err)
	}
	
	//route init
	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		data, err := getRecipeList(db)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})

	r.GET("/recipe/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		data, err := getIngredientsForRecipe(db, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})


	defer db.Close()
	
	r.Run(":8080")
}