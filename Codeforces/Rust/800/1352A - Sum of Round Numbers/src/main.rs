use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input).unwrap();
    let t: i32 = input.trim().parse().unwrap();

    for _ in 0..t {
        let mut n_str = String::new();
        let _ = io::stdin().read_line(&mut n_str).unwrap();
        let mut n: i32 = n_str.trim().parse().unwrap();

        let mut results = Vec::new();
        let mut multiplier = 1;

        while n > 0 {
            let digit = n % 10;
            if digit != 0 {
                results.push(digit * multiplier);
            }
            n /= 10;
            multiplier *= 10;
        }

        println!("{}", results.len());
        for (i, val) in results.iter().enumerate() {
            if i > 0 {
                print!(" ");
            }
            print!("{}", val);
        }
        println!();
    }
}
