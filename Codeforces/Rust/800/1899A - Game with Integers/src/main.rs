fn main() {
    let mut input = String::new();
    let _ = std::io::stdin().read_line(&mut input);

    let n: i64 = input.trim().parse().unwrap();

    for _ in 0..n {
        let mut input = String::new();
        let _ = std::io::stdin().read_line(&mut input);

        let n: i64 = input.trim().parse().unwrap();

        if n % 3 != 0 {
            println!("First")
        } else {
            println!("Second")
        }
    }
}
