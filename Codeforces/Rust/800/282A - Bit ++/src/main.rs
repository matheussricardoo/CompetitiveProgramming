use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin()
        .read_line(&mut input)
        .expect("It's not a string");

    let mut s = 0;

    let n: usize = input.trim().parse().expect("It's not a number");

    for _ in 0..n {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let instructions = input.trim();

        if instructions.contains("+") {
            s += 1;
        } else {
            s -= 1;
        }
    }
    println!("{}", s);
}
