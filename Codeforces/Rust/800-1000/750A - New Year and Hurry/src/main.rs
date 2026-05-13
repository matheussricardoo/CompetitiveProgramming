fn main() {
    let mut input = String::new();
    let _ = std::io::stdin().read_line(&mut input);

    let inputs: Vec<i64> = input
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let n: i64 = inputs[0];
    let k: i64 = inputs[1];

    let mut time_reamining = 240 - k;
    let mut solved_problems = 0;

    for i in 1..=n {
        let time_to_problem = 5 * i;

        if time_reamining >= time_to_problem {
            time_reamining -= time_to_problem;
            solved_problems += 1
        } else {
            break;
        }
    }
    println!("{}", solved_problems)
}
