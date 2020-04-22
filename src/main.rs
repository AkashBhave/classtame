use std::path::PathBuf;
use structopt::StructOpt;

#[derive(Debug, StructOpt)]
#[structopt(name = "classtame", about = "Manage classes with ease.")]
enum Opt {
    Info {},
    Open {},
    Mod {},
}

fn main() {
    let args = Opt::from_args();
    println!("{:?}", &args)
}
