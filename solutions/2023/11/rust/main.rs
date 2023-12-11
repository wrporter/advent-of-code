use std::fmt::Debug;
use crate::solution::{AbstractSolution, AnyDebug, Solution};

mod solution;
mod color;

#[derive(Debug, Copy, Clone)]
struct Point {
  x: usize,
  y: usize,
}

impl Point {
  fn new(x: usize, y: usize) -> Point {
      Point { x, y }
  }

  fn manhattan_distance(&self, other: &Point) -> usize {
      (self.x as isize - other.x as isize).abs() as usize
          + (self.y as isize - other.y as isize).abs() as usize
  }
}


impl AbstractSolution for Solution {
    fn year(&self) -> u32 { 2023 }
    fn day(&self) -> u32 { 11 }

    fn part_1(&self, input: &str) -> Box<dyn AnyDebug> {
        Box::new(get_total_distance(input, 2))
    }

    fn part_2(&self, input: &str) -> Box<dyn AnyDebug> {
        Box::new(get_total_distance(input, 1_000_000))
    }
}

fn get_total_distance(input: &str, gap_size: usize) -> usize {
    let galaxies = parse_input(input, gap_size - 1);
    let mut sum = 0;

    for i in 0..galaxies.len() {
        for j in (i + 1)..galaxies.len() {
            sum += galaxies[i].manhattan_distance(&galaxies[j]);
        }
    }

    sum
}

fn parse_input(input: &str, add_gap: usize) -> Vec<Point> {
    let image: Vec<&str> = input.lines().collect();
    let mut y_gaps = vec![0; image.len()];
    let mut y_gap = 0;

    for (y, line) in image.iter().enumerate() {
        let mut gap = add_gap;

        for x in 0..line.len() {
            if line.chars().nth(x).unwrap() == '#' {
                gap = 0;
            }
        }

        y_gap += gap;
        y_gaps[y] = y_gap;
    }

    let mut x_gaps = vec![0; image[0].len()];
    let mut x_gap = 0;

    for x in 0..image[0].len() {
        let mut gap = add_gap;

        for y in 0..image.len() {
            if image[y].chars().nth(x).unwrap() == '#' {
                gap = 0;
            }
        }

        x_gap += gap;
        x_gaps[x] = x_gap;
    }

    let mut galaxies = Vec::new();

    for (y, line) in image.iter().enumerate() {
        for (x, char) in line.chars().enumerate() {
            if char == '#' {
                galaxies.push(Point::new(x + x_gaps[x], y + y_gaps[y]));
            }
        }
    }

    galaxies
}

fn main() {
    let runner = solution::Runner::new();
    let s = Box::new(solution::Solution {});
    runner.run(s)
}
