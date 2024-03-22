/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   route.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:10 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 15:52:13 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/api/handler"
)

func RegisterRoutes(e *echo.Echo) {
	handler.RegisterUserRoutes(e)
	handler.RegisterItemRoutes(e)
}
