package main

type Ingredients struct {
	Id   int
	Name string
}

type Recipe struct {
	Id int
	Name string
	Description string
	Instructions []Instruction
}

type Instruction struct {
	Instruction string
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