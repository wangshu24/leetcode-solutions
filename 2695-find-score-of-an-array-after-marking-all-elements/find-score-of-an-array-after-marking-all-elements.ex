defmodule Solution do
  @spec find_score(nums :: [integer]) :: integer
  def find_score(nums) do
    Enum.chunk_while(nums, [], fn
      b, [a | tail] when a <= b -> {:cont, [a | tail], []}
      b, list -> {:cont, [b | list]}
    end,
    fn
      list -> {:cont, list, []}
    end)
    |> Enum.flat_map(&Enum.take_every(&1, 2))
    |> Enum.sum()
  end
end