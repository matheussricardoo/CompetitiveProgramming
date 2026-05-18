use std::io;

fn main() {
    let mut s = 0;
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: usize = input.trim().parse().expect("It's not a number");

    for _ in 0..n {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let mut iter = input.split_whitespace();

        let p: i32 = iter.next().unwrap().parse().expect("It's not a number");
        let v: i32 = iter.next().unwrap().parse().expect("It's not a number");
        let t: i32 = iter.next().unwrap().parse().expect("It's not a number");

        if p + v + t >= 2 {
            s += 1;
        }
    }

    println!("{}", s)
}
