package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
// ClientManager 来管理整个websocket连接
type ClientManager struct {
	// Clients is a map to find all client online
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
//Client 用来管理websocket用户
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is return msg
// Message 用来管理发送的信息
type Message struct {
	Sender    string `json:"sender,omitempty"`
	// Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
//Manager 指定一个websocket服务管理
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
}

// Start is  项目运行前, 协程开启start -> go Manager.Start()
func (manager *ClientManager) Start() {
	var message []byte
	for {
		log.Println("<---管道通信--->")
		select {
		case conn := <-Manager.Register:
			log.Printf("新用户加入:%v", conn.ID)
			message = []byte("新用户加入：" + conn.ID)
			Manager.Clients[conn.ID] = conn
			/*jsonMessage, _ := json.Marshal(&Message{Content: "Successful connection to socket service"})
			conn.Send <- jsonMessage*/
			for _, conn := range Manager.Clients {
				conn.Send <- message
				//  fmt.Println(conn.ID)
			}

		case conn := <-Manager.Unregister:
			log.Printf("用户离开:%v", conn.ID)
			message = []byte(conn.ID + "用户离开")
			for _, conn := range Manager.Clients {
				conn.Send <- message
			}

			if _, ok := Manager.Clients[conn.ID]; ok {
				jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected"})
				conn.Send <- jsonMessage
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case message := <-Manager.Broadcast:
			MessageStruct := Message{}
			json.Unmarshal(message, &MessageStruct)
			// msgSend(&message)
			for _, conn := range Manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
		}
	}
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		log.Printf("读取到客户端的信息:%s", string(message))
		Manager.Broadcast <- message
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("发送到到客户端的信息:%s", string(message))

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

//TestHandler socket 连接 中间件 作用:升级协议,用户验证,自定义信息等
func WsHandler(c *gin.Context) {
	// conn is a router of updated http(websocket) link
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	conn.PongHandler()
	_, ass, err := conn.ReadMessage()
	if err != nil {
		return
	}
	//可以添加用户信息验证
	// creat a client
	uid := string(ass)
	uid = uid[1 : len(uid)-1]
	client := &Client{
		ID:     uid,
		Socket: conn,
		Send:   make(chan []byte),
	}
	Manager.Register <- client
	go client.Read()
	go client.Write()
}
