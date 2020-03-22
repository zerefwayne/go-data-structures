## Stack

#### Usage

```
package main

import (
    "github.com/zerefwayne/go-data-structures/stack
)

func main() {

    newStack := new(stack.Stack)
    
    size, _ := newStack.Push(34)
    
    fmt.Println(size)
    
    size, _ = newStack.Pop()
    
    fmt.Println(size)
    
    size, err := newStack.Pop()
    
    if err != nil {
        fmt.Println(err.Error())
    }

    newStack.Push(24)
    newStack.Push(45)

    for !newStack.Empty() {
        top := newStack.Top()
        newStack.pop()
        fmt.Println(top)
    }   

}
```