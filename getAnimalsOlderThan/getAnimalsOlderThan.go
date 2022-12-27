package zoo

type Resident struct {
  Name string
  Sex  string
  Age  int
}

type Species struct {
  ID          string
  Name        string
  Popularity  int
  Location    string
  Availability []string
  Residents   []Resident
}

type Employee struct {
  ID          string
  FirstName   string
  LastName    string
  Managers    []string
  ResponsibleFor []string
}

type Hours struct {
  Open  int
  Close int
}

type Prices struct {
  Adult  float64
  Senior float64
  Child  float64
}

type Zoo struct {
  Species    []Species
  Employees  []Employee
  Hours      map[string]Hours
  Prices     Prices
}
func GetAnimalsOlderThan(zoo Zoo, speciesName string, minAge int) bool {
  // Percorra a lista de espécies do zoológico
  for _, s := range zoo.Species {
    // Verifique se o nome da espécie atual é o mesmo que o especificado
    if s.Name == speciesName {
      // Percorra a lista de animais da espécie atual
      for _, r := range s.Residents {
        // Verifique se o animal atual tem a idade mínima especificada
        if r.Age < minAge {
          // Se não tiver, retorne "false"
          return false
        }
      }
    }
  }

  // Se chegou até aqui, significa que todos os animais da espécie especificada têm a idade mínima especificada
  // Retorne "true"
  return true
}
