<p align="center">
  <img  alt="docker for dummies" height="128px" width="128px" src="https://miro.medium.com/max/1200/1*i2skbfmDsHayHhqPfwt6pA.png">
</p>

# Low-Level System Design in Go

Welcome to the **Low-Level System Design in Go** repository! This repository contains various low-level system design problems and their solutions implemented in Go. The primary aim is to demonstrate the design and architecture of systems through practical examples.

## Table of Contents

- [Overview](#overview)
- [Parking Lot System](#parking-lot-system)
- [Elevator System](#elevator-system)
- [Library Management System](#library-management-system)

## Overview

Low-level system design involves understanding the core concepts of system architecture and designing scalable, maintainable, and efficient systems. This repository will try to cover solutions of various problems and scenarios using Go.

## Parking Lot System

The first project in this repository is a **Parking Lot System**. This system simulates a parking lot where vehicles can be parked and unparked. It demonstrates:

- Singleton design pattern for managing the parking lot instance.
- Handling different types of vehicles (e.g., cars, trucks).
- Parking space management across multiple floors.
- Payment processing for parked vehicles.

### Features

- Add and remove vehicles from the parking lot.
- Display the availability of parking spots.
- Calculate parking charges based on the duration of stay.

## Elevator System

The second project in this repository is an Elevator System. This system simulates an elevator control system in a multi-floor building. It demonstrates:

- Multi-Elevator Management: Manages multiple elevators within a single building.
- Request Handling: Processes floor requests and assigns the most suitable elevator based on proximity and current direction.
- Direction Control: Dynamically changes elevator direction based on requests.
- Concurrent Operations: Uses Go routines and synchronization to handle multiple elevator requests simultaneously.

### Features

- Request an elevator from any floor, specifying the desired direction (up or down).
- Assign the nearest elevator to respond to requests.
- Manage elevator movement and optimize direction changes based on requests and destinations.

## Library Management System

The third project in this repository is a Library Management System. This system simulates a library where members can borrow and return books. It demonstrates:

- Singleton design pattern for managing the library instance.
- Book management with support for multiple copies of each book.
- Member management, including borrowing history and borrowing limits.
- Concurrency control to handle multiple borrow and return requests simultaneously.

### Features

- Add and remove books from the library collection.
- Allow members to borrow and return books, with automatic status updates.
- Check book availability and member borrowing quoto limit.
- Maintain a history of borrowed books for each member.
