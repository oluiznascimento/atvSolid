# Liskov Substitution Principle (Princípio da substituição de Liskov)

Este código em GoLang exemplifica o Liskov Substitution Principle (LSP), que afirma que as subclasses devem poder ser substituídas por suas superclasses sem afetar a integridade do programa.

O programa apresenta duas classes, `Duck` e `Ostrich`, ambas implementando a interface `Bird`. O método `MakeBirdFly` espera um tipo `Bird`. Mesmo que `Ostrich` não voe, o programa ainda compila e executa corretamente, demonstrando a substituição de `Ostrich` por `Duck`.

*Nota: O código está em inglês por convenção, mas os comentários estão em português para facilitar o entendimento dos colegas de classe.*
