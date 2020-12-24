package test

// 摩尔投票算法是用于找出数组中出现次数超过n/2的元素，
// 其核心思想在于遍历过程中不同元素之间两两抵消，
// 由于一个数组中，出现次数超过n/2最多只有一个，
// 那么遍历结束时，未被抵消掉的即是出现次数超过n/2的元素。
// 算法中核心变量有两个，一个是maj，用来保存目前未被抵消的元素，
// 一个是count，用来反映抵消maj元素所需要的元素数目。
// 也就是说，在数组中maj元素出现一次，count就自加一次，
// 如果出现了和maj不同的元素，说明maj可被抵消一次，count就自减一次，
// 如果count减为0，也就说明maj元素已经被抵消完了，maj元素也不可能是出现次数超过n/2次的元素，
// 因此就更新maj。显然，maj初始化可以赋予任何值，count初始化应当为0，程序如下：
func majorityElement(nums []int) int {
    major := 0
    count := 0

    for _, num := range nums {
        if count == 0 {
            major = num
        }
        if major == num {
            count++
        } else {
            count--
        }
    }
    
    return major
}