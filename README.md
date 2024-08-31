# Poc nats
Process messages with Nats

Simple messaging design
NATS makes it easy for applications to communicate by sending and receiving messages. These messages are addressed and identified by subject strings, and do not depend on network location.

Data is encoded and framed as a message and sent by a publisher. The message is received, decoded, and processed by one or more subscribers.

![Arch](Captura%20de%20tela%202024-08-31%20074905.png)

Reference: https://docs.nats.io/nats-concepts/what-is-nats#simple-messaging-design

Nats: https://natsbyexample.com/

# Usage

Run nats server:
`docker compose up nats`

Run published:
`docker compose up published`

Run consumer
`docker compose up consumer`
