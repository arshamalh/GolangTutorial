https://www.youtube.com/watch?v=YEKjSzIwAdA

Watch until 27:30 and I realized I'm understanding nothing, so I stopped this file after that.


## Concurrency in detail

- Group code (and data) by identifying independent tasks
- No race conditions
- No deadlocks
- More workers = faster execution

## Communicating Sequential Processes (CSP)

- Each process is built for sequential execution
- Data is communicated between processes via channels, No shared state! (Sharing can cause deadlocks and race conditions)
- Scale by adding more of the same.

## Go concurrency tool set

- Go routines
- Channels
- Select
- Sync Package

Channels are the backbone of CSP

Think of a bucket chain consist of 3 components: sender, buffer (optional), receiver

Image: concurrency 1 & 2

- Blocking breaks concurrency
- Blocking can lead to deadlocks
- Blocking can prevent scaling

## Closing channels
Close sends a special “closed” message
The receiver will at some point see “closed” and there is nothing more to do.
If you try to send more data to a closed channel: panic!

Think about channels standing between manager and worker
Manager sends work to worker
Manager determines when there is no work
So, sender always close the channel.

Image: concurrency 3

```go
c := make(chan int)
close(c)
fmt.Println(<-c)
// out: 0, false
```
0 as it is the zero value of int
false because of “there is no more data” like maps key access (ok).

————
## Select
- Like a switch statement on channel operations
- The order of cases doesn’t matter at all
- There is a default case too
- The first non-blocking case is chosen (send and/or receive)


Making channels non-blocking:
```go
func TryReceive(c <- chan int) (data int, more, ok) {
select {
case data, more = <- c:
	return data, more, true 
} 
default: // Processed when c is blocking 
	return 0, true, false 
}
```

```go
func TryReceiveWithTimeout(c <- chan int, duration time.Duration) (data int, more, ok) {
select {
case data, more = <- c:
	return data, more, true 
} 
case <- time.After(duration): // Processed when c is blocking 
	return 0, true, false 
} 
```

Shape your data flow:
- Channels are streams of data
- Dealing with multiple streams is the true power of select

Image: concurrency 4

## Fanout
```go
func Fanout(In <- chan int, OutA, OutB chan int) {
    for data := range In {
        select {
            case OutA <- data:
            case OutB <- data:
        }
    }
}
```

## Funnel

!!: My own implementation, I'm not quiet sure about it.
```go
func Funnel(InA, InB <- chan int, Out chan int) {
    var data int
    var moreA, moreB bool
    // Can't range because of iterating over two channels
    for {
        select { // Receives from first non-blocking
            case data, moreA = <- InA:
            case data, moreB = <- InB:
        }

        if !moreA || !moreB { 
            return
        }
        Out <- data
    }
}
```

## Turnout
```go
func Turnout(InA, InB <- chan int, OutA, OutB chan int) {
    var data int
    var more bool
    // Can't range because of iterating over two channels
    for {
        select { // Receives from first non-blocking
            case data, more = <- InA:
            case data, more = <- InB:
        }

        // Bad solution, 
        // if one channel is closed, 
        // then other channel will have data but we won't read it because we'll be out of this function.
        if !more { 
            return
        }

        select { // Sends to first non-blocking
            case OutA <- data:
            case OutB <- data:
        }
    }
}
```

```go
func Turnout(Quit <- chan int, InA, InB, OutA, OutB chan int) {
    var data int
    var more bool
    // Can't range because of iterating over two channels
    for {
        select { // Receives from first non-blocking
            case data <- InA:
            case data <- InB:
            case <- Quit: // Sender is asking receiver to close the channel, this is not actually an anti-patterns
                close(InA)
                close(InB)

                Fanout(InA, OutA, OutB) // Flush the remaining data
                Fanout(InB, OutA, OutB)
                return
        }
    }
}
```

## Where channels fail
- Channels can cause deadlocks when not used properly.
- Channels pass copies around, which can impact performance.
- Passing pointers to channels can create race conditions (you may change the underlying value of the pointer after sending it).
- Don't make "naturally shared" structures like caches or registries

## Three shades of code
- **Blocking**: Your program may get locked up (for undefined time)
- **Lock free**: At least one part of your program is always making progress
- **Wait free**: All parts of your program are always making progress

## Atomic operations
- sync.atomic package
- Store, Load, Add, Swap, and CompareAndSwap
- Mapped to thread-safe CPU instructions
- These instructions only work on integer types
- Only about 10 - 60x slower than their non-atomic counterparts

## Spinning CAS
- You need a state variable and a free constant
- Use CAS (CompareAndSwap) in a loop:
  - If state is not gree: try again until it is
  - If state is free: set it to something else
- If you managed to change the state, you "own" it.

## Spinning CAS: The most basic lock we can write
```go
type SpinLock struct {
    state *int32
} 

const free = int32(0)

func (l *SpinLock) Lock() {
    for !atomic.CompareAndSwapInt32(l.state, free, 42) { // 42 or any other number except 0
        runtime.Gosched() // Poke the scheduler
    }
}

func (l *SpinLock) Unlock() {
    atomic.StoreInt32(l.state, free) // Once atomic, always atomic.
}
```

## Ticket Storage
- We need an indexed data structure, and ticket and a done variable
- A function draws a new ticket by adding 1 to the ticket
- Every ticket number is unique as we never decrement
- Treat the ticket as an index to store your data
- Increase done to extend the "ready to read" range

```go
type TicketStore struct {
    ticket *uint64
    done *uint64
    slots []string // Imagine it has a infinite length
}

func (ts *TicketStore) Put(s string) {
    t := atomic.AddUint64(ts.ticket, 1) - 1             // Draw a ticket
    slots[t] = s                                        // Store your data
    for !atomic.CompareAndSwapUint64(ts.done, t, t+1) { // Increase done
        runtime.Gosched() 
    }
}

func (ts *TicketStore) GetDone() []string {
    return ts.slots[:atomic.LoadUint64(ts.done)+1] // read up to done
}
```


