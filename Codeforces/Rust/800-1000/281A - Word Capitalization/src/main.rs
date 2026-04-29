use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let world = input.trim();
    let mut ch = world.chars();
    let capitalized = match ch.next() {
        None => String::new(),
        Some(f) => f.to_uppercase().collect::<String>() + ch.as_str(),
    };

    println!("{}", capitalized);
}
