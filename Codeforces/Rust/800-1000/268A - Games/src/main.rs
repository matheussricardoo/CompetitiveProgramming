use std::io;

fn main() {
    let mut input = String::new();
    let _ = io::stdin().read_line(&mut input);

    let n: usize = input.trim().parse().unwrap();

    let mut uniforms = Vec::new();
    for _ in 0..n {
        let mut line = String::new();
        let _ = io::stdin().read_line(&mut line).unwrap();
        let cols: Vec<i32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();
        uniforms.push((cols[0], cols[1]));
    }

    let mut count = 0;
    for i in 0..n {
        for j in 0..n {
            if i != j {
                if uniforms[i].0 == uniforms[j].1 {
                    count += 1;
                }
            }
        }
    }
    println!("{}", count)
}
