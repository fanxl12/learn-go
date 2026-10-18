package main

import (
	u "LearnGo/ch09/user" // 导入的别名
	"LearnGo/ch10/user"   // 包的路径
	//. "LearnGo/ch10/user" //用的很少
	//_ "LearnGo/ch10/user" //只导入初始化包，不用包里面的内容
	"fmt"
)

func main() {
	c := user.Course{
		Age: 10,
	}
	fmt.Println(user.GetCourseAge(c))
	u9 := u.Course{
		Age: 18,
	}
	fmt.Println(u.GetCourseAge(u9))
}
