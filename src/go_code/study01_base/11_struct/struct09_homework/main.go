package main

import (
	"fmt"
	"os"
)

// 结构体实战 - 学员信息管理系统
// 1、添加学员信息
// 2、编辑学员信息
// 3、展示所有的学员信息
func main() {
	sm := newStudentMgr()
	for {

		// 1、打印系统菜单
		showMenu()
		// 2、等待用户要执行的选项
		var input int
		fmt.Scanf("%d\n", &input)     // &input 值接收结果
		fmt.Println("用户输入的是:", input) // 执行对应的操作
		// 3.执行用户选择的选项
		switch input {
		case 1:
			sm.saveStudent(getInput())
			break
		case 2:
			sm.updateStudent(getInput())
			break
		case 3:
			sm.showAllStudent()
			break
		case 4:
			sm.deleteStudent(getInput())
			break
		case 5:
			// 退出
			os.Exit(0)
		}

	}
}

// showMenu 打印系统菜单
func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.删除学员信息")
	fmt.Println("5.推出系统")
}

// getInput 获取用户输入的信息
func getInput() *student {
	var (
		id          int
		name, class string
	)
	fmt.Println("请按照要求输入学员信息")
	// 姓名、地址
	fmt.Println("id:")
	fmt.Scanf("%d\n", &id) // &input 值接收结果
	fmt.Println("姓名:")
	fmt.Scanf("%s\n", &name) // &input 值接收结果
	fmt.Println("班级:")
	fmt.Scanf("%s\n", &class) // &input 值接收结果

	stu := newStudent(id, name, class)
	return stu
}
