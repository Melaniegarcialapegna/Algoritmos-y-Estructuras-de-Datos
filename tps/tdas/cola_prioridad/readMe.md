# Heap

En el sitio de descargas se incluye el archivo `cola_prioridad.go` correspondiente al ejercicio de la cola de prioridad.

El trabajo grupal consiste en implementar el tipo de dato abstracto Cola de Prioridad, utilizando un Heap.

## Primitivas de la cola de prioridad

```go
type ColaPrioridad[T any] interface {
    // EstaVacia devuelve true si la cantidad de elementos en el heap es 0, 
    // false en caso contrario.
    EstaVacia() bool

    // Encolar agrega un elemento al heap.
    Encolar(T)

    // VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra 
    // en pánico con un mensaje "La cola está vacía".
    VerMax() T

    // Desencolar elimina el elemento con máxima prioridad, y lo devuelve. Si 
    // está vacía, entra en pánico con un mensaje "La cola está vacía".
    Desencolar() T

    // Cantidad devuelve la cantidad de elementos que hay en la cola de 
    // prioridad.
    Cantidad() int
}
```

Además, las primitivas de creación del Heap deberán ser:

```go
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T]
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T]
```

La función de comparación funciona igual que en el caso del ABB. La segunda primitiva de creación debe ejecutarse en tiempo lineal, permitiendo crear el heap con los elementos pasados por parámetro.

También deben implementar el ordenamiento heapsort sobre un arreglo genérico, y las pruebas unitarias de todas las primitivas implementadas.

```go
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int)
```

Como siempre, deben subir el código completo a la página de entregas de la materia.

No olviden revisar las preguntas frecuentes del heap.

## Bibliografía recomendada

- Weiss, Mark Allen, “Data Structures and Algorithm Analysis”: Capítulo 6: Priority Queues (Heaps).
- Cormen, Thomas H. “Introduction to Algorithms”: 6.5. Priority queues, 6.1. Heaps, 6.2. Maintaining the heap property.