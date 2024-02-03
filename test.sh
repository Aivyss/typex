#!/bin/bash

coverfile="cover.out"

total_lines=0
covered_lines=0

while IFS=, read -r range count; do
  if [[ $range == *"mode: set"* ]]; then
    continue
  fi

  # 라인별 커버리지 정보 추출 및 합산
  start=$(echo "$range" | cut -d "-" -f 1 | awk '{print int($1)}')
  end=$(echo "$range" | cut -d "-" -f 2 | awk '{print int($1)}')
  total_lines=$((total_lines + end - start + 1))
  covered_lines=$((covered_lines + count))

done < "$coverfile"

# 전체 라인 수가 0이 아닌 경우에만 커버리지 계산 수행
if [ $total_lines -gt 0 ]; then
  # 전체 커버리지 퍼센트 계산
  coverage_percent=$(echo "scale=2; $covered_lines * 100 / $total_lines" | bc)
  echo "Total Lines: $total_lines"
  echo "Covered Lines: $covered_lines"
  echo "Coverage Percentage: $coverage_percent%"
else
  echo "No coverage information available or total lines is 0."
fi
