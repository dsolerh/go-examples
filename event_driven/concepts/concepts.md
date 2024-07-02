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

### Event 
