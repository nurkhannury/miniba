package handlers

import (
	"encoding/json"
	"log"
	"mini_BA/db"
	"net/http"

	"mini_BA/api"
	"mini_BA/models"
	"mini_BA/prompts"
)

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "🚫 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var biz models.Business
	if err := json.NewDecoder(r.Body).Decode(&biz); err != nil {
		http.Error(w, "❌ Invalid JSON input", http.StatusBadRequest)
		log.Println("Decode error:", err)
		return
	}

	// 🧠 Генерируем промпт и получаем ответ от GPT
	prompt := prompts.GeneratePrompt(biz)
	response, err := api.GetGPTResponse(prompt)
	if err != nil {
		http.Error(w, "❌ Failed to get response from GPT", http.StatusInternalServerError)
		log.Println("GPT error:", err)
		return
	}

	// 💾 Сохраняем в базу данных (если email и password указаны)
	if biz.Email != "" && biz.Password != "" {
		if err := db.SaveReport(biz.Email, biz.Password, response); err != nil {
			log.Println("❌ Failed to save report:", err)
		}
	}

	// ✅ Возвращаем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"response": response,
	}); err != nil {
		log.Println("❌ Failed to write response:", err)
	}
}

func RecommendHandler(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только POST-запросы
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем входные данные формы в структуру бизнес-модели
	var biz models.Business
	if err := json.NewDecoder(r.Body).Decode(&biz); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Генерируем промпт для GPT
	prompt := prompts.GenerateRecommendationsPrompt(biz)

	// Получаем ответ от GPT
	response, err := api.GetGPTResponse(prompt)
	if err != nil {
		http.Error(w, "Failed to get GPT response", http.StatusInternalServerError)
		return
	}

	// Возвращаем сырой текст как JSON-поле "response"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"response": response,
	})
}
