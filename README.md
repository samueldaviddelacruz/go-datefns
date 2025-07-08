# 📆 go-datefns

A modern Go date/time utility library inspired by [`date-fns`](https://date-fns.org/).  
Simple, predictable, and idiomatic time manipulation helpers, written in Go.

---

## 🚀 Motivation

This project started as a personal learning exercise to port common utilities from JavaScript’s `date-fns` to Go. Over time, the goal is to build a lightweight, well-tested library with readable APIs for common date/time operations in Go.

---

## ✨ Features (so far)

- ✅ Pure functions — no mutation
- ✅ Tested with 100% coverage
- ✅ No dependencies
- ✅ Idiomatic Go naming & documentation

---

## 🧰 Currently Available Functions

### 📅 Date Arithmetic

- `AddMilliseconds(date, amount)`
- `AddMinutes(date, amount)`
- `AddHours(date, amount)`
- `AddDays(date, amount)`
- `AddMonths(date, amount)`
- `AddBusinessDays(date, amount)`

### 📊 Difference Calculations

- `DifferenceInExactDays(later, earlier)`  
  → `float64`, includes fractions
- `DifferenceInCalendarDays(later, earlier)`  
  → `int`, ignores time component
- `DifferenceInDaysWithRounding(later, earlier, method)`  
  → support for rounding (`RoundDown`, `RoundUp`, `RoundNearest`)

### 📌 Utilities

- `IsWeekend(date)`  
  → true if Saturday or Sunday

---

## 📦 Installation

```bash
go get github.com/samueldaviddelacruz/go-datefns
```

Then in your Go code:

```go
import "github.com/samueldaviddelacruz/go-datefns/datefns"
```

---

## 🧪 Example Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/samueldaviddelacruz/go-datefns/datefns"
)

func main() {
	now := time.Now()

	nextWeek := datefns.AddDays(now, 7)
	fmt.Println("Next week:", nextWeek)

	diff := datefns.DifferenceInExactDays(nextWeek, now)
	fmt.Println("Days apart:", diff)

	fmt.Println("Is weekend?", datefns.IsWeekend(nextWeek))
}
```

---

## 🧱 Project Structure

```
go-datefns/
├── datefns/
│   ├── datefns.go       # Core library functions
│   └── datefns_test.go  # Tests
├── go.mod
└── README.md
```

---

## 🛣️ Roadmap

- [ ] Add `StartOfDay`, `EndOfDay`, etc.
- [ ] Add `IsToday`, `IsTomorrow`, `IsSameDay`
- [ ] Add `Parse`, `Format` (with layout helpers?)
- [ ] Add week/month/year helpers
- [ ] Add duration formatters (`FormatDistance`, etc.)

---

## 🧠 Why Not Use the Standard Library?

Go's `time` package is powerful but sometimes verbose or low-level for common tasks.  
This library focuses on ergonomics and reusability, especially for:

- Applications that require frequent date math
- Testing time-based logic
- Simplifying your business logic

---

## 🙌 Contributing

Ideas, suggestions, and PRs are welcome! This project is intentionally simple and educational — feel free to fork it, clone it, or use it to learn Go better.

---

## 📄 License

MIT © 2025 Samuel David De La Cruz Portorreal
