# Coupling
- Degree of how much one module is dependent on another module
- We can’t remove coupling, we can simply make it loose

# Low Coupling
Example: E-commerce Application with Order Processing and Payment Modules

Order Processing Module: This module handles order creation, order updates, and order tracking. It doesn’t directly interact with the payment process; instead, it sends a message or event when an order is ready for payment.

Payment Module: This module handles payment processing, payment confirmation, and refund handling. It listens for the event or message from the Order Processing module to start processing the payment.

Ways to acheive decoupling:
- Dependency Inversion/Injection
- Event-Oriented Decoupling
  - In a user registration system, when a new user is registered, an event is published. Subscribers (like an email notification service) listen for the event and send a welcome email.
- Queue-Based Decoupling
  - In an order processing system, an order service places new orders in a message queue. A payment service reads the orders from the queue and processes payments independently.