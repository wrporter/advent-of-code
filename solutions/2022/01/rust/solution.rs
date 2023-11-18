use std::any::Any;
use std::cmp::max;
use std::fmt::Display;
use std::time::Instant;

use crate::color;

pub trait AnyDebug: Any + Display {}

impl<T> AnyDebug for T where T: Any + Display {}

pub trait AbstractSolution {
    fn year(&self) -> u32;
    fn day(&self) -> u32;
    fn part_1(&self, input: &str) -> Box<dyn AnyDebug>;
    fn part_2(&self, input: &str) -> Box<dyn AnyDebug>;
}

pub struct Runner;

pub struct Solution;

impl Runner {
    pub fn new() -> Self {
        Self {}
    }

    pub fn read_input(&self) -> &str {
        include_str!("../input.txt")
    }

    pub fn run(&self, solution: Box<dyn AbstractSolution>) {
        println!("🎄 {}{}{}{}: Day {}{}", color::GREEN, color::UNDERLINED, color::BOLD, solution.year(), solution.day(), color::RESET);

        let input = self.read_input();
        let start_total = Instant::now();

        let start1 = Instant::now();
        let answer1 = solution.part_1(input);
        let elapsed1 = start1.elapsed();

        let start2 = Instant::now();
        let answer2 = solution.part_2(input);
        let elapsed2 = start2.elapsed();

        let answer1_f = format!("{}", answer1);
        let answer2_f = format!("{}", answer2);
        let padding = max(answer1_f.len(), answer2_f.len());
        // let padding = ints.Max(len(fmt.Sprintf("%v", answer1)), len(fmt.Sprintf("%v", answer2)));

        println!("⭐  {}Part 1: {}{:<padding$}{} {}|{} 🕒 {}{:?}{}", color::GREEN, color::RED, answer1_f, color::RESET, color::CYAN, color::RESET, color::PURPLE, elapsed1, color::RESET);
        println!("⭐  {}Part 2: {}{:<padding$}{} {}|{} 🕒 {}{:?}{}", color::GREEN, color::RED, answer2_f, color::RESET, color::CYAN, color::RESET, color::PURPLE, elapsed2, color::RESET);

        let elapsed_total = start_total.elapsed();
        println!("🕒 {}Total: {:?}{}", color::BLUE, elapsed_total, color::RESET)
    }
}
