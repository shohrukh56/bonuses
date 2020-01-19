package main

func main()  {
}
func bonuses(sales []int) int {
	const salesborder =10_000
	const percent=5
	totalBonuses:=0
	for _, sale:=range sales{
		if sale > salesborder{
			totalBonuses += (sale-salesborder)*percent/100
		}
	}

	return totalBonuses
}
