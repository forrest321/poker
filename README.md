# Poker ♠️♥️♦️♣️🃏🃏

This project is a Go implementation of a poker game system. It currently includes the foundational elements for a Texas Hold'em poker game, with potential for extension to other poker variants.

## Project Structure

The project is organized into several packages, each focusing on a specific aspect of the game:

### card Package

- **Purpose**: Manages individual playing cards.
- **Files**:
  - `definition.go`: Defines the `Card` struct and constructors.
  - `types.go`: Defines the `Suit` and `Value` types, including support for both numeric and face card values.

### deck Package

- **Purpose**: Manages a deck of playing cards.
- **Files**:
  - `definition.go`: Defines the `Deck` struct and a constructor for creating a new deck.
  - `management.go`: Includes methods for shuffling and dealing cards from the deck.

### player Package

- **Purpose**: Manages player-specific information and actions.
- **Files**:
  - `definition.go`: Defines the `Player` struct and a constructor for creating new players.
  - `management.go`: Contains methods for managing player actions such as betting, folding, and hand management.

### betting Package

- **Purpose**: Manages the betting aspect of the game, including different betting actions and pot management.
- **Files**:
  - Various interface definitions for standard and advanced betting functionalities.

### holdem Package (Planned)

- **Purpose**: To implement the specific game logic and rules for Texas Hold'em.

## Current Features

- A complete representation of a standard deck of playing cards, including suits and values.
- Functionality to shuffle and deal cards from the deck.
- Basic player management, including tracking of player's chips and hand.
- An extensible betting system with interfaces for standard and advanced betting actions.

## Upcoming Features

- Implementation of Texas Hold'em gameplay logic.
- Integration of the betting system with game rules.
- Multi-round game management with support for blinds and betting rounds.

## Usage

To be detailed as more features are implemented and the system becomes operational.
