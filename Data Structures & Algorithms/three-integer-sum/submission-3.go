func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    lenNums := len(nums)

    finalVals := make([][]int, 0)

    for i := 0; i < lenNums-2; i++ {
        left := i+1
        right := lenNums-1

        if i == 0 || (nums[i] != nums[i-1]) {


            for left < right {
                sum := nums[i] + nums[left] + nums[right]

                if sum == 0 {
                    if left == i+1 || (nums[left] != nums[left-1]) {
                        if right == lenNums-1 || (nums[right] != nums[right+1]) {
                            finalVals = append(finalVals, []int{nums[i], nums[left], nums[right]})
                        }
                    }
                    left++
                    right--
                } else if sum < 0 {
                    left++ 
                } else {
                    right--
                }

            }

        }

    }

    return finalVals
}
