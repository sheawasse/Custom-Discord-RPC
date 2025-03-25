# Custom-Discord-RPC
RU:
✨ Основные функции:
🔗 Подключение к Discord

Отображает кастомный статус в профиле Discord.
Работает в фоне (не требует постоянного открытого окна).

📝 Полная настройка текста
Первая строка (Details) — например, название лаунчера.
Вторая строка (State) — доп. информация (версия, моды и т.д.).
Тултипы (LargeText, SmallText) — пояснения при наведении на иконки.

🖼 Поддержка картинок
Большая и маленькая иконки (загружаются в Discord Dev Portal).
Можно динамически менять изображения.

🛠 2 кнопки с ссылками
Например:
Download — скачать лаунчер.
Help — открыть гайд.

⚙ Автономная работа (работает до перезагрузки ПК). 

🔧 Технические детали:
Язык: Go (сборка в .exe).
Библиотека: rich-go для Discord RPC.
Платформа: Windows (можно адаптировать под macOS/Linux).
-----------------------------------------------------------------------
ENG:
✨ Core Features:
🔗 Discord Integration

Displays custom status in Discord profile.
Runs in the background (no need to keep the window open).

📝 Full Text Customization
First line (Details) — e.g., launcher name.
Second line (State) — additional info (version, mods, etc.).
Tooltips (LargeText, SmallText) — hover text for icons.

🖼 Image Support
Large + small icons (uploaded to Discord Dev Portal).
Dynamic image switching is possible.

🛠 2 Clickable Buttons
Example:
Download — direct download link.
Help — open documentation.

⚙ Standalone Operation (runs until PC reboot).

🔧 Technical Details:
Language: Go (compiles to .exe).
Library: rich-go for Discord RPC.
Platform: Windows (can be adapted for macOS/Linux).
-----------------------------------------------------------------------
🚀 Пример кода / Code Example
go

activity := client.Activity{
    Details:    "My Awesome Launcher",  // Основной текст / Main text
    State:      "v1.2.0 • With Mods",  // Доп. информация / Additional info
    LargeImage: "main_icon",           // Ключ картинки / Image key
    Buttons: []*client.Button{
        {Label: "Download", Url: "https://example.com/download"},
        {Label: "Help", Url: "https://example.com/docs"},
    },
}

🎨 Где использовать? / Use Cases:
Игровые лаунчеры / Game launchers.

Стримерские утилиты / Streamer tools.

Кастомные статусы / Custom Discord profiles.
