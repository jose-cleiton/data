package data


func IsManager(id string, zoo Zoo) bool {
  // Obtenha a lista de funcionários
  employees := zoo.Employees

  // Percorra a lista de funcionários
  for _, employee := range employees {
    // Verifique se o ID fornecido está incluído na lista de gerentes de um funcionário
    for _, managerID := range employee.Managers {
      if managerID == id {
        return true
      }
    }
  }

  // Se nenhum funcionário gerente responsável pelo ID fornecido for encontrado, retorne false
  return false
}
