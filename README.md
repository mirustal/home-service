# Home-Services

## Auth Service
![Build Status](https://github.com/mirustal/home-service/actions/workflows/auth-service.yml/badge.svg)
## Home Service
![Build Status](https://github.com/mirustal/home-service/actions/workflows/home-service.yml/badge.svg)
---
### Общий процент покрытия тестами
[![Overall Coverage](https://codecov.io/gh/mirustal/home-service/branch/develop/graph/badge.svg)](https://codecov.io/gh/mirustal/home-service/branch/develop)

# Тестовое задание для отбора на Backend Bootcamp

## Запуск сервиса

### Make
  1. make up-d - запустит контейнеры на фоне
  1. make up-b - запустит контейнеры

### Docker-compose 
  docker-compose up -d | docker-compose up --build


## Описание
В данном репозитории представлена реализация тестового задания для Backend Bootcamp. Реализация полностью соответствует предоставленному API и выполнена с использованием `grpc-gateway`.

**port gateway: 8080**
**port swagger: 8081/auth/**
**port swagger: 8081/home/**

## Архитектура

![Снимок экрана 2024-08-19 в 23 17 12](https://github.com/user-attachments/assets/f7141fa8-c12e-4321-a20d-6a45c3b15ec2)

Реализация состоит из следующих сервисов:

- **auth-service**: Отвечает за авторизацию пользователя. Использует JWT для создания access и refresh токенов. В текущей реализации используется Bearer токен в соответствии с API. Легко модифицируется для работы с различными secret_key.

- **home-service**: Обрабатывает все запросы, связанные с квартирами и домами.

- **sender-service**: Отвечает за отправку сообщений подписчикам. ! **Не RPC**

- **gateway**: Прокси-сервер, который перенаправляет запросы на RPC серверы.

## Улучшения и особенности реализации

- **Авторизация**: Используются JWT токены, включающие access и refresh токены. Несмотря на то что текущий API требует Bearer токены, сервис можно легко адаптировать для работы с различными методами авторизации и секретными ключами.

- **Регистрация**: Пароли пользователей хешируются с помощью bcrypt перед сохранением в базе данных.

- **Обновление квартиры**: При изменении информации о квартире проверяется таблица подписчиков, и все подписчики, которые интересуются этой квартирой, получают уведомления через брокер сообщений. Другой сервис асинхронно читает сообщения из брокера и отправляет их подписчикам.

- **Поле flat_num**: В таблицу `flats` было добавлено поле `flat_num`, которое предназначено для хранения номера квартиры в доме. В текущей реализации это поле не используется в API.

- **Поле flat_num**: Реализовано кэширование с помощью Redis на авторизацию, home-service будет сохранять состояние токена для того чтобы сократить запросы на сервис авторизации

## Возможные улучшения

### Масштабирование БД

- **Репликация БД**: Использование репликации для обеспечения отказоустойчивости и повышения доступности. Это может включать создание нескольких реплик базы данных для распределения нагрузки и обеспечения резервного копирования. 

- **Шардирование**: Деление базы данных на более мелкие части, Для  обеспечивания горизонтального масштабирования. Каждая шард хранит подмножество данных, что позволяет распределить нагрузку по нескольким серверам и повысить общую производительность системы.
- **API**: Дополнить API чтобы оно реализовывало полноценный CRUD 

### Метрики и мониторинг graphana/prometheus/open-telemerty

Добавить метрики к примеру:

- **Метрики производительности**: 
  - **Время отклика**: Измерение времени ответа сервиса на запросы.
  - **Пропускная способность**: Количество обработанных запросов в единицу времени.
  - **Ошибка запросов**: Процент запросов, завершившихся с ошибками.

- **Мониторинг системных ресурсов**:
  - **Использование CPU**: Наблюдение за загрузкой процессора.
  - **Использование памяти**: Отслеживание потребления оперативной памяти.
  - **Дисковое пространство**: Мониторинг использования дискового пространства и нагрузки на диск.

### Интеграционные тесты

- **Тестирование взаимодействия между сервисами**: Проверка корректности обмена данными и взаимодействия между различными сервисами в системе.

- **Тестирование API**: Проверка того, что API корректно обрабатывает запросы и возвращает ожидаемые ответы.

### Нагрузочное тестирование

- **Симуляция нагрузки**: Проведение тестов, имитирующих большой объем запросов для оценки производительности системы под нагрузкой.

- **Анализ масштабируемости**: Проверка, как система справляется с увеличением количества пользователей и запросов, и как можно улучшить её масштабируемость.



