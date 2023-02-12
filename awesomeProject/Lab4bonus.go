package main

// Bonus task for 4th week (Creating LIFO struct(stack))
import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top  *Node
	Size int
}

func (s *Stack) Empty() bool { // Simple function like in other languages like c++
	if s.Size != 0 {
		return false
	} else {
		return true
	}
}

func (s *Stack) Push(a int) {
	newNode := &Node{Value: a, Next: s.Top} // New node with value 'a' which will be on top of previous node
	s.Top = newNode
	s.Size++
}

func (s *Stack) Pop() int {
	if s.Empty() == true {
		return -1 // We consider that stack contains only natural numbers, so -1 means error
	}
	prevTop := s.Top     // prevTop is deleting element
	s.Top = prevTop.Next // Top element now must be next element after prevTop
	s.Size--
	return prevTop.Value //
}

func (s *Stack) Peek() int {
	if s.Empty() == true {
		return -1
	}
	return s.Top.Value
}

func (s *Stack) Clear() {
	s.Top = nil
	s.Size = 0
}

func (s *Stack) Contains(val int) bool {
	node := s.Top
	for node != nil { // We just linearly go through stack to find 'val' node
		if node.Value == val {
			return true
		} else {
			node = node.Next
		}
	}

	return false
}

func (s *Stack) Increment() {
	node := s.Top
	for node != nil {
		node.Value++
		node = node.Next
	}
}

func (s *Stack) Print() {
	node := s.Top
	for node != nil {
		fmt.Print(node.Value, " ") // Simply printing each element with space in between
		node = node.Next
	}
	fmt.Println() // Going to next line to prevent confusion
}

func (s *Stack) PrintReverse() {

	if s.Empty() == true { // Implementation with recursion
		return
	}
	s.PrintReverseRecursion(s.Top)
	fmt.Println()

	/*node := s.Top // Implementation without recursion
	ret := ""     // String that will contain our output
	for node != nil {
		ret = strconv.Itoa(node.Value) + " " + ret // Adding each new value in front of values added before
		node = node.Next
	}
	fmt.Println(ret)*/
}

func (s *Stack) PrintReverseRecursion(node *Node) { // Recursion function for reverse printing
	if node.Next == nil { // If next element doesn't exist we just print current value
		fmt.Print(node.Value, " ")
		return
	}
	s.PrintReverseRecursion(node.Next) // Go to next node first to print it before current one
	fmt.Print(node.Value, " ")
}

func main() {
	st := &Stack{}
	fmt.Println(st.Empty())
	st.Push(1)
	fmt.Println(st.Peek())
	fmt.Println(st.Pop())
	fmt.Println(st.Peek()) // Returning '-1' meaning that stack is empty
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Contains(1)) // Returning 'true' because stack contains '1'
	fmt.Println(st.Contains(4))
	st.Print()
	st.PrintReverse()
	st.Increment()
	st.Print()
	st.PrintReverse()
	st.Clear()
	fmt.Println("Size of stack: ", st.Size)
}
