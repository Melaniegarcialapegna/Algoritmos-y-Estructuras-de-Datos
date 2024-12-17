# TP1: Calculadora Polaca Inversa

El trabajo práctico número 1 tiene fecha de entrega para el día 26/09.

## Contenido
- Previo al enunciado
- Calculadora en notación posfija
- Funcionamiento
- Formato de entrada
- Condiciones de error
- Criterios de aprobación

## Previo al enunciado

Como se ha indicado en clase, esperamos para la elaboración de este trabajo práctico que ya tengan conocimiento pleno del uso de Go, lo cual incluye todo lo visto en clase y lo explicado en los videos sobre el lenguaje. Si no se ha visto alguno de los videos, es necesario que primero lo revisen porque este enunciado asume que esto es sabido.

También, recomendamos volver a revisar el video sobre cómo armar los módulos en Go, en particular para los TPs.

## Calculadora en notación posfija

Se pide implementar un programa `dc` que permita realizar operaciones matemáticas. La calculadora leerá exclusivamente de entrada estándar (no toma argumentos por línea de comandos), interpretando cada línea como una operación en notación polaca inversa (también llamada notación posfija, en inglés reverse Polish notation); para cada línea, se imprimirá por salida estándar el resultado del cálculo.

Ejemplo de varias operaciones, y su resultado:

```sh
$ cat oper.txt
5 3 +
5 3 -
5 3 /
3 5 8 + +
3 5 8 + -
3 5 - 8 +
2 2 + +
0 1 ?
1 -1 0 ?
5 sqrt

$ ./dc < oper.txt
8
2
1
16
-10
6
ERROR
ERROR
-1
2
```

## Funcionamiento

Todas las operaciones trabajarán con números enteros, y devolverán números enteros. Se recomienda usar el tipo de dato de Go `int64` para permitir operaciones de más de 32 bits.

El conjunto de operadores posibles es: suma (`+`), resta (`-`), multiplicación (`*`), división entera (`/`), raíz cuadrada (`sqrt`), exponenciación (`^`), logaritmo (`log`) en base arbitraria, y operador ternario (`?`).

Todos los operadores funcionan con dos operandos, excepto `sqrt` (toma un solo argumento) y el operador ternario (toma tres).

Tal y como se describe en la bibliografía enlazada, cualquier operación aritmética `a op b` se escribe en postfijo como `a b op` (por ejemplo, `3 - 2` se escribe en postfijo como `3 2 -`).

Para operaciones con un solo operando, el formato es obviamente `a op` (por ejemplo, `5 sqrt`). Por su parte, para el operador ternario, el ordenamiento de los argumentos seguiría el mismo principio, transformándose `a ? b : c` en `a b c ?`. Este operador ternario devuelve, si `a` es distinto a 0, el valor de `b`, y si es 0 el valor de `c`.

Ejemplos (nótese que toda la aritmética es entera, y el resultado siempre se trunca):

- `20 11 -` → `20-11 = 9`
- `20 -3 /` → `20/-3 = -6`
- `20 10 ^` → `20^10 = 10240000000000`
- `60 sqrt` → `√60 = 7`
- `256 4 ^ 2 log` → `log₂(256⁴) = 32`
- `1 -1 0 ?` → `1 ? -1 : 0 = -1` (funciona como en C)

## Formato de entrada

Cada línea de la entrada estándar representa una operación en su totalidad (produce un único resultado); y cada una de estas operaciones es independiente de las demás.

Los símbolos en la expresión pueden ser números, u operadores. Todos ellos estarán siempre separados por uno o más espacios; la presencia de múltiples espacios debe tenerse en cuenta a la hora de realizar la separación en tokens.

Nota adicional: puede haber también múltiples espacios al comienzo de la línea, antes del primer token; por otra parte, no necesariamente habrá un espacio entre el último token y el carácter salto de línea que le sigue.

El resultado final de cada operación debe imprimirse en una sola línea por salida estándar (stdout). En caso de error, debe imprimirse —para esa operación— la cadena `ERROR`, también por salida estándar, y sin ningún tipo de resultado parcial. Tras cualquier error en una operación, el programa continuará procesando el resto de líneas con normalidad.

Está permitido, para el cálculo de potencias, raíces y logaritmos, el uso de las funciones de la biblioteca estándar `math`.

## Condiciones de error

El mensaje `ERROR` debe imprimirse como resultado en cualquiera de las siguientes situaciones:

- Cantidad de operandos insuficiente (`1 +`).
- Al finalizar la evaluación, queda más de un valor resultante. Es decir, no se realizaron suficientes operaciones para terminar que quede un único resultado (por ejemplo `1 2 3 +`, o `+ 2 3 -`).

Errores propios de cada operación matemática, descritos a continuación:

- división por 0
- raíz de un número negativo
- base del logaritmo menor a 2
- argumento del logaritmo menor o igual a 0
- potencia con exponente negativo

## Criterios de aprobación

El código entregado debe ser claro y legible y ajustarse a las especificaciones de la consigna. Debe compilar sin advertencias y correr sin errores.

La entrega incluye, obligatoriamente, todos los archivos involucrados en la realización del TP (es decir, el módulo del trabajo en sí, que debe llamarse `dc`), así como el módulo `tdas` en caso de haber utilizado al menos alguno de los tipos de datos implementados anteriormente.

La entrega se realiza únicamente en forma digital a través del sistema de entregas, con todos los archivos mencionados en un único archivo ZIP.