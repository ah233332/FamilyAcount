package utils

import "fmt"

type FamilyAccount struct {
	acount float64
	shouru float64
	zhichu float64
	note   string
	detail string
	key    int
	check  bool
	end    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		acount: 0.0,
		shouru: 0.0,
		zhichu: 0.0,
		note:   "",
		detail: "收支\t账户金额\t收支金额\t说明\n",
		key:    0,
		check:  false,
		end:    true,
	}
}
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("------家庭收支记账软件------")
		fmt.Println("       1.收支明细")
		fmt.Println("       2.登记收入")
		fmt.Println("       3.登记支出")
		fmt.Println("       4.退出系统")

		fmt.Printf(" 请选择功能(1~4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case 1:
			this.checkdetail()
		case 2:
			this.incount()
		case 3:
			this.outcount()
		case 4:
			this.endexe()
		default:
			fmt.Println("输入错误，请重新输入！")
		}

		if !this.end {
			fmt.Println("成功退出！")
			break
		}
	}
}
func (this *FamilyAccount) checkdetail() {
	fmt.Println("------当前收支明细------")
	if !this.check {
		fmt.Print("还没有收支记录，请先登记一笔吧^_^\n\n")
	} else {
		fmt.Print(this.detail)
		fmt.Println()
	}
}
func (this *FamilyAccount) incount() {
	fmt.Printf("请输入收入金额：")
	_, err := fmt.Scan(&this.shouru)
	if err != nil {
		fmt.Println("输入错误，请重新输入！")
		return
	}
	this.acount += this.shouru
	fmt.Printf("请输入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("收入\t%v\t%v\t%v\n", this.acount, this.shouru, this.note)
	this.check = true
}
func (this *FamilyAccount) outcount() {
	fmt.Print("请输入支出金额：")
	_, err := fmt.Scan(&this.zhichu)
	if err != nil {
		fmt.Println("输入错误，请重新输入！")
		return
	}
	this.acount -= this.zhichu
	fmt.Print("请输入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("支出\t%v\t%v\t%v\n", this.acount, this.zhichu, this.note)
	this.check = true
}
func (this *FamilyAccount) endexe() {
	var queding = ""
	for {
		fmt.Print("您确定要推出吗(Y/N)?")
		fmt.Scan(&queding)
		if queding == "Y" || queding == "y" {
			this.end = false
			break
		} else if queding == "N" || queding == "n" {
			this.end = true
			fmt.Print("退出失败！\n")
			break
		}
	}
}
