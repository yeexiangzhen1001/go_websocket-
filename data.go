package main

//将连接中传输的数据抽象出对象
type Data struct{
	Ip string `json:"ip"`
	//标识信息的类型
	Type string `json:"type"`
	//代表那个用户说的
	From string `json:"from"`
	//内容
	Content string `json:"content"`
	//用户
	User string `json:"user"`
	//用户列表
	UserList []string `json:"user_list"`
}