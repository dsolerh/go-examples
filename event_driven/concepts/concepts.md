## Event Driven Concepts

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

-   Shares the same design that the [Event notifications](#event-notifications).
-   Some services may receive more data than needed.
