package _317_1

func getSheep(year int) int {
	num := 1 // init sheep num
	// note: init i = 1; i <= year
	for i := 1; i <= year; i++ {
		if i == 2 {
			num += getSheep(year-2) // 递归计算2年时出生新羊
		}else if i == 4 {
			num += getSheep(year-4) // 递归计算4年时出生新羊
		}else if i == 5 {
			num--
			break // 优化
		}
	}
	return num
}