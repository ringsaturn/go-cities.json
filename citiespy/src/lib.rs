use cities_json::get_random_cities;
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
}

#[pymethods]
impl City {
    fn __str__(slf: PyRef<'_, Self>) -> PyResult<String> {
        return Ok(format!(
            "City: country={country}, name={name}, lng={lng}, lat={lat}",
            country = slf.country,
            name = slf.name,
            lng = slf.lng,
            lat = slf.lat,
        ));
    }
}

#[pyfunction]
pub fn random_city() -> PyResult<City> {
    let raw_city = get_random_cities();
    let city = City {
        country: raw_city.country.to_string(),
        lat: raw_city.lat,
        lng: raw_city.lng,
        name: raw_city.name.to_string(),
    };
    return Ok(city);
}

#[pymodule]
fn citiespy(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(random_city, m)?)?;
    Ok(())
}
