package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



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

// retourne une liste des recettes
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


// retourne la liste des ingrédients d'une recette
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

// retourne l'id d'une recette à partir de son nom 
func getId(db *sqlx.DB, recipe string) (string, error) {
	var id string

	query := "SELECT id FROM recipe WHERE name = ?"

	err := db.Get(id, query, recipe)
	if err != nil {
		return "", err
	}
	return id, nil
}


// retourne une recette qui correspond le plus à la liste d'ingredients en parametre 
func getRecipeFromIngredients(db *sqlx.DB, ingredients []string) ([]Recipe, error) {
	var recipes []Recipe 

	query, args, err := sqlx.In(`
	SELECT DISTINCT recipe.name, recipe.description FROM recipe
	JOIN recipe_ingredients ON recipe.id = recipe_ingredients.id_recipe
	JOIN ingredients ON recipe_ingredients.id_ingredients = ingredients.id
	WHERE ingredients.name IN (?)`, ingredients)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Get(recipes, query, args...)
	if err != nil  {
		return nil , err
	}
	return recipes, nil}