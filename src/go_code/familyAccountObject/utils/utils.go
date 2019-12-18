package utils
import (
	"fmt"
)

type FamilyAccount struct {
	key string
	loop bool
	hasShouzhi bool
	balance float64
	money	float64
	note string
	details string
}

func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		key: "",
		loop: true,
		hasShouzhi:false,
		balance: 10000.0,
		money: 0.0,
		note: "",
		details:"",
	}
}
func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------------当前收支明细记录--------------------")
	if this.hasShouzhi {
		fmt.Println(this.details)
	} else {
		fmt.Println("没有收支情况！")
	}
}
func (this *FamilyAccount) save() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance,this.money,this.note)
	this.hasShouzhi = true
	this.showDetails()
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足！")
		//break
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明：")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
		this.hasShouzhi = true
		this.showDetails()
	}
}

func (this *FamilyAccount) exit() {
	fmt.Println("确定退出吗 y/n")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" {
		this.loop = false
		// break
	}
}
func (this *FamilyAccount) MainMenu() {
	//显示主菜单
	for {
		fmt.Println("\n--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1 收支明细")
		fmt.Println("                     2 登记收入")
		fmt.Println("                     3 登记支出")
		fmt.Println("                     4 退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.save()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("请正确输入正确选项")
		}
		if !this.loop {
			fmt.Println("你退出了家庭收支记账软件")
			break
		}
	}
}