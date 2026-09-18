/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	hashMap := make(map[*ListNode]bool)
	for head.Next != nil {
		if _, ok := hashMap[head.Next]; ok {
			return true
		}
		hashMap[head.Next] = true
		head = head.Next
	}
	return false
}
