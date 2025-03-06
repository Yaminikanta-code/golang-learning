# Memory Management

Memory allocation and deallocation happens automatically in Go.

new and make are used to allocate memory.

## new()

- Allocate memory but not initialize.
- You will get a memory address which we can reference using a pointer.
- zeroed storage. In zereoed storage, we can't put any data initially.

## make()

- Allocate memory and initialize.
- You will get a memory address which we can reference using a pointer.
- non-zeroed storage. In non-zeroed storage, we can put any data initially.

Garbage collection happens automatically in Go.

Anything which is oout of scope or becomes nil will be garbage collected.

## runtime package

https://pkg.go.dev/runtime

runtime is a package in itself which gives low level imformation about memory. (read more about it here: https://pkg.go.dev/runtime)
