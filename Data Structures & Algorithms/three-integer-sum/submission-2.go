func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    lenNums := len(nums)

    finalVals := make([][]int, 0)

    for i := 0; i < lenNums-2; i++ {

        if (i > 0 && nums[i] != nums[i-1]) || i == 0 {
            // add the 2 pointer logic here

            for j := i+1; j < lenNums-1; j++ {
                if (j > i+1 && nums[j] != nums[j-1]) || j == i+1 {
                    for k := lenNums-1; k > j; k-- {
                        if (k < lenNums-1 && nums[k] != nums[k+1]) || k == lenNums-1 {
                            if nums[j] + nums[k] + nums[i] == 0 {
                                finalVals = append(finalVals, []int{nums[i], nums[j], nums[k]})
                            }
                        }
                    }
                }
            }

        }

    }

    return finalVals
}
