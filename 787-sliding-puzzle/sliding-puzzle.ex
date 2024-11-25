defmodule Solution do
  @rows 2
  @cols 3

  defmodule Aux do
    @spec solvable_boards(rows, cols) :: %{board => steps}
          when rows: pos_integer(),
               cols: pos_integer(),
               board: [[non_neg_integer()]],
               steps: non_neg_integer()
    def solvable_boards(rows, cols) do
      tiles = rows * cols

      solved =
        for i <- 0..rows - 1,
            j <- 0..cols - 1,
            into: %{},
            do: {{i, j}, rem(i * cols + j + 1, tiles)}

      [{solved, 0}]
      |> :queue.from_list()
      |> bfs(%{})
      |> Map.new(fn {board, steps} ->
        board
        |> Enum.sort()
        |> Enum.map(&elem(&1, 1))
        |> Enum.chunk_every(cols)
        |> then(&{&1, steps})
      end)
    end

    defp bfs(unquote(Macro.escape(:queue.new())), seen), do: seen

    defp bfs(queue, seen) do
      {{:value, {board, steps}}, queue} = :queue.out(queue)

      seen = Map.put(seen, board, steps)
      {i0, j0} = board |> Enum.find(&elem(&1, 1) == 0) |> elem(0)

      queue =
        for {i, j} <- [{i0 - 1, j0}, {i0 + 1, j0}, {i0, j0 - 1}, {i0, j0 + 1}],
            val = board[{i, j}],
            new_board = %{board | {i, j} => 0, {i0, j0} => val},
            !Map.has_key?(seen, new_board),
            reduce: queue do
          queue -> :queue.in({new_board, steps + 1}, queue)
        end
        
      bfs(queue, seen)
    end
  end

  @solvable_boards Aux.solvable_boards(@rows, @cols) 

  @spec sliding_puzzle(board) :: non_neg_integer() | -1
        when board: [[non_neg_integer()]]
  def sliding_puzzle(board) do
    Map.get(@solvable_boards, board, -1)
  end
end