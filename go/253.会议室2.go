//用了暴力法解决，这里存在有个优化，用堆排序去排序一下各个room的截止时间（堆在进行非全量比较，只用比较某个极值的时候是比较好用的）

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	rooms := []int{intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		//check是否存在合适的room,如果存在则更新当前会议室的结束时间
		check := false
		for j := 0; j < len(rooms); j++ {
			if intervals[i][0] >= rooms[j] {
				rooms[j] = intervals[i][1]
				check = true
				break
			}
		}
		if !check {
			rooms = append(rooms, intervals[i][1])
		}
	}

	return len(rooms)
}