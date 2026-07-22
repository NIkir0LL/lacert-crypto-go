// Copyright (c) 2025 NIkir0LL
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package crypto

import "time"

// Экспортированные точки управления временем для тестов в других пакетах
// (например, scheduler/gateway), которым нужно детерминированно проверять
// тайм-аут атомарной ротации без реальных задержек. В обычной сборке эти
// функции не используются рабочим кодом.

// NowForTest возвращает текущее (возможно, подменённое) время сессии.
func NowForTest() time.Time { return timeNow() }

// SetNowForTest подменяет источник времени, используемый логикой сессии
// (BeginRotate/AbortIfStale и т.п.).
func SetNowForTest(fn func() time.Time) { timeNow = fn }

// ResetNowForTest восстанавливает реальное время.
func ResetNowForTest() { timeNow = timeDefault }
