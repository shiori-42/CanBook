/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ws_handler.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 14:57:46 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 17:45:44 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (

	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/ws"
)

func RegisterWebSocketRoutes(e *echo.Echo, hub *ws.Hub) {
    wsHandler := &ws.Handler{
        Hub: hub,
    }
    e.GET("/ws/:roomId", wsHandler.HandleWebSocket)
    e.POST("/ws/createRoom", wsHandler.CreateRoom)
}
