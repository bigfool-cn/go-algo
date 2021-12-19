package _6_linkedlist

/*
判断字符串是否是回文串
回文串：正读和反读一样的字符串
*/

/*
开辟一个栈空间存储前半段, 后半段再与前半段比较
时间复杂度O(n)
空间复杂度O(n)
*/
func IsPalindrome(l *LinkedList) bool {
	llen := l.length
	if llen == 0 {
		return false
	}

	if llen == 1 {
		return true
	}

	lslice := make([]string, 0, llen/2)
	cur := l.head
	for i := uint(1); i < llen; i++ {
		cur := cur.next
		if llen%2 != 0 && i == (llen/2+1) {
			continue
		}
		if i <= llen/2 {
			lslice = append(lslice, cur.GetValue().(string))
		} else {
			if lslice[llen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}

	return true
}

/*
找到中间节点，将前半段链表转置，再从中间往左右依次遍历比较，最后复原链表
时间复杂度O(n)
空间复杂度O(1)
*/
func IsPalindrome2(l *LinkedList) bool {
	llen := l.length
	if llen == 0 {
		return false
	}

	if llen == 1 {
		return true
	}

	var isPalindrome bool = true

	step := llen / 2
	var pre *ListNode = nil
	cur := l.head.next
	next := l.head.next.next
	// 转置前半段链表
	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}

	mid := cur

	var left, right *ListNode = pre, nil
	if llen%2 != 0 {
		right = mid.GetNext()
	} else {
		right = mid
	}

	// 从中间左右依次比较
	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	// 复原链表
	cur = pre
	pre = mid
	for cur != nil {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre
	return isPalindrome
}
