extern crate rand;
extern crate clap;
use clap::{Arg, App};

use std::io;
use std::cmp::Ordering;
use rand::Rng;

fn foo() {

    let secret_number = rand::thread_rng().gen_range(1, 101);

    let mut guess = String::new();

    io::stdin().read_line(&mut guess)
        .expect("failed to read line");

    let guess: i32 = guess.trim().parse()
        .expect("please type a number");

    match guess.cmp(&secret_number) {
        Ordering::Equal => println!("same"),
        Ordering::Greater => println!("too big"),
        Ordering::Less => println!("too small"),
    }
}

fn main() {

    let matches = App::new("asami")
        //.version("1.0")
        .author("Sebastian Schulz <contact@sesc.me>")
        .about("pixel sorter")
        .arg(Arg::with_name("input")
            .short("i")
            .long("input")
            .value_name("PATH")
            .help("path to input file")
            .takes_value(true)
            .required(true))
        .arg(Arg::with_name("output")
            .short("o")
            .long("output")
            .value_name("PATH")
            .help("path to output file")
            .takes_value(true)
            .required(true))
        /*.arg(Arg::with_name("v")
            .short("v")
            .multiple(true)
            .help("Sets the level of verbosity"))*/
        .get_matches();

    let input = matches.value_of("input").unwrap();
    let output = matches.value_of("output").unwrap();

    println!("input: {}, output: {}", input, output);

    // Vary the output based on how many times the user used the "verbose" flag
    // (i.e. 'myprog -v -v -v' or 'myprog -vvv' vs 'myprog -v'
    /*match matches.occurrences_of("v") {
        0 => println!("No verbose info"),
        1 => println!("Some verbose info"),
        2 => println!("Tons of verbose info"),
        3 | _ => println!("Don't be crazy"),
    }*/
}
