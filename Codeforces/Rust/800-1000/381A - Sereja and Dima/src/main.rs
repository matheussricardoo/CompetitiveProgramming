use std::io;

fn main() {
    let mut input_n = String::new();
    io::stdin().read_line(&mut input_n).unwrap();
    let n: usize = input_n.trim().parse().unwrap();

    let mut input_cards = String::new();
    io::stdin().read_line(&mut input_cards).unwrap();
    let cards: Vec<i32> = input_cards
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let (mut left, mut right) = (0, n - 1);
    let (mut sereja_score, mut dima_score) = (0, 0);
    let mut is_sereja_turn = true;

    while left <= right {
        let picked_card: i32;

        if cards[left] > cards[right] {
            picked_card = cards[left];
            left += 1;
        } else {
            picked_card = cards[right];
            if right > 0 {
                right -= 1;
            } else if left == right {
                right = 0;
                if is_sereja_turn {
                    sereja_score += picked_card;
                } else {
                    dima_score += picked_card;
                }
                break;
            }
        }

        if is_sereja_turn {
            sereja_score += picked_card;
        } else {
            dima_score += picked_card;
        }

        is_sereja_turn = !is_sereja_turn;

        if left > n - 1 || (left > right && n > 1) {
            break;
        }
    }

    println!("{} {}", sereja_score, dima_score);
}
