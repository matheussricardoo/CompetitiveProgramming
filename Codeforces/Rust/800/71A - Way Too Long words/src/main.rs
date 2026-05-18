use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let time: i32 = input.trim().parse().expect("It's not a number");

    for _ in 0..time {
        input.clear();
        io::stdin()
            .read_line(&mut input)
            .expect("Error read the word");

        let word = input.trim();
        let len = word.len();

        if len > 10 {
            let first = word.chars().next().unwrap();
            let last = word.chars().last().unwrap();

            println!("{}{}{}", first, len - 2, last);
        } else {
            println!("{}", word);
        }
    }
}
