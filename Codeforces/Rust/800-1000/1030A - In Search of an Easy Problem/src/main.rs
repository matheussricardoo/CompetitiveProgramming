use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let _n: i64 = input.trim().parse().expect("It's not number");

    input.clear();
    let _ = io::stdin().read_line(&mut input);

    if input.contains('1') {
        println!("HARD")
    } else {
        println!("EASY")
    }
}
