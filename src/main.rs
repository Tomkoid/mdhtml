use std::process::exit;

use clap::CommandFactory;
use clap::Parser;
use spinners::Spinners;

const SPINNER_TYPE: Spinners = Spinners::Dots6;

mod args;
mod convert;
mod http;
pub mod tools;

#[tokio::main]
async fn main() {
    let args = args::Args::parse();

    match args.command {
        Some(args::Commands::Convert(mut args)) => {
            // if raw and style is set, exit
            if args.raw && args.style.is_some() {
                eprintln!("Error: --raw and --style are mutually exclusive.");
                eprintln!("Please choose one or the other.");
                exit(1);
            }

            // if raw and no_external_libs is set, exit
            if args.raw && args.no_external_libs {
                eprintln!("Warning: --raw and --no-external-libs are mutually exclusive.");
                eprintln!("Will use raw mode instead.");
                args.no_external_libs = false;
            }

            if args.server {
                http::server::start_server(&args).await;
            }

            convert::convert(&args, true, None);
        }

        None => {
            args::Args::command().print_help().unwrap();
        }
    }
}
