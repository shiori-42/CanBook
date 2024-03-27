/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   client.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 14:57:43 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 17:34:23 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	SellerID string `json:"sellerId"`
	BuyerID  string `json:"buyerId"`
	Type     string `json:"type"` // "seller" or "buyer"
	RoomID   string
	hub      *Hub
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		m := &Message{
			Content:  string(message),
			RoomID:   c.RoomID,
			SenderID: c.ID,
		}
		hub.Broadcast <- m
	}
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := c.Conn.WriteJSON(message)
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
	}
}
