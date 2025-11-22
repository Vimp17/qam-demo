markdown
# 🔢 QAM-16 Modulation Demo

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://docker.com)

Проект демонстрирует работу модуляции **16-QAM (Quadrature Amplitude Modulation)** с добавлением гауссовского шума и подсчетом битовых ошибок (BER).

## 📋 Оглавление

- [Особенности](#особенности)
- [Структура проекта](#структура-проекта)
- [Быстрый старт](#быстрый-старт)
- [Использование](#использование)
- [Примеры](#примеры)
- [Технические детали](#технические-детали)
- [Разработка](#разработка)

## ✨ Особенности

- 🎯 **16-QAM модуляция/демодуляция** с Gray coding
- 📡 **Добавление гауссовского шума** с регулируемым уровнем
- 🔍 **Подсчет битовых ошибок** (Bit Error Rate)
- 🐳 **Docker-поддержка** для легкого развертывания
- 🔄 **Преобразование текст-биты-текст** с полным циклом обработки

## 📁 Структура проекта
qam-demo/
├── go.mod # Go модули и зависимости 
├── main.go # Основная программа
├── pkg/
│ ├── modulation/
│ │ ├── qam16.go # 16-QAM модуляция/демодуляция
│ │ └── noise.go # Добавление гауссовского шума
│ └── utils/
│ └── bit_utils.go # Утилиты для работы с битами
├── Dockerfile # Конфигурация Docker
└── README.md # Документация

text

## 🚀 Быстрый старт

### Предварительные требования

- **Go** 1.22 или выше
- **Docker** (опционально)

### Локальный запуск

```bash
# Клонирование репозитория
git clone https://github.com/Vimp17/qam-demo.git
cd qam-demo

# Сборка проекта
go build -o qam-demo
```

# Запуск
```bash
./qam-demo "Hello World" 0.1
Запуск через Docker
```

# Сборка Docker образа
```bash
docker build -t qam-demo .
```
# Запуск контейнера
```bash
docker run qam-demo "Тестовое сообщение" 0.2
```

💻 Использование
_Синтаксис команды_

bash
./qam-demo <сообщение> <уровень_шума>

_Параметры_
 - сообщение - Текст для модуляции (строка)

 -уровень_шума - Уровень гауссовского шума (число с плавающей точкой)

📊 Примеры
### Пример 1: Низкий уровень шума
bash```
./qam-demo "Hello QAM" 0.05
```
Ожидаемый вывод:

Original message: Hello QAM
Received message: Hello QAM
Bit errors: 0
BER: 0.00%

### Пример 2: Средний уровень шума
bash```
./qam-demo "Test Message" 0.3
```
Ожидаемый вывод:


Original message: Test Message
Received message: Test Mess ge
Bit errors: 3
BER: 2.34%
Пример 3: Высокий уровень шума
bash
./qam-demo "Demo" 1.0
Ожидаемый вывод:

text
Original message: Demo
Received message: D m 
Bit errors: 8
BER: 25.00%
🔧 Технические детали
Модуляция 16-QAM
Схема кодирования: Gray code

Сопоставление символов:

00 → +3 | 01 → +1 | 11 → -1 | 10 → -3

Количество бит на символ: 4 бита (2 для I, 2 для Q)

Модель канала
Тип шума: Аддитивный белый гауссовский шум (AWGN)

Параметр: Стандартное отклонение (σ)

Обработка данных
text
Текст → Биты → QAM-16 Символы → + Шум → Демодуляция → Биты → Текст
