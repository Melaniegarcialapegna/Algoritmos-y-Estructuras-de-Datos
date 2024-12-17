# Árbol Binario de Búsqueda

El trabajo grupal consiste en implementar el tipo de dato abstracto Árbol Binario de Búsqueda (ABB), una extensión del Diccionario de la entrega anterior, denominado DiccionarioOrdenado. Tanto el DiccionarioOrdenado como el ABB deben estar dentro del paquete `diccionario`. En el sitio de descargas se incluye el archivo `diccionario_ordenado.go`, descrito a continuación:

```go
type DiccionarioOrdenado[K comparable, V any] interface {
    Diccionario[K, V]

    IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)
    IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}
```

Todas las primitivas anteriores deben funcionar también, con el agregado de que tanto el iterador interno (`Iterar`) como el externo (`Iterador`) deben iterar en el orden correspondiente al Diccionario. Se agregan primitivas para iterar por rangos dados. Si `desde` es `nil`, se debe iterar desde la primera clave, y si `hasta` es `nil`, hasta la última. Si ambos son `nil`, se comporta como el iterador sin rango.

Además, la primitiva de creación del ABB será:

```go
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccinarioOrdenado[K, V]
```

La función de comparación recibe dos claves y devuelve:
- Un entero menor que 0 si la primera clave es menor que la segunda.
- Un entero mayor que 0 si la primera clave es mayor que la segunda.
- 0 si ambas claves son iguales.

Qué implica que una clave sea igual, mayor o menor que otra dependerá del usuario del TDA. Por ejemplo, `strings.Compare` cumple con esta especificación (si las claves son cadenas).

Como siempre, deben subir el código completo a la página de entregas de la materia.

## Estructura del módulo

Dado que el ABB es otra implementación de diccionario, simplemente deberán agregar los archivos al mismo paquete. Es decir:

```
tdas
   ├── pila
   ├── cola
   ├── lista
   ├── diccionario
   │   ├── hash.go
   │   ├── diccionario.go
   │   ├── diccionario_test.go
   │   ├── abb.go
   │   ├── diccionario_ordenado.go
   │   └── diccionario_ordenado_test.go
   └── go.mod
```

No olviden revisar las preguntas frecuentes del árbol binario de búsqueda.

## Bibliografía recomendada

- Weiss, Mark Allen, “Data Structures and Algorithm Analysis”: Capítulo 4: Árboles, en particular desde 4.3. The Search Tree ADT - Binary Search Trees.
- Cormen, Thomas H. “Introduction to Algorithms”: 12. Binary Search Trees.