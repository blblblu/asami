extern crate rand;
#[macro_use]
extern crate clap;
extern crate image;

use std::path::Path;
use rand::Rng;
use clap::{Arg, App};
use image::{DynamicImage, GenericImage, RgbaImage};

arg_enum!{
    enum Mode {
        Brute
    }
}

fn main() {
    let matches = App::new("asami")
        .version("0.1.0")
        .author("Sebastian Schulz <mail@sesc.me>")
        .about("pixel sorter")
        .args_from_usage(
            "<INPUT> 'The input file to use'
             <OUTPUT> 'The output file to use'")
        .arg(Arg::from_usage("-m --mode <MODE> 'The mode to use")
            .possible_values(&Mode::variants())
            .default_value("Brute"))
        .get_matches();

    let input = matches.value_of("INPUT").unwrap();
    let output = matches.value_of("OUTPUT").unwrap();

    let mode = matches.value_of("mode").unwrap();
    println!("{}", mode);

    handle_request(input, output);
}

fn handle_request(input: &str, output: &str) {
    println!("input: {}, output: {}", input, output);

    let ref in_img = image::open(&Path::new(input)).unwrap();

    println!("dimensions {:?}", in_img.dimensions());
    println!("{:?}", in_img.color());

    let out_img = brute_sort(in_img, 32, 64);

    let _ = out_img.save(output).unwrap();
}

fn brute_sort(input: &DynamicImage, chunk_min_length: u32, chunk_max_length: u32) -> RgbaImage {
    let mut output = RgbaImage::new(input.width(), input.height());

    let mut out_x = 0;
    let mut out_y = 0;

    let mut tmp = Vec::new();

    let mut threshold = calculate_chunk_threshold(chunk_min_length, chunk_max_length);

    for (_, _, pixel) in input.pixels() {
        tmp.push(pixel);
        if tmp.len() > threshold as usize {
            threshold = calculate_chunk_threshold(chunk_min_length, chunk_max_length);
            tmp.sort_by(|a, b| (a[3]).cmp(&b[1]));
            write_pixel_to_image(&mut output, &mut out_x, &mut out_y, &mut tmp);
        }
    }
    output
}

fn calculate_chunk_threshold(chunk_min_length: u32, chunk_max_length: u32) -> u32 {
    rand::thread_rng().gen_range(chunk_min_length, chunk_max_length + 1)
}

fn write_pixel_to_image(image: &mut RgbaImage, x: &mut u32, y: &mut u32, pixels: &mut Vec<image::Rgba<u8>>){
    while !pixels.is_empty() {
        image.put_pixel(*x, *y, pixels.pop().unwrap());
        *y = *y + (*x + 1) / image.width();
        *x = (*x + 1) % image.width();
    }
}
