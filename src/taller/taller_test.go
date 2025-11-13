package taller_test

import (
  "testing"
  "p2/main"
)

var defaults Taller
defaults.CrearMecanico("Pepe", 0, 0)
c := Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
v := Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
i := Incidencia{Id: 1, Tipo: 1, Prioridad: 1, Descripcion: "Luna delantera rota", Estado: 1}
v.CrearIncidencia(i)
c.CrearVehiculo(v)
v = Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
c.CrearVehiculo(v)
defaults.CrearCliente(c)
defaults.AsignarPlaza(c.Vehiculos[0])
defaults.AsignarPlaza(c.Vehiculos[1])

func Benchmark(b *testing.B){
  t := defaults
  t.Estado()
}
