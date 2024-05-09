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
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/repository"
)

type WSHandler interface {
	HandleWebSocket(c echo.Context) error
}

type wsHandler struct{}

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

func (h *wsHandler) HandleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	var token string
	var itemID uint
	for {
		var req map[string]interface{}
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Println(err)
			break
		}

		if token == "" {
			if tokenStr, ok := req["token"].(string); ok {
				token = tokenStr
				claims := jwt.MapClaims{}
				_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("SECRET")), nil
				})
				if err != nil {
					log.Println(err)
					break
				}
				userID := uint(claims["user_id"].(float64))
				c.Set("user_id", userID)
			}
		}

		if itemIDStr, ok := req["itemId"].(string); ok {
			itemIDUint64, err := strconv.ParseUint(itemIDStr, 10, 64)
			if err != nil {
				log.Println(err)
				break
			}
			itemID = uint(itemIDUint64)
		}

		if itemID != 0 {
			item, err := repository.GetItemByID(int(itemID))
			if err != nil {
				log.Println(err)
				break
			}

			client := &Client{
				UserID: c.Get("user_id").(uint),
				Conn:   conn,
			}
			clients[item.UserID] = client

			for {
				var message Message
				err := conn.ReadJSON(&message)
				if err != nil {
					log.Println(err)
					break
				}

				message.SenderID = strconv.FormatUint(uint64(client.UserID), 10)
				message.RecipientID = strconv.FormatUint(uint64(item.UserID), 10)

				recipientClient, ok := clients[item.UserID]
				if !ok {
					log.Printf("Recipient %d not found", item.UserID)
					continue
				}

				if err := recipientClient.Conn.WriteJSON(message); err != nil {
					log.Println(err)
				}
			}

			delete(clients, item.UserID)
		}
	}

	return nil
}

func RegisterWebSocketRoutes(e *echo.Echo) {
	h := &wsHandler{}
	e.GET("/ws/:itemId", h.HandleWebSocket, AuthMiddleware)
}
