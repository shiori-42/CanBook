package handler

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/repository"
)

func RegisterWebSocketRoutes(e *echo.Echo) {
	h := &wsHandler{}
	e.GET("/ws", h.HandleWebSocket)
	e.GET("/ws/:itemId", h.HandleWebSocket, AuthMiddleware)
}

type WSHandler interface {
	HandleWebSocket(c echo.Context) error
}

type wsHandler struct{}

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"messageType"`
}

type Client struct {
	UserID uint
	ItemID uint
	Conn   *websocket.Conn
}

type Message struct {
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
	Content     string `json:"content"`
}

var (
	clients  = make(map[uint]map[uint]*Client)
	mu       sync.Mutex
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// HandleWebSocket upgrades connection to WebSocket
func (h *wsHandler) HandleWebSocket(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	var client *Client
	defer func() {
		mu.Lock()	
		if client != nil {
			delete(clients[client.UserID], client.ItemID)
			if len(clients[client.UserID]) == 0 {
				delete(clients, client.UserID)
			}
		}
		mu.Unlock()	
		ws.Close()	
	}()


	var response WsJsonResponse
	response.Message = "Connected to websocket"
	if err = ws.WriteJSON(response); err != nil {
		return err
	}

	for {
		var req map[string]interface{}
		if err := ws.ReadJSON(&req); err != nil {
			log.Println("ReadJSON error:", err)
			return err
		}

		itemIDStr, ok := req["itemId"].(string)
		if !ok {
			log.Println("Invalid itemId")
			continue
		}

		itemID, err := strconv.Atoi(itemIDStr)
		if err != nil {
			log.Println("Invalid itemId format:", err)
			continue
		}

		item, err := repository.GetItemByID(itemID)
		if err != nil {
			log.Println("GetItemByID error:", err)
			break
		}

		client = &Client{
			UserID: userID,
			ItemID: uint(itemID),
			Conn:   ws,
		}

		mu.Lock()
		if clients[item.UserID] == nil {
			clients[item.UserID] = make(map[uint]*Client)
		}
		clients[item.UserID][uint(itemID)] = client
		mu.Unlock()

		for {
			var message Message
			if err := ws.ReadJSON(&message); err != nil {
				log.Println("ReadJSON message error:", err)
				break
			}

			message.SenderID = strconv.FormatUint(uint64(userID), 10)
			message.RecipientID = strconv.FormatUint(uint64(item.UserID), 10)

			mu.Lock()
			recipientClient, ok := clients[item.UserID][uint(itemID)]
			mu.Unlock()
			if !ok {
				log.Printf("Recipient %d not found for item %d", item.UserID, itemID)
				continue
			}

			if err := recipientClient.Conn.WriteJSON(message); err != nil {
				log.Println("WriteJSON error:", err)
			}
		}
	}
}

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// Home renders the home page
func Home(c echo.Context) error {
	err := renderPage(c, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
	return err
}

// renderPage renders a jet template
func renderPage(c echo.Context, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		return err
	}

	err = view.Execute(c.Response().Writer, data, nil)
	if err != nil {
		return err
	}

	return nil
}
