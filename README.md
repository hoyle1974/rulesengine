# rulesengine
Logic library for a game server rules engine

I'll be uisng this repo to play around with creating a rules engine for something like a game server.  The basic premise is
this library has gameobjects.

A gameobject can:

- Contain other game objects
- Be contained by game objects
- Apply actions, that may affect (in order) themselves, internal game objects, or parent game objects

This library doesn't do anything with networking.
