# lacert-crypto-go

[English](#english) · [Русский](#русский)

---

## English

**Post-quantum cryptographic core of the LACERT protocol, in Go.**

`lacert-crypto-go` is a self-contained library implementing the cryptographic
heart of LACERT — a fully local, cloud-free scheme for connecting IoT devices
to corporate networks with post-quantum security. The package has **no internal
dependencies on the rest of the project** and can be used on its own.

### What it does

- **Post-quantum key exchange** — ML-KEM-1024 (FIPS 203) for the shared secret,
  so traffic stays safe even against a future quantum adversary ("harvest now,
  decrypt later").
- **Mutually-authenticated handshake** built along the lines of the Noise_XX
  pattern, with the classic Diffie–Hellman step replaced by ML-KEM.
- **Continuous key rotation** — every session key is derived with a fresh
  post-quantum secret (BLAKE3), giving forward secrecy and post-compromise
  security.
- **Authenticated encryption** — ChaCha20-Poly1305 for the data channel.
- **Firmware integrity check** — ECDSA P-256 challenge–response.

### Design note

The signature is intentionally **classical (ECDSA P-256)**, not post-quantum.
Experiments showed that a hash-based post-quantum signature (SLH-DSA) is about
11 000× slower and 110× larger, which is unacceptable on constrained devices.
Post-quantum strength is therefore placed on the **key exchange**, where it is
cheap, and the signature stays fast.

### Install

```bash
go get github.com/NIkir0LL/lacert-crypto-go
```

Third-party dependencies (CIRCL, blake3, x/crypto) are fetched automatically by
Go modules — they are **not** vendored in this repository.

### Packages

| Package | Purpose |
|---------|---------|
| `crypto` | key exchange, handshake, rotation, AEAD, firmware check |
| `wire`   | message serialization for the protocol |

### License

Apache License 2.0 — see [LICENSE](LICENSE). You may use, modify and
redistribute it freely, including commercially.

---

## Русский

**Постквантовое криптографическое ядро протокола LACERT на языке Go.**

`lacert-crypto-go` — самостоятельная библиотека, реализующая криптографическую
основу LACERT: полностью локальной, без облака, схемы безопасного подключения
IoT-устройств к корпоративным сетям с постквантовой стойкостью. Пакет **не
зависит от остальной части проекта** и может использоваться отдельно.

### Что делает

- **Постквантовый обмен ключами** — ML-KEM-1024 (FIPS 203) для общего секрета,
  поэтому трафик защищён даже от будущего квантового противника (атака
  «перехвати сейчас — расшифруй потом»).
- **Взаимно-аутентифицированное рукопожатие** по образцу шаблона Noise_XX, где
  классический обмен Диффи–Хеллмана заменён на ML-KEM.
- **Непрерывная ротация ключей** — каждый сессионный ключ выводится со свежим
  постквантовым секретом (BLAKE3), что даёт прямую секретность и восстановление
  стойкости после компрометации.
- **Аутентифицированное шифрование** — ChaCha20-Poly1305 для канала данных.
- **Проверка целостности прошивки** — по схеме «запрос-ответ» на ECDSA P-256.

### О выборе алгоритмов

Подпись намеренно оставлена **классической (ECDSA P-256)**, а не постквантовой.
Эксперименты показали, что постквантовая хеш-подпись SLH-DSA примерно в 11 000
раз медленнее и в 110 раз объёмнее, что неприемлемо для устройств с
ограниченными ресурсами. Поэтому постквантовая стойкость обеспечена на уровне
**обмена ключами**, где она обходится дёшево, а подпись остаётся быстрой.

### Установка

```bash
go get github.com/NIkir0LL/lacert-crypto-go
```

Внешние зависимости (CIRCL, blake3, x/crypto) подтягиваются менеджером модулей
Go автоматически — в репозитории они **не хранятся**.

### Пакеты

| Пакет | Назначение |
|-------|------------|
| `crypto` | обмен ключами, рукопожатие, ротация, AEAD, проверка прошивки |
| `wire`   | сериализация сообщений протокола |

### Лицензия

Apache License 2.0 — см. [LICENSE](LICENSE). Разрешено свободно использовать,
изменять и распространять, в том числе в коммерческих целях.
