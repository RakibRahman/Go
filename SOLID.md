# S.O.L.I.D Principles

## S -> Single Responsibility Principle (SRP)
  One Class should have only one reponsibility. There should be one reason to change a class.
  Achieve SRP by not breaking DRY and KISS principle.

## O -> Open/Closed Principle (OCP)
> Class should be open for extension but closed for modification.

If a class is used in app and there some changes needed to add, dont modify the  existing class, create a new class (extend) the class.

# Dependency Inversion vs Dependency Injection

The `Dependency Inversion Principle (DIP)` is one of the SOLID principles of object-oriented design. It states: 
- High-level modules should not depend on low-level modules. Both should depend on abstractions(interfaces). 
- Abstractions should not depend on details. Details should depend on abstractions. 
Example: Instead of a high-level module (like NotificationService) directly depending on a low-level module (like EmailService), both should depend on an abstraction (e.g., an interface)

Goal → Decouple high-level modules from low-level modules, promoting a more flexible and maintainable codebase

`Dependency Injection (DI) ` is a design pattern that implements the Dependency Inversion Principle. It involves "injecting" dependencies (like `EmailService`) into a class or struct (like `NotificationService`) rather than having the class create its own dependencies.
Dependency injection is the concept that your code should explicitly
specify the functionality it needs to perform its task

Benefits →
- Helps in achieving loose coupling between classes and facilitates easier testing, maintenance, and reusability.
- Easily swap out implementations (e.g., replace EmailService with SMSService).