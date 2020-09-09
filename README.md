# Ordenes_api

El API ordenes, proporciona implementación con graphql y orm en postgres .

### Modelo de datos
![ordenes_db] 

* Lenguaje Golang Versión v1.14.6, (No Requerido con docker)

## Como ejecutar


### Requisitos previos :whale:
* Docker
* Docker Compose 


### Comenzar a desarrollar
#### Ejecutar con Docker

1. Clonar el repositorio
```sh
git clone -b dev https://github.com/davidmm07/ordenes_api
```

2. Moverse a la carpeta del repositorio
```sh
cd ordenes_api
```


3. Crear la network **back_end** para los contenedores
```sh
docker network create back_end
```

4. Ejecutar el compose del contenedor
```sh
docker-compose up --build
```

5. Comprobar que los contenedores estén en ejecución
```sh
docker ps 
```

6. Probar los endpoints del archivo en el directorio files

# Patrones Usados

- DAO(Data Access Object)

## Derechos de Autor

This program is free software: you can redistribute it 
and/or modify it under the terms of the GNU General Public 
License as published by the Free Software Foundation, either
version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

### David Alejandro Morales

### 2020
### 
