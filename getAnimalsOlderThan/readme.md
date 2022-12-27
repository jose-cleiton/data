# GetAnimalsOlderThan

Esta função, a partir do nome de uma espécie e uma idade mínima, verifica se todos os animais daquela espécie possuem a idade mínima especificada.

## Parâmetros

- `zoo`: um struct do tipo `Zoo` que representa o zoológico.
- `speciesName`: uma string contendo o nome da espécie a ser verificada.
- `minAge`: um inteiro contendo a idade mínima a ser verificada.

## Retorno

Um valor booleano indicando se todos os animais da espécie especificada possuem a idade mínima especificada.

## Exemplos

```go
import "zoo"

// Verifique se todos os leões do zoológico possuem pelo menos 12 anos
allLionsOlderThan12 := zoo.GetAnimalsOlderThan(zoo, "lions", 12)

// Verifique se todos os tigres do zoológico possuem pelo menos 5 anos
allTigersOlderThan5 := zoo.GetAnimalsOlderThan(zoo, "tigers", 5)
