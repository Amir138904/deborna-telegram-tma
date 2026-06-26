


# DeborWin 🎲
> **The Ultimate Play-to-Earn (P2E) Tombola/Bingo Telegram Mini App** > *اولین مینی‌اپلیکیشن درآمدی دبرنای آنلاین در بستر تلگرام*

---

## 🌍 Choose Your Language / انتخاب زبان
* 🇬🇧 [English](#english-version)
* 🇮🇷 [فارسی ](#persian-version)

---

<a id="english-version"></a>
## 🇬🇧 English Version

DeborWin is a high-performance, real-time Play-to-Earn (P2E) Tombola/Bingo game built as a Telegram Mini App (TMA). Powered by a fast, concurrent Go backend, it connects players into competitive multi-user rooms while managing secure, automated wallet ecosystems and instant, trustless winner payouts.

### ✨ Key Features
- **Scalable Game Rooms:** Supports instant-start matchmaking for rooms of **30, 50, and 100 players**.
- **Anti-Whale Mechanics:** Strict limit of **maximum 4 game cards per user** to maintain fairness and strategic balance.
- **Dedicated Web3/In-App Wallet:** Personalized wallet for every registered user to seamlessly deposit, hold, and track funds.
- **Fully Automated Financial Bot:**
  - Collects entry fees securely at the start of each match.
  - Automatically calculates and distributes the winner's share directly to their in-app wallet upon game completion.
  - Routes the platform's service fee directly to the admin/treasury system.

### 🛠️ Architecture & Tech Stack
- **Backend & Core Game Engine:** **Go (Golang)** — utilizing Goroutines and WebSockets for ultra-low latency real-time state synchronization across 100+ concurrent players.
- **Telegram Bot API & Financial Layer:** **Python / Go** — robust automation for bot interactions, wallet updates, and transactions.
- **Frontend (Mini App UI):** **TypeScript / React** — responsive, lightweight, and nostalgic visual board fully integrated into the Telegram client interface.

### 📂 Repository Structure
```text
deborna-telegram-tma/
├── apps/
│   ├── backend-core/      # Go-based real-time game server & game loop
│   ├── telegram-bot/      # Python/Go bot for commands and financial webhooks
│   └── frontend-tma/      # React/TS Telegram Mini App interface
├── deployment/            # Docker Compose & Kubernetes configurations
└── README.md

```

---
<a id="persian-version"></a>
## 🇮🇷 نسخه فارسی

پلتفرم **دبروین (DeborWin)** یک مینی‌اپلیکیشن تلگرامی (TMA) پیشرفته و همزمان (Real-time) در سبک بازی‌های کسب درآمد (Play-to-Earn) است. این پروژه با بهره‌گیری از هسته قدرتمند و فوق‌سریع زبان Go، پتانسیل نوستالژیک بازی دبرنا را به یک سیستم رقابتی گروهی همراه با مدیریت مالی کاملاً خودکار و آنی تبدیل کرده است.

### ✨ ویژگی‌های کلیدی

* **اتاق‌های بازی مقیاس‌پذیر:** قابلیت شروع سریع مسابقات در گروه‌های **۳۰، ۵۰ و ۱۰۰ نفره**.
* **قانون ضد انحصار (Anti-Whale):** محدودیت خرید **حداکثر ۴ برگه بازی برای هر کاربر** جهت حفظ عدالت و بالا نگه‌داشتن شانس همه بازیکنان.
* **کیف پول اختصاصی درون‌برنامه‌ای:** اختصاص یک پنل و کیف پول مجزا به هر کاربر برای واریز، نگهداری و برداشت امن موجودی.
* **ربات مالی کاملاً خودکار:**
* جمع‌آوری حق ورودیه مچ‌ها در شروع بازی از کیف پول کاربران.
* محاسبه و واریز اتوماتیک سهم برنده مسابقه به کیف پول اختصاصی او بلافاصله پس از اتمام بازی.
* انتقال خودکار کارمزد پلتفرم (سهم ادمین) به حساب مرجع.



### 🛠️ معماری و تکنولوژی‌های به‌کاررفته

* **بک‌اند و موتور اصلی بازی:** **زبان Go (Golang)** — استفاده از Goroutines و WebSocket برای همگام‌سازی لحظه‌ای شماره‌ها بین ۱۰۰ کاربر همزمان با کمترین تاخیر ممکن.
* **ربات تلگرام و لایه مالی:** **پایتون / گو** — پیاده‌سازی اتوماسیون فرآیندهای مالی، مدیریت کیف پول‌ها و ارسال نوتیفیکیشن‌ها.
* **فرانت‌اند (ظاهر مینی‌اپ):** **TypeScript / React** — طراحی یک رابط کاربری سبک، روان و نوستالژیک کاملاً سازگار با مرورگر داخلی تلگرام.

### 📂 ساختار ریپازیتوری

```text
deborna-telegram-tma/
├── apps/
│   ├── backend-core/      # سرور اصلی بازی و مدیریت وب‌سوکت‌ها (Go)
│   ├── telegram-bot/      # ربات تلگرام و مدیریت تراکنش‌های مالی
│   └── frontend-tma/      # رابط کاربری مینی‌اپ تلگرام (React/TS)
├── deployment/            # تنظیمات داکر و دپلویمنت سرور
└── README.md

```

---

Developed for [deborna-telegram-tma](https://github.com/ramtin-2024/deborna-telegram-tma) with 🎲 and 💻.

```

```
