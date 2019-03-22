package main

import(
	"fmt"
	"math"	
)


/*
Lo que se le debe alimentar a mis funciones es una cadena con la secuecia de libros que compro el usuario
por ejemplo 123451234123
De ahi saco las permutaciones posibles de eso
Luego formo los subconjutos
Y luego saco el precio de cada uno
*/


type Libro struct{
	isbn string
	nombre string
	editora string
	copiasDisponibles int
}

var possiblePermutations [] string
var bestPrice float64 //valor maximo que puede tener un int :v, no me dejaba setearlo de un solo

//Funcion para encontrar todas las permutaciones posibles
func findPermutations(original string, acum string){
	if(len(original) == 0){
		//Agrega cada permutacion posible al arreglo global
		possiblePermutations = append(possiblePermutations,acum)
		//fmt.Print(acum + " ")
		return
	}
	for i:= 0; i < len(original); i++ {
		ch := string(original[i])
		res := original[0:i]+original[i+1:]
		findPermutations(res,acum+ch)
	}
}

//Funcion para determinar el mejor precio de todas las permutaciones encontradas
func findBestPrice(){
	for i:= 0; i < len(possiblePermutations); i++{
		conjActual := possiblePermutations[i]
		var conjunto [] string
		acum := string(conjActual[0])

		for j:= 1;  j < len(conjActual)-1; j++{
			if(exists(string(conjActual[j]),acum)){
				conjunto = append(conjunto, acum)
				acum = 	string(conjActual[j])
			}else{
				acum += string(conjActual[j])
			}
		}
		if(exists(string(conjActual[len(conjActual)-1]), acum)){
			conjunto = append(conjunto,acum)
			conjunto = append(conjunto,string(conjActual[len(conjActual)-1]))
		}else{
			acum += string(conjActual[len(conjActual)-1])
			conjunto = append(conjunto,acum)
		}
		tempo:= findPrice(conjunto)
		if(tempo < bestPrice){
			//fmt.Println("El mejor precio se encontro con la permutación ",i)
			bestPrice = tempo
		}
		/*for k:= 0; k < len(conjunto); k++{
			fmt.Println("conjunto",i,",",k,")",conjunto[k],"-")
		}*/
	}
}

//Funcion para encontrar el precio de UNA permutacion especifica
func findPrice(set[] string) float64{
	retorno:=0.0
	for i:= 0; i < len(set); i++{
		actual := string(set[i])
		tempo1 := float64(len(actual))
		retorno += (tempo1*800.00)*findDisc(actual)
	}
	return retorno
}

//Funcion auxiliar para el porcentaje de descuento
func findDisc(set string) float64 {
	switch len(set){
		case 1:
			return 1
		case 2:
			return 0.95
		case 3:
			return 0.9
		case 4:
			return 0.8
		case 5: 
			return 0.75
	}
	return 1
}

//Funcion para verificar si un elemento está presente en una cadena
func exists(ind string, str string) bool{
	for i:= 0; i < len(str); i++{
		if(ind == string(str[i])){
			return true
		}
	}
	return false
}

func main(){
	bestPrice = math.MaxFloat32
	//libro:= Libro{isbn: "1111", nombre: "Hols", editora: "Guaymuras", copiasDisponibles: 12}
	//fmt.Println(libro.isbn)
	//fmt.Println(libro.nombre)
	//fmt.Println(libro.editora)
	//fmt.Println(libro.copiasDisponibles)
	s:= "11223345"
	findPermutations(s,"")
	//fmt.Println("Total number of permutations: ", len(possiblePermutations))
	findBestPrice()
	fmt.Println("")
	fmt.Println("besto prizo: ",bestPrice)
	//fmt.Println(len(possiblePermutations))
	//fmt.Println("go es estúpido :v")
	/*for i:=0; i < len(possiblePermutations); i++{
		fmt.Println(i,") ", possiblePermutations[i])
	}*/
}