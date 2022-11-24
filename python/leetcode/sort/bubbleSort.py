def bubble_sort(num_list):
    for i in range(len(num_list) - 1):
        exchange_flag = False
        for j in range(0, len(num_list) - 2):
            if num_list[j] > num_list[j + 1]:
                num_list[j], num_list[j + 1] = num_list[j + 1], num_list[j]
                # temp = num_list[j]
                # num_list[j] = num_list[j + 1]
                # num_list[j + 1] = temp
                exchange_flag = True
        if exchange_flag is False:
            return num_list
    return num_list


if __name__ == "__main__":
    num_list = []
    print("<Bubble Sort>\n Plz input numbers:")
    num_list = [int(x) for x in input().split()]
    num_list = bubble_sort(num_list)
    print(num_list)
