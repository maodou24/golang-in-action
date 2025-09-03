package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func initUploadTask(c *gin.Context) {
	var task Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// concurrent
	if err := task.Init(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taskId":    task.TaskId,
		"uploaded":  task.Uploaded,
		"chunkSize": task.ChunkSize,
		"fileSize":  task.FileSize,
		"filename":  task.Filename,
	})
}

func upload(c *gin.Context) {
	var meta ChunkMeta

	if err := c.ShouldBindJSON(&meta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chunkFileName := fmt.Sprintf("./upload/%v/%v", meta.TaskId, meta.ChunkId)
	if err := c.SaveUploadedFile(file, chunkFileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taskId":  meta.TaskId,
		"chunkId": meta.ChunkId})
}

func complete(c *gin.Context) {
	var task Task // must: fileName, taskId. totalChunkCount, uploadedChunk ?

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mergeChunks(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func mergeChunks(task *Task) error {
	file, err := os.OpenFile("./upload/"+task.Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	for _, chunk := range task.Uploaded {
		chunkFileName := fmt.Sprintf("./upload/%v/%v", task.TaskId, chunk)
		chunkFile, err := os.OpenFile(chunkFileName, os.O_RDONLY, 0666)
		if err != nil {
			return err
		}
		if _, err := io.Copy(file, chunkFile); err != nil {
			return err
		}
	}

	go func() {
		if err := os.RemoveAll("./upload/" + task.TaskId); err != nil {
			log.Printf("remove upload task err:%v", err)
		}
	}()

	return nil
}
