#[macro_use] extern crate rocket;
use rocket::{ serde::json::Json};
use rocket::post;
use reqwest::Client;
use serde::{Deserialize, Serialize};
use rocket_cors::{AllowedOrigins, CorsOptions};
use rocket::routes;

#[derive(Debug, Serialize, Deserialize)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String,
}

#[post("/data", data = "<data>")]
async fn send_data(data: Json<Data>) -> String {
    let client = Client::new();
    let server_url = "http://localhost:5050/datos";
    let response = client.post(server_url).json(&data.into_inner()).send().await;

    match response {
        Ok(_) => "Data sent successfully!".to_string(),
        Err(e) => format!("Failed to send data: {}", e),
    }
}

#[launch]
fn rocket() -> _ {
    let cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .to_cors()
        .expect("failed to create CORS fairing");

    println!("Client Listening On Port 5000");

    rocket::build()
        .configure(rocket::Config::figment()
            .merge(("port", 5000))
            .merge(("address", "0.0.0.0")))
        .attach(cors)
        .mount("/rust", routes![send_data])
}