use std::io::{self, BufRead};

fn read_line_tokens() {
    let stdin = io::stdin();
    let mut handle = stdin.lock();
    let mut raw_line = String::new();
    
    handle.read_line(&mut raw_line).unwrap();
    let mut tokens = raw_line.split_whitespace();
    
    if let Some(token) = tokens.next() {
        let number: i32 = token.parse().unwrap();
    }
}