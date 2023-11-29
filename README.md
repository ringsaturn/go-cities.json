# Multi languages binding for `cities.json`

> [!NOTE]
>
> It's used for geo related benchmarks, not for production usages.

- Go: <https://pkg.go.dev/github.com/ringsaturn/go-cities.json>
- Rust: <https://crates.io/crates/cities-json>
  ```bash
  cargo add cities-json
  ```
  ```rust
  use cities_json::CITIES;
  println!("first {:?}", CITIES.get(0).unwrap().name);
  ```
- Python: <https://pypi.org/project/citiespy/>
  ```bash
  pip install citiespy
  ```
  ```python
  from citiespy import all_cities, random_city

  city = random_city()
  print(city)

  cities = all_cities()
  print(len(cities))
  ```

## LICENSE

This project is open sourced under [The Unlicense](./LICENSE).

Upstream data is under
[Creative Commons Attribution 3.0 License](https://github.com/lutangar/cities.json#licence).
