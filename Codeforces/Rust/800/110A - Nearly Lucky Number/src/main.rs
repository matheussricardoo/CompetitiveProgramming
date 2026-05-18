use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let lucky_count = input
        .trim()
        .chars()
        .filter(|&c| c == '4' || c == '7')
        .count();

    if lucky_count == 4 || lucky_count == 7 {
        println!("YES");
    } else {
        print!("NO")
    }
}
