```go
// We have shared bank account between Arsham and Abtin
// Account balance is 1000
var balance = 1000

// Exactly at the same time
 
// Arsham deposit 10 to 1000
balance = balance + 10
// Abtin withdraw 10 from 1000
balance = balance - 10

// Arsham watches the balance and it's 990!
1010
1000
```