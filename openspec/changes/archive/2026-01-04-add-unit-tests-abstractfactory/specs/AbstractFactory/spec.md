## ADDED Requirements
### Requirement: Unit tests defined for AbstractFactory
The system MUST provide unit tests confirming that the AbstractFactory implementation behaves as intended.

#### Scenario: Verify concrete product creation
- **WHEN** a client requests a concrete component via the factory
- **THEN** the returned concrete component satisfies its interface and contains expected data
