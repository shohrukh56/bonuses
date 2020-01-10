package main

func main()  {
}
func bonuses(sales []int) int {
	const salesborder =10_000
	const percent=5
	totalbonuses:=0
	bonus:=0
	for _, sale:=range sales{
		if sale > salesborder{
			bonus=(sale-salesborder)*percent/100
			totalbonuses+=bonus
		}
	}
	return totalbonuses
}
