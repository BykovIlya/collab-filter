package models

func CreateInputProduct(products []Product) [][]float64 {
	input := make ([][]float64, 0)
	for i := 0; i < len(products) ; i++ {
		input = append(input, []float64 {
			float64(products[i].cathegory), products[i].price })
	}
	return input
}

func CreateTargetProduct(persons []Person, product Product, visitors []Visitor) [][]float64 {
	output := make ([][]float64, 0)
	for i := 0; i < len(persons); i++ {
		buf := FindProductInPerson(i, product.id, visitors)
		output = append(output, []float64 {
			convertBoolToFloat64(buf), convertBoolToFloat64(!buf)})
	}
	return output
}

/*
func CreateNeuralNetworkProduct(persons []Person, product Product, visitors []Visitor) {
	nn := gonn.DefaultNetwork(2,10,2, false)
	input := CreateInputPerson(persons, product)
	target := CreateTargetPerson(persons, product, visitors)
	nn.Train(input, target, 10000)
	gonn.DumpNN("gonnPerson", nn)
}*/