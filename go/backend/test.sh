#!/bin/bash

output_file="all_go_files.txt"  # 出力ファイルの名前を指定します
> $output_file  # 出力ファイルを空にして初期化します

# findコマンドを使用してカレントディレクトリ以下のすべての.goファイルを検索し、ループで処理します
find . -type f -name "*.go" | while read file; do
    echo "File: $file" >> $output_file  # ファイル名を書き込みます
    awk '{print FNR ": " $0}' "$file" >> $output_file  # 各行に行番号を付けてファイルの内容を書き込みます
    echo "" >> $output_file  # ファイル間に空行を追加します
done