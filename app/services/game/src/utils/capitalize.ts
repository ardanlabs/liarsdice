// Capitalizes a word.
export function capitalize(word: string) {
  return word.slice(0, 1).toUpperCase() + word.slice(1, word.length)
}
