# Kafka Clone

A simplified implementation of a message broker system, inspired by Apache Kafka. This project was created as a challenge from [Devgym](https://www.devgym.com.br).

## Overview

The Kafka Clone project is a basic message broker system that allows clients to publish and consume messages through a TCP connection. It provides a simplified version of the publish-subscribe model, allowing clients to create topics, publish messages to topics, and consume messages from topics.

## Flow Chart

![Flow Chart](./_docs/flow-chart.png)

## Features

- **Publish-Subscribe Model:** Clients can create topics, publish messages to topics, and consume messages from topics.

- **TCP Communication:** Communication between the server and clients is established over TCP connections.

- **Simple Command Line Interface:** Clients interact with the system using a simple command line interface.

## Getting Started

### Prerequisites

- [Go](https://go.dev/learn) (Golang) installed

- [Make](https://www.gnu.org/software/make/#download) (Optional)

### Usage

#### Server

1. Start the server:

    ```bash
    make server
    ```

2. The server will start listening on `localhost:8080`.

#### Client

1. Start a client:

    ```bash
    make client
    ```

2. Follow the command line prompts to publish or consume messages.

### Commands

- **PUBLISH:** Publish a message to a topic.
- **CONSUME:** Consume messages from a topic.
- **EXIT:** Exit the client application.

## How to Contribute

🎉 Thank you for considering contributing to Kafka Clone! 🎉

The following is a set of guidelines for contributing to this project. Please take a moment to review this document to make the contribution process smooth and effective for everyone.

1 - Fork the repository and clone it to your local machine.

```bash
git clone https://github.com/RianNegreiros/devgym-kafka-clone.git
cd kafka-clone
```

2 - Create a new branch for your contribution.

```bash
git checkout -b feature/your-feature
```

3 - Make your changes and commit them.

```bash
git add .
git commit -m "Add your meaningful commit message here"
```

4 - Push your changes to your fork.

```bash
git push origin feature/your-feature
```

Reporting Bugs

If you find a bug, please open an issue and provide a detailed description of the problem, including any error messages and steps to reproduce.

Suggesting Enhancements

If you have an idea for an enhancement, feel free to open an issue to discuss it. Provide as much detail as possible to help the community understand your suggestion.
