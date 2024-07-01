use std::{cell::RefCell, sync::{Arc, Mutex}};

use axum::{routing::get, Router};
use tokio::sync::broadcast::{self, Receiver, Sender};

use crate::args::Convert;

use super::{
    message::{ChanMessage, Messages},
    watch::watch,
};

#[derive(Clone)]
pub struct AppState {
    pub args: Convert,
    pub messages: Arc<Mutex<Vec<ChanMessage>>>,
    pub messages_object: Arc<Mutex<Messages>>,
    pub tx: Sender<ChanMessage>,
    //pub rx: Receiver<ChanMessage>,
}

pub async fn start_server(args: &Convert) {
    let messages = Messages::new();

    let (tx, rx): (Sender<ChanMessage>, Receiver<ChanMessage>) = broadcast::channel(16);

    let app_state = AppState {
        args: args.clone(),
        messages: Mutex::new(Vec::new()).into(),
        messages_object: Mutex::new(messages.clone()).into(),
        tx,
        //rx
    };

    let state = Arc::new(app_state);

    // build our application with a single route
    let app = Router::new()
        .route("/", get(super::index::get_index))
        .route("/update", get(super::message::Messages::send_update_async))
        .route("/messages", get(super::message::Messages::messages_html))
        .route("/ws", get(super::websocket::ws_handler))
        .with_state(state.clone());

    let args_temp = args.clone();

    tokio::spawn(async move {
        watch(&args_temp, axum::extract::State(state))
            .await
            .unwrap();
    });

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind(format!("{}:{}", args.hostname, args.port))
        .await
        .unwrap();
    axum::serve(listener, app).await.unwrap();
}
