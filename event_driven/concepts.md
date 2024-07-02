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
![image](https://github.com/dsolerh/go-examples/assets/55502191/1580cfb6-0df6-4239-af9f-87f4b1e47665)

### Event 
