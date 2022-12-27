package data


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
