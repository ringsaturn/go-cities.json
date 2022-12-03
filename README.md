# Multi languages binding for `cities.json`

- Go: <https://pkg.go.dev/github.com/ringsaturn/go-cities.json>
- Rust: <https://crates.io/crates/cities-json>
  ```rust
  use cities_json::CITIES;
  println!("first {:?}", CITIES.get(0).unwrap().name);
  ```
