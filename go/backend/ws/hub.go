/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   hub.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 14:57:31 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 17:44:02 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			roomID := generateRoomID(client.SellerID, client.BuyerID)
			if room, ok := h.Rooms[roomID]; ok {
				room.Clients[client.ID] = client
			} else {
				room = &Room{
					ID:      roomID,
					Clients: make(map[string]*Client),
				}
				room.Clients[client.ID] = client
				h.Rooms[roomID] = room
			}
		case client := <-h.Unregister:
			roomID := generateRoomID(client.SellerID, client.BuyerID)
			if room, ok := h.Rooms[roomID]; ok {
				delete(room.Clients, client.ID)
				close(client.Message)
				if len(room.Clients) == 0 {
					delete(h.Rooms, roomID)
				}
			}
		case message := <-h.Broadcast:
			if room, ok := h.Rooms[message.RoomID]; ok {
				for _, client := range room.Clients {
					if client.ID != message.SenderID {
						select {
						case client.Message <- message:
						default:
							close(client.Message)
							delete(room.Clients, client.ID)
						}
					}
				}
			}
		}
	}
}
func generateRoomID(sellerID, buyerID string) string {
	// Generate a unique room ID based on seller and buyer IDs
	// For example, concatenate seller and buyer IDs
	return sellerID + "_" + buyerID
}
