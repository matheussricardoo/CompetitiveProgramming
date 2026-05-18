use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let s = input.trim();

    let upper_count = s.chars().filter(|c| c.is_uppercase()).count();
    let lower_count = s.len() - upper_count;

    if upper_count > lower_count {
        println!("{}", s.to_uppercase())
    } else {
        println!("{}", s.to_lowercase())
    }
}
