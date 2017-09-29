#  :point_right: CVDI :wink:
Componente vital descentralizado de interconexion (sinapsis). 

ES el encargado del flujo de trabajo, siendo el medio principal de comunicación por donde pasaran todas las conexiones.

Esta compuesto por:

- Agentes.
- CEHDUN.
 
## Caracteristicas:

Construccion   :no_entry_sign: :construction:  :muscle:

## Tecnologias:
 
- Back-end:
  - [Go](https://golang.org/)
  - [OAuth2](https://oauth.net/2/)

- Front-end:
  - [React](https://facebook.github.io/react/) o [Angular](https://angularjs.org/)
  - [Materialize](http://materializecss.com/)

## Manejo de datos:
 
- [Nosql](https://es.wikipedia.org/wiki/NoSQL)
- [Arangodb](https://www.arangodb.com/)

## Diagramas

Comportamiento del CVDI:

![Image of CVDI](https://github.com/merakive/cvdi/blob/master/diagrams/cvdi.png)


## Instalacion

Instalacion de Glide	

	curl https://glide.sh/get | sh

Obtener el paquete

	go get github.com/merakiVE/CVDI

Instalar dependencias

	cd $GOPATH/src/github.com/merakiVE/CVDI

	glide install

Generar llave publica y privada

    openssl genrsa -des3 -out keys/private.pem 2048
    openssl rsa -in keys/private.pem -outform PEM -pubout -out keys/public.pem

Generar llave publica y privada con el CLI

    CVDI generate keys
    CVDI gen keys

Renombrar el archivo "cvdi.conf.exmaple" por "cvdi.conf" y colocar el archivo en la raiz del proyecto o "/etc/cvdi/cvdi.conf" o "/opt/cvdi/cvdi.conf"

## ¿Cómo puedo contribuir? 
Solo debes leer el archivo `contributing.md`, que encontraras en [este enlace](https://github.com/merakive/cvdi/blob/master/.github/CONTRIBUTING.md)


