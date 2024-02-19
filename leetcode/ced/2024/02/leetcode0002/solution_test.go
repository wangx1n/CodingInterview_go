package leetcode0002

import (
	"github.com/bytedance/mockey"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	mockey.PatchConvey("", t, func() {
		l1 := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}
		l2 := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}
		ret := addTwoNumbers(l1, l2)
		convey.So(ret, convey.ShouldResemble, &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		})
	})
}
