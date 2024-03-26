use cities_json::get_random_cities;
use cities_json::CITIES as RUST_CITIES;
use pyo3::prelude::*;

#[pyclass]
#[derive(Debug)]
pub struct City {
    #[pyo3(get)]
    pub country: String,
    #[pyo3(get)]
    pub name: String,
    #[pyo3(get)]
    pub lat: f64,
    #[pyo3(get)]
    pub lng: f64,
    #[pyo3(get)]
    pub admin1: String,
    #[pyo3(get)]
    pub admin2: String,
}

#[pymethods]
impl City {
    fn __repr__(slf: PyRef<'_, Self>) -> PyResult<String> {
        return Ok(format!(
            "City: country={country}, name={name}, lng={lng}, lat={lat}, admin1={admin1}, admin2={admin2}",
            country = slf.country,
            name = slf.name,
            lng = slf.lng,
            lat = slf.lat,
            admin1 = slf.admin1,
            admin2 = slf.admin2,
        ));
    }
}

pub type Cities = Vec<City>;

#[pyfunction]
pub fn random_city() -> PyResult<City> {
    let rust_city = get_random_cities();
    let city = City {
        country: rust_city.country.to_string(),
        lat: rust_city.lat,
        lng: rust_city.lng,
        name: rust_city.name.to_string(),
        admin1: rust_city.admin1.to_string(),
        admin2: rust_city.admin2.to_string(),
    };
    return Ok(city);
}

#[pyfunction]
pub fn all_cities() -> PyResult<Cities> {
    let mut cities: Cities = vec![];
    for rust_city in RUST_CITIES.iter(){
        cities.push(City {
            country: rust_city.country.to_string(),
            lat: rust_city.lat,
            lng: rust_city.lng,
            name: rust_city.name.to_string(),
            admin1: rust_city.admin1.to_string(),
            admin2: rust_city.admin2.to_string(),
        })
    }
    return Ok(cities);
}


#[pymodule]
fn citiespy(_py: Python, m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(random_city, m)?)?;
    m.add_function(wrap_pyfunction!(all_cities, m)?)?;
    Ok(())
}
