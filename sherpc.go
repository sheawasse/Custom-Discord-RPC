// This is a simple visual RPC by sheawasse : Golang
// This is a simple visual RPC by sheawasse : Golang
// This is a simple visual RPC by sheawasse : Golang

package main

import (
	"log"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func main() {
	// Подключение к Discord RPC
	err := client.Login("1354080587115266141") // Замените на ваш Client ID из Discord Dev Portal
	if err != nil {
		log.Fatalf("Ошибка подключения к Discord RPC: %v", err)
	}

	// Настройка активности
	err = client.SetActivity(client.Activity{
		Details:    "Plays with Horion Injector: custom dll",      // Основной текст (например, название лаунчера)
		State:      "Minecraft V. - 1.21.44 // DLL - HorionOld.dll",      // Дополнительный текст (например, версия)
		LargeImage: "horion",          // Ключ большой картинки
		SmallImage: "minecraft",          // Ключ маленькой картинки (опционально)
		LargeText:  "This is a simple visual RPC by sheawasse",   // Тултип для большой картинки
		SmallText:  "This is a simple visual RPC by sheawasse",            // Тултип для маленькой картинки
		Buttons: []*client.Button{
			{
				Label: "Download",            // Первая кнопка
				Url:   "https://horioninjector.com", // Замените на реальную ссылку
			},
			{
				Label: "Discord",                 // Вторая кнопка
				Url:   "https://discord.gg/dfA9SxXVeH",    // Замените на реальную ссылку
			},
		},
	})
	if err != nil {
		log.Fatalf("Ошибка установки статуса: %v", err)
	}

	log.Println("Custom RPC by sheawasse successfully launched! The program runs in the background.")

	// Бесконечный цикл для поддержания работы
	for {
		time.Sleep(1 * time.Hour)
	}
} 