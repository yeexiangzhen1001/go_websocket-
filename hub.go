package main

import (
	"encoding/json"
)

//将连接对象初始化
var h = hub {
	connection: make(map[*connection]bool),
	broadcast: make(chan []byte),
	register: make(chan *connection),
	unregister: make(chan *connection),

}


//处理ws的逻辑实现
func (h *hub) run() {
	//监听数据管道，在后端不断处理管道数据
	for {
		//根据不同的数据管道，处理不同的逻辑
		select {
		//注册
		case c := <- h.register:
			//标识注册了
			h.connection[c] = true
			//组装data数据
			c.data.Ip = c.ws.RemoteAddr().String()
			//更新类型
			c.data.Type = "handshake"
			//用户列表
			c.data.UserList = user_list
			data_b, _:= json.Marshal(c.data)
			//将数据放入数据管道
			c.send <- data_b
		case c := <- h.unregister :
			//判断map里是存在要删除的数据
			if _,ok := h.connection [c]; ok {
				delete(h.connection,c)
				close(c.send)
			}
		case data := <- h.broadcast:
			//处理数据流转,将数据同步到所有用户
			//c是具体的每一个连接
			for c := range h.connection{
				//将数据同步
				select {
				case c.send <- data:
				default:
					//防止死循环
					delete(h.connection,c)
					close(c.send)
				}
			}

		}
	}
}
