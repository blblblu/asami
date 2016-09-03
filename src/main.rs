// extern crate rand;
extern crate clap;
extern crate image;

use std::path::Path;
// use rand::Rng;
use clap::{Arg, App};
use image::{DynamicImage, GenericImage, RgbaImage};

fn main() {
    let matches = App::new("asami")
        .version("0.1.0")
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
        .get_matches();

    let input = matches.value_of("input").unwrap();
    let output = matches.value_of("output").unwrap();

    handle_request(input, output);
}

fn handle_request(input: &str, output: &str) {
    println!("input: {}, output: {}", input, output);

    let ref in_img = image::open(&Path::new(input)).unwrap();

    println!("dimensions {:?}", in_img.dimensions());
    println!("{:?}", in_img.color());

    let out_img = sort_them_pixels(in_img);

    let _ = out_img.save(output).unwrap();
}

fn sort_them_pixels(input: &DynamicImage) -> RgbaImage {
    let mut output = RgbaImage::new(input.width(), input.height());

    let mut x = 0;
    let mut y = 0;

    let mut tmp = Vec::new();

    for (_, _, pixel) in input.pixels() {
        tmp.push(pixel);
        if tmp.len() > 100 {
            tmp.sort_by(|a, b| (a[3]).cmp(&b[1]));
            while !tmp.is_empty() {
                output.put_pixel(x, y, tmp.pop().unwrap());
                y = y + (x + 1) / input.width();
                x = (x + 1) % input.width();
            }
        }
    }
    output
}
