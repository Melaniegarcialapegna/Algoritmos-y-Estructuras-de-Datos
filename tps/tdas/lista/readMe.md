# Lista Enlazada

En esta entrega les agregamos el requerimiento de escribir la documentación completa de las primitivas (`lista.go`), con sus correspondientes pre y post condiciones. Para esto, pueden usar de muestra los archivos de interfaces que ya utilizaron para la implementación de pilas y colas.

## Primitivas de la Lista

Estas son las primitivas de listas que tienen que implementar:

```go
type Lista[T any] interface {
    EstaVacia() bool
    InsertarPrimero(T)
    InsertarUltimo(T)
    BorrarPrimero() T
    VerPrimero() T
    VerUltimo() T
    Largo() int
    Iterar(visitar func(T) bool)
    Iterador() IteradorLista[T]
}
```

En caso que se invoque a `BorrarPrimero`, `VerPrimero` o `VerUltimo` sobre una lista vacía, todas deben entrar en pánico con un mensaje `La lista está vacía`.

Además, es necesario tener la primitiva de creación de la lista enlazada (`lista_enlazada.go`):

```go
func CrearListaEnlazada[T any]() Lista[T] {
    //...
}
```

## Primitiva del Iterador Interno

Como está indicado entre las primitivas, se debe implementar el iterador interno cuya firma es:

```go
Iterar(visitar func(T) bool)
```

Dicha función debe aplicarse a cada uno de los datos de la lista (de primero a último), hasta que la lista se termine o la función `visitar` devuelva `false` (lo que ocurra primero).

## Primitivas del Iterador Externo

La primitiva `Iterador` de la lista debe devolver un `IteradorLista`, cuyas primitivas son:

```go
type IteradorLista[T any] interface {
    VerActual() T
    HaySiguiente() bool
    Siguiente()
    Insertar(T)
    Borrar() T
}
```

En caso que se invoque a `VerActual`, `Siguiente` o `Borrar` sobre un iterador que ya haya iterado todos los elementos, debe entrar en pánico con un mensaje `El iterador terminó de iterar`.

## Pruebas

Recordar que el archivo de pruebas debe estar en el paquete `lista_test`, mientras que `lista_enlazada` debe estar en el paquete `lista`.

Considerar que todas las primitivas (exceptuando `Iterar`) deben funcionar en tiempo constante.

Las pruebas deben incluir los casos básicos de TDA similares a los contemplados para la pila y la cola, y adicionalmente debe verificar los siguientes casos del iterador externo:

- Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
- Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
- Insertar un elemento en el medio se hace en la posición correcta.
- Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
- Remover el último elemento con el iterador cambia el último de la lista.
- Verificar que al remover un elemento del medio, este no está.
- Otros casos borde que pueden encontrarse al utilizar el iterador externo.
- Casos del iterador interno, incluyendo casos con corte (la función `visitar` devuelve `false` eventualmente).

Al igual que en los casos anteriores, deberán entregar en formato digital, subiendo el código a la página de entregas de la materia, con el código completo.

No olviden revisar las preguntas frecuentes de la lista enlazada.

## Bibliografía Recomendada

- Weiss, Mark Allen, “Data Structures and Algorithm Analysis”: 3.2. The List ADT.
- Cormen, Thomas H. “Introduction to Algorithms”: 10.2. Linked lists.