/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ws_handler.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 23:23:35 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/28 21:30:56 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (
	"net/http"
	// "os"
	// "strconv"
	"github.com/CloudyKit/jet/v6"
	"log"
	
	// "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	// "github.com/shiori-42/textbook_change_app/go/backend/repository"
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
	Action  string `json:"action"`
	Message string `json:"message"`
	MessageType string `json:"messageType"`
}
type Client struct {
	UserID uint
	Conn   *websocket.Conn
}

type Message struct {
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
	Content     string `json:"content"`
}

var (
	clients  = make(map[uint]*Client)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//WsEndpoint upgrades connection to websocket
func (h *wsHandler) HandleWebSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	var response WsJsonResponse
	response.Message=`Connected to websocket`
	err=ws.WriteJSON(response)
	if err != nil {
		return err
	}

	userID := c.Get("user_id").(uint)

	for {
		var req map[string]interface{}
		err := ws.ReadJSON(&req)
		if err != nil {
			log.Println(err)
			return err
		}

		
		itemID:= req["itemId"].(string)
		

			item, err := repository.GetItemByID(int(itemID))
			if err != nil {
				log.Println(err)
				break
			}

			client := &Client{
				UserID:	userID,
				Conn:   ws,
			}
			clients[item.UserID] = client

			for {
				var message Message
				err := ws.ReadJSON(&message)
				if err != nil {
					log.Println(err)
					break
				}

				message.SenderID = strconv.FormatUint(uint64(userID), 10)
				message.RecipientID = strconv.FormatUint(uint64(item.UserID), 10)

				recipientClient, ok := clients[item.UserID]
				if !ok {
					log.Printf("Recipient %d not found", item.UserID)
					continue
				}

				if err := recipientClient.Conn.WriteJSON(message); err != nil {
					log.Println(err)
				}

			delete(clients, item.UserID)
		}
	}

	return nil
}

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// Home renders the home page
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
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

func Home(c echo.Context) error {
    data := make(jet.VarMap)
    data.Set("title", "Home Page")
    return renderPage(c, "home.jet", data)
}

e.GET("/", handler.Home)