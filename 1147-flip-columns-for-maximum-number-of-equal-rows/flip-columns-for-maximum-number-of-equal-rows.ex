defmodule Solution do
  @spec max_equal_rows_after_flips(matrix :: [[integer]]) :: integer
  def max_equal_rows_after_flips(matrix) do
     # Map to store the frequency of patterns
    bit_map =
      Enum.reduce(matrix, %{}, fn row, acc ->
        # Construct the pattern based on the first element of the row
        pattern =
          if hd(row) == 0 do
            Enum.map(row, &Integer.to_string(&1))
          else
            Enum.map(row, &Integer.to_string(1 - &1))
          end
          |> Enum.join("")

        # Update the frequency map
        Map.update(acc, pattern, 1, &(&1 + 1))
      end)

    # Find the maximum value in the frequency map
    bit_map
    |> Map.values()
    |> Enum.max()
  end
end