use actix_web::{web, App, HttpResponse, HttpServer, Result};
use actix_cors::Cors;
use rdkafka::config::ClientConfig;
use rdkafka::producer::{BaseProducer,BaseRecord, Producer};
use std::time::Duration;

use serde_json::json;

#[derive(serde::Deserialize)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String,
}
async fn productor(json:&str){
    let proc: BaseProducer= ClientConfig::new()
    .set("bootstrap.servers","my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092")
    .create()
    .expect("Error al crear el productor");

    proc.send(
        BaseRecord::to("mi-topic")
        .key("data")
        .payload(json)
    ).expect("Error al generar mensaje");

    for _ in 0..5 {
        proc.poll(Duration::from_millis(10));
    }
    proc.flush(Duration::from_secs(1)).expect("No se pudo encolar el mensaje");
}
async fn receive_data(datos: web::Json<Data>) -> Result<HttpResponse> {
    let received_data = datos.into_inner();
    let data_json =json!({
            "name": received_data.name,
            "album": received_data.album,
            "year": received_data.year,
            "rank": received_data.rank
        }
    ).to_string();
    println!("{data_json}");
    productor(&data_json).await;

    Ok(HttpResponse::Ok().json(json!({})))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        let cors = Cors::default()
            .allow_any_origin()
            .allow_any_method()
            .allow_any_header();

        App::new()
            .wrap(cors)
            .wrap(actix_web::middleware::Logger::default()) // Agregar el middleware Logger
            .route("/datos", web::post().to(receive_data))
    })
    .bind(("0.0.0.0",5050))?
    .run()
    .await
}