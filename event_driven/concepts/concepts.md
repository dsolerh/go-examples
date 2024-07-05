## Event Driven Concepts

### Core

In essence all event architecture is simply a producer/consumer in a queue. There can be many producers and many consumers. The next diagram represents this.

![image](./assets/proucer_consumer.png)

### Main Benefits

-   **Resilincy/Fault tolerance**: In a P2P comunication both the caller and the callee need to be available and not produce an error, otherwise the error will propagate and the entire operation will fail. Sometimes this can happen in a long chain of call, then the error end up reflecting in a different service than the one were was originally created, this makes it harder to find the root cause, increasing the cost of maintenance. In a EDA the different services are decoupled by the event broker, if one fails it's not reflected on the other one.
-   **Agility**: The different teams can work independently (to a certain degree) from each other as an extra benefit of the services being decoupled. New features can be experimented with more freedom.
-   **User Experience (UX)**: The users of services built with EDA have instant feedback/notification of the actions they perform or they want to be aware of.
-   **Analytics and Auditing**: Provided that all the important events are being stored the app will have the records to reconstruct action taken on it. Also good for gathering **Bussines Inteligence (BI)**.

### Chalenges of EDA

-   **Eventual Consistency**: Comes from being distributed, any change in the whole application state may take some time to be visible. The user may query some data that may not be fully acknoledge by the system, and get back information that is outdated.
-   **Dual writes**: local changes need to ensure to be transmitted to the other components of the system, otherwise can create inconsistency. There's a solution for this known as the **Outbox** pattern.
-   **Distributed and Asynchronous Workflows**: It's a problem related to the **Eventual Consistency**, in the sense that the UX team will have to work around this limitation to prevent the user frustration on no inmediate result. Not only the UX needs to account for this but also the services have to design an efficient comunication strategy to aliviate the issue.
    -   **Components Collaboration**: There are two main patterns to manage workflows:
        -   **Choreography**: Components know about the work they need to perform.
        -   **Orchestration**: Each component does their task whithout knowing of the overall, and are managed by an orchestrator.
-   **Debuggability**: A service produces an event but does not know if the event is going to be processed, and therefore can only log one side of the story, this makes it harder to debug an operation in the system.

### Event notifications

-   Tipically carries a minimun state.
-   An identifier for the event.
-   The components notified can decide if need more information and request it.
-   Payload example:

```go
type PaymentReceived struct {
    PaymentID string
}
```

-   Design example:

![image](./assets/event_notification.png)


### Event-carried state transfer

-   Similar to REST, but with a push model instead of a pull.
-   Contains more data than an [Event notifications](#event-notifications).
-   There's less need to request extra info from the service who originated the event.
-   Payload Example:

```go
type PaymentReceived struct {
    PaymentID   string
    CustomerID  string
    OrderID     string
    Amount      int
}
```

-   Design example:

![image](./assets/event_carried_state_transfer.png)

-   Some services may receive more data than needed.


### Event sourcing

-   The event won't contain the full data. Instead it'll be a collection of events in time that can be aggregated to obtain the final state.
-   Events are stored in an **event store**, instead of being sent to the other components.
-   There's no specific payload for this kind of events.
-   Design Example:

![image](./assets/event_sourcing.png)


### Queues

-   It's a `first-in-first-out` **FIFO** struct. In events domain this holds true __"most of"__ the cases.
-   Can be refered in a variety of terms: bus, channel, stream, topic, etc..

#### Message Queues

-   It's defining characteristic is it's lack of event retention. All of it's events have a limited lifetime.
-   It's useful for simple **publisher/subscriber (pub/sub)**, when the subscribers can retrieve the events quickly enough (Otherwhile they'll be lost).

#### Event Streams

-   In essence it's a **Message Queue** with retention.
-   The data is persisted and therefore it grows over time.
-   Needs some config over what's the max amount data to store, and how to delete/archive old data.

#### Event Stores

-   It's an append-only repository for events.
-   It could be compose of millions of individual **Event Streams**.
-   Provides **optimistic concurrency** controls to maintain a strong consistency on the **Event Streams**.
-   Not typically used for message communication.
-   Commonly used with [Event sourcing](#event-sourcing) to track changes to entities.


### Producers

-   Send event over some queue.
-   May include additional metadata with the event.
-   Does not know about the consumers of the events.

### Consumers

-   Subscribe to consume event from a source.
-   Can be organized into groups to share the load.
-   Can start reaing events from different points:
    -   The beginning of the stream.
    -   The new event from when they joined the stream.
    -   A cursor to pick up from where they left the stream.

## Hexagonal Architecture
