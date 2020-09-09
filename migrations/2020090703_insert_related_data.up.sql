INSERT INTO ordenes.solicitud
(descripcion, fecha, ticket, calificacion, tipo_solicitud, cliente, estado, servicio)
VALUES('los soportes estan muy sucios en la parte superior y se aflojaron dos tornillos', '2020-09-07', '#001', 3, 2, 2, 3, 1);
INSERT INTO ordenes.solicitud
(descripcion, fecha, ticket, calificacion, tipo_solicitud, cliente, estado, servicio)
VALUES('requerimos de instalación de nuevos equipos', '2020-09-08', '#002', 0, 1, 1, 2, 2);
INSERT INTO ordenes.solicitud
(descripcion, fecha, ticket, calificacion, tipo_solicitud, cliente, estado, servicio)
VALUES('mantenimiento preventivo de equipos usados', '2020-09-15', '#003', 4, 2, 3, 3, 3);

INSERT INTO ordenes.solicitud_tecnico
(solicitud, tecnico)
VALUES(1, 1);
INSERT INTO ordenes.solicitud_tecnico
(solicitud, tecnico)
VALUES(2, 1);
INSERT INTO ordenes.solicitud_tecnico
(solicitud, tecnico)
VALUES(3, 2);
