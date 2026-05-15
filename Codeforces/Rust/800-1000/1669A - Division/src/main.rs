use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let t: i64 = input.trim().parse().unwrap();

    for _ in 0..t {
        input.clear();
        let _ = io::stdin().read_line(&mut input);

        let rating: i64 = input.trim().parse().unwrap();

        if 1900 <= rating {
            println!("Division 1")
        } else if 1600 <= rating && rating <= 1899 {
            println!("Division 2")
        } else if 1400 <= rating && rating <= 1599 {
            println!("Division 3")
        } else if rating <= 1399 {
            println!("Division 4")
        }
    }
}
