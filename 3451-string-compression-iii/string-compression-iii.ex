defmodule Solution do
  @spec compressed_string(word :: String.t) :: String.t
  def compressed_string(word) do
    Regex.scan(~r/(.)\1{0,8}/, word, capture: :first, return: :index)
    |> Enum.map(fn [{i, len}] ->
      [?0 + len, :binary.at(word, i)]
    end)
    |> IO.iodata_to_binary()
  end
end