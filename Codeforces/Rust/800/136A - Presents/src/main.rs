use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let n: usize = input.trim().parse().expect("It's not number");

    let mut input_nums = String::new();
    let _ = io::stdin().read_line(&mut input_nums);

    let p_list: Vec<usize> = input_nums
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let mut result = vec![0; n + 1];

    for (i, &p) in p_list.iter().enumerate() {
        result[p] = i + 1;
    }

    for i in 1..=n {
        print!("{} ", result[i])
    }
    println!()
}
