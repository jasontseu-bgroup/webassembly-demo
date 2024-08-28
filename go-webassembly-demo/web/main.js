// Define functions to handle user interactions and display results
function handleClick() {
  const input1 = document.getElementById("input1").value;
  const input2 = document.getElementById("input2").value;
  const result = addNumbers(input1, input2);
  document.getElementById("output").textContent = result;
}

function addNumbers(input1, input2) {
  const result = addFunction(input1, input2);
  return result;
}

// Bind the functions to the window object
window.handleClick = handleClick;
window.addNumbers = addNumbers;