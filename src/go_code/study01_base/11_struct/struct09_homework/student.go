package main

import "fmt"

// 用来存放student相关信息

// student 学生信息
type student struct {
	id          int // 学号唯一的
	name, class string
}

// newStudent student的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// studentMgr 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

// newStudentMgr 学员管理的类型的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// saveStudent 添加学生studentMgr类型的的方法
func (s *studentMgr) saveStudent(student2 *student) {
	s.allStudents = append(s.allStudents, student2)
}

// updateStudent 修改学生 studentMgr类型的的方法
func (s *studentMgr) updateStudent(student2 *student) {
	// 遍历所有的学生
	for i, v := range s.allStudents {
		// 找到id相同的学生
		if v.id == student2.id {
			// 修改对应值
			s.allStudents[i] = student2
			return
		}
	}
	//
	fmt.Println("该学生不存在，无法修改")
}

// deleteStudent 删除学生 studentMgr类型的的方法
func (s *studentMgr) deleteStudent(student2 *student) {
	// 找到要删除的学生的索引位置
	// 定义索引位置
	index := -1
	// 遍历所有的学生
	for i, v := range s.allStudents {
		// 找到id相同的学生
		if v.id == student2.id {
			index = i
			break
		}
	}
	if index != -1 {
		s.allStudents = append(s.allStudents[:index], s.allStudents[index+1:]...)
	} else {
		fmt.Println("该学生不存在，无法删除")
	}
}

// showAllStudent 展示学生
func (s *studentMgr) showAllStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名: %s 班级 %s\n", v.id, v.name, v.class)
	}
}
