/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   message.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/27 23:21:02 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/28 01:52:47 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ws

type Message struct {
	SenderID    string `json:"senderID"`
	RecipientID string `json:"recipientID"`
	Content     string `json:"content"`
}
