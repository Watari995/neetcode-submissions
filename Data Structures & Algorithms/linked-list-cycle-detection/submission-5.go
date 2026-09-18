/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    valMap := make(map[*ListNode]bool)
	for head != nil {
	  if _, ok := valMap[head]; ok {
	// mapを作って再度同じところに来たらtrueとする
		return true
	  } else {
		valMap[head] = true
	  }
	  head = head.Next
	}
	return false
}
