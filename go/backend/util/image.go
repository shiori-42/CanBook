/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   image.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:44:01 by shiori0123        #+#    #+#             */
/*   Updated: 2024/07/13 19:03:33 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"log"
)

func SaveImage(src multipart.File, image *multipart.FileHeader) (string, error) {
	uploadDir := os.Getenv("UPLOAD_DIR")
	log.Println("UPLOAD_DIR is set to:", uploadDir)

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		log.Println("Failed to create directory:", err)
		return "", err
	}
	hasher := sha256.New()
	_, err = io.Copy(hasher, src)
	if err != nil {
		return "", err
	}
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fileName := hashString + filepath.Ext(image.Filename)

	filePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
