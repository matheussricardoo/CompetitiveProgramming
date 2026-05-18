use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let n: i64 = input.trim().parse().unwrap();

    let mut result = Vec::new();

    for i in 1..=n {
        if i % 2 != 0 {
            result.push("I hate");
        } else {
            result.push("I love")
        }
    }

    println!("{} it", result.join(" that "))
}
