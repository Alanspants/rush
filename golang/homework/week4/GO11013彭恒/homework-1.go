package main

func main() {
	//s1 := make([]int,3,4)
	//s2 := append(s1,1)

	// s1 切片，元素为3个，容量为4，元素默认都是int类型的零值，[0, 0, 0]；
	// s2 也是切片，是由s1切片追加而来，元素为4个，容量为4，前三个元素默认为int类型的零值，第四个为1，[0, 0, 0, 1]

	// s1修改元素会影响s2，因为s2追加元素后，并未超过s1切片的容量，所以共用底层数据，因此同样s2修改元素也会影响s1

	// s2再增加一个元素就会超过s1切片的容量，会在内存中创建一个新的底层数组，跟原来的数组分开，不共用了，
	// 追加后新创建的切片，元素为5个，容器为8（之前切片容量的2倍，版本1.18之前，扩容后的容量小于1024时，扩容翻倍，扩容后容量大于等于1024时，
	// 是之前的1.25倍；版本1.18以后，阈值变成了256，当扩容后的容量小于256时，扩容翻倍，当扩容后容量大于等于256时，
	// 公式为 newcap += (newcap + 3*threshold) / 4，即 newcap = newcap + newcap/4 + 192，即1.25倍再加192）
}

//批改意见
// 1. 追加元素后，s2最后一个元素修改时不会影响s1。
