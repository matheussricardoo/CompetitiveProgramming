use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let nums: Vec<i32> = input
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();

    let (n, m) = (nums[0], nums[1]);

    for i in 0..n {
        if i % 2 == 0 {
            println!("{}", "#".repeat(m as usize));
        } else {
            if i % 4 == 1 {
                println!("{}#", ".".repeat((m - 1) as usize));
            } else {
                println!("#{}", ".".repeat((m - 1) as usize))
            }
        }
    }
}
