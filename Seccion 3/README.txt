Para ejecutar el archivo go, se debe ocupar el siguiente comando en la consola: go run a.go

La funcion principal (main), es la encargada de definir los canales, las listas, la creacion de las cajas a traves de las gorutinas, la llamada de los clientes que fueron creados en la gorutina correspondiente a la generacion de clientes y la generacion de la atencion del cliente a traves de la gorutina.

Primero respecto a la generacion de clientes, se realiza en la gorutina correspondiente, en la cual se le asigna un ID a cada cliente de forma randomica para poder identificarlos. De la generacion de estos clientes, se genera una lista de los clientes en al funcion main. Esta creacion de la lista se realiza a traves de un canal que toma los clientes y los añade en una lista de la funcion main para posteriormente mandarselo a la caja para que este cliente sea atendido.

Para el control de la pausa o reanudacion de las gorutinas, crean canales de control donde de llaman para verificar su estado y dependiendo del numero que recibe, se pausa o se reanuda la gorutina.

La creacion de las cajas se realiza a través de la funcion go que crea la gorutina y además los canales necesarios para la transferencia de datos necesarios.

Se crea un for sin fin para el envio de los clientes a las cajas dependiendo de la lista de los clientes (obviamente con su canal de control) y que estos clientes sean atendidos con un tiempo randomico entre 1 a 5 segundos. El criterio de atencion al cliente depende del tiempo de la otra caja, es decir que si el timepo de llegada es mayor o igual al tiempo fin de la caja 1, se deberia atender al cliente en la caja 2, y viceversa.

Y por ultimo, al gorutina encargada de las cajas, solamente realiza el trabajo de atencion al cliente, obteniendo el ID del cliente, generando el numero random entre 1 y 5 (que es lo que se va a demorar el cliente en su atencion) y por ultimo retorna esos numeros para decir que la atencion al cliente fue completada exitosamente.