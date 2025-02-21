# Cohesion
- **Definition:** Cohesion refers to the degree to which the elements within a module belong together. It measures how closely related and focused the responsibilities of a single module are.
- **Goal:** Aim for high cohesion to ensure that a module has a single, well-defined task or responsibility. High cohesion makes the module easier to understand, maintain, and reuse.
- **Impact**: High cohesion usually results in better modularity and easier maintenance. It often leads to low coupling as well.
- The more cohesive your classes are, the less coupled they will be
- 
High Cohesion
Example: User Authentication Module

A User Authentication module that includes functions for user login, logout, and password management.

Each function within the module is directly related to the single responsibility of managing user authentication.

The module might include methods like login(userCredentials), logout(), resetPassword(email), etc.

Types of cohesion (worst to best):
Coincidental Cohesion: Unrelated functions put together. Should always avoid
Logical Cohesion: Functions are grouped based on a logical category
Temporal Cohesion: Functions executed at the same time
Procedural Cohesion: Functions that follow a sequence 
Communicational Cohesion: Functions that work on the same data 
Sequential Cohesion: Functions where the output of one is the input of the next
Functional Cohesion: All functions contribute to a single, well-defined task (best form)
