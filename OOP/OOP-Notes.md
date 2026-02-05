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
abstract classes are meant to be subclassed, not instantiated.


Encapsulation is kind of protection wrapper around data and behavior.
Encapsulation is the process of keeping property and method inside abstracted object(class) and controlling access.
 Encapsulation allows us to protect the object's internal state and provide controlled access to it. Additionally, encapsulation enables us to include logic within setter accessors, such as validating that the age is not set to a negative value.
 Encapsulation protects data integrity, enforces validation, and improves code maintainability.

Polymorphism allows objects of different types to be treated as objects of a common supertype. This is achieved through method overriding and interfaces.
Polymorphism enables a single interface to be used for a general class of actions, allowing for flexibility and the ability to switch between different implementations.
Polymorphism allows objects to be treated as instances of their parent class rather than their actual class. This provides a way to use a single interface to represent different types of objects.


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

# Subtype Polymorphism
Subtype polymorphism is the ability of a single function or method to work on objects of different types. It relies on inheritance and interfaces in OOP to achieve this flexibility.
- Common behavior is defined in a base class and reused by multiple subclasses.
- New subclasses can be added without changing existing code.






