Explanation: Explicacion:

What happens in memory when a func is called? 

Stack works like a pile struct, that allocate the funcs one above of another 
that its convert the piece of code that are isolate and that spaces it will be free
according it will be executing the func in that order ( the first that enter is the last
in exit and the last one in enter is the first in exit from the stack)

Que pasa en memoria cuando una func es llamada?

La stack funciona como una estructura de pila, la cual coloca las funciones una encima de la otra
lo cual las convierte en trozos de codigo que estan isolados y dichos espacios se van liberando
conforme se vayan ejecutando las funciones en ese orden (la primera que entra es la ultima en salir
y la ultima en entrar es la primera en salir del stack)

Difference beetween stack and heap allocations? 

The allocations in stack are address memory meanwhile are in execution code could be are access,
if the func execution ends the address memory, it lost in change in heap could maintain during
the execution time of all go program, depending just by garbage collector to optimize
that allocation or clean

Diferencias entre asignaciones stack y heap?

Las asignaciones en stack son direcciones de memoria que mientras se este en ejecucion el codigo
podran ser accedidas, si la ejecucion de la funcion termina la direccion de memoria, se pierde
en cambio en heap podemos mantenerla durante el tiempo de ejecucion de todo el programa go,
dependiendo solo del garbage collector para optimizar ese asignacion o limpieza

How pointers relate to any of that? 

The pointer are desing to could access to address memory instead in stack like in heap,
and made directly changes in the references of that variables, in place of 
generate copies to have control mutability from different functions that the program
could have. By another way we have copies of the same variable in different places
of memory and this could generate inmutability problems.

Como los punteros estan relacionados a todo esto?

Los punteros estan diseñados para poder acceder a la ubicacion en memoria tanto en stack como 
en heap, y hacer cambios directos en las referencias de esas variables, en lugar
de generar copias para tener un mutabilidad controlada desde diferentes funciones
que el programa pueda tener. De otra manera tendriamos copias de la misma variable
en diferentes puntos de la memoria y esto podria generar problemas de immutabilidad.