INSERT INTO ordenes.cliente
(usuario, "password")
VALUES('Albertwesk@gmail.com', 'jhquuqwdhhaG211h23jadhajs1');
INSERT INTO ordenes.cliente
(usuario, "password")
VALUES('davidmoy@gmail.com', 'jqwwqhkjhaasjkdhsakdjahdhdk2');
INSERT INTO ordenes.cliente
(usuario, "password")
VALUES('elizabethpersonal@gmail.com', 'qhwejkqwkhuisahdkjsa211kjasd3');

INSERT INTO ordenes.calificacion
(id,nombre)
VALUES(0,'sin calificar');
INSERT INTO ordenes.calificacion
(nombre)
VALUES('muy malo');
INSERT INTO ordenes.calificacion
(nombre)
VALUES('malo');
INSERT INTO ordenes.calificacion
(nombre)
VALUES('regular');
INSERT INTO ordenes.calificacion
(nombre)
VALUES('bueno');
INSERT INTO ordenes.calificacion
(nombre)
VALUES('excelente');

INSERT INTO ordenes.estado
(nombre)
VALUES('creado');
INSERT INTO ordenes.estado
(nombre)
VALUES('en espera');
INSERT INTO ordenes.estado
(nombre)
VALUES('finalizado');


INSERT INTO ordenes.servicio
(nombre, descripcion)
VALUES('reparación y limpieza de soportes', 'los soportes estan muy sucios en la parte superior y se aflojaron dos tornillos');
INSERT INTO ordenes.servicio
(nombre, descripcion)
VALUES('instalación de nuevos equipos', 'requerimos de instalación de nuevos equipos');
INSERT INTO ordenes.servicio
(nombre, descripcion)
VALUES('mantenimiento preventivo de equipos usados', 'quiero hacer un mantenimiento a mis equipos casi nuevos');

INSERT INTO ordenes.tecnico
(nombre)
VALUES('Julian Reus');

INSERT INTO ordenes.tecnico
(nombre)
VALUES('Mariana Lopez');

INSERT INTO ordenes.tipo_solicitud
(nombre)
VALUES('instalacion');

INSERT INTO ordenes.tipo_solicitud
(nombre)
VALUES('mantenimiento');

