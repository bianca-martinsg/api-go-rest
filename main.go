package main

import (
	// Correct import path

	"github.com/bianca-martinsg/go-rest-api/models"
	"github.com/bianca-martinsg/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Albert Einstein", History: "Albert Einstein was a theoretical physicist who developed the theory of relativity, one of the two pillars of modern physics. His work is also known for its influence on the philosophy of science."},
		{Id: 2, Name: "Isaac Newton", History: "Sir Isaac Newton was an English mathematician, physicist, astronomer, and author who is widely recognized as one of the most influential scientists of all time, and a key figure in the scientific revolution."},
	}

	database.db.connectToDB() // Ensure the function name starts with an uppercase letter
	routes.HandleRequest()
}
