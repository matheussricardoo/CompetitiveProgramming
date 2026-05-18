use std::io;

fn main() {
    let mut input = String::new();
    let mut _moves = 0;

    for i in 1..=5i32 {
        input.clear();
        let _ = io::stdin().read_line(&mut input);
        let mut j = 1;
        for val_str in input.split_whitespace() {
            let v: i32 = val_str.parse().expect("It's not a number");
            if v == 1 {
                _moves = (i - 3).abs() + (j as i32 - 3).abs();
                println!("{}", _moves);
                return;
            }
            j += 1;
        }
    }
}
