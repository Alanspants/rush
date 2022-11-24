def quick_sort(i, j, num):
    if i > j: return num
    pivot = num[i]
    low = i
    high = j
    while i < j:
        while i < j and num[i] <= pivot:
            i += 1
        num[i], num[j] = num[j], num[i]
        while i < j and num[j] >= pivot:
            j -= 1
        num[i], num[j] = num[j], num[i]
    quick_sort(low, i - 1, num)
    quick_sort(j + 1, high, num)
    return num


if __name__ == "__main__":
    num_list = []
    print("<quick sort>\nPlz input nums: ")
    num_list = [int(x) for x in input().split()]
    num_list = quick_sort(0, len(num_list) - 1, num_list)
    print(num_list)
