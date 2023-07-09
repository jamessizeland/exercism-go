package lasagnaMaster

func PreparationTime(layers []string, av_prep_time int) int {
	if av_prep_time <= 0 {
		av_prep_time = 2
	}
	return len(layers) * av_prep_time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles, sauces int
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 1
		case "sauce":
			sauces += 1
		}
	}
	return noodles * 50, float64(sauces) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64
	for _, item := range quantities {
		scaledQuantities = append(scaledQuantities, (item/2)*float64(portions))
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
