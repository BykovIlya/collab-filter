package models

import (
	"fmt"
	"time"
	"github.com/fxsjy/gonn/gonn"
)

func convertBoolToFloat64 (in bool) float64 {
	out := 0.0
	if (in) {
		out = 1.0
	}
	return out
}

func CreateInputPerson(persons []Person, product Product, visitors  [] Visitor) ([][]float64, [][]float64) {
	input := make ([][]float64, 0)
	output := make ([][]float64, 0)
	for i := 0; i < len(persons) ; i++ {
			input = append(input, []float64 {
				convertBoolToFloat64(persons[i].gender), float64(persons[i].age), float64(product.cathegory), product.price})
		buf := FindProductInPerson(i, product.id, visitors)
		output = append(output, []float64 {
			convertBoolToFloat64(buf), convertBoolToFloat64(!buf)})
	}
	return input, output
}

/*
func CreateTargetPerson(persons []Person, product Product, visitors []Visitor) [][]float64 {
	output := make ([][]float64, 0)
	for i := 0; i < len(persons); i++ {
		buf := FindProductInPerson(i, product.id, visitors)
		output = append(output, []float64 {
			convertBoolToFloat64(buf), convertBoolToFloat64(!buf)})
	}
	return output
}*/



var input, target [][]float64

func quickAppendIn(el []float64) {
	input = append(input,el)
}

func quickAppendOut(el []float64) {
	target = append(target, el)
}

func quickAppend(j int, in, o [][]float64) {
	//for i := j; i < j + 10; i++ {
		 quickAppendIn(in[j])
		 quickAppendOut(o[j])
//	}
}

/*func write(arr [][]float64, filename string) error{
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range arr {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}*/
/*func read(filename string) ([][] float64){
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines [][] float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	return lines
}*/

func CreateNeuralNetworkPerson(persons []Person, product []Product, visitors []Visitor) {
	nn := gonn.DefaultNetwork(4,10,2, false)
/*	fileIn, err := os.Open("in.csv")
	if err != nil {
		return
	}
	fileOut, err := os.Open("target.csv")
	if err != nil {
		return
	}*/
	if (!ImportInputNNToDB(input) && !ImportTargetNNToDB(target)) {

	} else {
		input = make([][]float64, 0)
		target = make([][] float64, 0)
		for i := 0; i < len(product); i++ {
			start := time.Now()
			in, o := CreateInputPerson(persons, product[i], visitors)
			for j := 0; j < len(persons); j++ {
				//fmt.Println(i," ")
				//input = append(input,in[j])
				//target = append(target,o[j])
				//quickAppendIn(in[j])
				//quickAppendOut(o[j])
				quickAppend(j, in, o)
				// fmt.Println(input)
				// fmt.Println(target)
			}
			t := time.Now()
			fmt.Println(i, ":", t.Sub(start))
		}
		/*write(input, "in.csv")
		write(target, "target.csv")*/

	}
	nn.Train(input, target, 1000)
	gonn.DumpNN("gonnPerson", nn)
}

func GetResult(output [] float64) int64{
	max := -99999.0
	pos := -1

	for i, val := range output {
		if (val > max) {
			max = val
			pos = i
		}
	}

	switch pos {
	case 0: return 1
	case 1: return 0

	}
	return -1
}
