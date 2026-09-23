package main

import "fmt"

var is = true

func main() {
	var acount, shouru, zhichu float64
	var note string
	detail := "收支\t账户金额\t收支金额\t说明\n"
	count := 0
	for {
		fmt.Println("------家庭收支记账软件------")
		fmt.Println("       1.收支明细")
		fmt.Println("       2.登记收入")
		fmt.Println("       3.登记支出")
		fmt.Println("       4.退出系统")

		var n int
		fmt.Printf(" 请选择功能(1~4):")
		fmt.Scan(&n)
		switch n {
		case 1:
			fmt.Println("------当前收支明细------")
			if count == 0 {
				fmt.Print("还没有收支记录，请先登记一笔吧^_^\n\n")
			} else {
				fmt.Print(detail)
				fmt.Println()
			}
		case 2:
			fmt.Printf("请输入收入金额：")
			_, err := fmt.Scan(&shouru)
			if err != nil {
				fmt.Println("输入错误，请重新输入！")
				break
			}
			acount += shouru
			fmt.Printf("请输入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("收入\t%v\t%v\t%v\n", acount, shouru, note)
			count++
		case 3:
			fmt.Print("请输入支出金额：")
			_, err := fmt.Scan(&zhichu)
			if err != nil {
				fmt.Println("输入错误，请重新输入！")
				break
			}
			acount -= zhichu
			fmt.Print("请输入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("支出\t%v\t%v\t%v\n", acount, zhichu, note)
			count++
		case 4:
			var queding = ""
			for {
				fmt.Print("您确定要推出吗(Y/N)?")
				fmt.Scan(&queding)
				if queding == "Y" || queding == "y" {
					is = false
					break
				} else if queding == "N" || queding == "n" {
					is = true
					fmt.Print("退出失败！\n")
					break
				}
			}

		default:
			fmt.Println("输入错误，请重新输入！")
		}

		if !is {
			break
		}
	}
	fmt.Println("成功退出！")
}
