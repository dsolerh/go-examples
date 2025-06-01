Design and implement a Frequency stack.
A Frequency Stack is a data structure that mimics the behavior of a regular stack but with a twist: when you perform a pop operation, instead of simply removing the last element added, it removes the element that appears most frequently in the stack. If there’s a tie (i.e., multiple elements have the same highest frequency), it pops the one that was most recently pushed.
The stack should accept integer numbers.
The following methods should be implemented
void push(value: int): Adds the value to the stack
int pop(): Removes the most frequent item from the stack and returns it. If there’s a tie, returns the most recently pushed

```
FreqStack = new FreqStack();

FreqStack.push(5);
FreqStack.push(7);
FreqStack.push(5);
FreqStack.push(7);
FreqStack.push(4);
FreqStack.push(5);

FreqStack.pop();   // returns 5 
FreqStack.pop();   // returns 7 
FreqStack.pop();   // returns 5
FreqStack.pop();   // returns 4 
FreqStack.pop();   // returns 7
FreqStack.pop();   // returns 5
```

```
map_values: {
    5: {2, 00003}, [00000, 00002]
    7: {2, 00002}, [00001, 00002]
    3: {1, 00004}, [00004]
}

map_values: {
    5: {2, 00003}, [00000]
    7: {2, 00003}, [00001, 00002]
    3: {1, 00004}, [00004]
}

map_values: {
    5: {2, 00003}, [00000]
    7: {2, 00002}, [00002]
    3: {1, 00004}, [00004]
}

map_values: {
    5: {2, [00000, 00003]},
    7: {2, [00001, 00002]}, 
    3: {1, [00004]},
}

```