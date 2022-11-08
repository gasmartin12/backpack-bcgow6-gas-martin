# Se requiere obtener las siguientes consultas:

#Seleccionar el nombre, el puesto y la localidad de los departamentos donde 
#trabajan los vendedores.
select nombre, puesto, localidad from departamento as dtp
join empleado as emp on dtp.depto_nro = emp.depto_nro;

#Visualizar los departamentos con más de cinco empleados.
select dep.nombre_depto from empleado emp join departamento dep
on dep.depto_nro = emp.depto_nro
group by emp.depto_nro
having count(emp.depto_nro) > 1;

#Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select emp.nombre, emp.apellido, emp.salario, dep.nombre_depto, emp.puesto from empleado emp 
join departamento dep on emp.depto_nro = dep.depto_nro 
where emp.puesto =
(select empleado.puesto 
	from empleado where empleado.nombre 
    like 'Mito' and empleado.apellido like 'Barchuk');

#Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
select emp.* from empleado emp 
join departamento dep on emp.depto_nro = dep.depto_nro
where dep.nombre_depto like 'Contabilidad'
order by emp.nombre;

#Mostrar el nombre del empleado que tiene el salario más bajo.
select nombre, apellido, salario from empleado order by salario limit 1;

#Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
select * from empleado emp
where emp.depto_nro = (select dep.depto_nro from departamento dep
						where dep.nombre_depto like 'Ventas')
and emp.salario = (select max(emp.salario) from empleado emp);