use std::io::{self, Read};

struct FastScanner {
    buffer: Vec<u8>,
    position: usize,
}

impl FastScanner {
    fn new() -> Self {
        let mut buffer = Vec::new();
        io::stdin().read_to_end(&mut buffer).unwrap();
        Self { buffer, position: 0 }
    }

    fn next_int(&mut self) -> i32 {
        while self.position < self.buffer.len() && self.buffer[self.position] <= b' ' {
            self.position += 1;
        }
        if self.position >= self.buffer.len() { return 0; }
        let mut sign = 1;
        if self.buffer[self.position] == b'-' {
            sign = -1;
            self.position += 1;
        }
        let mut res = 0;
        while self.position < self.buffer.len() && self.buffer[self.position] >= b'0' && self.buffer[self.position] <= b'9' {
            res = res * 10 + (self.buffer[self.position] - b'0') as i32;
            self.position += 1;
        }
        res * sign
    }

    fn next_string(&mut self) -> String {
        while self.position < self.buffer.len() && self.buffer[self.position] <= b' ' {
            self.position += 1;
        }
        let start = self.position;
        while self.position < self.buffer.len() && self.buffer[self.position] > b' ' {
            self.position += 1;
        }
        String::from_utf8_lossy(&self.buffer[start..self.position]).into_owned()
    }
}