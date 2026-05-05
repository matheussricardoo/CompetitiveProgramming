use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let mut y = input.trim().parse::<i64>().expect("It's not number");

    loop {
        y += 1;
        let a = y / 1000;
        let b = (y / 100) % 10;
        let c = (y / 10) % 10;
        let d = y % 10;

        if a != b && a != c && a != d && b != c && b != d && c != d {
            println!("{}", y);
            break;
        }
    }
}
