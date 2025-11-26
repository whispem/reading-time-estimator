# Reading Time Estimator (Go mini project)

A super simple project: paste a sentence or paragraph, and it estimates how long it would take to read
at an average adult speed (default: 200 words per minute).

You can also try different reading speeds!

---

## How to use

1. Open this folder in your terminal.
2. Run:
   ```sh
   go run main.go
   ```
3. Paste your text, then press Enter.
4. The program will count words and show the estimated reading time.
5. Want to try another reading speed? Enter a value (in words per minute), or just press Enter to quit.

---

## How it works

- Counts words based on spaces.
- Uses default reading speed of 200 wpm (words per minute) but you can change it.
- Outputs an estimated reading time as “X min Y sec”.
