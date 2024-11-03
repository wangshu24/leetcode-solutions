defmodule Solution do
  @spec rotate_string(s :: String.t, goal :: String.t) :: boolean
  def rotate_string(s, goal) do
    if String.length(s) != String.length(goal) do
        false
    else
        String.contains? s <> s, goal
    end
    end
end