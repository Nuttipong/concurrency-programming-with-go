## Concurrency Programming With Go

- How to use goroutines
- How to use waitgroup and mutex
- How to use sick package and channels

#### Concurrency and Parallelisym
    What is concurrency? for example, going to grocery store. So you've got to do your weekly grocery shopping. There's really a few ways that you can do that. The first is you only had to buy a single item say apples, So you had a single task to do, and you can execute that task easily on your own. The next thing, and probably what happends more often is that you go to the grocery store with a list of things to purchase. So in that case, you can concurrently get everyting on your list because you got multiple things. for example, you got apples, eggs, bread, and milk. Well, you can get any one of those items at any time. you can get the apples first, then get the eggs or you can get the eggs and then the apples.
    The point is, Yoy've got multiple tasks and you can get those in any order. So this is mean working concurrently. You've got multiple tasks that you can work one thing at a time but we had multiple things to get, and that concurrency.
    - Single task -> have a single task<br/>
    - Concurrency -> have multiple tasks<br/>
    - Parallelism ->  execute multiple tasks simultaneously<br/>
    (Note, you cannot have parallelism without concurrency becuase you need more than one thing to do to effectively act on that)

#### goroutines
- A thread represents the ability for operating system to run a task. Thread have own execution stack. So there's a list of instructions that are assigned to that thread, and that thread start to working on it. It's got own memory, and it has its own call. So it's know what it's doing as it's doing it. fixed stack space around 1 MB. Thread managed by operating system.
- A goroutines represents the ability for operating system to run a task.
  - have own execution stack
  - stack space 2KB(But if they do need to grow, the goroutines has ability to increase its stack space to take more and more operations, as needed)
  - goroutines managed by Go runtime. Now go is no difference than any other programming langauges. It has a threads as well, but go managed those threads for us. So when create goroutines, the go runtime is gonna manage mapping those goroutines onto OS threads. however, the runtime provides an interface, allowing a relatively small number of threads to work with all of those goroutines then the runtime schedules is goroutines under the threads as they have thinga to do.

#### sync package(golang.org/pkg/sync), WaitsGroups, and Mutexes 
- WaitsGroups(coordinating tasks), sync.WaitGroups, A WaitGroups waits for a collection of goroutines to finish.
- Mutexes(shared memory, protect memomry to shared between multiple goroutines)


