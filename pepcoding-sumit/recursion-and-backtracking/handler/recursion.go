package handler

func NewRecursion() *recursion {
	return &recursion{}
}

type recursion struct{}

func (re *recursion) FindFirstOccurrence(nums []int, idx, target int) int {
	// nums := []int{2, 4, 3, 8, 5, 2, 3, 5, 9, 6}
	if idx == len(nums)-1 {
		return -1
	}
	if nums[idx] == target {
		return idx
	}
	return re.FindFirstOccurrence(nums, idx+1, target)
}

func (re *recursion) FindLastOccurrence(nums []int, idx, target, ans int) int {
	// nums := []int{2, 4, 3, 8, 5, 2, 3, 5, 9, 6}
	if nums[idx] == target {
		ans = idx
	}
	if idx == len(nums)-1 {
		return ans
	}
	return re.FindLastOccurrence(nums, idx+1, target, ans)
}

func (re *recursion) FindAllOccurrence(nums, ans []int, idx, target int) []int {
	if nums[idx] == target {
		ans = append(ans, idx)
	}
	if idx == len(nums)-1 {
		return ans
	}
	return re.FindAllOccurrence(nums, ans, idx+1, target)
}
