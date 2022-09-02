// Filters player_order array
// Condition: item has length
export default function getActivePlayersLength(player_order: string[]) {
  return player_order.filter((player: string) => player.length).length
}
