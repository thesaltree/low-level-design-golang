
# Low-Level System Design in Go

Welcome to the **Low-Level System Design in Go** repository! This repository contains various low-level system design problems and their solutions implemented in Go. The primary aim is to demonstrate the design and architecture of systems through practical examples.

## Table of Contents

- [Overview](#overview)
- [Parking Lot System](#parking-lot-system)
- [Elevator System](#elevator-system)
- [Library Management System](#library-management-system)
- [Vending Machine System](#vending-machine-system)
- [Social Media Platform](#social-media-platform)
- [Meeting Scheduler](#meeting-scheduler)
- [Contributing Guidelines](#contributing-guidelines)

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

## Vending Machine System

The fourth project in this repository is a Vending Machine System. This system simulates a vending machine that accepts payments, dispenses products, and manages inventory. It demonstrates:

- State Management: Uses the state pattern to manage vending machine states (e.g., waiting for money, product selection, dispensing).
- Inventory Management: Keeps track of product stock and availability.
- Transaction Processing: Handles money insertion, change return, and product dispensing.
  
### Features
- Insert money and select a product for purchase.
- Dispense product if sufficient funds are provided, or return change if funds are insufficient.
- Automatically update inventory levels upon each transaction.
- Handle various states in the vending machine’s workflow, including error handling for out-of-stock items or insufficient funds.

## Social Media Platform

The fifth project in this repository is a **Social Media Platform**. This system simulates a basic social media platform where users can connect, create posts, interact with posts, and view activity feeds. It demonstrates:

- **Facade Pattern**: Simplifies user interaction with various social media functionalities (e.g., managing users, posts, friendships) through a unified interface.
- **User and Post Management**: Allows the creation and management of users, posts, comments, and likes, encapsulated in dedicated manager classes.
- **Concurrency Control**: Manages concurrent operations such as posting, commenting, and friend requests using synchronization techniques.
  
### Features

- **User Registration and Profile Management**: Users can register on the platform with profile details, add a profile bio, and edit their information.
- **Posting and Feed System**: Users can create posts and see posts from their friends in their feed.
- **Friendship System**: Users can send and accept friend requests, and posts from friends are included in their feeds.
- **Post Interactions**: Users can like and comment on posts from other users. The number of likes and all comments are viewable on each post.
- **Feed Management**: Users can view a personalized feed of posts from their friends, sorted to display recent posts first.
  
### Implementation Highlights

- **Concurrency with RWMutex**: Manages concurrent reads and writes to shared data (such as posts and users) to avoid race conditions.
- **Error Handling**: Provides informative error messages for actions such as attempting to interact with nonexistent users or posts.
- **Encapsulation of Features**: Each feature (users, posts, friendships) is encapsulated in a separate manager class, promoting modularity and ease of maintenance.


## Meeting Scheduler

The sixth project in this repository is a Meeting Scheduler, a system that simulates meeting scheduling for users. It demonstrates the following concepts:

- Observer pattern for notifying users about scheduled meetings.
- Book meeting for muliple users in particular time slot, so that room slot does not collide with each other.
- Room calendar management, tracking when a room is booked or available.
- Concurrency control, preventing scheduling conflicts in real time.

## ATM

The seventh project in this repository is a ATM machine, a system that simulates ATM. It demonstrates the following concepts:

- State design pattern for different state of ATM.
- Chain of responsibility principle for withdrawing cash.
- Template pattern for handling error of each place in one place.
- Concurrency control, preventing atm to conflict.

### Features

- Book and cancel meetings seamlessly.
- Prevent double booking, allowing reservations only if the time slot is available.
- Capacity-based booking, ensuring a room can only be booked if it has enough seats for all attendees.
- Automated notifications to inform users about meeting bookings or cancellations.
- Fetching all free rooms for a specific time slot.

## Contributing Guidelines
Contributions to expand this repository with more low-level system design problems and solutions are welcomed. Here’s how you can contribute:

### Steps to Contribute

- Choose a System Design Problem.
- Select a problem to implement, feel free to propose it in an issue.
- Open a new issue in the repository to let maintainers know about the system design problem you plan to add. Provide a brief overview of the problem and your approach.
- Follow Repository Standards. Use Go for the implementation. Follow the project structure and coding style used in the repository.
- Include a clear README section for the problem with the following details: problem description, key features, design patterns used, concurrency or error handling highlights (if any)
- Fork the repository, make your changes, and submit a pull request.
- Ensure your code is clean, well-documented, and tested.
- Your pull request will be reviewed by the maintainer. After addressing feedback, it will be merged into the repository.

### Contribution Tips
- Break the problem into smaller components and implement them incrementally.
- Document your solution clearly to help others understand your implementation.
  
Thank you for contributing to Low-Level System Design in Go!
