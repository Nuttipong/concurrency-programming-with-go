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


