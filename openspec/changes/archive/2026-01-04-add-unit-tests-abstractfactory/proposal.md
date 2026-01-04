# Change: Add unit tests for AbstractFactory pattern

## Why
The current repository demonstrates the AbstractFactory design pattern through executable examples but lacks unit tests. Adding unit tests will validate the pattern implementation, improve confidence, and serve as living documentation for future developers.

## What Changes
- Create `AbstractFactory/abstractfactory_test.go` containing suitable unit tests.
- No runtime changes; test file is compile‑time only.
- No breaking changes to existing code.

## Impact
- The new test file will live alongside the main program in the `AbstractFactory` directory.
- Existing behaviour remains unchanged; only provides additional documentation.
