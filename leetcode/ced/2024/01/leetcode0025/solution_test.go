package leetcode0025

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/bytedance/mockey"
)

func Test_reverseKGroup(t *testing.T) {
	mockey.PatchConvey("", t, func() {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		}
		ret := reverseKGroup(head, 2)
		convey.So(ret, convey.ShouldResemble, &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		})
	})
}
