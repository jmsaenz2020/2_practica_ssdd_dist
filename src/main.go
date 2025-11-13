package main

import (
  "2_practica_ssdd_dist/taller"
  "2_practica_ssdd_dist/utils"
)

func main(){
  var t taller.Taller
  
  menu := []string{
    "Menu principal",
    "Taller",
    "Clientes",
    "Mecánicos"}

  // INICIALIZAR
  t.CrearMecanico("Pepe", 0, 0)
  c := taller.Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
  v := taller.Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  i := taller.Incidencia{Id: 1, Tipo: 1, Prioridad: 1, Descripcion: "Luna delantera rota", Estado: 1}
  v.CrearIncidencia(i)
  c.CrearVehiculo(v)
  v = taller.Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  c.CrearVehiculo(v)
  t.CrearCliente(c)
  t.AsignarPlaza(c.Vehiculos[0])
  t.AsignarPlaza(c.Vehiculos[1])
  // FIN INICIALIZAR

  for{
    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          t.Menu()
        case 2:
          t.MenuClientes()
        case 3:
          t.MenuMecanicos()
      }
    } else if status == 2{
      break
    }
  }
  
}
