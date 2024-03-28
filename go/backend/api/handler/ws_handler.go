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
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/ws"
)

func RegisterWebSocketRoutes(e *echo.Echo) {
	wsHandler := &ws.Handler{}
	e.GET("/ws/:itemId", wsHandler.HandleWebSocket, AuthMiddleware)
}
