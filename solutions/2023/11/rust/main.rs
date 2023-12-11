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
    galaxies
        .iter()
        .enumerate()
        .flat_map(|(i, p1)| galaxies.iter().skip(i + 1).map(move |p2| p1.manhattan_distance(p2)))
        .sum()
}

fn parse_input(input: &str, add_gap: usize) -> Vec<Point> {
    let image: Vec<&str> = input.lines().collect();

    let y_gaps: Vec<usize> = image
        .iter()
        .scan(0, |y_gap, _| {
            let gap = add_gap;
            *y_gap += gap;
            Some(*y_gap)
        })
        .collect();

    let x_gaps: Vec<usize> = (0..image[0].len())
        .scan(0, |x_gap, _| {
            let gap = add_gap;
            *x_gap += gap;
            Some(*x_gap)
        })
        .collect();

    let galaxies: Vec<Point> = image
        .iter()
        .enumerate()
        .flat_map(|(y, line)| {
            line.chars().enumerate().filter_map(move |(x, char)| {
                if char == '#' {
                    Some(Point::new(x + x_gaps[x], y + y_gaps[y]))
                } else {
                    None
                }
            })
        })
        .collect();

    galaxies
}

fn main() {
    let runner = solution::Runner::new();
    let s = Box::new(solution::Solution {});
    runner.run(s)
}
