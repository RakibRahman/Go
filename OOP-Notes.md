Object-Oriented Programming (OOP)
Object-Oriented Programming is a paradigm that uses "objects" to design software. These objects are instances of classes, which bundle data (attributes) and methods (functions) that operate on the data. OOP focuses on modeling real-world entities and their interactions.

Key Features:

Class-based: Code is organized into classes, which define the blueprint for objects. Class - blueprint of object
Object - abstraction of entity

Encapsulation: Bundling of data and methods within classes to restrict access and protect state.

Inheritance: Mechanism to create new classes based on existing classes to promote code reuse.
Inheritance - Inherit all properties, behaviors and override or add some

Polymorphism: Ability to use a common interface for different underlying forms (methods).

Abstraction: Hiding complex implementation details and showing only essential features.



Encapsulation is kind of protection wrapper around data and behavior.
Encapsulation is the process of keeping property and method inside abstracted object(class) and controlling access.

Polymorphism allows objects of different types to be treated as objects of a common supertype. This is achieved through method overriding and interfaces.
Polymorphism enables a single interface to be used for a general class of actions, allowing for flexibility and the ability to switch between different implementations.


An interface defines a contract or a set of methods that implementing classes must provide. It specifies what methods a class must implement but does not provide the implementation details.
Interfaces promote decoupling and enhance code flexibility and reusability by ensuring that different classes can interact with each other through a common set of methods.
 An interface consists only unimplemented methods.


 #  principles of OOAD

- Abtraction -  Hiding the complex implementation details and exposing only the essential behavior.
- Encapsulation -Controlling access to properties and behaviors using access modifiers (private, protected, public).
- Decomposition - Breaking down a system into smaller, manageable objects.
- Generalization -  Identifying common behavior among objects and creating a base class to represent it.
- Association - Two objects exist independently but can interact with each other.
- Aggregation - A unidirectional association where one object contains another, but the contained object can exist independently.
- Composition -  A strong relationship where one object is composed of another, and the contained object cannot exist without the container.

  

# What is an Interface? 

An interface  in Go is a collection of method signatures. Any type that implements all the methods defined in the interface is said to "satisfy" or "implement" the interface. Interfaces allow you to define shared behavior across different types without tying them to a specific implementation. 


# Dependency Inversion vs Dependency Injection

`Dependency Inversion` is a design principle that guides the structure of your code.
A design principle that suggests high-level modules or classes should not depend on low-level modules but both should depend on abstractions. 
It also states that abstractions should not depend on details; details should depend on abstractions.

Goal → Decouple high-level modules from low-level modules, promoting a more flexible and maintainable codebase

`Dependency Injection` is a specific technique to implement Dependency Inversion by injecting dependencies from the outside rather than creating them internally.
A pattern that implements the Dependency Inversion principle. It involves supplying the dependencies of a class from the outside rather than creating them within the class. This can be achieved through constructor injection, method injection, or property injection.

Goal → Helps in achieving loose coupling between classes and facilitates easier testing, maintenance, and reusability.