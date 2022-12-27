package data



func GetSpeciesByIds(zoo Zoo, ids ...string) []Species {
  // Declare uma variável para armazenar as espécies encontradas
  var species []Species

  // Percorra a lista de espécies do zoológico
  for _, s := range zoo.Species {
    // Verifique se o ID da espécie atual está presente na lista de IDs fornecidos
    for _, id := range ids {
      if s.ID == id {
        // Se estiver, adicione a espécie ao array de espécies encontradas
        species = append(species, s)
      }
    }
  }

  // Retorne o array de espécies encontradas
  return species
}
