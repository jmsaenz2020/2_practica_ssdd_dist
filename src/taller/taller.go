package taller

import (
  "fmt"
  "2_practica_ssdd_dist/utils"
)

const PLAZAS_MECANICO = 2

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
    opt, status := utils.MenuFunc(menu)
    
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

    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          m.Inicializar()
          t.CrearMecanico(m.Nombre, m.Especialidad, m.Experiencia)
          if !m.Valido() {
            utils.ErrorMsg("No se ha creado el mecánico")
          }
        case 2:
          for {
            fmt.Println("Introduzca el ID del mecánico")
            utils.LeerInt(&id)
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

    opt, status := utils.MenuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          c.Inicializar()
          t.CrearCliente(c)
          if !c.Valido() {
            utils.ErrorMsg("No se ha creado el cliente")
          }
        case 2:
          for {
            fmt.Println("Introduzca el ID del cliente")
            utils.LeerInt(&id)
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
        utils.InfoMsg(msg)
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
    utils.BoldMsg("VEHICULOS DISPONIBLES")
    for _, m := range matriculas{
      fmt.Println(m)
    }
    fmt.Println("Escriba la matrícula del vehículo a asignar")
    utils.LeerInt(&num)
    for _, c := range t.Clientes{
      v = c.ObtenerVehiculoPorMatricula(num)
      if v.Valido(){
        t.AsignarPlaza(v)
      }
    }
  } else {
    utils.WarningMsg("No hay incidencias en el taller")
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
    utils.ErrorMsg("No se ha podido crear al mecanico")
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
    utils.ErrorMsg("No se pudo eliminar al mecánico")
  }
}

func (t *Taller)EliminarCliente(c Cliente){
  indice := t.ObtenerIndiceCliente(c)
    
  if indice >= 0{ // Eliminar
    lista := t.Clientes
    lista = lista[:indice+copy(lista[indice:], lista[indice+1:])]
    t.Clientes = lista
  } else {
    utils.ErrorMsg("No se pudo eliminar al mecánico")
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

  utils.BoldMsg("ID")
  utils.LeerInt(&c.Id)
  if c.Id == 0{
    exit = true
  }

  if !exit{
    utils.BoldMsg("Nombre")
    utils.LeerStr(&c.Nombre)
    if len(c.Nombre) == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Telefono")
    utils.LeerInt(&c.Telefono)
    if c.Telefono == 0{
      exit = true
    }
  }

  if !exit{
    utils.BoldMsg("Email")
    utils.LeerStr(&c.Email)
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
  fmt.Printf("%sID: %s%08d\n", utils.BOLD, utils.END, c.Id)
  fmt.Printf("%sNombre: %s%s\n", utils.BOLD, utils.END, c.Nombre)
  fmt.Printf("%sTeléfono: %s%09d\n", utils.BOLD, utils.END, c.Telefono)
  fmt.Printf("%sEmail: %s%s\n", utils.BOLD, utils.END, c.Email)
  fmt.Printf("%sVehiculos:%s\n", utils.BOLD, utils.END)
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

    opt, status := utils.MenuFunc(menu)

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
    utils.ErrorMsg("No se ha podido crear el vehículo")
  }
}

func (c Cliente)ListarVehiculos(){
  if len(c.Vehiculos) > 0{
    for _, v := range c.Vehiculos{
      fmt.Printf("  %s·%s%s\n", utils.BOLD, utils.END, v.Info())
    }
  } else {
    utils.BoldMsg("SIN VEHICULOS")
  }
}

func (c *Cliente)Menu(){
  menu := []string{
    "Menu de cliente",
    "Visualizar",
    "Modificar"}

  for{
    menu[0] = fmt.Sprintf("Menu de %s", c.Nombre)

    opt, status := utils.MenuFunc(menu)

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
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerInt(&num)
          c.Id = num
          utils.InfoMsg("ID modificado")
        case 2:
          utils.LeerStr(&buf)
          c.Nombre = buf
          utils.InfoMsg("Nombre modificado")
        case 3:
          utils.LeerInt(&num)
          c.Telefono = num
          utils.InfoMsg("Teléfono modificado")
        case 4:
          utils.LeerStr(&buf)
          c.Email = buf
          utils.InfoMsg("Email modificado")
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
  fmt.Printf("%sId: %s%03d\n", utils.BOLD, utils.END, i.Id)
  fmt.Printf("%sMarca: %s%d\n", utils.BOLD, utils.END, i.Tipo)
  fmt.Printf("%sModelo: %s%d\n", utils.BOLD, utils.END, i.Prioridad)
  fmt.Printf("%sFecha de entrada: %s%s\n", utils.BOLD, utils.END, i.Descripcion)
  fmt.Printf("%sFecha estimada de entrada: %s%d\n", utils.BOLD, utils.END, i.Estado)
  // Mecanicos
}

func (i *Incidencia)Menu(){
menu := []string{
  "Menu de incidencia",
  "Visualizar",
  "Modificar"}

for{
  menu[0] = fmt.Sprintf("Menu de %s", i.Info())

  opt, status := utils.MenuFunc(menu)

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

  utils.BoldMsg("ID")
  utils.LeerInt(&i.Id)
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
      opt, status := utils.MenuFunc(menu_esp)
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
  utils.BoldMsg("Prioridad")
    utils.LeerInt(&i.Prioridad)
    if i.Prioridad == 0{
      exit = true
    }
  }

  if !exit{
  utils.BoldMsg("Descripción")
    utils.LeerStr(&i.Descripcion)
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

    opt, status := utils.MenuFunc(menu)

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

  utils.BoldMsg("Nombre")
  utils.LeerStr(&m.Nombre)
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
      opt, status := utils.MenuFunc(menu_esp)
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
    utils.BoldMsg("Experiencia")
    utils.LeerInt(&m.Experiencia)
    m.Id = 1
  }
}

func (m Mecanico)Info() (string){
  return fmt.Sprintf("%s (%03d)", m.Nombre, m.Id)
}

func (m Mecanico)Visualizar(){
  fmt.Printf("%sID: %s%03d\n", utils.BOLD, utils.END, m.Id)
  fmt.Printf("%sNombre: %s%s\n", utils.BOLD, utils.END, m.Nombre)
  fmt.Printf("%sEspecialidad: %s%s\n", utils.BOLD, utils.END, m.ObtenerEspecialidad())
  fmt.Printf("%sExperiencia: %s%d años\n", utils.BOLD, utils.END, m.Experiencia)
  fmt.Printf("%s¿Está de alta? %s%t\n", utils.BOLD, utils.END, m.Alta)
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
    opt, status := utils.MenuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          utils.LeerStr(&buf)
          m.Nombre = buf
          utils.InfoMsg("Nombre modificado")
        case 2:
          menu_esp := []string{
            "Selecciona especialidad",
            "Mecánica",
            "Electrónica",
            "Carrocería"}
          opt, status = utils.MenuFunc(menu_esp)
          if status == 0{
            esp := m.ObtenerEspecialidad()
            m.Especialidad = opt - 1
            msg := fmt.Sprintf("Especialidad modificada: %s->%s", esp, m.ObtenerEspecialidad())
            utils.InfoMsg(msg)
          }
        case 3:
          utils.LeerInt(&num)
          m.Experiencia = num
          utils.InfoMsg("Experiencia modificada")
        case 4:
          m.Alta = !m.Alta
          utils.InfoMsg("Estado modificado")
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
