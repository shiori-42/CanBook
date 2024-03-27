/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   message.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 15:07:55 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 17:39:44 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

type Message struct {
    Content  string `json:"content"`
    SenderID string `json:"senderId"`
    RoomID   string `json:"roomId"`
}