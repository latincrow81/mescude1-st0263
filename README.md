# ST0263 Topicos Especiales En Telematica
#
## Estudiante(s): Mauricio Escudero Restrepo , mescude1@eafit.edu.co
#
## Profesor: Edwin Nelson Montoya Munera, emontoya@eafit.edu.co
#
### Reto No 2. Procesos comunicantes por API REST, RPC y MOM
#
 Diseño e implementación de 2 microservicios básicos que ofrecen ambos un servicio al API Gateway y que se deben 
 comunicar por un middleware RPC y por un middleware MOM.
#

#### Como ejecutar

Este proyecto requiere Docker para funcionar.

<code>docker-compose up</code>

El Api Gateway funciona en el puerto 5000 y provee una interfaz tipo swagger en localhost:5000/ui


#### referencias
* https://x-team.com/blog/set-up-rabbitmq-with-docker-compose/
* https://blog.logrocket.com/creating-a-web-server-with-golang/
* https://github.com/sahansera/go-grpc
* https://github.com/lakhinsu/rabbitmq-go-example
* https://medium.com/@rahulsamant_2674/python-rabbitmq-8c1c3b79ab3d

### Reto No 3. Aplicación Monolítica con Balanceo y Datos Distribuidos (BD y archivos)

Desplegar un CMS wordpress empleando la tecnología de contenedores, con su propio dominio y certificado SSL. 

El sitio se llama: reto3.azuloso.me.

Se desplego de la siguiente manera:
    - 2 instancias en servidores individuales de wordpress dockerizado (identicas). Inicializadas desde crontab con el script incluido server_start.sh
    - 1 instancia para el servidor NFS instalado manualmente, se intento dockerizar pero se encontraron problemas con la imagen usada. quiza con mas timepo se podria contruir un mejor dockerfile basado en ubutu ya que los pasos de la guia en la referencia son simples.
    - 1 instancia para el servidor de bases de datos dockerizada y con un volumen local, tambien iniciada con un script al inicio
    - 1 instancia para el balanceador de cargas

Se reservo una ip para el balanceador de cargas y se asigno a un subdominio en un dominio ya existente, y se genero un certificado manualemente usando freeSSL.org

Todos los archivos y guias de referencia necesarias estan incluidas en el repositorio y readme.

Preferente pero no necesariamente las instancias de base de datos y nfs deben iniciarse primero.

#### referencias
* https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nfs-mount-on-ubuntu-20-04
* https://www.digitalocean.com/community/tutorials/how-to-set-up-nginx-load-balancing-with-ssl-termination

##### versión README.md -> 1.0 (2023-Marzo)