# Competitive Programming

Personal Codeforces solutions in Go and Rust, organized by language, rating and contest. The repository is for algorithm practice, not evidence of a deployed backend in either language. Other platforms are future additions.

## Structure

- `Codeforces/Golang/800/`: standalone Go solutions.
- `Codeforces/Rust/800/` and `900/`: Rust solutions with Cargo manifests.
- `Codeforces/Rounds/`: contest-specific submissions.
- `Support functions/`: runnable input-reading examples, not problem solutions.
- `Cargo.toml`: workspace for the Rust programs.

## Run and validate

Requires Go and/or Rust. No cloud account or paid service is needed.

```bash
cargo check --workspace
cd 'Codeforces/Golang/800/4A - Watermelon'
printf '8\n' | go run .
```

Expected output for that example: YES. For Rust, run cargo in the selected problem directory, or choose its workspace package. Problem READMEs describe approaches and complexity.

The Go input helpers return EOF instead of looping indefinitely. The full-line helper supports lines longer than a reader buffer. Helpers are examples for well-formed contest input, not general parsing libraries.

## Scope of verification

Compilation and selected sample cases were checked during the portfolio review. That does not establish that every submission is accepted or every edge case is covered. There are 153 Go/Rust source files in the reviewed snapshot, including support examples; this is not a solved-problem count.

## Author and license

[Matheus Ricardo](https://github.com/matheussricardoo) · [LinkedIn](https://www.linkedin.com/in/matheussricardoo/). See [LICENSE](LICENSE).
