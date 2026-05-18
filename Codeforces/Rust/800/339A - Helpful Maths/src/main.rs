use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let s = input.trim();
    let mut parts: Vec<&str> = s.split('+').collect();
    parts.sort();
    let result = parts.join("+");
    if result.len() == 1 {
        println!("{}", result)
    } else {
        println!("{}", result)
    }
}
