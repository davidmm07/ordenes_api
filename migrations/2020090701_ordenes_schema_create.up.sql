-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.7.2
-- PostgreSQL version: 9.4
-- Project Site: pgmodeler.com.br
-- Model Author: ---

SET check_function_bodies = false;
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: organizacion | type: DATABASE --
-- -- DROP DATABASE organizacion;
-- CREATE DATABASE organizacion
-- ;
-- -- ddl-end --
-- 

-- object: ordenes | type: SCHEMA --
-- DROP SCHEMA ordenes;
CREATE SCHEMA ordenes;
-- ddl-end --

-- -- object: sesion | type: SCHEMA --
-- -- DROP SCHEMA sesion;
-- CREATE SCHEMA sesion;
-- -- ddl-end --
-- 
SET search_path TO pg_catalog,public,ordenes,sesion;
-- ddl-end --

-- object: ordenes.servicio | type: TABLE --
-- DROP TABLE ordenes.servicio;
CREATE TABLE ordenes.servicio(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.servicio ADD COLUMN id serial NOT NULL;
-- ddl-end --

-- object: nombre | type: COLUMN --
ALTER TABLE ordenes.servicio ADD COLUMN nombre varchar;
-- ddl-end --

-- object: descripcion | type: COLUMN --
ALTER TABLE ordenes.servicio ADD COLUMN descripcion varchar;
-- ddl-end --
-- object: pk_id_servicio | type: CONSTRAINT --
-- ALTER TABLE ordenes.servicio DROP CONSTRAINT pk_id_servicio;
ALTER TABLE ordenes.servicio ADD CONSTRAINT pk_id_servicio PRIMARY KEY (id);
-- ddl-end --


COMMENT ON COLUMN ordenes.servicio.id IS 'identificador tabla de servicio';
COMMENT ON COLUMN ordenes.servicio.nombre IS 'nombre del servicio';
COMMENT ON COLUMN ordenes.servicio.descripcion IS 'información acerca del servicio';
COMMENT ON CONSTRAINT pk_id_servicio ON ordenes.servicio IS 'primary key de la tabla de servicio';
-- ddl-end --

-- object: ordenes.solicitud | type: TABLE --
-- DROP TABLE ordenes.solicitud;
CREATE TABLE ordenes.solicitud(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN id serial NOT NULL;
-- ddl-end --

-- object: descripcion | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN descripcion varchar;
-- ddl-end --

-- object: fecha | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN fecha date;
-- ddl-end --

-- object: ticket | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN ticket varchar NOT NULL;
-- ddl-end --

-- object: calificacion | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN calificacion integer DEFAULT 0;
-- ddl-end --

-- object: tipo_solicitud | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN tipo_solicitud integer NOT NULL;
-- ddl-end --

-- object: cliente | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN cliente integer NOT NULL;
-- ddl-end --

-- object: estado | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN estado integer NOT NULL;
-- ddl-end --

-- object: servicio | type: COLUMN --
ALTER TABLE ordenes.solicitud ADD COLUMN servicio integer NOT NULL;
-- ddl-end --
-- object: pk_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud DROP CONSTRAINT pk_solicitud;
ALTER TABLE ordenes.solicitud ADD CONSTRAINT pk_solicitud PRIMARY KEY (id);
-- ddl-end --


COMMENT ON TABLE ordenes.solicitud IS 'entidad que gestiona las solicitudes creadas por los clientes';
COMMENT ON COLUMN ordenes.solicitud.id IS 'identificador tabla solicitud';
COMMENT ON COLUMN ordenes.solicitud.descripcion IS 'información detallada de la solicitud ';
COMMENT ON COLUMN ordenes.solicitud.ticket IS 'codigo generado a través del sistema';
COMMENT ON COLUMN ordenes.solicitud.calificacion IS 'calificacion de la orden, por defecto será 0';
COMMENT ON COLUMN ordenes.solicitud.cliente IS 'identificador del cliente asociado a la solicitud';
COMMENT ON COLUMN ordenes.solicitud.estado IS 'estado de la solicitud';
COMMENT ON CONSTRAINT pk_solicitud ON ordenes.solicitud IS 'primary key de la tabla solicitud';
-- ddl-end --

-- object: ordenes.cliente | type: TABLE --
-- DROP TABLE ordenes.cliente;
CREATE TABLE ordenes.cliente(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.cliente ADD COLUMN id serial NOT NULL;
-- ddl-end --

-- object: usuario | type: COLUMN --
ALTER TABLE ordenes.cliente ADD COLUMN usuario varchar(90) NOT NULL;
-- ddl-end --

-- object: password | type: COLUMN --
ALTER TABLE ordenes.cliente ADD COLUMN password varchar;
-- ddl-end --
-- object: pk_cliente | type: CONSTRAINT --
-- ALTER TABLE ordenes.cliente DROP CONSTRAINT pk_cliente;
ALTER TABLE ordenes.cliente ADD CONSTRAINT pk_cliente PRIMARY KEY (id);
-- ddl-end --


COMMENT ON TABLE ordenes.cliente IS 'tabla que maneja información de los clientes';
COMMENT ON CONSTRAINT pk_cliente ON ordenes.cliente IS 'primary key de la tabla cliente';
-- ddl-end --

-- object: ordenes.tecnico | type: TABLE --
-- DROP TABLE ordenes.tecnico;
CREATE TABLE ordenes.tecnico(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.tecnico ADD COLUMN id serial;
-- ddl-end --

-- object: nombre | type: COLUMN --
ALTER TABLE ordenes.tecnico ADD COLUMN nombre varchar(30);
-- ddl-end --
-- object: pk_tecnico | type: CONSTRAINT --
-- ALTER TABLE ordenes.tecnico DROP CONSTRAINT pk_tecnico;
ALTER TABLE ordenes.tecnico ADD CONSTRAINT pk_tecnico PRIMARY KEY (id);
-- ddl-end --


COMMENT ON TABLE ordenes.tecnico IS 'tabla que gestiona información de los tecnicos';
COMMENT ON COLUMN ordenes.tecnico.nombre IS 'nombre del tecnico';
COMMENT ON CONSTRAINT pk_tecnico ON ordenes.tecnico IS 'primary key de la tabla tecnico';
-- ddl-end --

-- object: ordenes.solicitud_tecnico | type: TABLE --
-- DROP TABLE ordenes.solicitud_tecnico;
CREATE TABLE ordenes.solicitud_tecnico(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.solicitud_tecnico ADD COLUMN id serial NOT NULL;
-- ddl-end --

-- object: solicitud | type: COLUMN --
ALTER TABLE ordenes.solicitud_tecnico ADD COLUMN solicitud integer NOT NULL;
-- ddl-end --

-- object: tecnico | type: COLUMN --
ALTER TABLE ordenes.solicitud_tecnico ADD COLUMN tecnico integer NOT NULL;
-- ddl-end --
-- object: pk_solicitud_tecnico | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud_tecnico DROP CONSTRAINT pk_solicitud_tecnico;
ALTER TABLE ordenes.solicitud_tecnico ADD CONSTRAINT pk_solicitud_tecnico PRIMARY KEY (id);
-- ddl-end --


COMMENT ON TABLE ordenes.solicitud_tecnico IS 'tabla compuesta de las tablas solicitud y tecnico';
COMMENT ON COLUMN ordenes.solicitud_tecnico.id IS 'identificador de la tabla compuesta de tecnico y solicitud';
COMMENT ON COLUMN ordenes.solicitud_tecnico.solicitud IS 'llave foranea a la tabla solicitud';
COMMENT ON COLUMN ordenes.solicitud_tecnico.tecnico IS 'llave foranea a la tabla tecnico';
COMMENT ON CONSTRAINT pk_solicitud_tecnico ON ordenes.solicitud_tecnico IS 'primary key de la tabla solicitud tecnico';
-- ddl-end --

-- object: ordenes.estado | type: TABLE --
-- DROP TABLE ordenes.estado;
CREATE TABLE ordenes.estado(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.estado ADD COLUMN id serial;
-- ddl-end --

-- object: nombre | type: COLUMN --
ALTER TABLE ordenes.estado ADD COLUMN nombre varchar;
-- ddl-end --
-- object: pk_estado | type: CONSTRAINT --
-- ALTER TABLE ordenes.estado DROP CONSTRAINT pk_estado;
ALTER TABLE ordenes.estado ADD CONSTRAINT pk_estado PRIMARY KEY (id);
-- ddl-end --


COMMENT ON TABLE ordenes.estado IS 'estados de la solicitud';
COMMENT ON CONSTRAINT pk_estado ON ordenes.estado IS 'primary key de la tabla estado';
-- ddl-end --

-- object: ordenes.tipo_solicitud | type: TABLE --
-- DROP TABLE ordenes.tipo_solicitud;
CREATE TABLE ordenes.tipo_solicitud(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.tipo_solicitud ADD COLUMN id serial NOT NULL;
-- ddl-end --

-- object: nombre | type: COLUMN --
ALTER TABLE ordenes.tipo_solicitud ADD COLUMN nombre varchar(30);
-- ddl-end --
-- object: pk_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.tipo_solicitud DROP CONSTRAINT pk_tipo_solicitud;
ALTER TABLE ordenes.tipo_solicitud ADD CONSTRAINT pk_tipo_solicitud PRIMARY KEY (id);
-- ddl-end --


COMMENT ON CONSTRAINT pk_tipo_solicitud ON ordenes.tipo_solicitud IS 'primary key de la tabla tipo solicitud';
-- ddl-end --

-- object: ordenes.calificacion | type: TABLE --
-- DROP TABLE ordenes.calificacion;
CREATE TABLE ordenes.calificacion(

);
-- ddl-end --

-- object: id | type: COLUMN --
ALTER TABLE ordenes.calificacion ADD COLUMN id serial;
-- ddl-end --

-- object: nombre | type: COLUMN --
ALTER TABLE ordenes.calificacion ADD COLUMN nombre varchar(25);
-- ddl-end --
-- object: pk_calificacion | type: CONSTRAINT --
-- ALTER TABLE ordenes.calificacion DROP CONSTRAINT pk_calificacion;
ALTER TABLE ordenes.calificacion ADD CONSTRAINT pk_calificacion PRIMARY KEY (id);
-- ddl-end --


COMMENT ON TABLE ordenes.calificacion IS 'tabla que contempla las posibles calificaciones y nomenclaturas de ellas';
COMMENT ON COLUMN ordenes.calificacion.id IS 'identificador de la tabla clasificacion';
COMMENT ON COLUMN ordenes.calificacion.nombre IS 'nombre de la clasificación';
COMMENT ON CONSTRAINT pk_calificacion ON ordenes.calificacion IS 'primary key de la tabla calificacion';
-- ddl-end --

-- object: fk_servicio_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud DROP CONSTRAINT fk_servicio_solicitud;
ALTER TABLE ordenes.solicitud ADD CONSTRAINT fk_servicio_solicitud FOREIGN KEY (servicio)
REFERENCES ordenes.servicio (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_tipo_solicitud_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud DROP CONSTRAINT fk_tipo_solicitud_solicitud;
ALTER TABLE ordenes.solicitud ADD CONSTRAINT fk_tipo_solicitud_solicitud FOREIGN KEY (tipo_solicitud)
REFERENCES ordenes.tipo_solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_estado_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud DROP CONSTRAINT fk_estado_solicitud;
ALTER TABLE ordenes.solicitud ADD CONSTRAINT fk_estado_solicitud FOREIGN KEY (estado)
REFERENCES ordenes.estado (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_cliente_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud DROP CONSTRAINT fk_cliente_solicitud;
ALTER TABLE ordenes.solicitud ADD CONSTRAINT fk_cliente_solicitud FOREIGN KEY (cliente)
REFERENCES ordenes.cliente (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_calificacion_solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud DROP CONSTRAINT fk_calificacion_solicitud;
ALTER TABLE ordenes.solicitud ADD CONSTRAINT fk_calificacion_solicitud FOREIGN KEY (calificacion)
REFERENCES ordenes.calificacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: solicitud | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud_tecnico DROP CONSTRAINT solicitud;
ALTER TABLE ordenes.solicitud_tecnico ADD CONSTRAINT solicitud FOREIGN KEY (solicitud)
REFERENCES ordenes.solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tecnico | type: CONSTRAINT --
-- ALTER TABLE ordenes.solicitud_tecnico DROP CONSTRAINT tecnico;
ALTER TABLE ordenes.solicitud_tecnico ADD CONSTRAINT tecnico FOREIGN KEY (tecnico)
REFERENCES ordenes.tecnico (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --



