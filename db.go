package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Ingredients struct {
	Id   int
	Name string
}

type Recipe struct {
	Id int
	Name string
	Description string
}

type Instruments struct {
	Id int
	Name string
}

type RecipeIngredients struct {
	Id int
	Name string
	Id_recipe int
	Id_ingredients int
	Quantity int 
	Unity string
}

func ConnectDB() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	
	DB_URI := os.Getenv("DB_URI")

	db, err := sqlx.Open("postgres", DB_URI)
	if err != nil {
		return nil, err
	}

	return db, nil

}


func getRecipeList(db *sqlx.DB)  ([]Recipe, error) {

	var recipe []Recipe

	query := `
	SELECT id, name, description FROM recipe`

	err := db.Select(&recipe, query)
	log.Print("getIngredients", err)
	if err != nil {
		return nil, err
	}

	return recipe, nil
	
}

func getIngredientsForRecipe(db *sqlx.DB, recipe_id string) ([]RecipeIngredients, error) {

	var ingredients []RecipeIngredients

	query := `
	SELECT ingredients.id, ingredients.name, recipe_ingredients.quantity, recipe_ingredients.unity
	FROM recipe_ingredients
	JOIN ingredients ON recipe_ingredients.id_ingredients = ingredients.id
	WHERE recipe_ingredients.id_recipe = ` + recipe_id

	err := db.Select(&ingredients, query)
	log.Print("getIngredientsForRecipe", err)
	if err != nil {
		return nil, err
	}

	return ingredients, nil

}


func getId(db *sqlx.DB, recipe string) (string, error) {
	var id string

	query := "SELECT id FROM recipe WHERE name = ?"

	err := db.Get(id, query, recipe)
	if err != nil {
		return "", err
	}
	return id, nil
}