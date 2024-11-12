defmodule Solution do
  @spec maximum_beauty(items :: [[integer]], queries :: [integer]) :: [integer]
  def maximum_beauty(items, queries) do
    items =
      items
      # Convert each item to {price, beauty}
      # because I think list is not the proper data structure
      # for storing heterogeneous data
      # (price and beauty are different things),
      # but tuples are exactly for such use cases.
      |> Enum.map(&List.to_tuple/1)
      # Sort by price (asc) then beauty (desc)
      # for the next step.
      |> Enum.sort_by(fn {price, beauty} -> {price, -beauty} end)
      # For each price, retain only the most beautiful item.
      |> Enum.uniq_by(&elem(&1, 0))
      # For each price, 
      # work out the maximum beauty of all the things with equal or lower price.
      |> Enum.scan({0, 0}, fn {price, beauty}, {_, max_beauty} ->
        {price, max(beauty, max_beauty)}
      end)
      # Convert the list to array
      # so that we can do binary search on it.
      # I'm not converting it to a tuple
      # because I think tuples should not be considered
      # as containers.
      |> :array.from_list()
    
    # For each query, do a binary search.
    Enum.map(queries, &search(items, 0, :array.size(items) - 1, &1))
  end

  defp search(items, i, j, query) when j <= i + 1 do
    [get(items, i), get(items, j)]
    |> Enum.reject(fn {price, _} -> price > query end)
    |> Enum.map(&elem(&1, 1))
    |> Enum.max(fn -> 0 end)
  end

  defp search(items, i, j, query) do
    m = div(i + j, 2)
    
    {price, beauty} = get(items, m)

    cond do
      query < price ->
        search(items, i, m - 1, query)

      query > price ->
        search(items, m, j, query)

      true ->
        beauty
    end
  end

  defp get(array, i) do
    :array.get(i, array)
  end
end