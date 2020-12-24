### lockless-counter

#### A dead simple threadsafe lockless-counter

#### Usage

```

	counter := NewLocklessCounter()
	
	taken := counter.Take() // 0
	taken  = counter.Take() // 1
	taken  = counter.Take() // 2

```


