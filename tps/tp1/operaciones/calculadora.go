package operaciones

import (
	"fmt"
	"math"
	"strconv"

	TDA "tdas/pila"
)

// Constantes que representan los valores de error y operadores
const (
	OP_SUMA           = "+"    //Simbolo de la operacion de suma
	OP_RESTA          = "-"    //Simbolo de la operacion de resta
	OP_MULTIPLICACION = "*"    //Simbolo de la operacion de multiplicacion
	OP_DIVISION       = "/"    //Simbolo de la operacion de division
	OP_POTENCIA       = "^"    //Simbolo de la operacion de potencia
	OP_LOG            = "log"  //Simbolo de la operacion de logaritmo
	OP_RAIZ           = "sqrt" //Simbolo de la operacion de raiz cuadrara
	OP_TERNARIO       = "?"    //Simbolo de la operacion ternaria
)

// Estructura de una operacion
type operacion struct {
	simbolo string
	aridad  int
	operar  func(operadores []int64) (int64, error)
}

var operaciones = []operacion{
	{OP_SUMA, 2, operarSuma},
	{OP_RESTA, 2, operarResta},
	{OP_MULTIPLICACION, 2, operarMultiplicacion},
	{OP_DIVISION, 2, operarDivision},
	{OP_POTENCIA, 2, operarPotencia},
	{OP_LOG, 2, operarLog},
	{OP_RAIZ, 1, operarSqrt},
	{OP_TERNARIO, 3, operarTernario},
}

// CalcularOperacion procesa los tokens que representan numeros y operadores
// Devuelve el resultado de realizar distintas operaciones entre los tokens
func CalcularOperacion(tokens []string) (int64, error) {
	pila := TDA.CrearPilaDinamica[int64]()
	for _, token := range tokens {
		operador, valido := buscarOperacion(token)

		if valido {
			numeros, err := tomarElementos(pila, operador.aridad)
			if detectorError(pila, 0, err, false) {
				return 0, fmt.Errorf("ERROR")
			}

			resultado, err := operador.operar(numeros)
			if detectorError(pila, resultado, err, true) {
				return 0, fmt.Errorf("ERROR")
			}

		} else {
			numero, err := strconv.Atoi(token)
			if detectorError(pila, int64(numero), err, true) {
				return 0, fmt.Errorf("ERROR")
			}
		}
	}
	//Verifica que el resultado este apilado
	if pila.EstaVacia() {
		return 0, fmt.Errorf("ERROR")
	}
	resultadoOperar := pila.Desapilar()
	//La pila debe quedar vacia despues de desapilar el resultado
	if !(pila.EstaVacia()) {
		return 0, fmt.Errorf("ERROR")
	}
	return resultadoOperar, nil
}

// buscarOperacion busca la operacion por su simbolo en el arreglo de operaciones
func buscarOperacion(token string) (operacion, bool) {
	for _, operacion := range operaciones {
		if token == operacion.simbolo {
			return operacion, true
		}
	}
	return operacion{}, false
}

// tomarElementos verifica que la pila tenga la cantidad de elementos deseados para la operacion solicitada
// desapila los elementos y los devuelve
func tomarElementos(pila TDA.Pila[int64], aridad int) ([]int64, error) {
	numeros := make([]int64, aridad)
	for i := aridad - 1; i >= 0; i-- {
		if pila.EstaVacia() {
			return nil, fmt.Errorf("ERROR")
		}
		numeros[i] = pila.Desapilar()
	}
	return numeros, nil
}

// operarSuma realiza la suma de dos numeros
func operarSuma(operadores []int64) (int64, error) {
	return operadores[0] + operadores[1], nil
}

// operarResta realiza la resta de dos numeros
func operarResta(operadores []int64) (int64, error) {
	return operadores[0] - operadores[1], nil
}

// operarMultiplicacion realiza la multiplicacion de dos numeros
func operarMultiplicacion(operadores []int64) (int64, error) {
	return operadores[0] * operadores[1], nil
}

// operarDivision realiza la division de dos numeros
func operarDivision(operadores []int64) (int64, error) {
	if operadores[1] == 0 { //Chequea que el divisor sea valido
		return 0, fmt.Errorf("ERROR")
	}
	return operadores[0] / operadores[1], nil
}

// operarPotencia calcula la potencia de un numero
func operarPotencia(operadores []int64) (int64, error) {
	if operadores[1] < 0 { //Chequea que la potencia sea valida
		return 0, fmt.Errorf("ERROR")
	}
	return (int64(math.Pow(float64(operadores[0]), float64(operadores[1])))), nil
}

// operarLog calcula el logaritmo de un numero y una base dada
func operarLog(operadores []int64) (int64, error) {
	if operadores[0] <= 0 || operadores[1] < 2 { //Chequea que la base del logaritmo sea mayor a dos y el argumento no sea menor o igual a cero
		return 0, fmt.Errorf("ERROR")
	}
	logA := math.Log(float64(operadores[0]))
	logB := math.Log(float64(operadores[1]))
	logBaseB := logA / logB // Calcula el logaritmo en base de B
	return int64(logBaseB), nil
}

// operarSqrt calcula la raiz cuadrada de un numero
func operarSqrt(operadores []int64) (int64, error) {
	if 0 > operadores[0] { //Chequea que la raiz sea mayor a 0 para que sea valida la operacion
		return 0, fmt.Errorf("ERROR")
	}
	return (int64(math.Sqrt(float64(operadores[0])))), nil
}

// operarTernario evalua una condicion y devuelve el valor correspondiente
func operarTernario(operadores []int64) (int64, error) {
	if operadores[0] != 0 { //Si la condicion es verdadera devuelve el segundo numero, caso contrario el tercero
		return operadores[1], nil
	}
	return operadores[2], nil
}

// detectorError maneja los errores, en caso de que no los haya apila el resultado en caso de ser necesario
// En caso de que haya algun error devuelve true, false en caso contrario
func detectorError(pila TDA.Pila[int64], numero int64, err error, apilar bool) bool {
	if err != nil {
		return true
	}
	if apilar {
		pila.Apilar(numero)
		return false
	}
	return false
}
