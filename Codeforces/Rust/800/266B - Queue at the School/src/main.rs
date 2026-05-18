use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);
    let parts: Vec<i32> = input
        .split_whitespace()
        .map(|x| x.parse().unwrap())
        .collect();
    let (n, t) = (parts[0] as usize, parts[1]);

    let mut s_input = String::new();
    let _ = io::stdin().read_line(&mut s_input);

    let mut queue: Vec<char> = s_input.trim().chars().collect();

    for _ in 0..t {
        let mut i = 0;
        while i < n - 1 {
            if queue[i] == 'B' && queue[i + 1] == 'G' {
                queue.swap(i, i + 1);
                i += 2
            } else {
                i += 1;
            }
        }
    }

    let result: String = queue.into_iter().collect();
    println!("{}", result);
}
