/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handler.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 23:20:41 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 23:48:15 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	clients  = make(map[string]*Client)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Handler struct{}

func (h *Handler) HandleWebSocket(c echo.Context) error {
	userID := c.QueryParam("userID")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := &Client{
		UserID: userID,
		Conn:   conn,
	}
	clients[userID] = client

	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			break
		}

		recipientClient, ok := clients[message.RecipientID]
		if !ok {
			log.Printf("Recipient %s not found", message.RecipientID)
			continue
		}

		if err := recipientClient.Conn.WriteJSON(message); err != nil {
			log.Println(err)
		}
	}

	delete(clients, userID)

	return nil
}
