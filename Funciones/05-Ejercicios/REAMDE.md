Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:

1. Perro 10 kg de alimento.
2. Gato 5 kg de alimento.
3. Hamster 250 g de alimento.
4. Tarántula 150 g de alimento.
Se solicita:

1. Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
2. Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:

const (

   dog    = "dog"

   cat    = "cat"

)


...


animalDog, msg := animal(dog)

animalCat, msg := animal(cat)


...


var amount float64

amount += animalDog(10)

amount += animalCat(10)