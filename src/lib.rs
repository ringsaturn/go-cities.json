use std::vec;
use lazy_static::lazy_static;
use rand::seq::SliceRandom;
use serde::{Deserialize, Serialize};

lazy_static! {
    pub static ref CITIES: Cities = get_cities();
}

/// ```rust
/// use cities_json::CITIES;
/// println!("first {:?}", CITIES.get(0).unwrap().name);
/// ```
pub type Cities = Vec<City>;

#[derive(Debug)]
pub struct City {
    pub country: String,
    pub name: String,
    pub lat: f64,
    pub lng: f64,
    pub admin1: String,
    pub admin2: String,
}

type RawCities = Vec<RawCity>;

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
struct RawCity {
    country: String,
    name: String,
    lat: String,
    lng: String,
    admin1: String,
    admin2: String,
}

fn load_json_data() -> Vec<u8> {
    include_bytes!("../upstream/cities.json").to_vec()
}

fn get_cities() -> Cities {
    let raw_cities: RawCities = serde_json::from_slice(&load_json_data()[..]).unwrap();
    let mut cities: Cities = vec![];
    for rawcity in raw_cities.iter() {
        cities.push(City {
            country: rawcity.country.to_owned(),
            name: rawcity.name.to_owned(),
            lat: rawcity.lat.parse::<f64>().unwrap(),
            lng: rawcity.lng.parse::<f64>().unwrap(),
            admin1: rawcity.admin1.to_owned(),
            admin2: rawcity.admin2.to_owned(),
        })
    }
    return cities;
}


/// ```rust
/// use cities_json::get_random_cities;
/// println!("random: {:?}", get_random_cities());
/// ```
pub fn get_random_cities() -> &'static City{
    return CITIES.choose(&mut rand::thread_rng()).unwrap();
}
