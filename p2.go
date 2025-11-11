package main

import (
  "fmt"
)

const PLAZAS_MECANICO = 2
const BOLD = "\033[1;37m"
const RED = "\033[1;31m"
const YELLOW = "\033[1;33m"
const BLUE = "\033[1;34m"
const END = "\033[0m"

type Taller struct{
  Clientes []Cliente
  Plazas []Vehiculo
  Mecanicos []Mecanico
  UltimoId int
}

func (t *Taller)Menu(){
  menu := []string{
    "Menu del taller",
    "Asignar vehiculo",
    "Estado del taller",
    "Listar incidencias",
    "Listar clientes con vehiculos en el taller",
    "Listar incidencias de mecánico",
    "Mecánicos disponibles"}

  for{
    opt, status := menuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          t.Asignar()
        case 2:
          t.Estado()
        case 3:
          incidencias := t.ObtenerIncidencias()

          for _, inc := range incidencias{
            fmt.Println(inc.Info())
          }
        case 4:
          clientes := t.ObtenerClientesEnTaller()

          for _, c := range clientes{
            fmt.Println(c.Info())
          }
        case 5:
          // Incidencias por mecánico
        case 6:
          t.MecanicosDisponibles()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (t *Taller)MenuMecanicos(){
  var menu []string
  var m Mecanico
  var id int

  for{
    menu = []string{
      "Selecciona un mecánico",
      "Crear Mecánico",
      "Eliminar Mecánico"}
    for _, m := range t.Mecanicos{
      menu = append(menu, m.Info())
    }

    opt, status := menuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          m.Inicializar()
          t.CrearMecanico(m.Nombre, m.Especialidad, m.Experiencia)
          if !m.Valido() {
            errorMsg("No se ha creado el mecánico")
          }
        case 2:
          for {
            fmt.Println("Introduzca el ID del mecánico")
            leerInt(&id)
            m = t.ObtenerMecanicoPorId(id)
            if m.Valido(){
              t.EliminarMecanico(m)
              break
            }
          }
        default:
          t.Mecanicos[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}

func (t *Taller)MenuClientes(){
  var menu []string
  var c Cliente
  var id int

  for{
    menu = []string{
      "Selecciona un cliente",
      "Crear Cliente",
      "Eliminar Cliente"}
    for _, c := range t.Clientes{
      menu = append(menu, c.Info())
    }

    opt, status := menuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          c.Inicializar()
          t.CrearCliente(c)
          if !c.Valido() {
            errorMsg("No se ha creado el cliente")
          }
        case 2:
          for {
            fmt.Println("Introduzca el ID del cliente")
            leerInt(&id)
            c = t.ObtenerClientePorId(id)
            if c.Valido(){
              t.EliminarCliente(c)
              break
            }
          }
        default:
          t.Clientes[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}

func (t Taller)ListarIncidencias(){
}

func (t Taller)ObtenerMatriculaVehiculos() ([]int){
  var matriculas []int

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      if len(v.Incidencias) > 0{
        matriculas = append(matriculas, v.Matricula)
      }
    }
  }

  return matriculas
}

func(t Taller)HayEspacio() (bool){
  return len(t.Plazas) <= PLAZAS_MECANICO*len(t.Mecanicos)
}

func (t *Taller)AsignarPlaza(v Vehiculo){
  if t.HayEspacio(){
    for i, p := range t.Plazas{ // Plaza libre
      if !p.Valido(){
        t.Plazas[i] = v
        msg := fmt.Sprintf("Vehiculo asignado a la plaza %d", i + 1)
        infoMsg(msg)
        return
      }
    }
  }
}

func (t Taller)Estado(){
  var v Vehiculo

  for i := 0; i < PLAZAS_MECANICO*len(t.Mecanicos); i++{
    fmt.Printf("%d.- ", i + 1)
    v = t.Plazas[i]
    if v.Valido(){
      fmt.Print(v.Info())
    }
    fmt.Println() 
  }
}

func (t *Taller)Asignar(){
  matriculas := t.ObtenerMatriculaVehiculos()
  var num int
  var v Vehiculo

  if len(matriculas) > 0{
    fmt.Println(BOLD + "VEHICULOS DISPONIBLES" + END)
    for _, m := range matriculas{
      fmt.Println(m)
    }
    fmt.Println("Escriba la matrícula del vehículo a asignar")
    leerInt(&num)
    for _, c := range t.Clientes{
      v = c.ObtenerVehiculoPorMatricula(num)
      if v.Valido(){
        t.AsignarPlaza(v)
      }
    }
  } else {
    warningMsg("No hay incidencias en el taller")
  }
}

func (t *Taller)CrearMecanico(nombre string, especialidad int, experiencia int){
  var m Mecanico
  var v Vehiculo // plazas vacias

  m.Nombre = nombre
  m.Especialidad = especialidad
  m.Experiencia = experiencia
  m.Id = t.UltimoId + 1

  if m.Valido() && t.ObtenerIndiceMecanico(m) == -1{
    t.UltimoId++
    m.Id = t.UltimoId
    m.Alta = true
    t.Mecanicos = append(t.Mecanicos, m)
    t.Plazas = append(t.Plazas, v)
    t.Plazas = append(t.Plazas, v)
  } else {
    errorMsg("No se ha podido crear al mecanico")
  }
}

func (t *Taller)CrearCliente(c Cliente){
  if c.Valido(){
    t.Clientes = append(t.Clientes, c)
  }
}

func (t *Taller)EliminarMecanico(m Mecanico){
  
  indice := t.ObtenerIndiceMecanico(m)
    
  if indice >= 0{ // Eliminar
    lista := t.Mecanicos
    lista = lista[:indice+copy(lista[indice:], lista[indice+1:])]
    t.Mecanicos = lista
  } else {
    errorMsg("No se pudo eliminar al mecánico")
  }
}

func (t *Taller)EliminarCliente(c Cliente){
  indice := t.ObtenerIndiceCliente(c)
    
  if indice >= 0{ // Eliminar
    lista := t.Clientes
    lista = lista[:indice+copy(lista[indice:], lista[indice+1:])]
    t.Clientes = lista
  } else {
    errorMsg("No se pudo eliminar al mecánico")
  }
}

func (t Taller)ObtenerIndiceMecanico(m_in Mecanico) (int){
  var res int = -1

  for i, m := range t.Mecanicos{
    if m.Igual(m_in){
      res = i
    }
  }

  return res
}

func (t Taller)ObtenerMecanicoPorId(id int) (Mecanico){
  var res Mecanico

  for i, m := range t.Mecanicos{
    if m.Id == id{
      res = t.Mecanicos[i]
    }
  }

  return res
}

func (t Taller)ObtenerClientePorId(id int) (Cliente){
  var res Cliente

  for i, m := range t.Clientes{
    if m.Id == id{
      res = t.Clientes[i]
    }
  }

  return res
}

func (t Taller)ObtenerIndiceCliente(c_in Cliente) (int){
  var res int = -1

  for i, c := range t.Clientes{
    if c.Igual(c_in){
      res = i
    }
  }

  return res
}

func (t Taller)ObtenerMecanicosDisponibles() ([]Mecanico){
  var mecanicos []Mecanico  

  for _, m := range t.Mecanicos{
    if m.Alta{
      mecanicos = append(mecanicos, m)
    }
  }

  return mecanicos
}

func (t Taller)ObtenerIncidencias() ([]Incidencia){
  var incidencias []Incidencia

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      for _, inc := range v.Incidencias{
        incidencias = append(incidencias, inc)
      }
    }
  }

  return incidencias
}

func (t Taller)ObtenerClientesEnTaller() ([]Cliente){
  var clientes []Cliente

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      for _, p := range t.Plazas{
        if v.Igual(p) && !c.ExisteCliente(clientes){
          clientes = append(clientes, c)
          break // Se ha encontrado el vehiculo del cliente
        }
      }
    }
  }

  return clientes
}

func (t Taller)MecanicosDisponibles(){
  for _, m := range t.Mecanicos{
    if m.Alta{
      fmt.Println(m.Info())
    }
  }
}

func (t *Taller)ModificarMecanico(modif Mecanico){
  for i, m := range t.Mecanicos{
    if m.Igual(modif){
      t.Mecanicos[i] = modif
    }
  }
}


type Cliente struct{
  Id int
  Nombre string
  Telefono int
  Email string
  Vehiculos []Vehiculo
}

func (c *Cliente)Inicializar(){
  var exit bool = false

  fmt.Printf("%sID%s\n", BOLD, END)
  leerInt(&c.Id)
  if c.Id == 0{
    exit = true
  }

  if !exit{
    fmt.Printf("%sNombre%s\n", BOLD, END)
    leerStr(&c.Nombre)
    if len(c.Nombre) == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sTeléfono%s\n", BOLD, END)
    leerInt(&c.Telefono)
    if c.Telefono == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sEmail%s\n", BOLD, END)
    leerStr(&c.Email)
    if len(c.Email) == 0{
      exit = true
    }
  }

  // Ya está creado el cliente base (sin vehículos)
  if !exit{
    c.MenuVehiculos()
  }
}

func (c Cliente)Info() (string){
  return fmt.Sprintf("%s (%08d)", c.Nombre, c.Id)
}

func (c Cliente)Visualizar(){
  fmt.Printf("%sID: %s%08d\n", BOLD, END, c.Id)
  fmt.Printf("%sNombre: %s%s\n", BOLD, END, c.Nombre)
  fmt.Printf("%sTeléfono: %s%09d\n", BOLD, END, c.Telefono)
  fmt.Printf("%sEmail: %s%s\n", BOLD, END, c.Email)
  fmt.Printf("%sVehiculos:%s\n", BOLD, END)
  c.ListarVehiculos()
}

func (c *Cliente)MenuVehiculos(){
  var v Vehiculo  
  menu := []string{
    "Seleccione un vehículo",
    "Crear vehículo",
    "Eliminar vehículo"}

  for{
    menu = []string{
      "Seleccione un vehículo",
      "Crear vehículo",
      "Eliminar vehículo"}
    for _, v := range c.Vehiculos{
      menu = append(menu, v.Info())
    }

    opt, status := menuFunc(menu)

    if status == 0{
      if opt == 1{
        v.Inicializar()
        c.CrearVehiculo(v)
      } else if opt == 2{
        // Eliminar vehiculo
      } else {
        c.Vehiculos[opt - 3].Menu()
      }
    } else if status == 2{
      break
    }
  }
}

func (c *Cliente)CrearVehiculo(v Vehiculo){
  if v.Valido() && c.ObtenerIndiceVehiculo(v) == -1{
    c.Vehiculos = append(c.Vehiculos, v)
  } else {
    errorMsg("No se ha podido crear el vehículo")
  }
}

func (c Cliente)ListarVehiculos(){
  if len(c.Vehiculos) > 0{
    for _, v := range c.Vehiculos{
      fmt.Printf("  %s·%s%s\n", BOLD, END, v.Info())
    }
  } else {
    fmt.Println(BOLD + "  SIN VEHICULOS" + END)
  }
}

func (c *Cliente)Menu(){
  menu := []string{
    "Menu de cliente",
    "Visualizar",
    "Modificar"}

  for{
    menu[0] = fmt.Sprintf("Menu de %s", c.Nombre)

    opt, status := menuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          c.Visualizar()
        case 2:
          c.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (c *Cliente)Modificar(){
  menu := []string{
    "Modificar datos de cliente",
    "ID",
    "Nombre",
    "Teléfono",
    "Email",
    "Vehiculos"}
  var buf string
  var num int

  for{
    menu[0] = fmt.Sprintf("Modificar datos de %s", c.Nombre)
    opt, status := menuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          leerInt(&num)
          c.Id = num
          infoMsg("ID modificado")
        case 2:
          leerStr(&buf)
          c.Nombre = buf
          infoMsg("Nombre modificado")
        case 3:
          leerInt(&num)
          c.Telefono = num
          infoMsg("Teléfono modificado")
        case 4:
          leerStr(&buf)
          c.Email = buf
          infoMsg("Email modificado")
        case 5:
          c.MenuVehiculos()
      }
    } else if status == 2{
      break
    }
  }
}

func (c Cliente)Valido() (bool){
  return c.Id > 0 && len(c.Nombre) > 0 && c.Telefono > 0 && len(c.Email) > 0
}

func (c1 Cliente)Igual(c2 Cliente) (bool){
  return c1.Id == c2.Id
}

func (c Cliente)ObtenerIndiceVehiculo(v_in Vehiculo) (int){
  var res int = -1

  for i, v := range c.Vehiculos{
    if v.Igual(v_in){
      res = i
    }
  }

  return res
}

func (c Cliente)ObtenerVehiculoPorMatricula(matricula int) (Vehiculo){
  var res Vehiculo  

  for _, v := range c.Vehiculos{
    if v.Matricula == matricula{
      res = v
    }
  }

  return res
}

func (c_in Cliente)ExisteCliente(clientes []Cliente) (bool){
  var existe bool = false

  for _, c := range clientes{
    if c.Igual(c_in){
      existe = true
    }
  }

  return existe
}


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
  fmt.Printf("%sMatricula: %s%05d\n", BOLD, END, v.Matricula)
  fmt.Printf("%sMarca: %s%s\n", BOLD, END, v.Marca)
  fmt.Printf("%sModelo: %s%s\n", BOLD, END, v.Modelo)
  fmt.Printf("%sFecha de entrada: %s%s\n", BOLD, END, v.FechaEntrada)
  fmt.Printf("%sFecha estimada de entrada: %s%s\n", BOLD, END, v.FechaSalida)
  // Incidencias
}

func (v *Vehiculo)Menu(){
  menu := []string{
    "Menu de vehiculo",
    "Visualizar",
    "Modificar"}

  for{
    menu[0] = fmt.Sprintf("Menu de %s", v.Info())

    opt, status := menuFunc(menu)

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

  fmt.Printf("%sMatricula%s\n", BOLD, END)
  leerInt(&v.Matricula)
  if v.Matricula == 0{
    exit = true
  }

  if !exit{
    fmt.Printf("%sMarca%s\n", BOLD, END)
    leerStr(&v.Marca)
    if len(v.Marca) == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sModelo%s\n", BOLD, END)
    leerStr(&v.Modelo)
    if len(v.Modelo) == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sFecha de entrada%s\n", BOLD, END)
    leerStr(&v.FechaEntrada)
    if len(v.FechaEntrada) == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sFecha estimada de salida%s\n", BOLD, END)
    leerStr(&v.FechaSalida)
    if len(v.FechaSalida) == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sIncidencias:%s\n", BOLD, END)
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
    opt, status := menuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          leerInt(&num)
          v.Matricula = num
          infoMsg("Matricula modificada")
        case 2:
          leerStr(&buf)
          v.Marca = buf
          leerStr(&buf)
          v.Modelo = buf
          infoMsg("Marca y modelo modificado")
        case 3:
          leerFecha(&v.FechaEntrada)
          infoMsg("Fecha de entrada modificada")
        case 4:
          leerFecha(&v.FechaSalida)
          infoMsg("Fecha estimada de salida modificada")
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

    opt, status := menuFunc(menu)

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
    errorMsg("No se ha podido crear el vehículo")
  }
}

func (v Vehiculo)Valido() (bool){
  return v.Matricula > 0 && len(v.Marca) > 0 && len(v.Modelo) > 0
}

func (v1 Vehiculo)Igual(v2 Vehiculo) (bool){
  return v1.Matricula == v2.Matricula
}


type Incidencia struct{
  Id int
  Mecanicos []Mecanico
  Tipo int // 1 (Mecánica), 2 (Electrónica), 3(Carrocería)
  Prioridad int // 1 a 3 (Alta a baja)
  Descripcion string
  Estado int // 0 (Cerrado), 1 (Abierta), 2 (En proceso)
}

func (i Incidencia)Info() (string){
  return fmt.Sprintf("%s (%03d)", i.Descripcion, i.Id)
}

func (i Incidencia)Visualizar(){
  fmt.Printf("%sId: %s%03d\n", BOLD, END, i.Id)
  fmt.Printf("%sMarca: %s%d\n", BOLD, END, i.Tipo)
  fmt.Printf("%sModelo: %s%d\n", BOLD, END, i.Prioridad)
  fmt.Printf("%sFecha de entrada: %s%s\n", BOLD, END, i.Descripcion)
  fmt.Printf("%sFecha estimada de entrada: %s%d\n", BOLD, END, i.Estado)
  // Mecanicos
}

func (i *Incidencia)Menu(){
menu := []string{
  "Menu de incidencia",
  "Visualizar",
  "Modificar"}

for{
  menu[0] = fmt.Sprintf("Menu de %s", i.Info())

  opt, status := menuFunc(menu)

  if status == 0{
    switch opt{
      case 1:
        i.Visualizar()
      case 2:
        i.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (i *Incidencia)Inicializar(){
  var exit bool = false

  fmt.Printf("%sID%s\n", BOLD, END)
  leerInt(&i.Id)
  if i.Id == 0{
    exit = true
  }

  if !exit{
    for{
      menu_esp := []string{
        "Selecciona tipo",
        "Mecánica",
        "Electrónica",
        "Carrocería"}
      opt, status := menuFunc(menu_esp)
      if status == 0{
        i.Tipo = opt - 1
        break
      } else if status == 2{
        exit = true
        break
      }
    }
  }

  if !exit{
    fmt.Printf("%sPrioridad%s\n", BOLD, END)
    leerInt(&i.Prioridad)
    if i.Prioridad == 0{
      exit = true
    }
  }

  if !exit{
    fmt.Printf("%sDescripción%s\n", BOLD, END)
    leerStr(&i.Descripcion)
    if len(i.Descripcion) == 0{
      exit = true
    } else {
      i.Estado = 1
    }
  }

}

func (i *Incidencia)Modificar(){

}

func (i Incidencia)Valido() (bool){
  return i.Id > 0
}

func (i1 Incidencia)Igual(i2 Incidencia) (bool){
  return i1.Id == i2.Id
}


type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecanica, Electrica, Carroceria
  Experiencia int
  Alta bool
}

func (m *Mecanico)Menu(){
  menu := []string{
    "Menu de mecánico",
    "Visualizar",
    "Modificar"}
  
  for{
    menu[0] = fmt.Sprintf("Menu de %s", m.Nombre)

    opt, status := menuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          m.Visualizar()
        case 2:
          m.Modificar()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (m *Mecanico)Inicializar(){
  var exit bool = false  

  fmt.Printf("%sNombre%s\n", BOLD, END)
  leerStr(&m.Nombre)
  if len(m.Nombre) == 0{
    exit = true
  }

  if !exit{
    menu_esp := []string{
    "Selecciona especialidad",
    "Mecánica",
    "Electrónica",
    "Carrocería"}
    for{
      opt, status := menuFunc(menu_esp)
      if status != 1{
        if status == 0{
          m.Especialidad = opt - 1
        } else {
          exit = true
        }
        break
      }
    }
  }

  if !exit{
    fmt.Printf("%sExperiencia%s\n", BOLD, END)
    leerInt(&m.Experiencia)
    m.Id = 1
  }
}

func (m Mecanico)Info() (string){
  return fmt.Sprintf("%s (%03d)", m.Nombre, m.Id)
}

func (m Mecanico)Visualizar(){
  fmt.Printf("%sID: %s%03d\n", BOLD, END, m.Id)
  fmt.Printf("%sNombre: %s%s\n", BOLD, END, m.Nombre)
  fmt.Printf("%sEspecialidad: %s%s\n", BOLD, END, m.ObtenerEspecialidad())
  fmt.Printf("%sExperiencia: %s%d años\n", BOLD, END, m.Experiencia)
  fmt.Printf("%s¿Está de alta? %s%t\n", BOLD, END, m.Alta)
}

func (m Mecanico)Valido() (bool){

  return m.Id > 0 && m.Id <= 999 && len(m.Nombre) > 0 && m.Experiencia >= 0 && m.Especialidad >= 0 && m.Especialidad <= 2
}

func (m1 Mecanico)Igual(m2 Mecanico) (bool){
  return m1.Id == m2.Id
}

func (m *Mecanico)Modificar(){
  menu := []string{
    "Modificar datos de mecánico",
    "Nombre",
    "Especialidad",
    "Experiencia",
    "Dar de baja"}
  var buf string
  var num int

  for{
    if !m.Alta{
      menu[len(menu) - 1] = "Dar de alta"
    } else {
      menu[len(menu) - 1] = "Dar de baja"
    }
    menu[0] = fmt.Sprintf("Modificar datos de %s", m.Nombre)
    opt, status := menuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          leerStr(&buf)
          m.Nombre = buf
          infoMsg("Nombre modificado")
        case 2:
          menu_esp := []string{
            "Selecciona especialidad",
            "Mecánica",
            "Electrónica",
            "Carrocería"}
          opt, status = menuFunc(menu_esp)
          if status == 0{
            esp := m.ObtenerEspecialidad()
            m.Especialidad = opt - 1
            msg := fmt.Sprintf("Especialidad modificada: %s->%s", esp, m.ObtenerEspecialidad())
            infoMsg(msg)
          }
        case 3:
          leerInt(&num)
          m.Experiencia = num
          infoMsg("Experiencia modificada")
        case 4:
          m.Alta = !m.Alta
          infoMsg("Estado modificado")
      }
    } else if status == 2{
      break
    }
  }
}

func (m Mecanico)ObtenerEspecialidad() (string){
  switch m.Especialidad{
    case 0:
      return "Mecánica"
    case 1:
      return "Electrónica"
    case 2:
      return "Carrocería"
    default:
      return "Sin especialidad"
  }
}


func errorMsg(msg string){
  fmt.Printf("%s%s%s\n", RED, msg, END)
}

func warningMsg(msg string){
  fmt.Printf("%s%s%s\n", YELLOW, msg, END)
}

func infoMsg(msg string){
  fmt.Printf("%s%s%s\n", BLUE, msg, END)
}

func leerFecha(aux *string){
  var dia int
  var mes int
  var anyo int

  for{
    fmt.Println("Día")
    leerInt(&dia)
    fmt.Println("Mes")
    leerInt(&mes)
    fmt.Println("Año")
    leerInt(&anyo)
    
    if (dia > 0 && dia <= 31 && mes > 0 && mes <= 12 && anyo > 0){
      *aux = fmt.Sprintf("%d-%d-%d", dia, mes, anyo)
      return
    } else if (dia == 0 && mes == 0 && anyo == 0){
      return
    }
  }
}

func leerInt(i *int){
  for{
    fmt.Print("> ")
    fmt.Scanf("%d", i)
    if *i >= 0{
      break
    } else {
      warningMsg("Valor entero inválido")
    }
  }
}

func leerStr(str *string){
  for{
    fmt.Print("> ")
    fmt.Scanf("%s", str)
    if len(*str) > 0{
      break
    } else {
      warningMsg("Cadena de texto inválida")
    }
  }
}

func menuFunc(menu []string) (int, int){
  var opt int

  menu = append(menu, "Salir")
  fmt.Printf("%s%s%s\n", BOLD, menu[0], END) // Menu title

  for i:= 1; i < len(menu); i++{
    fmt.Printf("%d.- %s\n", i, menu[i])
  }

  leerInt(&opt)

  if opt > 0 && opt < len(menu) - 1{
    return opt, 0
  } else if opt == len(menu) - 1{
    return opt, 2
  }
  return 0, 1
}


func main(){
  var t Taller
  
  menu := []string{
    "Menu principal",
    "Taller",
    "Clientes",
    "Mecánicos"}

  // INICIALIZAR
  t.CrearMecanico("Pepe", 0, 0)
  c := Cliente{Id: 1, Nombre: "Laura", Telefono: 1, Email: "laura27@mail.com"}
  v := Vehiculo{Matricula: 1234, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  i := Incidencia{Id: 1, Tipo: 1, Prioridad: 1, Descripcion: "Luna delantera rota", Estado: 1}
  v.CrearIncidencia(i)
  c.CrearVehiculo(v)
  v = Vehiculo{Matricula: 1235, Marca: "Toyota", Modelo: "Camry", FechaEntrada: "14-04-2009", FechaSalida: "19-04-2009"}
  c.CrearVehiculo(v)
  t.CrearCliente(c)
  t.AsignarPlaza(c.Vehiculos[0])
  t.AsignarPlaza(c.Vehiculos[1])
  // FIN INICIALIZAR

  for{
    opt, status := menuFunc(menu)
    
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
