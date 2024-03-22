/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   image.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:44:01 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 18:43:11 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package util

import (
	"fmt"
	"mime/multipart"
	"strings"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

func SaveImage(src io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".jpg" && ext != ".png" {
		return "", fmt.Errorf("unsupported file type: %v", ext)
	}

	hasher := sha256.New()
	if _, err := io.Copy(hasher, src); err != nil {
		return "", err
	}
	hashedFileName := hex.EncodeToString(hasher.Sum(nil)) + ext
	hashedFilePath := filepath.Join("images", hashedFileName)
	file, err := os.Create(hashedFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(file, src); err != nil {
		return "", err
	}

	return hashedFileName, nil
}
