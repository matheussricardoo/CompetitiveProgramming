use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let mut iter = input.split_whitespace();

    let a: i32 = iter.next().unwrap().parse().expect("It's not a number");
    let mut b: i32 = iter.next().unwrap().parse().expect("It's not a number");

    let mut count = 0;
    let mut i = a;
    while i <= b {
        i *= 3;
        b *= 2;
        count += 1
    }
    println!("{}", count);
}
