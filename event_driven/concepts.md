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

![image](https://github.com/dsolerh/go-examples/assets/55502191/85a4f87f-043d-42a0-b72e-3d36242b4abf)

### Event 
