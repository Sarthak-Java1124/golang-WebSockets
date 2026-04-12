package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		Hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured in binding json"})
		return
	}
	h.Hub.Rooms[req.ID] = &Rooms{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}
	c.JSON(http.StatusOK, req)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error in connecting the websocket"})
		return
	}
	roomId := c.Param("room_id")
	clientId := c.Query("user_id")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		RoomID:   roomId,
		Username: username,
		Message:  make(chan *Message),
		ID:       clientId,
	}
	m := &Message{
		RoomID:   roomId,
		Username: username,
		Content:  "A new user has joined the room",
	}
	h.Hub.Register <- cl
	h.Hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.Hub)

}
