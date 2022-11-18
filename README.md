# parcial2-tlf

# El lenguaje para el que realizamos las expresiones regulares es scilab

# El lenguaje utilizado para construir el analizador es Golang

# Funcionamiento

- El programa recibe el nombre del archivo a leer como argumento al momento de la ejecución
- Se abre el archivo y se extrae su contenido en un arreglo de bytes
- La totalidad del arreglo es analizado por todas las expresiones regulares que generamos para scilab
- Los tokens y sus categorías son almacenadas en otro arreglo de bytes el cual es enviado al archivo de salida

# Los patrones

Los patrones son una estructura que contiene un puntero a una expresión regular compilada, así como una cadena
la cual es tratada como categoría

Luego, los patrones son compilados cuando se ejecuta el programa, y se añaden a un arreglo en orden, de la expresión más simple
(Palabras reservadas, operadores...) a la más compleja (Comentarios, nombres de variables...

# Los tokens

Los token son una estructura que contiene el valor (lexema) y la categoría a la cual pertenece (operador lógico, relacional, nombre de variable...)

# El análisis

Cuando se inicia el análisis, la función toma la totalidad de datos que existan en el archivo y comienza a evaluar cada uno de los patrones sobre este contenido, desde el más simple hasta el más complejo. Cuando existe una coincidencia (En cualquier parte del arreglo de bytes) se construye un nuevo token tomando el lexema y se asigna la categoría

Una vez construído el token, la ocurrencia es eliminada del arreglo de bytes y el proceso de análisis se repite hasta que el arreglo esté vacío

Es de vital importancia tener en cuenta que sólo son consideradas aquellas ocurrencias que ocurren en la posición 0 del arreglo, es decir, al comienzo del archivo... De esta manera conservamos el orden original del código fuente

Decidimos remover todos los espacios y tabulaciones del archivo antes de ser analizado. De esta manera es más fácil construir las expresiones regulares, los saltos de línea son ignorados

# Ejemplo (Cuando digo cadena es en realidad un arreglo de bytes)

Código fuente original:
variable1 = 2 + 4;

Cadena al momento de eliminar los espacios y tabulaciones
variable1=2+4;

Primera iteración del analizador
variable1 -> Identificador de Variable
Cadena:=2+4;

Segunda iteración
= -> Operador de asignación
Cadena=2+4;

Tercera iteración
2 -> Número Entero
Cadena=+4;

Cuarta iteración
+-> Operador Aritmético
Cadena:4;

Quinta iteración
4 -> Número Entero
Cadena:;

Sexta iteración
; -> Separador de Sentencia
Cadena:
