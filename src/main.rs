use std::path::PathBuf;
use structopt::StructOpt;

#[derive(Debug, StructOpt)]
#[structopt(name = "classtame", about = "Manage classes with ease.")]
struct Opt {
    #[structopt(name = "command")]
    command: String,
    #[structopt(required_if("command", "info"))]
    class: Option<String>,
    #[structopt(
        default_value = "~/.classtame.yaml",
        short = "c",
        long,
        parse(from_os_str)
    )]
    config: PathBuf,
}

fn main() {
    let args = Opt::from_args();
}
