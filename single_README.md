# Single Responsibility Principle (Princípio da responsabilidade única)

Este código em GoLang ilustra o Single Responsibility Principle (SRP), que é um dos princípios SOLID. O SRP afirma que uma classe deve ter apenas uma razão para mudar, ou seja, ela deve ter apenas uma responsabilidade.

O código consiste em um programa que envia emails e notificações por meio de uma `NotificationManager`, que coordena o envio. Cada componente - `EmailSender` e `NotificationService` - tem uma responsabilidade única de enviar emails e notificações, respectivamente. Isso permite que cada componente seja alterado independentemente, facilitando a manutenção e evolução do sistema.

*Nota: O código está em inglês por convenção, mas os comentários estão em português para facilitar o entendimento dos colegas de classe.*
