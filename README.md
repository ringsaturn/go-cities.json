# Multi languages binding for `cities.json`

It's used for geo releted benchmarks, not for production usage.

- Go: <https://pkg.go.dev/github.com/ringsaturn/go-cities.json>
- Rust: <https://crates.io/crates/cities-json>
  ```rust
  use cities_json::CITIES;
  println!("first {:?}", CITIES.get(0).unwrap().name);
  ```
- Python: coming soon

## LICENSE

This project is open sourced under [The Unlicense](./LICENSE).

Upstream data is under [Creative Commons Attribution 3.0 License](https://github.com/lutangar/cities.json#licence).
