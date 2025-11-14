package taller

import (
  "fmt"
  "2_practica_ssdd_dist/utils"
)


type Vehiculo struct{
  Matricula int
  Marca string
  Modelo string
  FechaEntrada string
  FechaSalida string
  Incidencias []Incidencia
}

func (v Vehiculo)Info() (string){
  return fmt.Sprintf("%s %s (%05d)", v.Marca, v.Modelo, v.Matricula)
}

func (v Vehiculo)Visualizar(){
  fmt.Printf("%sMatricula: %s%05d\n", utils.BOLD, utils.END, v.Matricula)
  fmt.Printf("%sMarca: %s%s\n", utils.BOLD, utils.END, v.Marca)
  fmt.Printf("%sModelo: %s%s\n", utils.BOLD, utils.END, v.Modelo)
  fmt.Printf("%sFecha de entrada: %s%s\n", utils.BOLD, utils.END, v.FechaEntrada)
  fmt.Printf("%sFecha estimada de entrada: %s%s\n", utils.BOLD, utils.END, v.FechaSalida)
  // Incidencias
}

func (v *Vehiculo)Menu(){
  menu := []string{
    "Menu de vehiculo",
    "Visualizar",
    "Modificar"}

  for{
    menu[0] = fmt.Sprintf("Menu de %s", v.Info())

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          v.Visualizar()
        case 2:
          v.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (v *Vehiculo)Inicializar(){
  var exit bool = false

  utils.BoldMsg("Matrícula")
  utils.LeerInt(&v.Matricula)
  if v.Matricula == 0{
    exit = true
  }

  if !exit{
    utils.BoldMsg("Marca")
    utils.LeerStr(&v.Marca)
    if len(v.Marca) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Modelo")
    utils.LeerStr(&v.Modelo)
    if len(v.Modelo) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Fecha de entrada")
    utils.LeerStr(&v.FechaEntrada)
    if len(v.FechaEntrada) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Fecha de estimada de salida")
    utils.LeerStr(&v.FechaSalida)
    if len(v.FechaSalida) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Incidencias")
    v.MenuIncidencias()
  }
}

func (v *Vehiculo)Modificar(){

  menu := []string{
    "Modificar datos de vehículo",
    "Matricula",
    "Marca y modelo",
    "Fecha de entrada",
    "Fecha estimada de salida",
    "Incidencias"}
  var buf string
  var num int

  for{
    menu[0] = fmt.Sprintf("Modificar datos de %s", v.Info())
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerInt(&num)
          v.Matricula = num
          utils.InfoMsg("Matricula modificada")
        case 2:
          utils.LeerStr(&buf)
          v.Marca = buf
          utils.LeerStr(&buf)
          v.Modelo = buf
          utils.InfoMsg("Marca y modelo modificado")
        case 3:
          utils.LeerFecha(&v.FechaEntrada)
          utils.InfoMsg("Fecha de entrada modificada")
        case 4:
          utils.LeerFecha(&v.FechaSalida)
          utils.InfoMsg("Fecha estimada de salida modificada")
        case 5:
          v.MenuIncidencias()
      }
    } else if status == 2{
      break
    }
  }
}

func (v *Vehiculo)MenuIncidencias(){
  var i Incidencia  
  menu := []string{
    "Seleccione una incidencia",
    "Crear incidencia",
    "Eliminar incidencia"}

  for{
    menu = []string{
      "Seleccione una incidencia",
      "Crear incidencia",
      "Eliminar incidencia"}
    for _, i := range v.Incidencias{
      menu = append(menu, i.Info())
    }

    opt, status := utils.MenuFunc(menu)

    if status == 0{
      if opt == 1{
        i.Inicializar()
        v.CrearIncidencia(i)
      } else if opt == 2{
        // Eliminar incidencia
      } else {
        v.Incidencias[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}

func (v Vehiculo)ObtenerIndiceIncidencia(i_in Incidencia) (int){
  var res int = -1

  for i, inc := range v.Incidencias{
    if inc.Igual(i_in){
      res = i
    }
  }

  return res
}

func (v *Vehiculo)CrearIncidencia(i Incidencia){
  if i.Valido() && v.ObtenerIndiceIncidencia(i) == -1{
    v.Incidencias = append(v.Incidencias, i)
  } else {
    utils.ErrorMsg("No se ha podido crear el vehículo")
  }
}

func (v Vehiculo)Valido() (bool){
  return v.Matricula > 0 && len(v.Marca) > 0 && len(v.Modelo) > 0
}

func (v1 Vehiculo)Igual(v2 Vehiculo) (bool){
  return v1.Matricula == v2.Matricula
}

