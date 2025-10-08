## Behavioral Patterns

### How objects communicate and manage workflow.

```sh
| Pattern                     | Used For                                    | Real-world Example                                     |
| --------------------------- | ------------------------------------------- | ------------------------------------------------------ |
| **Observer (Pub/Sub)**      | Event-driven updates                        | Kafka consumers, Redis Pub/Sub, WebSockets             |
| **Strategy**                | Switch between algorithms dynamically       | Retry strategies, payment routing, compression formats |
| **Command**                 | Encapsulate actions as objects              | Job queues (Celery, Sidekiq, RabbitMQ tasks)           |
| **Chain of Responsibility** | Sequential request processing               | Express middleware, Gin middlewares                    |
| **State**                   | Manage state transitions cleanly            | Order/payment workflow, user session states            |
| **Mediator**                | Centralize communication between components | Chat rooms, game engines, event buses                  |
| **Template Method**         | Base workflow with pluggable steps          | Base HTTP handlers, pipeline frameworks                |

```