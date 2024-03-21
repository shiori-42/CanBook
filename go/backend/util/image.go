/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   image.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:44:01 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 16:44:03 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

func SaveImage(src io.Reader) (string, error) {
	tempFile, err := os.CreateTemp("images", "*.jpg")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	hasher := sha256.New()
	if _, err := io.Copy(io.MultiWriter(tempFile, hasher), src); err != nil {
		return "", err
	}

	hashedFileName := hex.EncodeToString(hasher.Sum(nil)) + ".jpg"
	if err := os.Rename(tempFile.Name(), filepath.Join("images", hashedFileName)); err != nil {
		return "", err
	}

	return hashedFileName, nil
}