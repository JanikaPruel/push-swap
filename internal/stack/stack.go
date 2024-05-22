package stack

// Реализация стека с необходимыми методами.


// type Stack struct {
//     elements []int
// }

// func New(elements []int) *Stack {
//     return &Stack{elements: elements}
// }

// func (s *Stack) Push(value int) {
//     s.elements = append(s.elements, value)
// }

// func (s *Stack) Pop() (int, bool) {
//     if len(s.elements) == 0 {
//         return 0, false
//     }
//     value := s.elements[len(s.elements)-1]
//     s.elements = s.elements[:len(s.elements)-1]
//     return value, true
// }

// func (s *Stack) IsEmpty() bool {
//     return len(s.elements) == 0
// }

// func (s *Stack) IsSorted() bool {
//     for i := 1; i < len(s.elements); i++ {
//         if s.elements[i-1] > s.elements[i] {
//             return false
//         }
//     }
//     return true
// }
