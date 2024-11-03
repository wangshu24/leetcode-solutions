defmodule Solution do
  @spec make_fancy_string(s :: String.t) :: String.t
  # Regex:
  def make_fancy_string(s) do
    String.replace(s, ~r/(.)\1+/, ~S"\1\1")
  end
end