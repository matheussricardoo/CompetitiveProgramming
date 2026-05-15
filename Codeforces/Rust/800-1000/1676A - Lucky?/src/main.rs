use std::io;

fn main() {
    let mut input_t = String::new();
    io::stdin().read_line(&mut input_t).unwrap();
    let t: usize = input_t.trim().parse().unwrap();

    for _ in 0..t {
        let mut ticket = String::new();
        io::stdin().read_line(&mut ticket).unwrap();
        let ticket = ticket.trim();

        let digits: Vec<u32> = ticket.chars().map(|c| c.to_digit(10).unwrap()).collect();

        let sum_first: u32 = digits[0..3].iter().sum();
        let sum_last: u32 = digits[3..6].iter().sum();

        if sum_first == sum_last {
            println!("YES");
        } else {
            println!("NO");
        }
    }
}
