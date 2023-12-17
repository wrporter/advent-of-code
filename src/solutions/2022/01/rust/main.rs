use std::any::Any;
use std::fmt::Debug;
use crate::solution::{AbstractSolution, AnyDebug, Solution};

mod solution;
mod color;


impl AbstractSolution for Solution {
    fn year(&self) -> u32 { 2022 }
    fn day(&self) -> u32 { 1 }

    fn part_1(&self, input: &str) -> Box<dyn AnyDebug> {
        let max = input.split("\n\n")
            .map(|elf| elf.lines()
                .map(|item| item.parse::<u32>().unwrap()).sum::<u32>())
            .max()
            .unwrap();

        Box::new(max)
    }

    fn part_2(&self, input: &str) -> Box<dyn AnyDebug> {
        let mut elves: Vec<u32> = input.split("\n\n")
            .map(|elf| elf.lines().map(|item| item.parse::<u32>().unwrap()).sum())
            .collect();

        elves.sort_by(|a, b| b.cmp(a));
        let top3: u32 = elves.into_iter().take(3).sum();

        Box::new(top3)
    }
}

fn main() {
    let runner = solution::Runner::new();
    let s = Box::new(solution::Solution {});
    runner.run(s)
}
