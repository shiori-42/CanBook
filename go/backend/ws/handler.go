/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handler.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 23:20:41 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/28 03:49:02 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	clients  = make(map[uint]*Client)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Handler struct{}

func (h *Handler) HandleWebSocket(c echo.Context) error {
    userID, ok := c.Get("user_id").(uint)
    if !ok {
        c.Logger().Error("Invalid user ID")
        return echo.NewHTTPError(http.StatusInternalServerError, "Invalid user ID")
    }
    c.Logger().Info("User ID in HandleWebSocket:", userID)

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

        recipientID, err := strconv.ParseUint(message.RecipientID, 10, 64)
        if err != nil {
            log.Printf("Invalid recipient ID: %v", err)
            continue
        }

        recipientClient, ok := clients[uint(recipientID)]
        if !ok {
            log.Printf("Recipient %d not found", recipientID)
            continue
        }

        if err := recipientClient.Conn.WriteJSON(message); err != nil {
            log.Println(err)
        }
    }
	delete(clients, userID)

	return nil
}
