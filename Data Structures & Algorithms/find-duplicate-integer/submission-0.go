func findDuplicate(nums []int) int {
    // find the meeting point inside the cycle
    slow, fast := nums[0], nums[0]

    for {
        slow = nums[slow]
        fast = nums[nums[fast]]

        if slow == fast {
            break
        }
    }

    // find the start of the cycle
    slow = nums[0]
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}
