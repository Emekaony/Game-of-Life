function randomState(height, width) {
  const arr = [];
  const RAND_THRESHOLD = 0.4;

  for (let i = 0; i < height; i++) {
    arr[i] = [];
    for (let j = 0; j < width; j++) {
      if (Math.random() >= RAND_THRESHOLD) {
        arr[i][j] = 0;
      } else {
        arr[i][j] = 1;
      }
    }
  }

  return arr;
}

// function render(board) {
//   const ALIVE_CHAR = "*";
//   const DEAD_CHAR = " ";
//   console.log(board)

//   // for (let i = 0; i < board[1]; i++) {
//   //   console.log(i);
//   // }

//   // for (let i = 0; i < board.length; i++) {
//   //   console.log('-')
//   //   for (let j = 0; j < board[i].length; j++) {
//   //     console.log(j);
//   //   }
//   // }
// }

// render(randomState(10, 10));
const board = randomState(50, 10);

board.forEach((array) => {
  console.log("[" + array.join(" ") + "]");
});
