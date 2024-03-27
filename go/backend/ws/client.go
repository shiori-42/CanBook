/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   client.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 23:19:34 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 23:19:38 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

import "github.com/gorilla/websocket"

type Client struct {
	UserID string
	Conn   *websocket.Conn
}
