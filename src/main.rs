//extern crate rand;
extern crate clap;
extern crate image;

//use std::io;
//use std::cmp::Ordering;
//use std::fs::File;
use std::path::Path;
//use rand::Rng;
use clap::{Arg, App};
use image::{DynamicImage, GenericImage, RgbaImage};

/*fn foo() {

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
}*/

fn main() {
    let matches = App::new("asami")
        //.version("1.0")
        .author("Sebastian Schulz <mail@sesc.me>")
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

    handle_request(input, output);

    // Vary the output based on how many times the user used the "verbose" flag
    // (i.e. 'myprog -v -v -v' or 'myprog -vvv' vs 'myprog -v'
    /*match matches.occurrences_of("v") {
        0 => println!("No verbose info"),
        1 => println!("Some verbose info"),
        2 => println!("Tons of verbose info"),
        3 | _ => println!("Don't be crazy"),
    }*/
}

fn handle_request(input: &str, output: &str) {
    println!("input: {}, output: {}", input, output);

    let ref in_img = image::open(&Path::new(input)).unwrap();

    // The dimensions method returns the images width and height
    println!("dimensions {:?}", in_img.dimensions());

    // The color method returns the image's ColorType
    println!("{:?}", in_img.color());

    let ref mut out_img = RgbaImage::new(in_img.width(), in_img.height());

    sort_them_pixels(in_img, out_img);

    // Write the contents of this image to the Writer in PNG format.
    let _ = out_img.save(output).unwrap();
}

fn sort_them_pixels(input: &DynamicImage, output: &mut RgbaImage){
    for (x, y, pixel) in input.pixels() {
        output.put_pixel(x, y, pixel);
    }
}
