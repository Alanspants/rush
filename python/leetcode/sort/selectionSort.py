import sys


def selection_sort(nums):
    for i in range(0, len(nums)):
        temp_min = sys.maxsize
        temp_min_index = 0
        for j in range(i, len(nums)):
            if nums[j] < temp_min:
                temp_min = nums[j]
                temp_min_index = j
        nums[i], nums[temp_min_index] = nums[temp_min_index], nums[i]
    return nums


if __name__ == "__main__":
    num_list = []
    print("<Selection Sort\nPlz input nums: ")
    num_list = [int(x) for x in input().split()]
    num_list = selection_sort(num_list)
    print(num_list)