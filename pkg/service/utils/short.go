package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Short struct {
	hashMap map[string]string
}

func New() *Short {
	return &Short{
		hashMap: make(map[string]string),
	}
}

func (s *Short) CreateIntMap(keys []string) {
	for _, key := range keys {
		shortURL := s.GenerateURL()
		s.hashMap[shortURL] = "vgang.io/retailer/list-of-products/" + key
	}
}

func (s *Short) GenerateURL() string {
	rand.Seed(time.Now().UnixNano())
	const chs = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var st strings.Builder
	for k := 0; k < 8; k++ {
		st.WriteByte(chs[rand.Intn(len(chs))])
	}
	return st.String()
}

func (s *Short) CreateURL(c *fiber.Ctx) error {
	realURL := c.FormValue("url")
	if realURL == "" {
		return fmt.Errorf("no url")
	}

	// Check if the realURL already has a shortURL
	existingShortURL, exists := s.findShortURL(realURL)
	if exists {
		return fmt.Errorf(fmt.Sprintf("Short URL: http://%s", existingShortURL))
	}

	shortURL := s.GenerateURL()

	s.hashMap[shortURL] = realURL
	return fmt.Errorf(fmt.Sprintf("Short URL: http://%s", shortURL))
}

func (s *Short) Redirect(c *fiber.Ctx) error {
	miniURL := strings.TrimPrefix(c.Params("shortURL"), "/")
	fmt.Println(miniURL)
	if miniURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("NO URL")
	}
	originalURL, ok := s.hashMap[miniURL]
	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("URL not found")
	}
	fmt.Println("originalURL", originalURL)
	return c.Redirect("https://"+originalURL, fiber.StatusSeeOther)
}

func (s *Short) GetAll(c *fiber.Ctx) error {
	values := make([]string, 0, len(s.hashMap))
	for key, value := range s.hashMap {
		values = append(values, "short:"+key+" real:"+value)
	}
	return c.Status(fiber.StatusOK).JSON(values)
}

func (s *Short) findShortURL(realURL string) (string, bool) {
	for shortURL, url := range s.hashMap {
		if url == realURL {
			return shortURL, true
		}
	}
	return "", false
}
