package handlers

import (
	"bufio"
	"fmt"
	"time"

	"encoding/json"
	"inoUwU/pinu/app/domain/port"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"github.com/valyala/fasthttp"
)

// SSEHandler はSSE用エンドポイントを提供します
type SSEHandler struct {
	sseBroker port.SSEBroker
}

func NewSSEHandler(i *do.Injector) (*SSEHandler, error) {
	broker := do.MustInvoke[port.SSEBroker](i)
	return &SSEHandler{sseBroker: broker}, nil
}

func (h *SSEHandler) Route(router fiber.Router) error {
	fmt.Println("Registering sse routes")
	// user := router.Group("/user")
	router.Get("/sse", h.SSEStream)
	router.Put("/publish", h.Publish)
	return nil
}

// SSEStream SSEストリームを返却します
func (h *SSEHandler) SSEStream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Cache-Control")

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			msg, exists := h.sseBroker.Consume()
			var jsonData []byte
			var err error

			if exists {
				// メッセージがキューにある場合
				data := map[string]interface{}{
					"type":      "message",
					"data":      msg,
					"timestamp": time.Now().Unix(),
				}
				jsonData, err = json.Marshal(data)
			} else {
				// ハートビート
				data := map[string]interface{}{
					"type":      "heartbeat",
					"timestamp": time.Now().Unix(),
				}
				jsonData, err = json.Marshal(data)
			}

			if err != nil {
				fmt.Printf("JSON marshaling error: %v\n", err)
				continue
			}

			fmt.Fprintf(w, "data: %s\n\n", string(jsonData))

			flushErr := w.Flush()
			if flushErr != nil {
				fmt.Printf("フラッシュ中にエラー発生: %v. HTTP接続を閉じます。\n", flushErr)
				break
			}
			time.Sleep(2 * time.Second)
		}
	}))

	return nil
}

// Publish キューにメッセージを追加します
func (h *SSEHandler) Publish(c *fiber.Ctx) error {
	var payload struct {
		Message string `json:"message"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	h.sseBroker.Publish(payload.Message)
	println("Message published:", payload.Message)
	return c.SendString("Message added to queue")
}
