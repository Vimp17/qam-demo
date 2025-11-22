# QAM-16 Modulation Demo

Проект демонстрирует работу модуляции 16-QAM (Quadrature Amplitude Modulation) с добавлением гауссовского шума и подсчетом битовых ошибок.

## Структура проекта

qam-demo/
├── go.mod
├── main.go
├── pkg/
│ ├── modulation/
│ │ ├── qam16.go # 16-QAM модуляция/демодуляция
│ │ └── noise.go # Добавление гауссовского шума
│ └── utils/
│ └── bit_utils.go # Утилиты для работы с битами
├── Dockerfile
└── README.md



## Функциональность

- Преобразование текста в битовую последовательность
- Модуляция 16-QAM с Gray coding
- Добавление гауссовского шума с заданным SNR
- Демодуляция и восстановление текста
- Подсчет битовых ошибок (BER)

## Требования

- Go 1.22 или выше
- Docker (опционально)

## Быстрый старт

### Локальный запуск

```bash
# Клонирование репозитория
git clone https://github.com/your-username/qam-demo.git
cd qam-demo

# Сборка и запуск
go build -o qam-demo
./qam-demo "Hello World" 0.1