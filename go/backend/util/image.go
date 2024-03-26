/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   image.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:44:01 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/27 00:03:48 by shiori0123       ###   ########.fr       */
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
)

func SaveImage(src multipart.File, image *multipart.FileHeader) (string, error) {
    uploadDir := "../images/"

    hasher := sha256.New()
    _, err := io.Copy(hasher, src)
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