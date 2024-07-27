package lasagna

// PreparationTime returns the estimatedpreparation time of lasagna based on layers count and the preparation time of each layer
func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		return len(layers) * 2
	}

	return len(layers) * preparationTime
}

// Quantities returns the computed amounts of noodles and sauce needed to prepare lasagna
func Quantities(layers []string) (int, float64) {
	noodleGrams := 0
	sauceLiters := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodleGrams += 50
		case "sauce":
			sauceLiters += 0.2
		}
	}

	return noodleGrams, sauceLiters
}

// AddSecretIngredient returns a modified myList with a secret ingredient from friendsList
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe returns calculated amounts for different numbers of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i, quantity := range quantities {
		scaledQuantities[i] = (quantity / 2) * float64(portions)
	}

	return scaledQuantities
}
