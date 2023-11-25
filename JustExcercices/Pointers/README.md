Explanation: 

What happens in memory when a func is called?

The stack works like a stack structure, which places functions on top of each other, making them
isolated chunks of code which turns them into isolated chunks of code, and those spaces are released 
as the functions are executed in that order as the functions are executed in that order (the first 
one to enter is the last one to leave and the last one to enter is the last one to exit). 

Differences between stack and heap allocations?

Stack allocations are memory addresses that can be accessed while the code is executing.
can be accessed, if the execution of the function ends the memory address is lost, while 
in heap we can keep it during the execution time of the whole go program.
Instead in heap we can keep it during the execution time of the whole go program,
depending only on the garbage collector to optimize that allocation or cleanup.

How do the copies of variables work according to the type of variable?

In strings, that are called from another function passing it as parameter a copy will be generated, to make a change in the string from another function you would have to return the string and assign it to another function.
to make a change in the string from another function would have to return this string and assign it in the result of calling
in the result of calling that function or by sending pointers as a reference.

In arrays, which are called from another function would be something similar to the strings, in fact there will be no possibility of appends 
possibility to make appends because again the array would be a copy.

In maps if you can affect the existing positions from external functions, even adding
new positions

How are pointers related to all this?

Pointers are designed to be able to access the memory location in both stack and heap, and make 
direct changes to the references to those variables, rather than to the heap, and make direct 
changes to the references of those variables, instead of generating copies to instead of generating 
copies to have a controlled mutability from different functions that the program may have.
that the program may have. Otherwise, we would have copies of the same variable in different points
of memory in different points of the memory and this could generate problems of immutability.

Explicacion:

Que pasa en memoria cuando una func es llamada?

La stack funciona como una estructura de pila, la cual coloca las funciones una encima de la otra
lo cual las convierte en trozos de codigo que estan isolados y dichos espacios se van liberando
conforme se vayan ejecutando las funciones en ese orden (la primera que entra es la ultima en salir
y la ultima en entrar es la primera en salir del stack)

Diferencias entre asignaciones stack y heap?

Las asignaciones en stack son direcciones de memoria que mientras se este en ejecucion el codigo
podran ser accedidas, si la ejecucion de la funcion termina la direccion de memoria se pierde,
en cambio en heap podemos mantenerla durante el tiempo de ejecucion de todo el programa go,
dependiendo solo del garbage collector para optimizar ese asignacion o limpieza

Como los punteros estan relacionados a todo esto?

Los punteros estan diseñados para poder acceder a la ubicacion en memoria tanto en stack como 
en heap, y hacer cambios directos en las referencias de esas variables, en lugar
de generar copias para tener un mutabilidad controlada desde diferentes funciones
que el programa pueda tener. De otra manera tendriamos copias de la misma variable
en diferentes puntos de la memoria y esto podria generar problemas de immutabilidad.

Como funcionan las copias de variables acorde al tipo de variable?

En strings, que se llaman desde otra funcion pasandolo como parametro se generara una copia, para
hacer un cambio en la string desde otra funcion tendria que retornarse dicha string y asignarla
en el resultado de llamar esa funcion o mediante el envio de punteros como referencia

En arrays, que se llaman desde otra funcion seria algo similar a las strings, de hecho no abra 
posibilidad de realizar appends pues nuevamente el array seria una copia

En maps si se puede afectar las posiciones existentes desde funciones externas, incluso agregar
nuevas posiciones